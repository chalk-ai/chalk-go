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
