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

// GetImmutableFields returns the list of field paths (including nested paths like "address.street")
// that are marked as IMMUTABLE for the given message type. Results are cached for performance.
func GetImmutableFields(msg proto.Message) []string {
	msgReflect := msg.ProtoReflect()
	msgType := msgReflect.Descriptor().FullName()

	// Check cache first
	if cached, ok := immutableFieldsCache.Load(msgType); ok {
		return cached.([]string)
	}

	// Not in cache, compute it
	visited := make(map[protoreflect.FullName]bool)
	immutableFields := computeImmutableFields(msgReflect.Descriptor(), "", visited)

	// Store in cache atomically
	cached, _ := immutableFieldsCache.LoadOrStore(msgType, immutableFields)

	return cached.([]string)
}

// computeImmutableFields recursively finds all fields marked with IMMUTABLE field_behavior,
// including nested fields. The prefix parameter is used to build nested paths like "address.street".
// The visited map prevents infinite recursion on circular message references.
func computeImmutableFields(msgDesc protoreflect.MessageDescriptor, prefix string, visited map[protoreflect.FullName]bool) []string {
	// Prevent infinite recursion on circular references
	msgType := msgDesc.FullName()
	if visited[msgType] {
		return nil
	}
	visited[msgType] = true
	defer delete(visited, msgType)

	var immutableFields []string

	fields := msgDesc.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		fieldName := string(field.Name())

		// Build the full path for this field
		var fieldPath string
		if prefix == "" {
			fieldPath = fieldName
		} else {
			fieldPath = prefix + "." + fieldName
		}

		// Check if this field is immutable
		opts := field.Options()
		if opts != nil && proto.HasExtension(opts, annotations.E_FieldBehavior) {
			// Get the field_behavior values
			behaviors := proto.GetExtension(opts, annotations.E_FieldBehavior).([]annotations.FieldBehavior)

			// Check if IMMUTABLE is in the behaviors
			if slices.Contains(behaviors, annotations.FieldBehavior_IMMUTABLE) {
				immutableFields = append(immutableFields, fieldPath)
			}
		}

		// If this is a message field, recursively check its nested fields
		if field.Kind() == protoreflect.MessageKind && field.Message() != nil {
			nestedFields := computeImmutableFields(field.Message(), fieldPath, visited)
			immutableFields = append(immutableFields, nestedFields...)
		}
	}

	return immutableFields
}

// fieldPathValue represents a value retrieved from a nested field path
type fieldPathValue struct {
	value  protoreflect.Value
	exists bool // false if the path couldn't be fully resolved
}

// getValueAtPath retrieves the value at the given field path (e.g., "address.street").
// Returns a fieldPathValue with exists=false if the path cannot be resolved.
func getValueAtPath(msg protoreflect.Message, path string) fieldPathValue {
	parts := splitFieldPath(path)
	if len(parts) == 0 {
		return fieldPathValue{exists: false}
	}

	currentMsg := msg

	// Iteratively navigate through the path
	for i, fieldName := range parts {
		fieldDesc := currentMsg.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDesc == nil {
			return fieldPathValue{exists: false}
		}

		isLastPart := i == len(parts)-1

		if isLastPart {
			// We've reached the target field
			return fieldPathValue{
				value:  currentMsg.Get(fieldDesc),
				exists: true,
			}
		}

		// Not the last part - need to navigate deeper
		// Can't navigate through non-message fields, repeated fields, or maps
		if fieldDesc.Kind() != protoreflect.MessageKind ||
			fieldDesc.Cardinality() == protoreflect.Repeated ||
			fieldDesc.IsMap() {
			return fieldPathValue{exists: false}
		}

		value := currentMsg.Get(fieldDesc)
		if !value.Message().IsValid() {
			// Nested message doesn't exist
			return fieldPathValue{exists: false}
		}

		currentMsg = value.Message()
	}

	return fieldPathValue{exists: false}
}

// splitFieldPath splits a field path by dots (e.g., "address.street" -> ["address", "street"])
func splitFieldPath(path string) []string {
	if path == "" {
		return nil
	}

	parts := []string{}
	current := ""
	for _, ch := range path {
		if ch == '.' {
			if current != "" {
				parts = append(parts, current)
				current = ""
			}
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}
	return parts
}

// ValidateImmutableFields checks that no immutable fields have changed between oldMsg and newMsg.
// Only fields specified in the fieldMaskPaths are checked.
// Supports nested field paths (e.g., "address.street").
// Returns an error if any immutable field was changed.
func ValidateImmutableFields(
	oldMsg proto.Message,
	newMsg proto.Message,
	fieldMaskPaths []string,
) error {
	if oldMsg == nil || newMsg == nil {
		return nil
	}

	// Get all immutable fields for this message type (cached)
	immutableFields := GetImmutableFields(newMsg)
	if len(immutableFields) == 0 {
		return nil // No immutable fields to check
	}

	// Create a set of immutable field paths for fast lookup
	immutableSet := make(map[string]bool, len(immutableFields))
	for _, field := range immutableFields {
		immutableSet[field] = true
	}

	oldReflect := oldMsg.ProtoReflect()
	newReflect := newMsg.ProtoReflect()

	for _, path := range fieldMaskPaths {
		// Check if this path or any parent path is immutable
		// For "address.street", check both "address.street" and "address"
		immutablePath := ""
		if immutableSet[path] {
			immutablePath = path
		} else {
			// Check parent paths
			parts := splitFieldPath(path)
			for i := len(parts) - 1; i > 0; i-- {
				parentPath := ""
				for j := 0; j < i; j++ {
					if j > 0 {
						parentPath += "."
					}
					parentPath += parts[j]
				}
				if immutableSet[parentPath] {
					immutablePath = parentPath
					break
				}
			}
		}

		// Skip if this path and no parent paths are immutable
		if immutablePath == "" {
			continue
		}

		// Get values at this path for both messages
		oldPathValue := getValueAtPath(oldReflect, path)
		newPathValue := getValueAtPath(newReflect, path)

		// Skip if path doesn't exist in either message
		if !oldPathValue.exists || !newPathValue.exists {
			continue
		}

		// Field is immutable - check if it changed
		if !oldPathValue.value.Equal(newPathValue.value) {
			return fmt.Errorf(
				"field %q is marked as IMMUTABLE and cannot be changed (attempted to change from %v to %v)",
				immutablePath,
				oldPathValue.value.Interface(),
				newPathValue.value.Interface(),
			)
		}
	}

	return nil
}
