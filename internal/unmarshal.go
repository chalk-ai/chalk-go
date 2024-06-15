package internal

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"reflect"
	"time"
)

var tableReaderChunkSize = 10_000

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

var skipUnmarshalFeatureNames = map[string]bool{
	"__chalk_observed_at__": true,
}

var skipUnmarshalFields = map[string]bool{
	"__ts__": true,
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

func GetValueFromArrowArray(a arrow.Array, idx int) (any, error) {
	if a.IsNull(idx) {
		return nil, nil
	}
	switch arr := a.(type) {
	case *array.LargeList:
		newSlice := make([]any, 0)
		for ptr := arr.Offsets()[idx]; ptr < arr.Offsets()[idx+1]; ptr++ {
			anyVal, valueErr := GetValueFromArrowArray(arr.ListValues(), int(ptr))
			if valueErr != nil {
				return nil, fmt.Errorf("error getting value for LargeList column: %w", valueErr)
			}
			newSlice = append(newSlice, anyVal)
		}
		return newSlice, nil
	case *array.List:
		newSlice := make([]any, 0)
		for ptr := arr.Offsets()[idx]; ptr < arr.Offsets()[idx+1]; ptr++ {
			anyVal, valueErr := GetValueFromArrowArray(arr.ListValues(), int(ptr))
			if valueErr != nil {
				return nil, fmt.Errorf("error getting value for List column: %w", valueErr)
			}
			newSlice = append(newSlice, anyVal)
		}
		return newSlice, nil
	case *array.Struct:
		newMap := map[string]any{}
		structType, typeOk := arr.DataType().(*arrow.StructType)
		if !typeOk {
			return nil, fmt.Errorf("error getting struct type")
		}
		for k := 0; k < arr.NumField(); k++ {
			anyVal, valueErr := GetValueFromArrowArray(arr.Field(k), idx)
			if valueErr != nil {
				return nil, fmt.Errorf("error getting value for Struct column: %w", valueErr)
			}
			newMap[structType.Field(k).Name] = anyVal
		}
		return newMap, nil
	case *array.String:
		return arr.Value(idx), nil
	case *array.LargeString:
		return arr.Value(idx), nil
	case *array.Uint8:
		return arr.Value(idx), nil
	case *array.Uint16:
		return arr.Value(idx), nil
	case *array.Uint32:
		return arr.Value(idx), nil
	case *array.Uint64:
		return arr.Value(idx), nil
	case *array.Int16:
		return arr.Value(idx), nil
	case *array.Int32:
		return arr.Value(idx), nil
	case *array.Int64:
		return arr.Value(idx), nil
	case *array.Float64:
		return arr.Value(idx), nil
	case *array.Boolean:
		return arr.Value(idx), nil
	case *array.Date32:
		return arr.Value(idx).ToTime(), nil
	case *array.Date64:
		return arr.Value(idx).ToTime(), nil
	case *array.Timestamp:
		timeUnit := arr.DataType().(*arrow.TimestampType).TimeUnit()
		return arr.Value(idx).ToTime(timeUnit), nil
	default:
		return nil, fmt.Errorf("unsupported array type: %T", arr)
	}
}

func ExtractFeaturesFromTable(table arrow.Table) ([]map[string]any, error) {
	res := make([]map[string]any, 0)
	reader := array.NewTableReader(table, int64(tableReaderChunkSize))
	defer reader.Release()
	for reader.Next() {
		record := reader.Record()
		for i := 0; i < int(record.NumRows()); i++ {
			m := map[string]any{}
			for j, col := range record.Columns() {
				name := record.ColumnName(j)
				if _, ok := skipUnmarshalFields[name]; ok {
					continue
				}
				if _, ok := skipUnmarshalFeatureNames[getFeatureNameFromFqn(name)]; ok {
					continue
				}
				value, valueErr := GetValueFromArrowArray(col, i)
				if valueErr != nil {
					return nil, fmt.Errorf("error getting value from arrow array: %w", valueErr)
				}
				m[name] = value
			}
			res = append(res, m)
		}
	}
	return res, nil
}

func SliceAppend(slicePtr any, value reflect.Value) {
	slicePtrValue := reflect.ValueOf(slicePtr)
	sliceValue := slicePtrValue.Elem()
	sliceValue.Set(reflect.Append(sliceValue, value))
}

func ValidatePointer(value any, typ reflect.Type) error {
	if typ.Kind() != reflect.Ptr {
		return fmt.Errorf("expected field to be a pointer, found %s", typ.Kind().String())
	}
	if IsTypeDataclass(typ.Elem()) && reflect.TypeOf(value).Kind() == reflect.Slice {
		return nil
	}
	if typ.Elem() != reflect.TypeOf(value) {
		return fmt.Errorf(
			"expected type '%s', got '%s'",
			typ.Elem().Kind().String(),
			reflect.TypeOf(value).Kind().String(),
		)
	}
	return nil
}

func GetReflectValueNonPtr(value any, typ reflect.Type) (*reflect.Value, error) {
	if IsTypeDataclass(typ) {
		structValue := reflect.New(typ).Elem()
		if slice, isSlice := value.([]any); isSlice {
			if len(slice) != structValue.NumField() {
				return nil, fmt.Errorf(
					"error unmarshalling value for struct %s"+
						": expected %d fields, got %d",
					structValue.Type().Name(),
					structValue.NumField(),
					len(slice),
				)
			}
			for idx, memberValue := range slice {
				memberFieldMeta := structValue.Type().Field(idx)
				memberField := structValue.Field(idx)
				pythonName := ChalkpySnakeCase(memberFieldMeta.Name)
				if memberField == (reflect.Value{}) {
					return nil, fmt.Errorf(
						"field %s not found in struct %s",
						pythonName, structValue.Type().Name(),
					)
				}
				if err := ValidatePointer(memberValue, memberField.Type()); err != nil {
					return nil, err
				}
				rVal, err := GetReflectValueNonPtr(memberValue, memberField.Type().Elem())
				if err != nil {
					return nil, fmt.Errorf(
						"error unmarshalling struct value for field '%s' in struct '%s': %w",
						pythonName, structValue.Type().Name(), err,
					)
				}
				ptrToVal := reflect.New(rVal.Type())
				ptrToVal.Elem().Set(*rVal)
				memberField.Set(ptrToVal)
			}
			return &structValue, nil
		} else if mapz, isMap := value.(map[string]any); isMap {
			nameToField := make(map[string]reflect.Value)
			for i := 0; i < structValue.NumField(); i++ {
				nameToField[ChalkpySnakeCase(structValue.Type().Field(i).Name)] = structValue.Field(i)
			}
			for k, v := range mapz {
				memberField, fieldOk := nameToField[k]
				if !fieldOk {
					return nil, fmt.Errorf(
						"field %s not found in struct %s",
						k, structValue.Type().Name(),
					)
				}
				if err := ValidatePointer(v, memberField.Type()); err != nil {
					return nil, err
				}
				rVal, err := GetReflectValueNonPtr(v, memberField.Type().Elem())
				if err != nil {
					return nil, fmt.Errorf(
						"error unmarshalling struct value '%s' for struct '%s': %w",
						k, structValue.Type().Name(), err,
					)
				}
				memberField.Elem().Set(*rVal)
			}
			return &structValue, nil
		} else {
			return nil, fmt.Errorf(
				"struct value is not an `any` slice or a `map[string]any`",
			)
		}
	} else if typ == reflect.TypeOf(time.Time{}) {
		// Datetimes have already been unmarshalled into time.Time in bulk online query
		if reflect.TypeOf(value) == typ {
			if timeValue, ok := value.(time.Time); ok {
				// Need to cast to time type, otherwise
				// reflect.ValueOf(&timeValue) will give
				// us a reflect value of the pointer to
				// an interface.
				return Ptr(reflect.ValueOf(timeValue)), nil
			} else {
				return nil, fmt.Errorf(
					"error getting reflect value: expected `time.Time`, got %s",
					reflect.TypeOf(value),
				)
			}
		}

		// Datetimes are returned as strings in online query (non-bulk)
		stringValue := reflect.ValueOf(value).String()
		timeValue, timeErr := time.Parse(time.RFC3339, stringValue)
		if timeErr == nil {
			return Ptr(reflect.ValueOf(timeValue)), nil
		}

		// Dates are returned as strings in online query (non-bulk)
		dateValue, dateErr := time.Parse("2006-01-02", stringValue)
		if dateErr != nil {
			// Return original datetime parsing error
			return nil, fmt.Errorf("error parsing date string: %w", timeErr)
		}
		return Ptr(reflect.ValueOf(dateValue)), nil
	} else if typ.Kind() == reflect.Slice {
		actualSlice := reflect.ValueOf(value)
		newSlice := reflect.MakeSlice(typ, 0, actualSlice.Len())
		for i := 0; i < actualSlice.Len(); i++ {
			rVal, err := GetReflectValueNonPtr(actualSlice.Index(i).Interface(), typ.Elem())
			if err != nil {
				return nil, fmt.Errorf("error getting reflect value for slice: %w", err)
			}
			newSlice = reflect.Append(newSlice, *rVal)
		}
		return &newSlice, nil
	} else {
		rVal := reflect.ValueOf(value)
		if rVal.Kind() != typ.Kind() {
			return nil, fmt.Errorf(
				"expected reflect value of kind '%s', got '%s'",
				typ.Kind(),
				reflect.ValueOf(value).Kind(),
			)
		}
		return Ptr(rVal), nil
	}
}

func GetReflectValue(value any, elemType reflect.Type) (reflect.Value, error) {
	value, convErr := convertIfNumber(value, elemType.Kind())
	if convErr != nil {
		return reflect.Value{}, fmt.Errorf("error getting reflect value: %w", convErr)
	}
	if elemType == reflect.TypeOf(time.Time{}) {
		// Datetimes have already been unmarshalled into time.Time in bulk online query
		if reflect.TypeOf(value) == elemType {
			if timeValue, ok := value.(time.Time); ok {
				// Need to cast to time type, otherwise
				// reflect.ValueOf(&timeValue) will give
				// us a reflect value of the pointer to
				// an interface.
				return reflect.ValueOf(timeValue), nil
			} else {
				return reflect.Value{}, fmt.Errorf(
					"error getting reflect value: expected `time.Time`, got %s",
					reflect.TypeOf(value),
				)
			}
		}

		// Datetimes are returned as strings in online query (non-bulk)
		stringValue := reflect.ValueOf(value).String()
		timeValue, timeErr := time.Parse(time.RFC3339, stringValue)
		if timeErr == nil {
			return reflect.ValueOf(timeValue), nil
		}

		// Dates are returned as strings in online query (non-bulk)
		dateValue, dateErr := time.Parse("2006-01-02", stringValue)
		if dateErr != nil {
			// Return original datetime parsing error
			return reflect.Value{}, timeErr
		}
		return reflect.ValueOf(dateValue), nil
	} else if elemType.Kind() == reflect.Slice || elemType.Kind() == reflect.Array {
		// TODO: This can entirely be made redundant by consolidating
		//       with new handling of slices in setFeatureSingle.
		value, convErr = convertSlice(elemType.Elem().Kind(), value)
		if convErr != nil {
			return reflect.Value{}, fmt.Errorf("error getting reflect value: %w", convErr)
		}
		return getPointerToCopied(elemType, value), nil
	} else {
		rVal := reflect.ValueOf(value)
		if rVal.Kind() != elemType.Kind() {
			return reflect.Value{}, fmt.Errorf(
				"expected reflect value of kind '%s', got '%s'",
				elemType.Kind(),
				reflect.ValueOf(value).Kind(),
			)
		}
		return rVal, nil
	}
}

func IsTypeDataclass(typ reflect.Type) bool {
	if typ.Kind() == reflect.Struct {
		for i := 0; i < typ.NumField(); i++ {
			fieldMeta := typ.Field(i)
			if fieldMeta.Tag.Get("dataclass_field") == "true" {
				return true
			}
		}
	}
	return false
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
