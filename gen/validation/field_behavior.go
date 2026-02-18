package validation

import (
	"fmt"
	"slices"
	"sync"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// immutableFieldsCache stores the cached immutable field names for each message type
var immutableFieldsCache sync.Map // map[protoreflect.FullName][]string

// GetImmutableFields returns the list of field names that are marked as IMMUTABLE
// for the given message type. Results are cached for performance.
func GetImmutableFields(msg proto.Message) []string {
	msgReflect := msg.ProtoReflect()
	msgType := msgReflect.Descriptor().FullName()

	// Check cache first
	if cached, ok := immutableFieldsCache.Load(msgType); ok {
		return cached.([]string)
	}

	// Not in cache, compute it
	immutableFields := computeImmutableFields(msgReflect.Descriptor())

	// Store in cache atomically (returns existing value if another goroutine stored first)
	cached, _ := immutableFieldsCache.LoadOrStore(msgType, immutableFields)

	return cached.([]string)
}

// computeImmutableFields uses reflection to find all fields marked with IMMUTABLE field_behavior
func computeImmutableFields(msgDesc protoreflect.MessageDescriptor) []string {
	var immutableFields []string

	fields := msgDesc.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)

		// Get the field options
		opts := field.Options()
		if opts == nil {
			continue
		}

		// Check if field has field_behavior extension
		if !proto.HasExtension(opts, annotations.E_FieldBehavior) {
			continue
		}

		// Get the field_behavior values
		behaviors := proto.GetExtension(opts, annotations.E_FieldBehavior).([]annotations.FieldBehavior)

		// Check if IMMUTABLE is in the behaviors
		if slices.Contains(behaviors, annotations.FieldBehavior_IMMUTABLE) {
			immutableFields = append(immutableFields, string(field.Name()))
		}
	}

	return immutableFields
}

// ValidateImmutableFields checks that no immutable fields have changed between oldMsg and newMsg.
// Only fields specified in the fieldMaskPaths are checked.
// Returns an error if any immutable field was changed.
func ValidateImmutableFields(
	oldMsg proto.Message,
	newMsg proto.Message,
	fieldMaskPaths []string,
) error {
	if oldMsg == nil || newMsg == nil {
		return nil
	}

	// Get immutable fields for this message type
	immutableFields := GetImmutableFields(newMsg)
	if len(immutableFields) == 0 {
		return nil // No immutable fields to check
	}

	// Create a set of immutable field names for fast lookup
	immutableSet := make(map[string]bool, len(immutableFields))
	for _, field := range immutableFields {
		immutableSet[field] = true
	}

	// Check each field in the update mask
	oldReflect := oldMsg.ProtoReflect()
	newReflect := newMsg.ProtoReflect()

	for _, path := range fieldMaskPaths {
		// Skip if not an immutable field
		if !immutableSet[path] {
			continue
		}

		// Get field descriptors from each message (needed for dynamic messages where descriptors may differ)
		oldFieldDesc := oldReflect.Descriptor().Fields().ByName(protoreflect.Name(path))
		newFieldDesc := newReflect.Descriptor().Fields().ByName(protoreflect.Name(path))
		if oldFieldDesc == nil || newFieldDesc == nil {
			continue // Field not found (shouldn't happen with valid field mask)
		}

		// Get old and new values using their respective field descriptors
		oldValue := oldReflect.Get(oldFieldDesc)
		newValue := newReflect.Get(newFieldDesc)

		// Check if values are different
		if !oldValue.Equal(newValue) {
			return fmt.Errorf(
				"field %q is marked as IMMUTABLE and cannot be changed (attempted to change from %v to %v)",
				path,
				oldValue.Interface(),
				newValue.Interface(),
			)
		}
	}

	return nil
}
