package chalk

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type fqnToField map[string]reflect.Value

var FieldNotFoundError = errors.New("field not found")

type Numbers interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func convertSliceyNumbers[T Numbers](anySlice []any) []T {
	// TODO: Possibly unmarshal float numbers as the correct type into FeatureResult,
	// instead of converting them here.
	typedSlice := make([]T, len(anySlice))

	for i, v := range anySlice {
		switch typedV := v.(type) {
		case float64:
			typedSlice[i] = T(typedV)
		default:
			typedSlice[i] = v.(T)
		}
	}
	return typedSlice
}

func convertSliceyNonNumbers[T any](anySlice []any) []T {
	typedSlice := make([]T, len(anySlice))
	for i, v := range anySlice {
		typedSlice[i] = v.(T)
	}
	return typedSlice
}

func fieldSetConvertedValue(field reflect.Value, value any) {
	copied := reflect.New(reflect.TypeOf(value))
	copied.Elem().Set(reflect.ValueOf(value))
	castedPointer := reflect.NewAt(field.Type().Elem(), copied.UnsafePointer())
	field.Set(castedPointer)
}

func (t fqnToField) setFeature(fqn string, value any) error {
	field, fieldFound := t[fqn]
	if !fieldFound {
		return FieldNotFoundError
	}

	if field.Type().Elem().Kind() == reflect.Int {
		switch typed := value.(type) {
		case float64:
			value = int(typed)
		}
	}
	if field.Type().Elem().String() == "time.Time" {
		stringValue := reflect.ValueOf(value).String()
		timeValue, timeErr := time.Parse(time.RFC3339, stringValue)
		if timeErr == nil {
			field.Set(reflect.ValueOf(&timeValue))
			return nil
		}

		dateValue, dateErr := time.Parse("2006-01-02", stringValue)
		if dateErr != nil {
			// Return original datetime parsing error
			return timeErr
		}
		field.Set(reflect.ValueOf(&dateValue))
	} else if field.Type().Elem().Kind() == reflect.Slice || field.Type().Elem().Kind() == reflect.Array {
		elementKind := field.Type().Elem().Elem().Kind()
		anySlice := value.([]any)

		switch elementKind {
		case reflect.Int8:
			typedSlice := convertSliceyNumbers[int8](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Int16:
			typedSlice := convertSliceyNumbers[int16](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Int32:
			typedSlice := convertSliceyNumbers[int32](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Int64:
			typedSlice := convertSliceyNumbers[int64](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Uint8:
			typedSlice := convertSliceyNumbers[uint8](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Uint16:
			typedSlice := convertSliceyNumbers[uint16](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Uint32:
			typedSlice := convertSliceyNumbers[uint32](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Uint64:
			typedSlice := convertSliceyNumbers[uint64](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Float32:
			typedSlice := convertSliceyNumbers[float32](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Float64:
			typedSlice := convertSliceyNumbers[float64](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.String:
			typedSlice := convertSliceyNonNumbers[string](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		case reflect.Bool:
			typedSlice := convertSliceyNonNumbers[bool](anySlice)
			fieldSetConvertedValue(field, typedSlice)
		default:
			panic(fmt.Sprintf("unsupported slice type: %s", field.Type().Elem().Elem().String()))
		}
	} else {
		fieldSetConvertedValue(field, value)
	}

	return nil
}

func (result *OnlineQueryResult) unmarshal(t any) *ClientError {
	fieldMap := make(fqnToField)
	structValue := reflect.ValueOf(t).Elem()

	// Has a side effect: fieldMap will be populated.
	initFeatures(structValue, "", make(map[string]bool), fieldMap)

	for _, featureResult := range result.Data {
		fqn := featureResult.Field
		err := fieldMap.setFeature(fqn, featureResult.Value)
		if err != nil {
			structName := structValue.Type().String()
			outputNamespace := "unknown namespace"
			sections := strings.Split(fqn, ".")
			if len(sections) > 0 {
				outputNamespace = sections[0]
			}
			if errors.Is(err, FieldNotFoundError) {
				fieldError := fmt.Sprintf("Error unmarshaling feature '%s' into the struct '%s'. ", fqn, structName)
				fieldError += fmt.Sprintf("First, check if you are passing a pointer to a struct that represents the output namespace '%s'. ", outputNamespace)
				fieldError += fmt.Sprintf("Also, make sure the feature name can be traced to a field in the struct '%s' and or its nested structs.", structName)
				return &ClientError{Message: fieldError}
			} else {
				return &ClientError{Message: fmt.Sprintf("Unknown error unmarshaling feature '%s' into the struct '%s'.", fqn, structName)}
			}

		}
	}
	for _, expectedOutput := range result.expectedOutputs {
		if field, ok := fieldMap[expectedOutput]; ok {
			if field.IsNil() {
				// TODO: Handle optional fields
				return &ClientError{Message: fmt.Sprintf("Unexpected error unmarshaling output feature '%s'. Feature is still nil after unmarshaling", expectedOutput)}
			}
		}
	}
	return nil
}
