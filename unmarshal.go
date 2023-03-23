package chalk

import (
	"errors"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type fqnToField map[string]reflect.Value

var FieldNotFoundError = errors.New("field not found")

type Numbers interface {
	int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func convertNumber[T Numbers](anyNumber any) T {
	// TODO: Possibly unmarshal numbers as the correct type (instead of float64)
	// into FeatureResult, instead of converting them here.
	switch typedNumber := anyNumber.(type) {
	case float64:
		return T(typedNumber)
	default:
		castedNumber, ok := anyNumber.(T)
		if !ok {
			var t T
			panic(fmt.Sprintf("exception occurred while unmarshaling online query result: cannot cast the number '%s' of type '%s' to the specified type '%s'", anyNumber, reflect.TypeOf(typedNumber), reflect.TypeOf(t)))
		}
		return castedNumber
	}
}

func convertSliceyNumbers[T Numbers](anySlice []any) []T {
	typedSlice := make([]T, len(anySlice))
	for i, v := range anySlice {
		typedSlice[i] = convertNumber[T](v)
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

func getPointerToCopied(elemType reflect.Type, value any) reflect.Value {
	copied := reflect.New(reflect.TypeOf(value))
	copied.Elem().Set(reflect.ValueOf(value))
	castedPointer := reflect.NewAt(elemType, copied.UnsafePointer())
	return castedPointer
}

func convertIfNumber(value any, kind reflect.Kind) any {
	// TODO: Figure out if we could possibly
	// do the equivalent by creating a new
	// reflect.Value with reflect.New
	// and reflect.Value.Set.
	switch kind {
	case reflect.Int8:
		value = convertNumber[int8](value)
	case reflect.Int16:
		value = convertNumber[int16](value)
	case reflect.Int32:
		value = convertNumber[int32](value)
	case reflect.Int64:
		value = convertNumber[int64](value)
	case reflect.Uint8:
		value = convertNumber[uint8](value)
	case reflect.Uint16:
		value = convertNumber[uint16](value)
	case reflect.Uint32:
		value = convertNumber[uint32](value)
	case reflect.Uint64:
		value = convertNumber[uint64](value)
	case reflect.Float32:
		value = convertNumber[float32](value)
	case reflect.Float64:
		value = convertNumber[float64](value)
	}

	return value
}

func convertNumberSlice(sliceElemKind reflect.Kind, value any) any {
	anySlice := value.([]any)
	switch sliceElemKind {
	case reflect.Int8:
		return convertSliceyNumbers[int8](anySlice)
	case reflect.Int16:
		return convertSliceyNumbers[int16](anySlice)
	case reflect.Int32:
		return convertSliceyNumbers[int32](anySlice)
	case reflect.Int64:
		return convertSliceyNumbers[int64](anySlice)
	case reflect.Uint8:
		return convertSliceyNumbers[uint8](anySlice)
	case reflect.Uint16:
		return convertSliceyNumbers[uint16](anySlice)
	case reflect.Uint32:
		return convertSliceyNumbers[uint32](anySlice)
	case reflect.Uint64:
		return convertSliceyNumbers[uint64](anySlice)
	case reflect.Float32:
		return convertSliceyNumbers[float32](anySlice)
	case reflect.Float64:
		return convertSliceyNumbers[float64](anySlice)
	case reflect.String:
		return convertSliceyNonNumbers[string](anySlice)
	case reflect.Bool:
		return convertSliceyNonNumbers[bool](anySlice)
	default:
		// TODO: Support non-primitive types?
		panic(fmt.Sprintf("unsupported slice type: %s", sliceElemKind))
	}
}

func getReflectValue(value any, elemType reflect.Type) (reflect.Value, error) {
	value = convertIfNumber(value, elemType.Kind())
	if elemType.String() == "time.Time" {
		stringValue := reflect.ValueOf(value).String()
		timeValue, timeErr := time.Parse(time.RFC3339, stringValue)
		if timeErr == nil {
			return reflect.ValueOf(&timeValue), nil
		}

		dateValue, dateErr := time.Parse("2006-01-02", stringValue)
		if dateErr != nil {
			// Return original datetime parsing error
			return reflect.Value{}, timeErr
		}
		return reflect.ValueOf(&dateValue), nil
	} else if elemType.Kind() == reflect.Slice || elemType.Kind() == reflect.Array {
		value = convertNumberSlice(elemType.Elem().Kind(), value)
		return getPointerToCopied(elemType, value), nil
	} else {
		return getPointerToCopied(elemType, value), nil
	}
}

func (t fqnToField) setFeature(fqn string, value any) error {
	if bucketDuration, baseFeatureField := getWindowedPseudofeatureMeta(fqn, t); bucketDuration != nil && baseFeatureField != nil {
		tagValue := reflect.ValueOf(internal.FormatBucketDuration(*bucketDuration))

		if baseFeatureField.Kind() != reflect.Map {
			panic(fmt.Sprintf("exception setting windowed feature '%s'", fqn))
		}

		reflectValue, err := getReflectValue(value, baseFeatureField.Type().Elem().Elem())
		if err != nil {
			return fmt.Errorf("error unmarshalling value for windowed feature %s: %w", fqn, err)
		}

		baseFeatureField.SetMapIndex(tagValue, reflectValue)
	} else {
		field, fieldFound := t[fqn]
		if !fieldFound {
			return FieldNotFoundError
		}
		reflectValue, err := getReflectValue(value, field.Type().Elem())
		if err != nil {
			return fmt.Errorf("error unmarshalling value for feature %s: %w", fqn, err)
		}
		field.Set(reflectValue)
	}

	return nil
}

func getWindowedPseudofeatureMeta(fqn string, fieldMap fqnToField) (*int, *reflect.Value) {
	sections := strings.Split(fqn, ".")
	lastSection := sections[len(sections)-1]

	lastSectionSplit := strings.Split(lastSection, "__")
	if len(lastSectionSplit) < 2 {
		return nil, nil
	}
	secondsStr := lastSectionSplit[1]
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return nil, nil
	}

	baseFeatureFqn := strings.Join(sections[:len(sections)-1], ".") + "." + lastSectionSplit[0]
	baseFeatureField, ok := fieldMap[baseFeatureFqn]
	if !ok {
		return nil, nil
	}

	return &seconds, &baseFeatureField
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
