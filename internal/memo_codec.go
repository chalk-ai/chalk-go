package internal

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/cockroachdb/errors"
	"reflect"
)

type CodecMemoT = MemosT[string, Codec]

var CodecMemo = &CodecMemoT{}

type Codec func(structValue reflect.Value, arr arrow.Array, arrIdx int) error

var codecNoOp Codec = func(structValue reflect.Value, arr arrow.Array, arrIdx int) error {
	return nil
}

type InitFeatureFunc func(structValue reflect.Value) (leafStructValue reflect.Value, err error)

var initFeatureNoOp InitFeatureFunc = func(structValue reflect.Value) (reflect.Value, error) { return reflect.Value{}, nil }

type GetValueFunc func(arr arrow.Array, arrIdx int) (reflect.Value, error)

type SetMapFunc func(mapVal reflect.Value, entryValue reflect.Value)

type namespaceMeta struct {
	codec           *Codec
	rootStructIndex int
}

func generateUnmarshalValueCodec(structType reflect.Type, arrowType arrow.DataType, allMemo *NamespaceMemosT, fqn string, fqnParts []string) (*Codec, error) {
	var initFeatureFunc InitFeatureFunc
	if len(fqnParts) > 2 {
		// Is has-one feature, init all structs on its path
		initFunc, leafStructType, err := generateInitFeatureFunc(fqnParts, structType, allMemo)
		if err != nil {
			return nil, errors.Wrap(err, "generating function to initialize has-one feature")
		}
		initFeatureFunc = initFunc
		structType = leafStructType
	}

	memo, err := allMemo.LoadOrStore(structType)
	if err != nil {
		return nil, errors.Wrapf(err, "loading namespace memo for struct '%s'", structType.Name())
	}

	reflectFieldIndices, ok := memo.ResolvedFieldNameToIndices[fqnParts[len(fqnParts)-1]]
	if !ok || len(reflectFieldIndices) == 0 {
		// This happens when we unmarshal new feature fields into old codegen structs
		// We no-op here to allow for backcompat.
		return &codecNoOp, nil
	}

	fieldType := structType.Field(reflectFieldIndices[0]).Type

	var setMapFunc SetMapFunc
	if fieldType.Kind() == reflect.Map {
		bucket, err := GetBucketFromFqn(fqn)
		if err != nil {
			return nil, errors.Wrapf(err, "getting bucket from field name '%s'", fqn)
		}
		setMapFunc = generateSetMapFunc(bucket)

		// Unwrap map type to get value type
		fieldType = fieldType.Elem()
	}

	getValueFunc, err := generateGetValueFunc(fieldType, arrowType, allMemo)
	if err != nil {
		return nil, errors.Wrap(err, "generating function to get unmarshalled value from arrow array")
	}

	var codec Codec = func(structValue reflect.Value, arr arrow.Array, arrIdx int) error {
		if initFeatureFunc != nil {
			leafStructValue, err := initFeatureFunc(structValue)
			if err != nil {
				return errors.Wrap(err, "initializing has-one feature")
			}
			structValue = leafStructValue
		}
		reflectValue, err := getValueFunc(arr, arrIdx)
		if err != nil {
			return errors.Wrap(err, "getting reflect value")
		}
		if reflectValue.IsValid() {
			for _, fieldIdx := range reflectFieldIndices {
				if fieldIdx >= structValue.NumField() {
					return errors.Newf("field index %d out of range for struct '%s'", fieldIdx, structValue.Type())
				}
				field := structValue.Field(fieldIdx)
				if setMapFunc == nil {
					field.Set(reflectValue)
				} else {
					setMapFunc(field, reflectValue)
				}
			}
		}
		return nil

	}
	return &codec, nil
}

func generateInitFeatureFunc(fqnParts []string, structType reflect.Type, allMemo *NamespaceMemosT) (InitFeatureFunc, reflect.Type, error) {
	memo, err := allMemo.LoadOrStore(structType)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "loading namespace memo for struct '%s'", structType.Name())
	}

	structFieldsLen := len(fqnParts) - 2 // first is namespace, last is a scalar feature
	fieldIndexChain := make([]int, structFieldsLen)
	fieldIsPointer := make([]bool, structFieldsLen)

	currStructType := structType
	for i, fqnPart := range fqnParts[1 : len(fqnParts)-1] { // first is namespace, last is a scalar feature
		indices, ok := memo.ResolvedFieldNameToIndices[fqnPart]
		if !ok || len(indices) == 0 {
			// Reaching here might mean that a new feature field was
			// not yet added to the codegen'd structs. Here we return
			// an no-op function for back-compat.
			return initFeatureNoOp, nil, nil
		}

		firstFieldType := currStructType.Field(indices[0]).Type
		if firstFieldType.Kind() == reflect.Ptr {
			fieldIsPointer[i] = true
			firstFieldType = firstFieldType.Elem()
		}

		if !(IsStruct(firstFieldType) && !IsTypeDataclass(firstFieldType)) {
			// Not a has-one feature
			return nil, nil, errors.Newf(
				"field '%s' in struct '%s' is not a has-one feature",
				fqnPart, currStructType.Name(),
			)
		}

		if len(indices) > 1 {
			// A has-one feature is never versioned, hence never has multiple fields
			fieldNames := make([]string, len(indices))
			for j, idx := range indices {
				fieldNames[j] = currStructType.Field(idx).Name
			}
			return nil, nil, errors.Newf(
				"has-one feature '%s' in struct '%s' unexpectedly corresponds to multiple fields: %v",
				fqnPart, currStructType.Name(), fieldNames,
			)
		}

		currStructType = firstFieldType
		fieldIndexChain[i] = indices[0]
	}

	return func(structValue reflect.Value) (reflect.Value, error) {
		currValue := structValue
		for i, fieldIndex := range fieldIndexChain {
			innerStruct := currValue.Field(fieldIndex)
			if fieldIsPointer[i] {
				if innerStruct.IsNil() {
					innerStruct.Set(reflect.New(currStructType))
				}
				innerStruct = innerStruct.Elem()
			}
			currValue = innerStruct
		}
		return currValue, nil
	}, currStructType, nil
}

func generateSetMapFunc(bucket string) SetMapFunc {
	key := reflect.ValueOf(bucket)
	return func(mapVal reflect.Value, entryValue reflect.Value) {
		if mapVal.IsNil() {
			mapVal.Set(reflect.MakeMap(mapVal.Type()))
		}
		mapVal.SetMapIndex(key, entryValue)
	}
}

func generateGetSliceFunc(sliceReflectType reflect.Type, elemArrowType arrow.DataType, allMemo *NamespaceMemosT) (func(arr arrow.Array, startIdx int, endIdx int) (reflect.Value, error), error) {
	var sliceType reflect.Type
	isPointer := sliceReflectType.Kind() == reflect.Ptr
	if isPointer {
		sliceType = sliceReflectType.Elem()
	} else {
		sliceType = sliceReflectType
	}

	if sliceType.Kind() != reflect.Slice {
		return nil, errors.New("sliceReflectType must be a slice type, found: " + sliceType.Kind().String())
	}

	codec, err := generateGetValueFunc(sliceType.Elem(), elemArrowType, allMemo)
	if err != nil {
		return nil, errors.Wrap(err, "generating getValue function for array element")
	}

	return func(arr arrow.Array, startIdx int, endIdx int) (reflect.Value, error) {
		length := endIdx - startIdx
		newSlice := reflect.MakeSlice(sliceType, length, length)
		newSliceIdx := 0
		for currIdx := startIdx; currIdx < endIdx; currIdx++ {
			val, err := codec(arr, currIdx)
			if err != nil {
				return reflect.Value{}, errors.Wrapf(
					err,
					"building slice value for row at underlying index %d",
					currIdx,
				)
			}
			if val.IsValid() {
				newSlice.Index(newSliceIdx).Set(val)
			}
			newSliceIdx += 1
		}

		if isPointer {
			slicePtr := reflect.New(newSlice.Type())
			slicePtr.Elem().Set(newSlice)
			return slicePtr, nil
		} else {
			return newSlice, nil
		}
	}, nil

}

func generateGetValueFunc(fieldType reflect.Type, arrowType arrow.DataType, allMemo *NamespaceMemosT) (GetValueFunc, error) {
	codec, err := generateGetValueFuncInner(fieldType, arrowType, allMemo)
	if err != nil {
		return nil, err
	}
	return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
		if arr.IsNull(arrIdx) {
			return reflect.Value{}, nil
		}
		return codec(arr, arrIdx)
	}, nil
}

func generateGetValueFuncInner(fieldType reflect.Type, arrowType arrow.DataType, allMemo *NamespaceMemosT) (GetValueFunc, error) {
	switch castArrType := arrowType.(type) {
	case *arrow.LargeListType:
		getSliceFunc, err := generateGetSliceFunc(fieldType, castArrType.Elem(), allMemo)
		if err != nil {
			return nil, errors.Wrap(err, "generating getSlice function for lage list")
		}
		return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
			// FIXME: Check if ok everywhere to avoid panic
			castArr := arr.(*array.LargeList)
			return getSliceFunc(castArr.ListValues(), int(castArr.Offsets()[arrIdx]), int(castArr.Offsets()[arrIdx+1]))
		}, err
	case *arrow.ListType:
		getSliceFunc, err := generateGetSliceFunc(fieldType, castArrType.Elem(), allMemo)
		if err != nil {
			return nil, errors.Wrap(err, "generating getSlice function for list")
		}
		return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
			castArr := arr.(*array.List)
			return getSliceFunc(castArr.ListValues(), int(castArr.Offsets()[arrIdx]), int(castArr.Offsets()[arrIdx+1]))
		}, err
	case *arrow.StructType:
		var structType reflect.Type
		isPointer := fieldType.Kind() == reflect.Ptr
		if isPointer {
			structType = fieldType.Elem()
		} else {
			structType = fieldType
		}

		memo, err := allMemo.LoadOrStore(structType)
		if err != nil {
			return nil, errors.Wrapf(err, "loading namespace memo for struct '%s'", structType.Name())
		}

		arrowFieldToGetValueFunc := make([]GetValueFunc, castArrType.NumFields())
		arrowFieldToSetMapFunc := make([]SetMapFunc, castArrType.NumFields())
		arrowFieldToReflectFieldIndices := make([][]int, castArrType.NumFields())
		for k := 0; k < castArrType.NumFields(); k++ {
			fieldName := castArrType.Field(k).Name
			reflectFieldIndices, ok := memo.ResolvedFieldNameToIndices[fieldName]
			if !ok {
				// For backcompat with old codegen'd structs, because server may
				// return new features that are not yet in the codegen'd structs.
				continue
			}

			if len(reflectFieldIndices) == 0 {
				return nil, errors.Newf(
					"no indices found for field '%s' in struct '%s'",
					fieldName, structType.Name(),
				)
			}

			// When we have multiple indices, it means the field is a versioned feature.
			// Different versions all have the same type, so just use the first type.
			firstFieldType := structType.Field(reflectFieldIndices[0]).Type

			if firstFieldType.Kind() == reflect.Map {
				bucket, err := GetBucketFromFqn(fieldName)
				if err != nil {
					return nil, errors.Wrapf(
						err,
						"getting bucket from field name '%s'",
						fieldName,
					)
				}
				arrowFieldToSetMapFunc[k] = generateSetMapFunc(bucket)

				// Unwrap the map type to get value type
				firstFieldType = firstFieldType.Elem()
			}

			getValueFunc, err := generateGetValueFunc(firstFieldType, castArrType.Field(k).Type, allMemo)
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"getting codec for struct '%s' field: %s",
					structType.Name(), castArrType.Field(k).Name,
				)
			}

			arrowFieldToGetValueFunc[k] = getValueFunc
			arrowFieldToReflectFieldIndices[k] = reflectFieldIndices
		}

		return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
			newStructPtr := reflect.New(structType)
			for k := 0; k < castArrType.NumFields(); k++ {
				getValueFunc := arrowFieldToGetValueFunc[k]
				if getValueFunc == nil {
					return reflect.Value{}, errors.Newf(
						"no get value function found for struct '%s' field: %s",
						structType.Name(), castArrType.Field(k).Name,
					)
				}
				reflectFieldIndices := arrowFieldToReflectFieldIndices[k]

				value, err := getValueFunc(arr.(*array.Struct).Field(k), arrIdx)
				if err != nil {
					return reflect.Value{}, errors.Wrapf(
						err,
						"getting value for struct '%s' field: %s",
						structType.Name(), castArrType.Field(k).Name,
					)
				}

				if value.IsValid() {
					setMapFunc := arrowFieldToSetMapFunc[k]
					for _, fieldIdx := range reflectFieldIndices {
						if setMapFunc == nil {
							newStructPtr.Elem().Field(fieldIdx).Set(value)
						} else {
							setMapFunc(newStructPtr.Elem().Field(fieldIdx), value)
						}
					}
				}
			}

			if isPointer {
				return newStructPtr, nil
			} else {
				return newStructPtr.Elem(), nil
			}
		}, nil
	case *arrow.StringType:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.String).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.String).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.LargeStringType:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.LargeString).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.LargeString).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Uint8Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Uint8).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Uint8).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Uint16Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Uint16).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Uint16).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Uint32Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Uint32).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Uint32).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Uint64Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Uint64).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Uint64).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Int8Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Int8).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Int8).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Int16Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Int16).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Int16).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Int32Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Int32).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Int32).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Int64Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Int64).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Int64).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Float32Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Float32).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Float32).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.Float64Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Float64).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Float64).Value(arrIdx)), nil
			}, nil
		}
	case *arrow.BooleanType:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				val := arr.(*array.Boolean).Value(arrIdx)
				return reflect.ValueOf(&val), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Boolean).Value(arrIdx)), nil
			}, nil
		}

	case *arrow.Date32Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				timeVal := arr.(*array.Date32).Value(arrIdx).ToTime()
				return reflect.ValueOf(&timeVal), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Date32).Value(arrIdx).ToTime()), nil
			}, nil
		}
	case *arrow.Date64Type:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				timeVal := arr.(*array.Date64).Value(arrIdx).ToTime()
				return reflect.ValueOf(&timeVal), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Date64).Value(arrIdx).ToTime()), nil
			}, nil
		}
	case *arrow.TimestampType:
		if fieldType.Kind() == reflect.Ptr {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				timeVal := arr.(*array.Timestamp).Value(arrIdx).ToTime(castArrType.TimeUnit())
				return reflect.ValueOf(&timeVal), nil
			}, nil
		} else {
			return func(arr arrow.Array, arrIdx int) (reflect.Value, error) {
				return reflect.ValueOf(arr.(*array.Timestamp).Value(arrIdx).ToTime(castArrType.TimeUnit())), nil
			}, nil
		}
	default:
		return nil, errors.Newf("unsupported array type: %T", arrowType)
	}
}
