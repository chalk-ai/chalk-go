package validation

import (
	"fmt"
	"slices"
	"strings"
	"sync"

	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// immutableFieldsCache stores the cached immutable field names for each message type
var immutableFieldsCache sync.Map // map[protoreflect.FullName][]string

// outputOnlyFieldsCache stores the cached output-only field names for each message type
var outputOnlyFieldsCache sync.Map // map[protoreflect.FullName][]string

// GetImmutableFields returns the list of field paths (including nested paths like "address.street")
// that are marked as IMMUTABLE for the given message type. Results are cached for performance.
func GetImmutableFields(msg proto.Message) []string {
	msgReflect := msg.ProtoReflect()
	msgType := msgReflect.Descriptor().FullName()

	if cached, ok := immutableFieldsCache.Load(msgType); ok {
		return cached.([]string)
	}

	visited := make(map[protoreflect.FullName]bool)
	immutableFields := computeFieldsByBehavior(msgReflect.Descriptor(), "", visited, annotations.FieldBehavior_IMMUTABLE)

	cached, _ := immutableFieldsCache.LoadOrStore(msgType, immutableFields)
	return cached.([]string)
}

// GetOutputOnlyFields returns the list of field paths (including nested paths like "address.street")
// that are marked as OUTPUT_ONLY for the given message type. Results are cached for performance.
func GetOutputOnlyFields(msg proto.Message) []string {
	msgReflect := msg.ProtoReflect()
	msgType := msgReflect.Descriptor().FullName()

	if cached, ok := outputOnlyFieldsCache.Load(msgType); ok {
		return cached.([]string)
	}

	visited := make(map[protoreflect.FullName]bool)
	outputOnlyFields := computeFieldsByBehavior(msgReflect.Descriptor(), "", visited, annotations.FieldBehavior_OUTPUT_ONLY)

	cached, _ := outputOnlyFieldsCache.LoadOrStore(msgType, outputOnlyFields)
	return cached.([]string)
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
				parentPath := strings.Join(parts[:i], ".")
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

// ValidateOutputOnlyFieldsNotSet returns an error if any OUTPUT_ONLY field is set in msg.
// Used in Create RPCs to reject caller-supplied output-only fields.
func ValidateOutputOnlyFieldsNotSet(msg proto.Message) error {
	if msg == nil {
		return nil
	}
	outputOnlyFields := GetOutputOnlyFields(msg)
	if len(outputOnlyFields) == 0 {
		return nil
	}

	msgReflect := msg.ProtoReflect()
	for _, path := range outputOnlyFields {
		if isFieldSet(msgReflect, path) {
			return fmt.Errorf("field %q is OUTPUT_ONLY and must not be set by the caller", path)
		}
	}
	return nil
}

// ValidateOutputOnlyFieldsNotInMask returns an error if any path in fieldMaskPaths refers to a
// field marked OUTPUT_ONLY. Used in Update RPCs to reject attempts to update output-only fields.
func ValidateOutputOnlyFieldsNotInMask(fieldMaskPaths []string, msg proto.Message) error {
	if msg == nil {
		return nil
	}
	outputOnlyFields := GetOutputOnlyFields(msg)
	if len(outputOnlyFields) == 0 {
		return nil
	}

	outputOnlySet := make(map[string]bool, len(outputOnlyFields))
	// outputOnlyByAncestor maps each ancestor path of an OUTPUT_ONLY field to that field,
	// allowing O(1) detection of mask paths that cover an OUTPUT_ONLY descendant.
	outputOnlyByAncestor := make(map[string]string)
	for _, f := range outputOnlyFields {
		outputOnlySet[f] = true
		parts := splitFieldPath(f)
		for i := 1; i < len(parts); i++ {
			outputOnlyByAncestor[strings.Join(parts[:i], ".")] = f
		}
	}

	for _, path := range fieldMaskPaths {
		if outputOnlySet[path] {
			return fmt.Errorf("field %q is OUTPUT_ONLY and cannot be updated", path)
		}
		// Check parent paths (e.g. if "address" is OUTPUT_ONLY, reject "address.street")
		parts := splitFieldPath(path)
		for i := len(parts) - 1; i > 0; i-- {
			parentPath := strings.Join(parts[:i], ".")
			if outputOnlySet[parentPath] {
				return fmt.Errorf("field %q is OUTPUT_ONLY and cannot be updated", parentPath)
			}
		}
		// Check child paths (e.g. if "address.street" is OUTPUT_ONLY, reject "address")
		if f, ok := outputOnlyByAncestor[path]; ok {
			return fmt.Errorf("field %q is OUTPUT_ONLY and cannot be updated", f)
		}
	}
	return nil
}

// computeFieldsByBehavior recursively finds all fields marked with the given field_behavior,
// including nested fields. When a field itself has the behavior, its subtree is not recursed
// (the top-level path covers it). The visited map prevents infinite recursion.
func computeFieldsByBehavior(msgDesc protoreflect.MessageDescriptor, prefix string, visited map[protoreflect.FullName]bool, behavior annotations.FieldBehavior) []string {
	msgType := msgDesc.FullName()
	if visited[msgType] {
		return nil
	}
	visited[msgType] = true
	defer delete(visited, msgType)

	var matchedFields []string

	fields := msgDesc.Fields()
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		fieldName := string(field.Name())

		var fieldPath string
		if prefix == "" {
			fieldPath = fieldName
		} else {
			fieldPath = prefix + "." + fieldName
		}

		opts := field.Options()
		if opts != nil && proto.HasExtension(opts, annotations.E_FieldBehavior) {
			behaviors := proto.GetExtension(opts, annotations.E_FieldBehavior).([]annotations.FieldBehavior)
			if slices.Contains(behaviors, behavior) {
				matchedFields = append(matchedFields, fieldPath)
				// Don't recurse — the whole subtree is covered by this path
				continue
			}
		}

		if field.Kind() == protoreflect.MessageKind && field.Message() != nil {
			nestedFields := computeFieldsByBehavior(field.Message(), fieldPath, visited, behavior)
			matchedFields = append(matchedFields, nestedFields...)
		}
	}

	return matchedFields
}

// fieldPathValue represents a value retrieved from a nested field path
type fieldPathValue struct {
	value  protoreflect.Value
	exists bool // false if the path couldn't be fully resolved
}

// navigateToField walks msg following each element of parts, returning the containing
// message and field descriptor of the final element. Returns (nil, nil) if the path
// cannot be resolved.
func navigateToField(msg protoreflect.Message, parts []string) (protoreflect.Message, protoreflect.FieldDescriptor) {
	currentMsg := msg
	for i, fieldName := range parts {
		fieldDesc := currentMsg.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
		if fieldDesc == nil {
			return nil, nil
		}
		if i == len(parts)-1 {
			return currentMsg, fieldDesc
		}
		if fieldDesc.Kind() != protoreflect.MessageKind ||
			fieldDesc.Cardinality() == protoreflect.Repeated ||
			fieldDesc.IsMap() {
			return nil, nil
		}
		nested := currentMsg.Get(fieldDesc)
		if !nested.Message().IsValid() {
			return nil, nil
		}
		currentMsg = nested.Message()
	}
	return nil, nil
}

// getValueAtPath retrieves the value at the given field path (e.g., "address.street").
// Returns a fieldPathValue with exists=false if the path cannot be resolved.
func getValueAtPath(msg protoreflect.Message, path string) fieldPathValue {
	parts := splitFieldPath(path)
	if len(parts) == 0 {
		return fieldPathValue{exists: false}
	}
	parent, fieldDesc := navigateToField(msg, parts)
	if parent == nil {
		return fieldPathValue{exists: false}
	}
	return fieldPathValue{value: parent.Get(fieldDesc), exists: true}
}

// isFieldSet returns true if the field at the given path is set to a non-default value.
func isFieldSet(msg protoreflect.Message, path string) bool {
	parts := splitFieldPath(path)
	if len(parts) == 0 {
		return false
	}
	parent, fieldDesc := navigateToField(msg, parts)
	if parent == nil {
		return false
	}
	return parent.Has(fieldDesc)
}

// splitFieldPath splits a field path by dots (e.g., "address.street" -> ["address", "street"])
func splitFieldPath(path string) []string {
	if path == "" {
		return nil
	}
	return strings.Split(path, ".")
}
