package internal

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	errors "github.com/pkg/errors"
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

func isTypeDataclass(typ reflect.Type) bool {
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
	return isTypeDataclass(field.Type())
}

func GetValueFromArrowArray(a arrow.Array, idx int) (any, error) {
	if a.IsNull(idx) {
		return nil, nil
	}
	switch arr := a.(type) {
	case *array.LargeList:
		newSlice := make([]any, 0)
		for ptr := arr.Offsets()[idx]; ptr < arr.Offsets()[idx+1]; ptr++ {
			anyVal, err := GetValueFromArrowArray(arr.ListValues(), int(ptr))
			if err != nil {
				return nil, errors.Wrap(err, "error getting value for LargeList column")
			}
			newSlice = append(newSlice, anyVal)
		}
		return newSlice, nil
	case *array.List:
		newSlice := make([]any, 0)
		for ptr := arr.Offsets()[idx]; ptr < arr.Offsets()[idx+1]; ptr++ {
			anyVal, err := GetValueFromArrowArray(arr.ListValues(), int(ptr))
			if err != nil {
				return nil, errors.Wrap(err, "error getting value for List column")
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
			anyVal, err := GetValueFromArrowArray(arr.Field(k), idx)
			if err != nil {
				return nil, errors.Wrap(err, "error getting value for Struct column")
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
				value, err := GetValueFromArrowArray(col, i)
				if err != nil {
					return nil, errors.Wrap(err, "error getting value from arrow array")
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
	if reflect.TypeOf(value).Kind() == reflect.Ptr {
		value = reflect.ValueOf(value).Elem().Interface()
	}
	valType := reflect.TypeOf(value)
	if isTypeDataclass(typ.Elem()) && (valType.Kind() == reflect.Slice || valType.Kind() == reflect.Map) {
		return nil
	}
	if typ.Elem().Kind() != valType.Kind() {
		return fmt.Errorf(
			"expected type '%s', got '%s'",
			typ.Elem().Kind().String(),
			valType.Kind().String(),
		)
	}
	return nil
}

func ReflectPtr(value reflect.Value) reflect.Value {
	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)
	return ptr
}

// GetReflectValue returns a reflect.Value of the given type from the given non-reflect value.
func GetReflectValue(value any, typ reflect.Type) (*reflect.Value, error) {
	if value == nil {
		return Ptr(reflect.Zero(typ)), nil
	}
	if reflect.ValueOf(value).Kind() == reflect.Ptr && typ.Kind() == reflect.Ptr {
		indirectValue, err := GetReflectValue(reflect.ValueOf(value).Elem().Interface(), typ.Elem())
		if err != nil {
			return nil, errors.Wrap(err, "error getting reflect value for pointed to value")
		}
		return Ptr(ReflectPtr(*indirectValue)), nil
	}
	if isTypeDataclass(typ) {
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
				//if err := ValidatePointer(memberValue, memberField.Type()); err != nil {
				//	return nil, err
				//}
				//rVal, err := GetReflectValue(memberValue, memberField.Type().Elem())
				rVal, err := GetReflectValue(&memberValue, memberField.Type())
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"error unmarshalling struct value for field '%s' in struct '%s'",
						pythonName, structValue.Type().Name(),
					)
				}
				//ptrToVal := reflect.New(rVal.Type())
				//ptrToVal.Elem().Set(*rVal)
				//memberField.Set(ptrToVal)
				memberField.Set(*rVal)
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
				//if err := ValidatePointer(v, memberField.Type()); err != nil {
				//	return nil, err
				//}
				//rVal, err := GetReflectValue(v, memberField.Type().Elem())
				rVal, err := GetReflectValue(&v, memberField.Type())
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"error unmarshalling struct value '%s' for struct '%s'",
						k, structValue.Type().Name(),
					)
				}
				//memberField.Set(ReflectPtr(*rVal))
				memberField.Set(*rVal)
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
			return nil, errors.Wrap(timeErr, "error parsing date string")
		}
		return Ptr(reflect.ValueOf(dateValue)), nil
	} else if typ.Kind() == reflect.Slice {
		// TODO: Nil value handling
		// 1. Check whether slice can have nullable values.
		// 1a. If yes, return a slice of pointers.
		// 1b. If no, return a slice of values.
		// 2. How do we determine whether a slice can have nullable values?
		// 2a. Check the type of the field.
		// 2aa. But GetReflectValue is independent of the field. No it is not.
		// 2ab. Find out why we don't get a slice of a poitner to a field.
		// 2b. This means we need to accept a pointer to the slice,
		//     or we introduce a param that says whether the slice
		//     is nullable.
		actualSlice := reflect.ValueOf(value)
		newSlice := reflect.MakeSlice(typ, 0, actualSlice.Len())
		for i := 0; i < actualSlice.Len(); i++ {
			actualValue := actualSlice.Index(i).Interface()
			if typ.Elem().Kind() == reflect.Ptr && actualValue != nil {
				if actualSlice.Index(i).Kind() == reflect.Interface {
					actualValue = ReflectPtr(actualSlice.Index(i).Elem()).Interface()
				}
			}
			rVal, err := GetReflectValue(actualValue, typ.Elem())
			if err != nil {
				return nil, errors.Wrap(err, "error getting reflect value for slice")
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
		return &rVal, nil
	}
}
