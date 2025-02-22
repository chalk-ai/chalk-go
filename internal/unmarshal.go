package internal

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/apache/arrow/go/v16/arrow/array"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const tableReaderChunkSizeKey = "CHALK_TABLE_READER_CHUNK_SIZE"
const defaultTableReaderChunkSize = 10_000
const metadataPrefix = "__chalk__.__result_metadata__."
const pkeyField = "__id__"

type ResultMetadataSourceType string

const (
	SourceTypeOnlineStore ResultMetadataSourceType = "online_store"
)

var tableReaderChunkSize = defaultTableReaderChunkSize

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type NamespaceMemo struct {
	// Root and non-root FQN as keys
	ResolvedFieldNameToIndices map[string][]int
	// Non-root FQN as keys only
	StructFieldsSet map[string]bool
}

type AllNamespaceMemoT sync.Map

var AllNamespaceMemo = &AllNamespaceMemoT{}

func init() {
	if chunkSizeStr := os.Getenv(tableReaderChunkSizeKey); chunkSizeStr != "" {
		if newChunkSize, err := strconv.Atoi(chunkSizeStr); err == nil {
			tableReaderChunkSize = newChunkSize
		}
	}
}

func NewNamespaceMemo() *NamespaceMemo {
	return &NamespaceMemo{
		ResolvedFieldNameToIndices: map[string][]int{},
		StructFieldsSet:            map[string]bool{},
	}
}

type NamespaceMutex struct {
	mu   sync.RWMutex
	memo *NamespaceMemo
}

func (m *AllNamespaceMemoT) Load(key reflect.Type) (*NamespaceMemo, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if !ok {
		return nil, false
	}
	namespaceMutex := value.(*NamespaceMutex)
	namespaceMutex.mu.RLock()
	defer namespaceMutex.mu.RUnlock()
	return namespaceMutex.memo, true
}

func (m *AllNamespaceMemoT) LoadOrStoreLockedMutex(key reflect.Type, memo *NamespaceMemo) (*NamespaceMutex, bool) {
	namespaceMutex := &NamespaceMutex{
		mu:   sync.RWMutex{},
		memo: memo,
	}
	// The mutex should always be inserted in a locked state.
	// Otherwise, there is an albeit slim possibility where
	// concurrent readers of the memo retrieve and obtain an
	// RLock on the mutex before the writer obtains a write
	// lock.
	namespaceMutex.mu.Lock()
	v, loaded := (*sync.Map)(m).LoadOrStore(key, namespaceMutex)
	return v.(*NamespaceMutex), loaded
}

func (m *AllNamespaceMemoT) Keys() []reflect.Type {
	keys := []reflect.Type{}
	(*sync.Map)(m).Range(func(key, _ any) bool {
		keys = append(keys, key.(reflect.Type))
		return true
	})
	return keys
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

func IsFeaturesClass(typ reflect.Type) bool {
	return IsStruct(typ) &&
		!IsTypeDataclass(typ) &&
		typ.Name() != "" // CHA-5430
}

func getInnerSliceFromArray(arr arrow.Array, offsets []int64, idx int, timeAsString bool) (any, error) {
	newSlice := make([]any, offsets[idx+1]-offsets[idx])
	newSliceIdx := 0
	for ptr := offsets[idx]; ptr < offsets[idx+1]; ptr++ {
		anyVal, err := GetValueFromArrowArray(arr, int(ptr), timeAsString)
		if err != nil {
			return nil, errors.Wrap(err, "error getting value for LargeList column")
		}
		newSlice[newSliceIdx] = anyVal
		newSliceIdx += 1
	}
	return newSlice, nil
}

func setValue(field *reflect.Value, a arrow.Array, idx int) error {
	if a.IsNull(idx) {
		return nil
	}
	switch arr := a.(type) {
	//case *array.LargeList:
	//	return getInnerSliceFromArray(arr.ListValues(), arr.Offsets(), idx, timeAsString)
	//case *array.List:
	//	o32 := arr.Offsets()
	//	o64 := make([]int64, len(o32))
	//	for i := 0; i < len(o32); i++ {
	//		o64[i] = int64(arr.Offsets()[i])
	//	}
	//	return getInnerSliceFromArray(arr.ListValues(), o64, idx, timeAsString)
	//case *array.Struct:
	//	newMap := map[string]any{}
	//	structType, typeOk := arr.DataType().(*arrow.StructType)
	//	if !typeOk {
	//		return nil, fmt.Errorf("error getting struct type")
	//	}
	//	for k := 0; k < arr.NumField(); k++ {
	//		anyVal, err := GetValueFromArrowArray(arr.Field(k), idx, timeAsString)
	//		if err != nil {
	//			return nil, errors.Wrap(err, "error getting value for Struct column")
	//		}
	//		newMap[structType.Field(k).Name] = anyVal
	//	}
	//	return newMap, nil
	//case *array.Dictionary:
	//	return GetValueFromArrowArray(arr.Dictionary(), arr.GetValueIndex(idx), timeAsString)
	case *array.String:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.SetString(val)
			return nil
		}
	case *array.LargeString:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.SetString(val)
			return nil
		}
	case *array.Uint8:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Uint16:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Uint32:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Uint64:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Int16:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Int32:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Int64:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Float64:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Boolean:
		val := arr.Value(idx)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&val))
			return nil
		} else {
			field.Set(reflect.ValueOf(val))
			return nil
		}
	case *array.Date32:
		timeVal := arr.Value(idx).ToTime()
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&timeVal))
			return nil
		} else {
			field.Set(reflect.ValueOf(timeVal))
			return nil
		}
	case *array.Date64:
		timeVal := arr.Value(idx).ToTime()
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&timeVal))
			return nil
		} else {
			field.Set(reflect.ValueOf(timeVal))
			return nil
		}
	case *array.Timestamp:
		timeUnit := arr.DataType().(*arrow.TimestampType).TimeUnit()
		timeVal := arr.Value(idx).ToTime(timeUnit)
		if field.Kind() == reflect.Ptr {
			field.Set(reflect.ValueOf(&timeVal))
			return nil
		} else {
			field.Set(reflect.ValueOf(timeVal))
			return nil
		}
	default:
		return errors.Newf("unsupported array type: %T", arr)
	}
}

func GetValueFromArrowArray(a arrow.Array, idx int, timeAsString bool) (any, error) {
	if a.IsNull(idx) {
		return nil, nil
	}
	switch arr := a.(type) {
	case *array.LargeList:
		return getInnerSliceFromArray(arr.ListValues(), arr.Offsets(), idx, timeAsString)
	case *array.List:
		o32 := arr.Offsets()
		o64 := make([]int64, len(o32))
		for i := 0; i < len(o32); i++ {
			o64[i] = int64(arr.Offsets()[i])
		}
		return getInnerSliceFromArray(arr.ListValues(), o64, idx, timeAsString)
	case *array.Struct:
		newMap := map[string]any{}
		structType, typeOk := arr.DataType().(*arrow.StructType)
		if !typeOk {
			return nil, fmt.Errorf("error getting struct type")
		}
		for k := 0; k < arr.NumField(); k++ {
			anyVal, err := GetValueFromArrowArray(arr.Field(k), idx, timeAsString)
			if err != nil {
				return nil, errors.Wrap(err, "error getting value for Struct column")
			}
			newMap[structType.Field(k).Name] = anyVal
		}
		return newMap, nil
	case *array.Dictionary:
		return GetValueFromArrowArray(arr.Dictionary(), arr.GetValueIndex(idx), timeAsString)
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
		timeVal := arr.Value(idx).ToTime()
		if timeAsString {
			return timeVal.Format(time.RFC3339), nil
		}
		return timeVal, nil
	case *array.Date64:
		timeVal := arr.Value(idx).ToTime()
		if timeAsString {
			return timeVal.Format(time.RFC3339), nil
		}
		return timeVal, nil
	case *array.Timestamp:
		timeUnit := arr.DataType().(*arrow.TimestampType).TimeUnit()
		timeVal := arr.Value(idx).ToTime(timeUnit)
		if timeAsString {
			return timeVal.Format(time.RFC3339), nil
		}
		return timeVal, nil

	default:
		return nil, fmt.Errorf("unsupported array type: %T", arr)
	}
}

type ChunkResult struct {
	chunkIdx int
	rows     []map[string]any
	meta     []map[string]FeatureMeta
	err      error
}

type FeatureMeta struct {
	SourceType  *string
	SourceId    *string
	ResolverFqn *string
	Pkey        any
}

func MapTableToStructs(
	table arrow.Table,
	structs *reflect.Value,
	allMemo *AllNamespaceMemoT,
) error {
	numRows, err := Int64ToInt(table.NumRows())
	if err != nil {
		return errors.Wrapf(err, "table too large, found %d rows", table.NumRows())
	}
	reader := array.NewTableReader(table, int64(tableReaderChunkSize))
	defer reader.Release()

	for reader.Next() {
		record := reader.Record()
		var featureColumnIdxs []int
		metaColumnFqnToIdx := make(map[string]int)
		for j := range record.Columns() {
			colName := record.ColumnName(j)
			if strings.HasPrefix(colName, metadataPrefix) || colName == pkeyField {
				metaColumnFqnToIdx[strings.TrimPrefix(colName, metadataPrefix)] = j
			} else if colName == "__ts__" || colName == "__index__" || strings.HasPrefix(colName, "__chalk__.") || strings.HasSuffix(colName, ".__chalk_observed_at__") {
				continue
			} else {
				featureColumnIdxs = append(featureColumnIdxs, j)
			}
		}
		mapRecordToStructs(record, structs, featureColumnIdxs, 0, int64(numRows), 0, allMemo)
	}
	return nil
}

func mapRecordToStructs(
	record arrow.Record,
	structs *reflect.Value,
	featureColumnIdxs []int,
	chunkStart int64,
	chunkEnd int64,
	chunkIdx int,
	allMemo *AllNamespaceMemoT,
) error {
	chunkEndInt, err := Int64ToInt(chunkEnd)
	if err != nil {
		return errors.Wrapf(err, "chunk too large, found %d rows", chunkEnd)
	}
	chunkStartInt := int(chunkStart)
	structType := structs.Type().Elem()
	memo, ok := allMemo.Load(structType)
	if !ok {
		return errors.Newf("memo not found for struct type %s, found keys: %v", structType, allMemo.Keys())
	}
	for rowIdx := chunkStartInt; rowIdx < chunkEndInt; rowIdx++ {
		reflectStruct := structs.Index(rowIdx)
		for _, colIdx := range featureColumnIdxs {
			columnArray := record.Column(colIdx)
			fqn := record.ColumnName(colIdx)
			fieldIdxs, ok := memo.ResolvedFieldNameToIndices[fqn]
			if !ok {
				return errors.Newf(
					"feature '%s' not found in memo for struct type %s, found: %v",
					fqn, structType, colls.Keys(memo.ResolvedFieldNameToIndices),
				)
			}
			for _, fieldIdx := range fieldIdxs {
				field := reflectStruct.Field(fieldIdx)
				if err := setValue(&field, columnArray, rowIdx); err != nil {
					return errors.Wrapf(err, "setting value for field '%s' row %d", fqn, rowIdx)
				}
			}
		}
	}
	return nil
}

func extractFeatures(
	record arrow.Record,
	featureColumnIdxs []int,
	metaColumnFqnToIdx map[string]int,
	chunkStart int64,
	chunkEnd int64,
	chunkIdx int,
	resChan chan<- *ChunkResult,
	wg *sync.WaitGroup,
	timeAsString bool,
) {
	defer wg.Done()

	var featureRes []map[string]any
	chunkEndInt, err := Int64ToInt(chunkEnd)
	if err != nil {
		resChan <- &ChunkResult{chunkIdx: chunkIdx, err: err}
		return
	}
	chunkStartInt := int(chunkStart)
	for i := chunkStartInt; i < chunkEndInt; i++ {
		m := make(map[string]any, len(featureColumnIdxs))
		for j := range featureColumnIdxs {
			name := record.ColumnName(j)
			value, err := GetValueFromArrowArray(record.Column(j), i, timeAsString)
			if err != nil {
				resChan <- &ChunkResult{
					chunkIdx: chunkIdx,
					err:      errors.Wrapf(err, "getting value from arrow array for feature '%s'", name),
				}
				return
			}
			m[name] = value
		}
		featureRes = append(featureRes, m)
	}

	if len(metaColumnFqnToIdx) == 0 {
		resChan <- &ChunkResult{chunkIdx: chunkIdx, rows: featureRes, meta: nil}
		return
	}

	var metaRes []map[string]FeatureMeta
	for i := chunkStartInt; i < chunkEndInt; i++ {
		m := make(map[string]FeatureMeta)

		var resolvedPkey any
		if idx, ok := metaColumnFqnToIdx[pkeyField]; ok {
			value, err := GetValueFromArrowArray(record.Column(idx), i, timeAsString)
			if err != nil {
				resChan <- &ChunkResult{
					chunkIdx: chunkIdx,
					err:      errors.Wrap(err, "getting primary key from arrow array"),
				}
				return
			}
			resolvedPkey = value
			delete(metaColumnFqnToIdx, pkeyField)
		}

		for fqn, j := range metaColumnFqnToIdx {
			featureMeta := FeatureMeta{
				Pkey: resolvedPkey,
			}

			value, err := GetValueFromArrowArray(record.Column(j), i, timeAsString)
			if err != nil {
				resChan <- &ChunkResult{
					chunkIdx: chunkIdx,
					err:      errors.Wrapf(err, "getting metadata from arrow array for feature '%s'", fqn),
				}
				return
			}
			metaCast, ok := value.(map[string]any)
			if !ok {
				resChan <- &ChunkResult{
					chunkIdx: chunkIdx,
					err:      fmt.Errorf("casting metadata into map for feature '%s'", fqn),
				}
				return
			}

			if sourceType, ok := metaCast["source_type"]; ok && sourceType != nil {
				val, ok := sourceType.(string)
				if !ok {
					resChan <- &ChunkResult{
						chunkIdx: chunkIdx,
						err:      fmt.Errorf("casting source_type into string for feature '%s'", fqn),
					}
					return
				}
				featureMeta.SourceType = &val
			}
			if sourceId, ok := metaCast["source_id"]; ok && sourceId != nil {
				val, ok := sourceId.(string)
				if !ok {
					resChan <- &ChunkResult{
						chunkIdx: chunkIdx,
						err:      fmt.Errorf("casting source_id into string for feature '%s'", fqn),
					}
					return
				}
				featureMeta.SourceId = &val
			}
			if resolverFqn, ok := metaCast["resolver_fqn"]; ok && resolverFqn != nil {
				val, ok := resolverFqn.(string)
				if !ok {
					resChan <- &ChunkResult{
						chunkIdx: chunkIdx,
						err:      fmt.Errorf("casting resolver_fqn into string for feature '%s'", fqn),
					}
					return
				}
				featureMeta.ResolverFqn = &val
			}

			m[fqn] = featureMeta
		}

		metaRes = append(metaRes, m)
	}

	resChan <- &ChunkResult{chunkIdx: chunkIdx, rows: featureRes, meta: metaRes}
}

func ExtractFeaturesFromTable(
	table arrow.Table,
	timeAsString bool, // CHA-5430
) ([]map[string]any, []map[string]FeatureMeta, error) {
	numRows, err := Int64ToInt(table.NumRows())
	if err != nil {
		return nil, nil, errors.Wrapf(err, "table too large, found %d rows", table.NumRows())
	}
	featureRes := make([]map[string]any, 0, numRows)
	metaRes := make([]map[string]FeatureMeta, 0, numRows)
	reader := array.NewTableReader(table, int64(tableReaderChunkSize))
	defer reader.Release()

	for reader.Next() {
		record := reader.Record()
		var featureColumnIdxs []int
		metaColumnFqnToIdx := make(map[string]int)
		for j := range record.Columns() {
			colName := record.ColumnName(j)
			if strings.HasPrefix(colName, metadataPrefix) || colName == pkeyField {
				metaColumnFqnToIdx[strings.TrimPrefix(colName, metadataPrefix)] = j
			} else if colName == "__ts__" || colName == "__index__" || strings.HasPrefix(colName, "__chalk__.") || strings.HasSuffix(colName, ".__chalk_observed_at__") {
				continue
			} else {
				featureColumnIdxs = append(featureColumnIdxs, j)
			}
		}

		var wg sync.WaitGroup
		numWorkers := runtime.NumCPU()
		resChan := make(chan *ChunkResult, numWorkers)
		chunkSize := (record.NumRows() / int64(numWorkers)) + 1

		chunkIdx := 0
		for chunkStart := int64(0); chunkStart < record.NumRows(); chunkStart += chunkSize {
			chunkEnd := min(chunkStart+chunkSize, record.NumRows())
			wg.Add(1)
			go extractFeatures(
				record,
				featureColumnIdxs,
				metaColumnFqnToIdx,
				chunkStart,
				chunkEnd,
				chunkIdx,
				resChan,
				&wg,
				timeAsString,
			)
			chunkIdx += 1
		}

		go func() {
			wg.Wait()
			close(resChan)
		}()

		var allChunks []*ChunkResult
		for chunkResult := range resChan {
			allChunks = append(allChunks, chunkResult)
		}

		sort.Slice(allChunks, func(i, j int) bool {
			return allChunks[i].chunkIdx < allChunks[j].chunkIdx
		})

		for _, chunkResult := range allChunks {
			if chunkResult.err != nil {
				return nil, nil, chunkResult.err
			}
			featureRes = append(featureRes, chunkResult.rows...)
			metaRes = append(metaRes, chunkResult.meta...)
		}
	}
	return featureRes, metaRes, nil
}

func ReflectPtr(value reflect.Value) reflect.Value {
	ptr := reflect.New(value.Type())
	ptr.Elem().Set(value)
	return ptr
}

// GetReflectValue returns a reflect.Value of the given type from the given non-reflect value.
func GetReflectValue(value any, typ reflect.Type, allMemo *AllNamespaceMemoT) (*reflect.Value, error) {
	if value == nil {
		return ptr.Ptr(reflect.Zero(typ)), nil
	}
	if reflect.ValueOf(value).Kind() == reflect.Ptr && typ.Kind() == reflect.Ptr {
		indirectValue, err := GetReflectValue(reflect.ValueOf(value).Elem().Interface(), typ.Elem(), allMemo)
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
				rVal, err := GetReflectValue(&memberValue, memberField.Type(), allMemo)
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
			memo, ok := allMemo.Load(structValue.Type())
			if !ok {
				return nil, fmt.Errorf(
					"namespace memo not found for struct '%s' - found %v",
					structValue.Type().Name(),
					allMemo.Keys(),
				)
			}
			if memo.ResolvedFieldNameToIndices == nil {
				return nil, fmt.Errorf(
					"resolved field name to index map not found for struct '%s'",
					structValue.Type().Name(),
				)
			}
			for k, v := range mapz {
				memberFieldIndices, fieldOk := memo.ResolvedFieldNameToIndices[k]
				if !fieldOk {
					// For forward compatibility, i.e. when clients add
					// more fields to their dataclasses in chalkpy, we want
					// to default to not erring when trying to deserialize
					// a new field that does not yet exist in the Go struct.
					// Eventually we might consider exposing a flag.
					continue
				}
				for _, memberFieldIdx := range memberFieldIndices {
					memberField := structValue.Field(memberFieldIdx)
					if v == nil {
						continue
					}

					if memberField.Type().Kind() == reflect.Map {
						bucket, err := GetBucketFromFqn(k)
						if err != nil {
							return nil, errors.Wrapf(err, "error extracting bucket value for feature '%s'", k)
						}
						if err := SetMapEntryValue(memberField, bucket, v, allMemo); err != nil {
							return nil, errors.Wrapf(
								err,
								"error setting map entry value for field '%s' in struct '%s'",
								k, structValue.Type().Name(),
							)
						}
					} else {
						rVal, err := GetReflectValue(&v, memberField.Type(), allMemo)
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
			rVal, err := GetReflectValue(actualValue, typ.Elem(), allMemo)
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
func SetMapEntryValue(mapValue reflect.Value, key string, value any, allMemo *AllNamespaceMemoT) error {
	if mapValue.IsNil() {
		mapType := mapValue.Type()
		newMap := reflect.MakeMap(mapType)
		mapValue.Set(newMap)
	}
	rVal, err := GetReflectValue(value, mapValue.Type().Elem().Elem(), allMemo)
	if err != nil {
		return errors.Wrap(err, "error getting reflect value for map entry")
	}
	mapValue.SetMapIndex(reflect.ValueOf(key), ReflectPtr(*rVal))
	return nil
}

/*  PopulateAllNamespaceMemo populates a memo to make bulk-unmarshalling and has-many unmarshalling efficient.
 *  i.e. Don't need to do the same work for the same features class multiple times. Given:
 *  type User struct {
 *      Id *string
 *      Transactions *[]Transactions `has_many:"id,user_id"`
 *      Grade   *int `versioned:"default(2)"`
 *      GradeV1 *int `versioned:"true"`
 *      GradeV2 *int `versioned:"true"`
 *  }
 *  type Transactions struct {
 *      Id *string
 *      UserId *string
 *      Amount *float64
 *  }
 *  The namespace memo will be:
 *  {
 *      "User": {
 *          ResolvedFieldNameToIndices: {
 *              "id": [0],
 *              "user.id": [0],
 *              "grade@2": [2, 4],
 *              "user.grade@2": [2, 4],
 *              "grade": [3],
 *              "user.grade": [3],
 *              "transactions": [1],
 *              "user.transactions": [1],
 *          }
 *      },
 *      "Transactions": {
 *          ResolvedFieldNameToIndices: {
 *              "id": [0],
 *              "transactions.id": [0],
 *              "user_id": [1],
 *              "transactions.user_id": [1],
 *              "amount": [2],
 *              "transactions.amount": [2],
 *          }
 *      }
 *  }
 */
func PopulateAllNamespaceMemo(typ reflect.Type) error {
	allMemo := AllNamespaceMemo
	if typ.Kind() == reflect.Ptr {
		return PopulateAllNamespaceMemo(typ.Elem())
	} else if typ.Kind() == reflect.Struct && typ != reflect.TypeOf(time.Time{}) {
		structName := typ.Name()
		namespace := ChalkpySnakeCase(structName)
		nsMutex, loaded := allMemo.LoadOrStoreLockedMutex(typ, NewNamespaceMemo())
		if loaded {
			// Prevent infinite loops and processing the same struct more than once.
			return nil
		}
		nsMemo := nsMutex.memo
		for fieldIdx := 0; fieldIdx < typ.NumField(); fieldIdx++ {
			fm := typ.Field(fieldIdx)
			resolvedName, err := ResolveFeatureName(fm)
			if err != nil {
				return errors.Wrapf(err, "error resolving feature name: %s", fm.Name)
			}
			nsMemo.ResolvedFieldNameToIndices[resolvedName] = append(nsMemo.ResolvedFieldNameToIndices[resolvedName], fieldIdx)
			// Has-many features come back as a list of structs whose keys are namespaced FQNs.
			// Here we map those keys to their respective indices in the struct, so that we
			// don't have to do any string manipulation to deprefix the FQN when unmarshalling.
			rootFqn := namespace + "." + resolvedName
			nsMemo.ResolvedFieldNameToIndices[rootFqn] = append(nsMemo.ResolvedFieldNameToIndices[rootFqn], fieldIdx)

			// Handle exploding windowed features
			if fm.Type.Kind() == reflect.Map {
				// Is a windowed feature
				intTags, err := GetWindowBucketsSecondsFromStructTag(fm)
				if err != nil {
					return errors.Wrapf(
						err,
						"error getting window buckets for field '%s' in struct '%s'",
						fm.Name,
						structName,
					)
				}
				for _, tag := range intTags {
					bucketFqn := resolvedName + "__" + strconv.Itoa(tag) + "__"
					nsMemo.ResolvedFieldNameToIndices[bucketFqn] = append(nsMemo.ResolvedFieldNameToIndices[bucketFqn], fieldIdx)
					rootBucketFqn := namespace + "." + bucketFqn
					nsMemo.ResolvedFieldNameToIndices[rootBucketFqn] = append(nsMemo.ResolvedFieldNameToIndices[rootBucketFqn], fieldIdx)
				}
			} else {
				if err := PopulateAllNamespaceMemo(fm.Type); err != nil {
					return err
				}
			}

			if fm.Type.Kind() == reflect.Ptr && IsStruct(fm.Type.Elem()) && !IsTypeDataclass(fm.Type.Elem()) {
				nsMemo.StructFieldsSet[resolvedName] = true
			}
		}
		nsMutex.mu.Unlock()
	} else if typ.Kind() == reflect.Slice {
		return PopulateAllNamespaceMemo(typ.Elem())
	}
	return nil
}
