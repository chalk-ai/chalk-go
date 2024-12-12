package internal

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"reflect"
	"time"
)

var tableReaderChunkSize = 10_000

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type NamespaceMemoItem struct {
	ResolvedFieldNameToIndex map[string]int
}

type NamespaceMemo map[string]*NamespaceMemoItem

var skipUnmarshalFeatureNames = map[string]bool{
	"__chalk_observed_at__": true,
}

var skipUnmarshalFields = map[string]bool{
	"__ts__":    true,
	"__index__": true,
}

var SkipUnmarshalFqnRoots = map[string]bool{
	"__chalk__": true,
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
	return IsTypeDataclass(field.Type())
}

func IsStruct(typ reflect.Type) bool {
	if typ.Kind() != reflect.Struct {
		// Not a dataclass nor a has-many feature.
		return false
	}

	if typ == reflect.TypeOf(time.Time{}) {
		return false
	}

	return true
}

func IsOrUnderlyingFeaturesClass(typ reflect.Type) bool {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return IsStruct(typ) && !IsTypeDataclass(typ)
}

func IsOrUnderlyingHasMany(typ reflect.Type) bool {
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}
	return typ.Kind() == reflect.Slice && IsOrUnderlyingFeaturesClass(typ.Elem())
}

func getInnerSliceFromArray(arr arrow.Array, offsets []int64, idx int) (any, error) {
	newSlice := make([]any, offsets[idx+1]-offsets[idx])
	newSliceIdx := 0
	for ptr := offsets[idx]; ptr < offsets[idx+1]; ptr++ {
		anyVal, err := GetValueFromArrowArray(arr, int(ptr))
		if err != nil {
			return nil, errors.Wrap(err, "error getting value for LargeList column")
		}
		newSlice[newSliceIdx] = anyVal
		newSliceIdx += 1
	}
	return newSlice, nil
}

func GetValueFromArrowArray(a arrow.Array, idx int) (any, error) {
	if a.IsNull(idx) {
		return nil, nil
	}
	switch arr := a.(type) {
	case *array.LargeList:
		return getInnerSliceFromArray(arr.ListValues(), arr.Offsets(), idx)
	case *array.List:
		o32 := arr.Offsets()
		o64 := make([]int64, len(o32))
		for i := 0; i < len(o32); i++ {
			o64[i] = int64(arr.Offsets()[i])
		}
		return getInnerSliceFromArray(arr.ListValues(), o64, idx)
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

func ReflectPtr(value reflect.Value) reflect.Value {
	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)
	return ptr
}

// GetReflectValue returns a reflect.Value of the given type from the given non-reflect value.
func GetReflectValue(value any, typ reflect.Type, nsMemo NamespaceMemo) (*reflect.Value, error) {
	if value == nil {
		return ptr.Ptr(reflect.Zero(typ)), nil
	}
	if reflect.ValueOf(value).Kind() == reflect.Ptr && typ.Kind() == reflect.Ptr {
		indirectValue, err := GetReflectValue(reflect.ValueOf(value).Elem().Interface(), typ.Elem(), nsMemo)
		if err != nil {
			return nil, errors.Wrap(err, "error getting reflect value for pointed to value")
		}
		return ptr.Ptr(ReflectPtr(*indirectValue)), nil
	}
	if IsStruct(typ) {
		structValue := reflect.New(typ).Elem()
		if slice, isSlice := value.([]any); isSlice {
			// Dataclasses come back as either slices or structs.
			// This is the slices case.
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
				resolvedName, err := ResolveFeatureName(memberFieldMeta)
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"error resolving name for field '%s' in struct '%s'",
						memberFieldMeta.Name,
						structValue.Type().Name(),
					)
				}
				if memberField == (reflect.Value{}) {
					return nil, fmt.Errorf(
						"member field '%s' not found in struct '%s'",
						resolvedName, structValue.Type().Name(),
					)
				}
				rVal, err := GetReflectValue(&memberValue, memberField.Type(), nsMemo)
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"error unmarshalling struct value for field '%s' in struct '%s'",
						resolvedName, structValue.Type().Name(),
					)
				}
				memberField.Set(*rVal)
			}
			return &structValue, nil
		} else if mapz, isMap := value.(map[string]any); isMap {
			// This could be either a dataclass or a feature class.
			memo, ok := nsMemo[structValue.Type().Name()]
			if !ok {
				return nil, fmt.Errorf(
					"namespace memo not found for struct '%s' - found %v",
					structValue.Type().Name(),
					colls.Keys(nsMemo),
				)
			}
			if memo.ResolvedFieldNameToIndex == nil {
				return nil, fmt.Errorf(
					"resolved field name to index map not found for struct '%s'",
					structValue.Type().Name(),
				)
			}
			for k, v := range mapz {
				memberFieldIdx, fieldOk := memo.ResolvedFieldNameToIndex[k]
				if !fieldOk {
					// For forward compatibility, i.e. when clients add
					// more fields to their dataclasses in chalkpy, we want
					// to default to not erring when trying to deserialize
					// a new field that does not yet exist in the Go struct.
					// Eventually we might consider exposing a flag.
					continue
				}
				memberField := structValue.Field(memberFieldIdx)
				if v == nil {
					continue
				}

				if memberField.Type().Kind() == reflect.Map {
					bucket, err := GetBucketFromFqn(k)
					if err != nil {
						return nil, errors.Wrapf(err, "error extracting bucket value for feature '%s'", k)
					}
					if err := SetMapEntryValue(memberField, bucket, v, nsMemo); err != nil {
						return nil, errors.Wrapf(
							err,
							"error setting map entry value for field '%s' in struct '%s'",
							k, structValue.Type().Name(),
						)
					}
				} else {
					rVal, err := GetReflectValue(&v, memberField.Type(), nsMemo)
					if err != nil {
						return nil, errors.Wrapf(
							err,
							"error unmarshalling struct value '%s' for struct '%s'",
							k, structValue.Type().Name(),
						)
					}
					memberField.Set(*rVal)
				}

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
				return ptr.Ptr(reflect.ValueOf(timeValue)), nil
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
			return ptr.Ptr(reflect.ValueOf(timeValue)), nil
		}

		// Dates are returned as strings in online query (non-bulk)
		dateValue, dateErr := time.Parse("2006-01-02", stringValue)
		if dateErr != nil {
			// Return original datetime parsing error
			return nil, errors.Wrap(timeErr, "error parsing date string")
		}
		return ptr.Ptr(reflect.ValueOf(dateValue)), nil
	} else if typ.Kind() == reflect.Slice {
		actualSlice := reflect.ValueOf(value)
		newSlice := reflect.MakeSlice(typ, 0, actualSlice.Len())
		for i := 0; i < actualSlice.Len(); i++ {
			actualValue := actualSlice.Index(i).Interface()
			if typ.Elem().Kind() == reflect.Ptr && actualValue != nil {
				if actualSlice.Index(i).Kind() == reflect.Interface {
					actualValue = ReflectPtr(actualSlice.Index(i).Elem()).Interface()
				} else {
					return nil, fmt.Errorf(
						"expected reflect value of kind 'interface', got '%s'",
						actualSlice.Index(i).Kind(),
					)
				}
			}
			rVal, err := GetReflectValue(actualValue, typ.Elem(), nsMemo)
			if err != nil {
				return nil, errors.Wrap(err, "error getting reflect value for slice")
			}
			newSlice = reflect.Append(newSlice, *rVal)
		}
		return &newSlice, nil
	} else {
		rVal := reflect.ValueOf(value)
		if rVal.Kind() != typ.Kind() {
			if rVal.Type().ConvertibleTo(typ) {
				rVal = rVal.Convert(typ)
			} else {
				return nil, KindMismatchError(typ.Kind(), rVal.Kind())
			}
		}
		return &rVal, nil
	}
}

// SetMapEntryValue exists as a separate special setter function because
// while all other fields are settable and can be passed into GetReflectValue
// to be set, map field values are not settable, and the entire map has to
// be passed instead.
func SetMapEntryValue(mapValue reflect.Value, key string, value any, nsMemo NamespaceMemo) error {
	if mapValue.IsNil() {
		mapType := mapValue.Type()
		newMap := reflect.MakeMap(mapType)
		mapValue.Set(newMap)
	}
	rVal, err := GetReflectValue(value, mapValue.Type().Elem().Elem(), nsMemo)
	if err != nil {
		return errors.Wrap(err, "error getting reflect value for map entry")
	}
	mapValue.SetMapIndex(reflect.ValueOf(key), ReflectPtr(*rVal))
	return nil
}
