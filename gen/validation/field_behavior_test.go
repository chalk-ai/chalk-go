package validation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

// Helper functions to create dynamic test messages for field behavior testing

type testField struct {
	name      string
	fieldType descriptorpb.FieldDescriptorProto_Type
	immutable bool
}

// createTestMessageDescriptor creates a dynamic message descriptor with specified fields and behaviors
func createTestMessageDescriptor(messageName string, fields []testField) (protoreflect.MessageDescriptor, error) {
	fieldDescriptors := make([]*descriptorpb.FieldDescriptorProto, len(fields))
	for i, f := range fields {
		fieldNum := int32(i + 1)
		fieldDescriptors[i] = &descriptorpb.FieldDescriptorProto{
			Name:   new(f.name),
			Number: new(fieldNum),
			Type:   f.fieldType.Enum(),
			Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		}

		if f.immutable {
			if fieldDescriptors[i].Options == nil {
				fieldDescriptors[i].Options = &descriptorpb.FieldOptions{}
			}

			behaviors := []annotations.FieldBehavior{annotations.FieldBehavior_IMMUTABLE}
			proto.SetExtension(fieldDescriptors[i].Options, annotations.E_FieldBehavior, behaviors)
		}
	}

	msgDesc := &descriptorpb.DescriptorProto{
		Name:  new(messageName),
		Field: fieldDescriptors,
	}

	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:        new(fmt.Sprintf("test_%s.proto", messageName)),
		Package:     new("test"),
		MessageType: []*descriptorpb.DescriptorProto{msgDesc},
		Dependency:  []string{"google/api/field_behavior.proto"},
	}

	fd, err := protodesc.NewFile(fileDesc, protoregistry.GlobalFiles)
	if err != nil {
		return nil, err
	}

	return fd.Messages().Get(0), nil
}

func getFieldBehaviorTestMessage1() (proto.Message, error) {
	desc, err := createTestMessageDescriptor("FieldBehaviorTestMessage", []testField{
		{name: "foo", fieldType: descriptorpb.FieldDescriptorProto_TYPE_STRING, immutable: true},
		{name: "bar", fieldType: descriptorpb.FieldDescriptorProto_TYPE_INT32, immutable: true},
		{name: "baz", fieldType: descriptorpb.FieldDescriptorProto_TYPE_STRING, immutable: false},
		{name: "qux", fieldType: descriptorpb.FieldDescriptorProto_TYPE_INT32, immutable: false},
	})
	if err != nil {
		return nil, err
	}
	return dynamicpb.NewMessage(desc), nil
}

func getFieldBehaviorTestMessage2() (proto.Message, error) {
	desc, err := createTestMessageDescriptor("AnotherFieldBehaviorTestMessage", []testField{
		{name: "foo", fieldType: descriptorpb.FieldDescriptorProto_TYPE_STRING, immutable: false},
		{name: "baz", fieldType: descriptorpb.FieldDescriptorProto_TYPE_STRING, immutable: true},
		{name: "quux", fieldType: descriptorpb.FieldDescriptorProto_TYPE_STRING, immutable: true},
	})
	if err != nil {
		return nil, err
	}
	return dynamicpb.NewMessage(desc), nil
}

func setField(msg proto.Message, fieldName string, value any) {
	m := msg.ProtoReflect()
	fd := m.Descriptor().Fields().ByName(protoreflect.Name(fieldName))
	if fd == nil {
		return
	}

	switch v := value.(type) {
	case string:
		m.Set(fd, protoreflect.ValueOfString(v))
	case int32:
		m.Set(fd, protoreflect.ValueOfInt32(v))
	}
}

func TestGetImmutableFields(t *testing.T) {
	t.Parallel()

	// Create dynamic test message with immutable fields
	msg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)

	immutableFields := GetImmutableFields(msg)

	// We expect exactly 2 immutable fields (foo and bar)
	expectedFields := []string{"foo", "bar"}

	assert.ElementsMatch(t, expectedFields, immutableFields, "Should identify all IMMUTABLE fields")
}

func TestGetImmutableFieldsCachingPerMessageType(t *testing.T) {
	t.Parallel()

	// Create two different dynamic message types
	msg1, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)

	msg2, err := getFieldBehaviorTestMessage2()
	require.NoError(t, err)

	// Get immutable fields for first message type
	fields1 := GetImmutableFields(msg1)
	require.ElementsMatch(t, []string{"foo", "bar"}, fields1)

	// Get immutable fields for second message type
	fields2 := GetImmutableFields(msg2)
	require.ElementsMatch(t, []string{"baz", "quux"}, fields2)

	// Get again for first message type - should return cached result
	fields1Again := GetImmutableFields(msg1)
	assert.Equal(t, fields1, fields1Again, "Should return cached result for first message type")

	// Get again for second message type - should return cached result
	fields2Again := GetImmutableFields(msg2)
	assert.Equal(t, fields2, fields2Again, "Should return cached result for second message type")

	// Verify they're still different (field "foo" is immutable in msg1, mutable in msg2)
	assert.NotEqual(t, fields1, fields2, "Different message types should have different immutable fields")
}

func TestValidateImmutableFieldsNoChanges(t *testing.T) {
	t.Parallel()

	oldMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(oldMsg, "foo", "test")
	setField(oldMsg, "bar", int32(42))
	setField(oldMsg, "baz", "old")
	setField(oldMsg, "qux", int32(100))

	newMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(newMsg, "foo", "test")
	setField(newMsg, "bar", int32(42))
	setField(newMsg, "baz", "new")      // Changed
	setField(newMsg, "qux", int32(200)) // Changed

	// Updating only mutable fields
	fieldMask := []string{"baz", "qux"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	assert.NoError(t, err, "Should allow changing mutable fields")
}

func TestValidateImmutableFieldsImmutableFieldChanged(t *testing.T) {
	t.Parallel()

	oldMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(oldMsg, "foo", "old-value")
	setField(oldMsg, "bar", int32(42))

	newMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(newMsg, "foo", "new-value") // Changed immutable field
	setField(newMsg, "bar", int32(42))

	// Attempting to update foo (immutable)
	fieldMask := []string{"foo"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	require.ErrorContains(t, err, "foo")
}

func TestValidateImmutableFieldsImmutableFieldUnchanged(t *testing.T) {
	t.Parallel()

	oldMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(oldMsg, "foo", "same")
	setField(oldMsg, "bar", int32(42))
	setField(oldMsg, "baz", "old")

	newMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(newMsg, "foo", "same")    // IMMUTABLE, but unchanged
	setField(newMsg, "bar", int32(42)) // IMMUTABLE, but unchanged
	setField(newMsg, "baz", "new")     // mutable, changed

	// Field mask includes IMMUTABLE fields, but they haven't changed
	fieldMask := []string{"foo", "bar", "baz"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	assert.NoError(t, err, "Should allow unchanged immutable fields in mask")
}

func TestValidateImmutableFieldsDifferentMessageTypes(t *testing.T) {
	t.Parallel()

	// In FieldBehaviorTestMessage: foo is IMMUTABLE, baz is MUTABLE
	// In AnotherFieldBehaviorTestMessage: foo is MUTABLE, baz is IMMUTABLE
	// This tests that validation uses the correct message type's annotations

	t.Run("FirstMessageType_FooIsImmutable", func(t *testing.T) {
		oldMsg, err := getFieldBehaviorTestMessage1()
		require.NoError(t, err)
		setField(oldMsg, "foo", "old")
		setField(oldMsg, "baz", "value")

		newMsg, err := getFieldBehaviorTestMessage1()
		require.NoError(t, err)
		setField(newMsg, "foo", "new") // IMMUTABLE in this message type
		setField(newMsg, "baz", "value")

		err = ValidateImmutableFields(oldMsg, newMsg, []string{"foo"})
		require.ErrorContains(t, err, "foo")
	})

	t.Run("SecondMessageType_FooIsMutable", func(t *testing.T) {
		oldMsg, err := getFieldBehaviorTestMessage2()
		require.NoError(t, err)
		setField(oldMsg, "foo", "old")
		setField(oldMsg, "baz", "value")

		newMsg, err := getFieldBehaviorTestMessage2()
		require.NoError(t, err)
		setField(newMsg, "foo", "new") // MUTABLE in this message type
		setField(newMsg, "baz", "value")

		err = ValidateImmutableFields(oldMsg, newMsg, []string{"foo"})
		assert.NoError(t, err, "Should allow changing foo in AnotherFieldBehaviorTestMessage")
	})

	t.Run("SecondMessageType_BazIsImmutable", func(t *testing.T) {
		oldMsg, err := getFieldBehaviorTestMessage2()
		require.NoError(t, err)
		setField(oldMsg, "foo", "value")
		setField(oldMsg, "baz", "old")

		newMsg, err := getFieldBehaviorTestMessage2()
		require.NoError(t, err)
		setField(newMsg, "foo", "value")
		setField(newMsg, "baz", "new") // IMMUTABLE in this message type

		err = ValidateImmutableFields(oldMsg, newMsg, []string{"baz"})
		require.ErrorContains(t, err, "baz")
	})
}

func TestValidateImmutableFieldsEmptyFieldMask(t *testing.T) {
	t.Parallel()

	oldMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(oldMsg, "foo", "old")

	newMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(newMsg, "foo", "new")

	// Empty field mask - no fields being updated
	err = ValidateImmutableFields(oldMsg, newMsg, []string{})
	assert.NoError(t, err, "Should allow empty field mask")
}

func TestValidateImmutableFieldsNilMessages(t *testing.T) {
	t.Parallel()

	err := ValidateImmutableFields(nil, nil, []string{"foo"})
	assert.NoError(t, err, "Should handle nil messages gracefully")

	msg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(msg, "foo", "test")

	err = ValidateImmutableFields(msg, nil, []string{"foo"})
	assert.NoError(t, err, "Should handle nil new message gracefully")

	err = ValidateImmutableFields(nil, msg, []string{"foo"})
	assert.NoError(t, err, "Should handle nil old message gracefully")
}

func TestValidateImmutableFieldsNonExistentField(t *testing.T) {
	t.Parallel()

	oldMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(oldMsg, "foo", "test")

	newMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(newMsg, "foo", "test")

	// Field mask with non-existent field
	fieldMask := []string{"non_existent_field"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	assert.NoError(t, err, "Should gracefully handle non-existent fields in mask")
}

// Helper to create a message with nested fields
func getNestedFieldBehaviorTestMessage() (proto.Message, error) {
	// Create nested message field descriptors
	nestedFieldDescriptors := []*descriptorpb.FieldDescriptorProto{
		{
			Name:    proto.String("street"),
			Number:  proto.Int32(1),
			Type:    descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Options: &descriptorpb.FieldOptions{},
		},
		{
			Name:   proto.String("city"),
			Number: proto.Int32(2),
			Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		},
		{
			Name:    proto.String("postal_code"),
			Number:  proto.Int32(3),
			Type:    descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Options: &descriptorpb.FieldOptions{},
		},
	}

	// Mark street and postal_code as immutable
	streetBehaviors := []annotations.FieldBehavior{annotations.FieldBehavior_IMMUTABLE}
	proto.SetExtension(nestedFieldDescriptors[0].Options, annotations.E_FieldBehavior, streetBehaviors)
	postalBehaviors := []annotations.FieldBehavior{annotations.FieldBehavior_IMMUTABLE}
	proto.SetExtension(nestedFieldDescriptors[2].Options, annotations.E_FieldBehavior, postalBehaviors)

	nestedMsgDesc := &descriptorpb.DescriptorProto{
		Name:  proto.String("NestedMessage"),
		Field: nestedFieldDescriptors,
	}

	// Create parent message field descriptors
	parentFieldDescriptors := []*descriptorpb.FieldDescriptorProto{
		{
			Name:    proto.String("id"),
			Number:  proto.Int32(1),
			Type:    descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Options: &descriptorpb.FieldOptions{},
		},
		{
			Name:     proto.String("address"),
			Number:   proto.Int32(2),
			Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
			TypeName: proto.String(".test.NestedMessage"),
			Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		},
		{
			Name:   proto.String("name"),
			Number: proto.Int32(3),
			Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		},
	}

	// Mark "id" as immutable
	idBehaviors := []annotations.FieldBehavior{annotations.FieldBehavior_IMMUTABLE}
	proto.SetExtension(parentFieldDescriptors[0].Options, annotations.E_FieldBehavior, idBehaviors)

	parentMsgDesc := &descriptorpb.DescriptorProto{
		Name:  proto.String("ParentMessage"),
		Field: parentFieldDescriptors,
	}

	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:        proto.String("test_nested.proto"),
		Package:     proto.String("test"),
		MessageType: []*descriptorpb.DescriptorProto{nestedMsgDesc, parentMsgDesc},
		Dependency:  []string{"google/api/field_behavior.proto"},
	}

	fd, err := protodesc.NewFile(fileDesc, protoregistry.GlobalFiles)
	if err != nil {
		return nil, err
	}

	return dynamicpb.NewMessage(fd.Messages().Get(1)), nil
}

func setNestedField(msg proto.Message, path string, value interface{}) {
	m := msg.ProtoReflect()

	// Handle nested paths like "address.street"
	parts := []string{}
	current := ""
	for _, ch := range path {
		if ch == '.' {
			parts = append(parts, current)
			current = ""
		} else {
			current += string(ch)
		}
	}
	if current != "" {
		parts = append(parts, current)
	}

	// Navigate to the nested field
	for i, part := range parts {
		fd := m.Descriptor().Fields().ByName(protoreflect.Name(part))
		if fd == nil {
			return
		}

		if i == len(parts)-1 {
			// Last part - set the value
			switch v := value.(type) {
			case string:
				m.Set(fd, protoreflect.ValueOfString(v))
			case int32:
				m.Set(fd, protoreflect.ValueOfInt32(v))
			}
		} else {
			// Navigate deeper
			if !m.Has(fd) {
				// Create the nested message if it doesn't exist
				m.Set(fd, protoreflect.ValueOfMessage(dynamicpb.NewMessage(fd.Message()).ProtoReflect()))
			}
			m = m.Get(fd).Message()
		}
	}
}

func TestValidateImmutableFieldsNestedFieldChanged(t *testing.T) {
	t.Parallel()

	oldMsg, err := getNestedFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(oldMsg, "id", "123")
	setNestedField(oldMsg, "address.street", "123 Old St")
	setNestedField(oldMsg, "address.city", "Boston")

	newMsg, err := getNestedFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(newMsg, "id", "123")
	setNestedField(newMsg, "address.street", "456 New St") // Changed immutable nested field
	setNestedField(newMsg, "address.city", "Boston")

	// Attempting to update nested immutable field
	fieldMask := []string{"address.street"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	require.Error(t, err, "Should not allow changing nested immutable field")
	assert.Contains(t, err.Error(), "street", "Error should mention the field name")
	assert.Contains(t, err.Error(), "IMMUTABLE", "Error should mention immutability")
}

func TestValidateImmutableFieldsNestedFieldUnchanged(t *testing.T) {
	t.Parallel()

	oldMsg, err := getNestedFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(oldMsg, "id", "123")
	setNestedField(oldMsg, "address.street", "123 Main St")
	setNestedField(oldMsg, "address.city", "Boston")

	newMsg, err := getNestedFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(newMsg, "id", "123")
	setNestedField(newMsg, "address.street", "123 Main St") // Unchanged immutable field
	setNestedField(newMsg, "address.city", "Cambridge")     // Changed mutable field

	// Field mask includes both changed mutable field and unchanged immutable field
	fieldMask := []string{"address.street", "address.city"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	assert.NoError(t, err, "Should allow unchanged nested immutable field")
}

func TestValidateImmutableFieldsNestedMutableFieldChanged(t *testing.T) {
	t.Parallel()

	oldMsg, err := getNestedFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(oldMsg, "address.street", "123 Main St")
	setNestedField(oldMsg, "address.city", "Boston")

	newMsg, err := getNestedFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(newMsg, "address.street", "123 Main St")
	setNestedField(newMsg, "address.city", "Cambridge") // Changed mutable nested field

	// Updating only mutable nested field
	fieldMask := []string{"address.city"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	assert.NoError(t, err, "Should allow changing nested mutable field")
}

func TestValidateImmutableFieldsMixedTopLevelAndNested(t *testing.T) {
	t.Parallel()

	oldMsg, err := getNestedFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(oldMsg, "id", "123")
	setNestedField(oldMsg, "name", "Alice")
	setNestedField(oldMsg, "address.street", "123 Main St")
	setNestedField(oldMsg, "address.postal_code", "02101")

	newMsg, err := getNestedFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(newMsg, "id", "123")                     // Unchanged immutable
	setNestedField(newMsg, "name", "Bob")                   // Changed mutable
	setNestedField(newMsg, "address.street", "123 Main St") // Unchanged nested immutable
	setNestedField(newMsg, "address.postal_code", "02101")  // Unchanged nested immutable

	// Mix of top-level and nested fields
	fieldMask := []string{"id", "name", "address.street", "address.postal_code"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	assert.NoError(t, err, "Should allow changes to mutable fields while leaving immutable fields unchanged")
}

// Helper to create a message with repeated fields
func getRepeatedFieldBehaviorTestMessage() (proto.Message, error) {
	// Create nested message descriptor
	nestedFieldDescriptors := []*descriptorpb.FieldDescriptorProto{
		{
			Name:    proto.String("street"),
			Number:  proto.Int32(1),
			Type:    descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Options: &descriptorpb.FieldOptions{},
		},
	}

	// Mark street as immutable
	streetBehaviors := []annotations.FieldBehavior{annotations.FieldBehavior_IMMUTABLE}
	proto.SetExtension(nestedFieldDescriptors[0].Options, annotations.E_FieldBehavior, streetBehaviors)

	nestedMsgDesc := &descriptorpb.DescriptorProto{
		Name:  proto.String("AddressItem"),
		Field: nestedFieldDescriptors,
	}

	// Create parent message with repeated field
	parentFieldDescriptors := []*descriptorpb.FieldDescriptorProto{
		{
			Name:     proto.String("addresses"),
			Number:   proto.Int32(1),
			Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
			TypeName: proto.String(".test.AddressItem"),
			Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		},
	}

	parentMsgDesc := &descriptorpb.DescriptorProto{
		Name:  proto.String("PersonWithAddresses"),
		Field: parentFieldDescriptors,
	}

	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:        proto.String("test_repeated.proto"),
		Package:     proto.String("test"),
		MessageType: []*descriptorpb.DescriptorProto{nestedMsgDesc, parentMsgDesc},
		Dependency:  []string{"google/api/field_behavior.proto"},
	}

	fd, err := protodesc.NewFile(fileDesc, protoregistry.GlobalFiles)
	if err != nil {
		return nil, err
	}

	return dynamicpb.NewMessage(fd.Messages().Get(1)), nil
}

func TestValidateImmutableFieldsWithRepeatedFieldPath(t *testing.T) {
	t.Parallel()

	oldMsg, err := getRepeatedFieldBehaviorTestMessage()
	require.NoError(t, err)

	newMsg, err := getRepeatedFieldBehaviorTestMessage()
	require.NoError(t, err)

	// Try to use a field mask that navigates through a repeated field
	// This should not panic, but should handle it gracefully
	fieldMask := []string{"addresses.street"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	// Should not panic and should handle gracefully (either skip or return error)
	assert.NoError(t, err, "Should gracefully handle field mask paths through repeated fields")
}

// Helper to create a message with map fields
func getMapFieldBehaviorTestMessage() (proto.Message, error) {
	// Create the map entry message descriptor (key-value pair)
	// For map<string, string>, we need a synthetic entry message
	mapEntryFields := []*descriptorpb.FieldDescriptorProto{
		{
			Name:   proto.String("key"),
			Number: proto.Int32(1),
			Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:  descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		},
		{
			Name:    proto.String("value"),
			Number:  proto.Int32(2),
			Type:    descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Options: &descriptorpb.FieldOptions{},
		},
	}

	// Mark the value as immutable
	valueBehaviors := []annotations.FieldBehavior{annotations.FieldBehavior_IMMUTABLE}
	proto.SetExtension(mapEntryFields[1].Options, annotations.E_FieldBehavior, valueBehaviors)

	mapEntryDesc := &descriptorpb.DescriptorProto{
		Name:    proto.String("AttributesEntry"),
		Field:   mapEntryFields,
		Options: &descriptorpb.MessageOptions{MapEntry: proto.Bool(true)},
	}

	// Create parent message with map field
	parentFieldDescriptors := []*descriptorpb.FieldDescriptorProto{
		{
			Name:     proto.String("attributes"),
			Number:   proto.Int32(1),
			Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
			TypeName: proto.String(".test.ResourceWithMap.AttributesEntry"),
			Label:    descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum(),
		},
	}

	parentMsgDesc := &descriptorpb.DescriptorProto{
		Name:       proto.String("ResourceWithMap"),
		Field:      parentFieldDescriptors,
		NestedType: []*descriptorpb.DescriptorProto{mapEntryDesc},
	}

	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:        proto.String("test_map.proto"),
		Package:     proto.String("test"),
		MessageType: []*descriptorpb.DescriptorProto{parentMsgDesc},
		Dependency:  []string{"google/api/field_behavior.proto"},
	}

	fd, err := protodesc.NewFile(fileDesc, protoregistry.GlobalFiles)
	if err != nil {
		return nil, err
	}

	return dynamicpb.NewMessage(fd.Messages().Get(0)), nil
}

func TestValidateImmutableFieldsWithMapFieldPath(t *testing.T) {
	t.Parallel()

	oldMsg, err := getMapFieldBehaviorTestMessage()
	require.NoError(t, err)

	newMsg, err := getMapFieldBehaviorTestMessage()
	require.NoError(t, err)

	// Try to use a field mask that navigates through a map field
	// This should not panic, but should handle it gracefully
	fieldMask := []string{"attributes.data"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	// Should not panic and should handle gracefully (skip validation for unresolvable path)
	assert.NoError(t, err, "Should gracefully handle field mask paths through map fields")
}

func TestValidateImmutableFieldsChangedToEmptyValue(t *testing.T) {
	t.Parallel()

	oldMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(oldMsg, "foo", "some-value")
	setField(oldMsg, "bar", int32(42))

	newMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(newMsg, "foo", "") // Changed immutable field to empty string
	setField(newMsg, "bar", int32(42))

	// Attempting to clear/empty immutable field
	fieldMask := []string{"foo"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	require.Error(t, err, "Should not allow changing immutable field to empty value")
	assert.Contains(t, err.Error(), "foo")
	assert.Contains(t, err.Error(), "IMMUTABLE")
	assert.Contains(t, err.Error(), "some-value")
}

func TestValidateImmutableFieldsChangedFromEmptyValue(t *testing.T) {
	t.Parallel()

	oldMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(oldMsg, "foo", "")
	setField(oldMsg, "bar", int32(0))

	newMsg, err := getFieldBehaviorTestMessage1()
	require.NoError(t, err)
	setField(newMsg, "foo", "new-value") // Changed immutable field from empty
	setField(newMsg, "bar", int32(0))

	// Attempting to change immutable field from empty to value
	fieldMask := []string{"foo"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	require.Error(t, err, "Should not allow changing immutable field from empty value")
	assert.Contains(t, err.Error(), "foo")
	assert.Contains(t, err.Error(), "IMMUTABLE")
}

// Helper to create a message with circular/recursive nesting
func getRecursiveFieldBehaviorTestMessage() (proto.Message, error) {
	// Create a message that references itself (like a linked list node)
	fieldDescriptors := []*descriptorpb.FieldDescriptorProto{
		{
			Name:    proto.String("id"),
			Number:  proto.Int32(1),
			Type:    descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Options: &descriptorpb.FieldOptions{},
		},
		{
			Name:    proto.String("value"),
			Number:  proto.Int32(2),
			Type:    descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
			Label:   descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Options: &descriptorpb.FieldOptions{},
		},
		{
			Name:     proto.String("next"),
			Number:   proto.Int32(3),
			Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
			TypeName: proto.String(".test.RecursiveNode"),
			Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		},
	}

	// Mark "id" as immutable
	idBehaviors := []annotations.FieldBehavior{annotations.FieldBehavior_IMMUTABLE}
	proto.SetExtension(fieldDescriptors[0].Options, annotations.E_FieldBehavior, idBehaviors)

	msgDesc := &descriptorpb.DescriptorProto{
		Name:  proto.String("RecursiveNode"),
		Field: fieldDescriptors,
	}

	fileDesc := &descriptorpb.FileDescriptorProto{
		Name:        proto.String("test_recursive.proto"),
		Package:     proto.String("test"),
		MessageType: []*descriptorpb.DescriptorProto{msgDesc},
		Dependency:  []string{"google/api/field_behavior.proto"},
	}

	fd, err := protodesc.NewFile(fileDesc, protoregistry.GlobalFiles)
	if err != nil {
		return nil, err
	}

	return dynamicpb.NewMessage(fd.Messages().Get(0)), nil
}

func TestGetImmutableFieldsWithRecursiveMessage(t *testing.T) {
	t.Parallel()

	msg, err := getRecursiveFieldBehaviorTestMessage()
	require.NoError(t, err)

	// This should not cause infinite recursion or panic
	immutableFields := GetImmutableFields(msg)

	// For recursive structures, we only enumerate top-level immutable fields
	// to avoid infinite enumeration (next.id, next.next.id, next.next.next.id, ...)
	assert.Len(t, immutableFields, 1, "Should find only top-level immutable fields in recursive structures")
	assert.Contains(t, immutableFields, "id", "Should find top-level immutable field")
}

func TestValidateImmutableFieldsWithRecursiveMessage(t *testing.T) {
	t.Parallel()

	oldMsg, err := getRecursiveFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(oldMsg, "id", "node1")
	setNestedField(oldMsg, "value", "old-value")

	newMsg, err := getRecursiveFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(newMsg, "id", "node1")        // Unchanged immutable
	setNestedField(newMsg, "value", "new-value") // Changed mutable

	// Update mutable field
	fieldMask := []string{"value"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	assert.NoError(t, err, "Should allow changing mutable fields in recursive structure")
}

func TestValidateImmutableFieldsWithRecursiveMessageImmutableChanged(t *testing.T) {
	t.Parallel()

	oldMsg, err := getRecursiveFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(oldMsg, "id", "old-id")

	newMsg, err := getRecursiveFieldBehaviorTestMessage()
	require.NoError(t, err)
	setNestedField(newMsg, "id", "new-id") // Changed top-level immutable

	fieldMask := []string{"id"}

	err = ValidateImmutableFields(oldMsg, newMsg, fieldMask)
	require.Error(t, err, "Should not allow changing top-level immutable field in recursive structure")
	assert.Contains(t, err.Error(), "id")
	assert.Contains(t, err.Error(), "IMMUTABLE")
}
