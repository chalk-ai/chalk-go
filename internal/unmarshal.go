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

var TableReaderChunkSize = defaultTableReaderChunkSize

type Numbers interface {
	int | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func init() {
	if chunkSizeStr := os.Getenv(tableReaderChunkSizeKey); chunkSizeStr != "" {
		if newChunkSize, err := strconv.Atoi(chunkSizeStr); err == nil {
			TableReaderChunkSize = newChunkSize
		}
	}
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

type InitScope struct {
	Children map[string]*InitScope
}

func (s *InitScope) addStr(fqn string) {
	s.add(strings.Split(fqn, "."))
}

func (s *InitScope) add(fqnParts []string) {
	if len(fqnParts) == 0 {
		return
	}
	firstPart := fqnParts[0]
	if s.Children == nil {
		s.Children = map[string]*InitScope{}
	}
	if _, found := s.Children[firstPart]; !found {
		s.Children[firstPart] = &InitScope{}
	}
	s.Children[firstPart].add(fqnParts[1:])
}

func InitRemoteFeatureMap(
	remoteFeatureMap map[string][]reflect.Value,
	structValue reflect.Value,
	cumulativeFqn string,
	visited map[string]bool,
	scope *InitScope,
	allMemo *NamespaceMemosT,
	scopeToJustStructs bool,
) error {
	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf(
			"feature initialization function argument must be a reflect.Value"+
				" of the kind reflect.Struct, found %s instead",
			structValue.Kind().String(),
		)
	}

	structName := structValue.Type().Name()
	if isVisited, ok := visited[structName]; ok && isVisited {
		// Found a cycle. Just return.
		return nil
	}
	visited[structName] = true
	defer func() {
		visited[structName] = false
	}()

	memo, err := allMemo.LoadOrStore(structValue.Type())
	if err != nil {
		return errors.Wrapf(err, "loading memo for struct '%s'", structName)
	}

	var fieldNames []string
	if scopeToJustStructs {
		fieldNames = colls.Keys(memo.StructFieldsSet)
	} else {
		fieldNames = colls.Keys(scope.Children)
	}

	for _, resolvedFieldName := range fieldNames {
		nextScope, inScope := scope.Children[resolvedFieldName]
		if !inScope {
			continue
		}
		updatedFqn := cumulativeFqn + "." + resolvedFieldName
		fieldIndices, ok := memo.ResolvedFieldNameToIndices[resolvedFieldName]
		if !ok {
			// We arrive here when chalk-go receives a response that contains a feature
			// newly added to one of their has-one feature classes. They have not updated
			// their codegen'd structs yet, so we simply skip unmarshalling this new
			// feature to ensure forward compatibility.
			continue
		}

		for _, fieldIdx := range fieldIndices {
			f := structValue.Field(fieldIdx)

			if _, isStruct := memo.StructFieldsSet[resolvedFieldName]; isStruct {
				if !f.CanSet() {
					continue
				}
				if f.IsNil() {
					featureSet := reflect.New(f.Type().Elem())
					f.Set(featureSet)
				}
				if err := InitRemoteFeatureMap(
					remoteFeatureMap,
					f.Elem(),
					updatedFqn,
					visited,
					nextScope,
					allMemo,
					false,
				); err != nil {
					return err
				}
			} else {
				remoteFeatureMap[updatedFqn] = append(remoteFeatureMap[updatedFqn], f)
			}
		}
	}
	return nil
}

func BuildScope(fqns []string) (*InitScope, error) {
	root := &InitScope{}
	for _, fqn := range fqns {
		root.addStr(fqn)
	}
	return root, nil
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
	reader := array.NewTableReader(table, int64(TableReaderChunkSize))
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
func GetReflectValue(value any, typ reflect.Type, allMemo *NamespaceMemosT) (*reflect.Value, error) {
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
			memo, err := allMemo.LoadOrStore(structValue.Type())
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"loading namespace memo for struct '%s'",
					structValue.Type().Name(),
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
func SetMapEntryValue(mapValue reflect.Value, key string, value any, allMemo *NamespaceMemosT) error {
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

func UnmarshalTableInto(table arrow.Table, resultHolders any) (returnErr error) {
	defer func() {
		if panicContents := recover(); panicContents != nil {
			detail := "details irretrievable"
			switch typedContents := panicContents.(type) {
			case *reflect.ValueError:
				detail = typedContents.Error()
			case string:
				detail = typedContents
			}
			returnErr = errors.Newf("exception occurred while unmarshalling result: %s", detail)
		}
	}()

	numRows, err := Int64ToInt(table.NumRows())
	if err != nil {
		return errors.Newf("table too large to unmarshal, found %d rows", table.NumRows())
	}

	slicePtr := reflect.ValueOf(resultHolders)
	if slicePtr.Kind() != reflect.Ptr {
		return fmt.Errorf(
			"result holder should be a pointer to a slice of structs, "+
				"got '%s' instead",
			slicePtr.Kind(),
		)
	}

	structs := reflect.Indirect(slicePtr)
	if structs.Kind() != reflect.Slice {
		return fmt.Errorf(
			"result holder should be a pointer to a slice of structs, "+
				"got '%s' instead",
			structs.Kind(),
		)
	}

	structType := structs.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return fmt.Errorf(
			"result holder should be a pointer to a slice of structs, "+
				"got a pointer to a slice of '%s' instead",
			structType.Kind(),
		)
	}

	codecMemo := CodecMemo
	allMemo := NamespaceMemos

	if err := PopulateAllNamespaceMemo(structType, nil); err != nil {
		return errors.Wrap(err, "building namespace memo")
	}

	if structs.Len() != numRows {
		structs.Set(reflect.MakeSlice(structs.Type(), numRows, numRows))
	}

	reader := array.NewTableReader(table, int64(TableReaderChunkSize))
	defer reader.Release()

	_, err = Int64ToInt(table.NumRows())
	if err != nil {
		return errors.Wrapf(err, "table too large")
	}

	fields := table.Schema().Fields()
	colFqnParts := make([][]string, len(fields))
	namespaceToColIndices := map[string][]int{}
	for i, field := range table.Schema().Fields() {
		colFqnParts[i] = strings.Split(field.Name, ".")
		if field.Name == "__ts__" || field.Name == "__index__" || field.Name == "__id__" || strings.HasPrefix(field.Name, "__chalk__.") || strings.HasSuffix(field.Name, ".__chalk_observed_at__") {
			continue
		} else {
			namespace := colFqnParts[i][0]
			if _, ok := namespaceToColIndices[namespace]; !ok {
				namespaceToColIndices[namespace] = make([]int, 0, len(fields))
			}
			namespaceToColIndices[namespace] = append(namespaceToColIndices[namespace], i)
		}
	}

	var rowOp func(structValue reflect.Value, record arrow.Record, rowIdx int) error

	if len(namespaceToColIndices) == 1 {
		// Single-namespace unmarshalling
		includedColIndices := namespaceToColIndices[colls.Keys(namespaceToColIndices)[0]]
		colToCodec := make([]Codec, len(includedColIndices))
		for k, colIdx := range includedColIndices {
			column := fields[colIdx]
			codec, err := codecMemo.LoadOrStore(column.Name, func() (*Codec, error) {
				return generateUnmarshalValueCodec(
					// Taking the first field's type because multiple field indices for the same column
					// means they are just versioned features, which are all the same type.
					structType,
					column.Type,
					allMemo,
					column.Name,
					colFqnParts[colIdx],
				)
			})
			if err != nil {
				return errors.Wrapf(err, "loading memo for column '%s'", column.Name)
			}

			colToCodec[k] = *codec
		}

		rowOp = func(structValue reflect.Value, record arrow.Record, rowIdx int) error {
			for k, colIdx := range includedColIndices {
				if err := colToCodec[k](structValue, record.Column(colIdx), rowIdx); err != nil {
					return errors.Wrapf(err, "running codec for column '%s'", fields[colIdx].Name)
				}
			}
			return nil
		}
	} else {
		// Multi-namespace unmarshalling
		rootMemo, err := allMemo.LoadOrStore(structType)
		if err != nil {
			return errors.Wrapf(err, "loading namespace memo for struct: %s", structType.Name())
		}

		colIdxToNamespaceMeta := make([]namespaceMeta, len(fields))
		for childNamespace, includedColIndices := range namespaceToColIndices {
			namespaceFieldIndices, ok := rootMemo.ResolvedFieldNameToIndices[childNamespace]
			if !ok || len(namespaceFieldIndices) == 0 {
				return errors.Newf(
					"Attempted multi-namespace unmarshalling - please make sure to pass in a list of structs. "+
						"The struct should contain an inner struct that corresponds to the namespace '%s'. Found only "+
						"inner structs for these namespaces: %v",
					childNamespace,
					colls.Keys(namespaceToColIndices),
				)
			} else if len(namespaceFieldIndices) > 1 {
				var foundFieldNames []string
				for _, fieldIdx := range namespaceFieldIndices {
					foundFieldNames = append(foundFieldNames, structType.Field(fieldIdx).Name)
				}
				return errors.Newf(
					"namespace '%s' corresponds to multiple fields in the struct, but only one field is allowed: %v",
					childNamespace,
					foundFieldNames,
				)
			}

			rootStructFieldIdx := namespaceFieldIndices[0]
			innerStructType := structType.Field(rootStructFieldIdx).Type
			for _, colIdx := range includedColIndices {
				column := fields[colIdx]
				codec, err := codecMemo.LoadOrStore(column.Name, func() (*Codec, error) {
					return generateUnmarshalValueCodec(
						// Taking the first field's type because multiple field indices for the same column
						// means they are just versioned features, which are all the same type.
						innerStructType,
						column.Type,
						allMemo,
						column.Name,
						colFqnParts[colIdx],
					)
				})
				if err != nil {
					return errors.Wrapf(err, "loading codec for column '%s'", column.Name)
				}
				colIdxToNamespaceMeta[colIdx] = namespaceMeta{
					codec:           codec,
					rootStructIndex: rootStructFieldIdx,
				}
			}
		}

		rowOp = func(structValue reflect.Value, record arrow.Record, rowIdx int) error {
			for _, includedColIndices := range namespaceToColIndices {
				for _, colIdx := range includedColIndices {
					memo := colIdxToNamespaceMeta[colIdx]
					if memo.codec == nil {
						return errors.Newf("codec not found for column '%s'", fields[colIdx].Name)
					}
					if err := (*(memo.codec))(structValue.Field(memo.rootStructIndex), record.Column(colIdx), rowIdx); err != nil {
						return errors.Wrapf(err, "running codec for column '%s'", fields[colIdx].Name)
					}
				}
			}
			return nil
		}
	}

	rowOffset := 0
	batchIdx := 0
	for reader.Next() {
		record := reader.Record()
		recordRows := int(record.NumRows())
		for rowIdx := 0; rowIdx < recordRows; rowIdx++ {
			if err := rowOp(structs.Index(rowOffset+rowIdx), record, rowIdx); err != nil {
				return errors.Wrapf(err, "unmarshalling record batch %d", batchIdx)
			}
		}
		rowOffset += recordRows
		batchIdx += 1
	}

	return nil
}
