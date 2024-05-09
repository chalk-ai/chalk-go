package internal

import (
	"fmt"
	"reflect"
	"time"
)

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func convertNumber[T Numbers](anyNumber any) (T, error) {
	// TODO: Possibly unmarshal numbers as the correct type (instead of float64)
	// into FeatureResult, instead of converting them here.
	switch typedNumber := anyNumber.(type) {
	case float64:
		return T(typedNumber), nil
	default:
		castedNumber, ok := anyNumber.(T)
		if !ok {
			var t T
			return t, fmt.Errorf("cannot cast the number '%s' of type '%s' to the specified type '%s'", anyNumber, reflect.TypeOf(typedNumber), reflect.TypeOf(t))
		}
		return castedNumber, nil
	}
}

func convertSliceyNumbers[T Numbers](anySlice []any) ([]T, error) {
	typedSlice := make([]T, len(anySlice))
	for i, v := range anySlice {
		convRes, convErr := convertNumber[T](v)
		if convErr != nil {
			return nil, fmt.Errorf("error converting number-slice element: %w", convErr)
		}
		typedSlice[i] = convRes
	}
	return typedSlice, nil
}

func convertSliceyNonNumbers[T any](anySlice []any) ([]T, error) {
	typedSlice := make([]T, len(anySlice))
	for i, v := range anySlice {
		castRes, ok := v.(T)
		if !ok {
			var t T
			return []T{}, fmt.Errorf("cannot cast the slice element '%s' of type '%s' to the specified type '%s'", v, reflect.TypeOf(v), reflect.TypeOf(t))
		}
		typedSlice[i] = castRes
	}
	return typedSlice, nil
}

func convertIfNumber(value any, kind reflect.Kind) (any, error) {
	// TODO: Figure out if we could possibly
	// do the equivalent by creating a new
	// reflect.Value with reflect.New
	// and reflect.Value.Set.
	var err error
	switch kind {
	case reflect.Int8:
		value, err = convertNumber[int8](value)
	case reflect.Int16:
		value, err = convertNumber[int16](value)
	case reflect.Int32:
		value, err = convertNumber[int32](value)
	case reflect.Int64:
		value, err = convertNumber[int64](value)
	case reflect.Uint8:
		value, err = convertNumber[uint8](value)
	case reflect.Uint16:
		value, err = convertNumber[uint16](value)
	case reflect.Uint32:
		value, err = convertNumber[uint32](value)
	case reflect.Uint64:
		value, err = convertNumber[uint64](value)
	case reflect.Float32:
		value, err = convertNumber[float32](value)
	case reflect.Float64:
		value, err = convertNumber[float64](value)
	}

	return value, err
}

func convertSlice(sliceElemKind reflect.Kind, value any) (any, error) {
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
		return nil, fmt.Errorf("unsupported slice type '%s' when converting slice", sliceElemKind)
	}
}

func getPointerToCopied(elemType reflect.Type, value any) reflect.Value {
	copied := reflect.New(reflect.TypeOf(value))
	copied.Elem().Set(reflect.ValueOf(value))
	castedPointer := reflect.NewAt(elemType, copied.UnsafePointer())
	return castedPointer
}

func GetReflectValue(value any, elemType reflect.Type) (reflect.Value, error) {
	value, convErr := convertIfNumber(value, elemType.Kind())
	if convErr != nil {
		return reflect.Value{}, fmt.Errorf("error getting reflect value: %w", convErr)
	}
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
		// FIXME: This needs to recurse
		// The element type might be:
		// - primitive
		// - dataclass
		// - NOT FeaturesClass
		value, convErr = convertSlice(elemType.Elem().Kind(), value)
		if convErr != nil {
			return reflect.Value{}, fmt.Errorf("error getting reflect value: %w", convErr)
		}
		return getPointerToCopied(elemType, value), nil
	} else {
		// Do a switch on the type of the value and then validate it with the elemType
		// to ensure that the value is of the correct type.
		// FIXME: Validate for all types
		switch elemType.Kind() {
		case reflect.String:
			strVal, ok := value.(string)
			if !ok {
				return reflect.Value{}, fmt.Errorf("error getting reflect value: expected string, got %s", reflect.TypeOf(value))
			}
			return getPointerToCopied(elemType, strVal), nil
		}
		return getPointerToCopied(elemType, value), nil
	}
}

func IsDataclass(field reflect.Value) bool {
	if field.Kind() == reflect.Struct {
		if field.NumField() == 0 {
			return false
		}
		for i := 0; i < field.NumField(); i++ {
			fieldMeta := field.Type().Field(i)
			if fieldMeta.Tag.Get("dataclass_field") != "true" {
				return false
			}
		}
		return true
	}

	return false
}

func IsDataclassPointer(field reflect.Value) bool {
	if field.Kind() == reflect.Ptr && IsDataclass(field.Elem()) {
		return true
	}
	return false
}
