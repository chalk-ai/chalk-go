package chalk

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/chalk-ai/chalk-go/internal/ptr"
	"github.com/cockroachdb/errors"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

var FieldNotFoundError = errors.New("field not found")

func setFeatureSingle(field reflect.Value, fqn string, value any, allMemo *internal.AllNamespaceMemoT) error {
	if field.Type().Kind() == reflect.Ptr {
		if reflect.TypeOf(value) == field.Type().Elem() {
			// Shortcut. Faster than GetReflectValue.
			switch castValue := value.(type) {
			case string:
				field.Set(reflect.ValueOf(&castValue))
			case int:
				field.Set(reflect.ValueOf(&castValue))
			case int8:
				field.Set(reflect.ValueOf(&castValue))
			case int16:
				field.Set(reflect.ValueOf(&castValue))
			case int32:
				field.Set(reflect.ValueOf(&castValue))
			case int64:
				field.Set(reflect.ValueOf(&castValue))
			case float32:
				field.Set(reflect.ValueOf(&castValue))
			case float64:
				field.Set(reflect.ValueOf(&castValue))
			case bool:
				field.Set(reflect.ValueOf(&castValue))
			case time.Time:
				field.Set(reflect.ValueOf(&castValue))
			default:
				return fmt.Errorf("unsupported type for feature '%s': %T", fqn, value)
			}
			return nil
		}
		rVal, err := internal.GetReflectValue(&value, field.Type(), allMemo)
		if err != nil {
			return errors.Wrapf(err, "error getting reflect value for feature '%s'", fqn)
		}
		field.Set(*rVal)
		return nil
	} else if field.Kind() == reflect.Map {
		bucket, err := internal.GetBucketFromFqn(fqn)
		if err != nil {
			return errors.Wrapf(err, "error extracting bucket value for feature '%s'", fqn)
		}
		if err := internal.SetMapEntryValue(field, bucket, value, allMemo); err != nil {
			return errors.Wrapf(err, "error setting map entry value for feature '%s'", fqn)
		}
		return nil
	} else {
		return fmt.Errorf("expected a pointer type for feature '%s', found %s", fqn, field.Type().Kind())
	}
}
func convertIfHasManyMap(value any) (any, error) {
	// For has-many values, we get this back:
	//
	// {
	//   "columns": ["user.id", "user.email"],
	//   "values": [
	//     ["id1", "id2"],
	//     ["email1@geemail.com", "email2@geemail.com"]
	//   ]
	// }
	//
	// We want to convert this to:
	//
	// [
	//   {"user.id": "id1", "user.email": "email1@geemail.com"},
	//   {"user.id": "id2", "user.email": "email2@geemail.com"}
	// ]
	//
	hasMany, ok := value.(map[string]any)
	if !ok {
		return value, nil
	}

	columnsRaw, hasColumns := hasMany["columns"]
	valuesRaw, hasValues := hasMany["values"]
	if !hasColumns || !hasValues {
		return value, nil
	}

	columnsAny, ok := columnsRaw.([]any)
	if !ok {
		return nil, errors.New("failed to convert columns to []any")
	}

	columns := make([]string, len(columnsAny))
	for i, column := range columnsAny {
		columns[i], ok = column.(string)
		if !ok {
			return nil, errors.Newf("failed to convert column '%v' to string", column)
		}
	}

	valuesAny, ok := valuesRaw.([]any)
	if !ok {
		return nil, errors.New("failed to convert values to [][]any")
	}

	values := make([][]any, len(valuesAny))
	for i, row := range valuesAny {
		values[i], ok = row.([]any)
		if !ok {
			return nil, errors.Newf("failed to convert row '%v' to []any", row)
		}
	}

	if len(values) == 0 {
		return nil, errors.New("values of has-many results is empty")
	}
	numRows := len(values[0])

	newValues := make([]map[string]any, numRows)
	for rowIdx := 0; rowIdx < numRows; rowIdx++ {
		newRow := make(map[string]any)
		for colIdx, colName := range columns {
			newRow[colName] = values[colIdx][rowIdx]
		}
		newValues[rowIdx] = newRow
	}
	return newValues, nil
}

func (result *OnlineQueryResult) unmarshal(resultHolder any) (returnErr *ClientError) {
	fqnToValue := make(map[Fqn]any, len(result.Data))
	for _, featureResult := range result.Data {
		convertedValue, err := convertIfHasManyMap(featureResult.Value)
		if err != nil {
			return &ClientError{Message: errors.Wrapf(err, "error converting feature '%s' value", featureResult.Field).Error()}
		}
		fqnToValue[featureResult.Field] = convertedValue
	}
	return UnmarshalInto(resultHolder, fqnToValue, result.expectedOutputs)
}

type ChunkResult struct {
	chunkIdx int
	rows     []reflect.Value
	err      error
}

func unmarshalTableInto(table arrow.Table, resultHolders any) (returnErr error) {
	defer func() {
		if panicContents := recover(); panicContents != nil {
			detail := "details irretrievable"
			switch typedContents := panicContents.(type) {
			case *reflect.ValueError:
				detail = typedContents.Error()
			case string:
				detail = typedContents
			}
			returnErr = fmt.Errorf("exception occurred while unmarshalling result: %s", detail)
		}
	}()

	numRows, err := internal.Int64ToInt(table.NumRows())
	if err != nil {
		return &ClientError{Message: fmt.Sprintf("table too large to unmarshal, found %d rows", table.NumRows())}
	}

	slicePtr := reflect.ValueOf(resultHolders)
	if slicePtr.Kind() != reflect.Ptr {
		return fmt.Errorf(
			"result holder should be a pointer to a slice of structs, "+
				"got '%s' instead",
			slicePtr.Kind(),
		)
	}

	slice := reflect.Indirect(slicePtr)
	if slice.Kind() != reflect.Slice {
		return fmt.Errorf(
			"result holder should be a pointer to a slice of structs, "+
				"got '%s' instead",
			slice.Kind(),
		)
	}

	sliceElemType := slice.Type().Elem()
	if sliceElemType.Kind() != reflect.Struct {
		return fmt.Errorf(
			"result holder should be a pointer to a slice of structs, "+
				"got a pointer to a slice of '%s' instead",
			sliceElemType.Kind(),
		)
	}

	rows, _, scalarsErr := internal.ExtractFeaturesFromTable(table, false)
	if scalarsErr != nil {
		return scalarsErr
	}
	if len(rows) == 0 {
		return nil
	}

	scope, err := buildScope(colls.Keys(rows[0]))
	if err != nil {
		return errors.Wrap(err, "building deserialization scope")
	}

	allMemo := internal.AllNamespaceMemo
	if err := internal.PopulateAllNamespaceMemo(sliceElemType); err != nil {
		return errors.Wrap(err, "building namespace memo")
	}

	structName := sliceElemType.Name()
	namespace := internal.ChalkpySnakeCase(structName)
	nsScope := scope.children[namespace]

	var rowToStruct func(map[Fqn]any) (*reflect.Value, error)
	if nsScope != nil {
		// single namespace unmarshalling
		nsMemo, ok := allMemo.Load(sliceElemType)
		if !ok {
			return errors.Newf("namespace '%s' not found in memo, found keys: %v", structName, allMemo.Keys())
		}

		rowToStruct = func(row map[Fqn]any) (*reflect.Value, error) {
			featuresStruct := reflect.New(sliceElemType)
			if err := thinUnmarshalInto(
				featuresStruct.Elem(),
				row,
				namespace,
				nil,
				nsScope,
				nsMemo,
				allMemo,
			); err != nil {
				return nil, err
			}
			return ptr.Ptr(featuresStruct.Elem()), nil
		}
	} else {
		// Multi namespace unmarshalling
		type namespaceMetaT struct {
			fieldIdx  int
			namespace string
			scope     *scopeTrie
			memo      *internal.NamespaceMemo
		}

		namespaceMeta := []namespaceMetaT{}
		for i := 0; i < sliceElemType.NumField(); i++ {
			fieldMeta := sliceElemType.Field(i)
			if fieldMeta.Type.Kind() != reflect.Struct {
				return errors.Newf(
					"If attempting single namespace unmarshalling, please make sure you're unmarshalling into the correct struct. "+
						"Attempted single namespace unmarshalling into struct '%s', but results are from these namespaces: %v. "+
						"If attempting multi-namespace unmarshalling, please pass in a pointer to a struct whose fields are all "+
						"structs (not struct pointers) corresponding to the output namespaces. The problematic field is '%s' of type '%s'.",
					structName,
					colls.Keys(scope.children),
					fieldMeta.Name,
					fieldMeta.Type.Name(),
				)
			}

			fieldNamespace := internal.ChalkpySnakeCase(fieldMeta.Type.Name())

			fieldScope := scope.children[fieldNamespace]
			if fieldScope == nil {
				return errors.Newf(
					"Please make sure you're unmarshalling into the correct struct. Attempted single namespace "+
						"unmarshalling into struct '%s', and attempted multi-namespace unmarshalling into the field '%s' "+
						"of type '%s', but results are from these namespaces: %v",
					structName,
					fieldMeta.Name,
					fieldMeta.Type.Name(),
					colls.Keys(scope.children),
				)
			}

			fieldMemo, ok := allMemo.Load(fieldMeta.Type)
			if !ok {
				return errors.Newf("namespace '%s' not found in memo, found keys: %v", structName, allMemo.Keys())
			}

			namespaceMeta = append(namespaceMeta, namespaceMetaT{
				fieldIdx:  i,
				namespace: fieldNamespace,
				scope:     fieldScope,
				memo:      fieldMemo,
			})
		}

		rowToStruct = func(row map[Fqn]any) (*reflect.Value, error) {
			rootStruct := reflect.New(sliceElemType)
			for _, meta := range namespaceMeta {
				if err := thinUnmarshalInto(
					rootStruct.Elem().Field(meta.fieldIdx),
					row,
					meta.namespace,
					nil,
					meta.scope,
					meta.memo,
					allMemo,
				); err != nil {
					return nil, errors.Wrapf(
						err,
						"error unmarshalling into field '%s'",
						sliceElemType.Field(meta.fieldIdx).Name,
					)
				}
			}
			return ptr.Ptr(rootStruct.Elem()), nil
		}
	}

	var wg sync.WaitGroup
	numWorkers := runtime.NumCPU()
	resChan := make(chan ChunkResult, numWorkers)
	chunkSize := (len(rows) / numWorkers) + 1

	for chunkIdx := 0; (chunkIdx * chunkSize) < len(rows); chunkIdx += 1 {
		wg.Add(1)
		go func(routineChunkIdx int) {
			defer wg.Done()
			chunkStart := routineChunkIdx * chunkSize
			chunkEnd := chunkStart + chunkSize
			chunkRows := rows[chunkStart:min(chunkEnd, len(rows))]
			results := make([]reflect.Value, len(chunkRows))
			for rowIdx, row := range chunkRows {
				res, err := rowToStruct(row)
				if err != nil {
					resChan <- ChunkResult{chunkIdx: routineChunkIdx, err: err}
					return
				} else {
					results[rowIdx] = *res
				}
			}
			resChan <- ChunkResult{chunkIdx: routineChunkIdx, rows: results}
		}(chunkIdx)
	}

	wg.Wait()
	close(resChan)

	var allChunks []ChunkResult
	for chunkResult := range resChan {
		allChunks = append(allChunks, chunkResult)
	}

	sort.Slice(allChunks, func(i, j int) bool {
		return allChunks[i].chunkIdx < allChunks[j].chunkIdx
	})

	newSlice := reflect.MakeSlice(slice.Type(), numRows, numRows)
	rowIdx := 0
	for _, chunkResult := range allChunks {
		if chunkResult.err != nil {
			return chunkResult.err
		}
		for _, row := range chunkResult.rows {
			newSlice.Index(rowIdx).Set(row)
			rowIdx += 1
		}
	}
	slice.Set(newSlice)

	return nil
}

// UnmarshalTableInto unmarshals the given Arrow table into the given result holders.
// The result holders should be a pointer to a slice of structs.
//
// Usage:
//
//	func printNumRelatives(chalkClient chalk.Client) {
//		result, _ := chalkClient.OnlineQueryBulk(chalk.OnlineQueryParams{}.WithOutputs(
//			Features.User.Relatives,
//		).WithInput(Features.User.Id, []int{1, 2}), nil)
//
//		relatives := make([]Relative, 0)
//		result.UnmarshalInto(&relatives)
//
//		feature, _ := chalk.UnwrapFeature(Features.User.Relatives)
//		fmt.Println("Number of relatives for all users: ", len(result.GroupsTable[feature.Fqn]))
//
//	}
func UnmarshalTableInto(table arrow.Table, resultHolders any) *ClientError {
	if err := unmarshalTableInto(table, resultHolders); err != nil {
		return &ClientError{err.Error()}
	}
	return nil
}

func buildScope(fqns []string) (*scopeTrie, error) {
	root := &scopeTrie{}
	for _, fqn := range fqns {
		root.addStr(fqn)
	}
	return root, nil
}

/*
UnmarshalInto unmarshals a map with keys being FQNs and values being the value for a
singular feature (rather than a list of values for multiple pkeys) into a struct whose
fields correspond to the FQNs. An illustration:

	type FinancialMetric struct {
		Id *string
		BusinessId *string
		MetricDate *time.Time
	}

	func main() {
		fqnToValue := map[Fqn]any{
			"FinancialMetric.Id": "id1",
			"FinancialMetric.BusinessId": "business_id1",
			"FinancialMetric.MetricDate": time.Now(),
		}
		fm := FinancialMetric{}
		if err := UnmarshalInto(&fm, fqnToValue, nil); err != (*ClientError)(nil) {
			fmt.Println(err)
		} else {
			fmt.Println(fm)
		}
	}

To ensure fast unmarshals, see `WarmUpUnmarshaller`.
*/
func UnmarshalInto(resultHolder any, fqnToValue map[Fqn]any, expectedOutputs []string) (returnErr *ClientError) {
	allMemo := internal.AllNamespaceMemo
	if err := internal.PopulateAllNamespaceMemo(reflect.ValueOf(resultHolder).Elem().Type()); err != nil {
		return &ClientError{errors.Wrap(err, "error building namespace memo").Error()}
	}
	scope, err := buildScope(colls.Keys(fqnToValue))
	if err != nil {
		return &ClientError{
			errors.Wrap(err, "error building scope for initializing result holder struct").Error(),
		}
	}

	holderValue := reflect.ValueOf(resultHolder)
	structValue := holderValue.Elem()
	structName := structValue.Type().Name()
	namespace := internal.ChalkpySnakeCase(structName)
	nsScope := scope.children[namespace]

	if nsScope != nil {
		// Single namespace unmarshalling
		if nsScope == nil {
			return &ClientError{
				errors.Newf(
					"Attempted to unmarshal into the feature struct '%s', "+
						"but results are from these feature class(es) '%v'",
					structName,
					colls.Keys(scope.children),
				).Error(),
			}
		}

		nsMemo, ok := allMemo.Load(holderValue.Elem().Type())
		if !ok {
			return &ClientError{errors.Newf("namespace '%s' not found in memo", structName).Error()}
		}

		return thinUnmarshalInto(
			holderValue.Elem(),
			fqnToValue,
			namespace,
			expectedOutputs,
			nsScope,
			nsMemo,
			allMemo,
		)
	}

	// Multi-namespace unmarshalling
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldMeta := structValue.Type().Field(i)
		if field.Type().Kind() != reflect.Struct {
			return &ClientError{
				Message: fmt.Sprintf(
					"If attempting single namespace unmarshalling, please make sure you're unmarshalling into the correct struct. "+
						"Attempted single namespace unmarshalling into struct '%s', but results are from these namespaces: %v. "+
						"If attempting multi-namespace unmarshalling, please pass in a pointer to a struct whose fields are all "+
						"structs (not struct pointers) corresponding to the output namespaces. The problematic field is '%s' of type '%s'.",
					structName,
					colls.Keys(scope.children),
					fieldMeta.Name,
					field.Type().Name(),
				),
			}
		}
		fieldNamespace := internal.ChalkpySnakeCase(field.Type().Name())

		fieldNsScope := scope.children[fieldNamespace]
		if fieldNsScope == nil {
			return &ClientError{
				Message: fmt.Sprintf(
					"Please make sure you're unmarshalling into the correct struct. Attempted single namespace "+
						"unmarshalling into struct '%s', and attempted multi-namespace unmarshalling into the field '%s' "+
						"of type '%s', but results are from these namespaces: %v",
					structName,
					fieldMeta.Name,
					field.Type().Name(),
					colls.Keys(scope.children),
				),
			}
		}

		fieldNsMemo, ok := allMemo.Load(field.Type())
		if !ok {
			return &ClientError{
				fmt.Sprintf(
					"namespace for struct '%s' of field '%s' not found in memo", field.Type().Name(), fieldMeta.Name,
				),
			}
		}
		if err := thinUnmarshalInto(
			field,
			fqnToValue,
			fieldNamespace,
			expectedOutputs,
			fieldNsScope,
			fieldNsMemo,
			allMemo,
		); err != nil {
			return &ClientError{Message: errors.Wrapf(err, "unmarshalling field '%s': %w", fieldMeta.Name).Error()}
		}
	}

	return nil
}

// thinUnmarshalInto is called per row. Any operation that can be
// done outside of this function must be done outside of this function.
func thinUnmarshalInto(
	structValue reflect.Value,
	fqnToValue map[Fqn]any,
	namespace string,
	expectedOutputs []string,
	namespaceScope *scopeTrie,
	namespaceMemo *internal.NamespaceMemo,
	allMemo *internal.AllNamespaceMemoT,
) (returnErr *ClientError) {
	remoteFeatureMap := map[string][]reflect.Value{}
	if err := initRemoteFeatureMap(
		remoteFeatureMap,
		structValue,
		namespace,
		map[string]bool{},
		namespaceScope,
		allMemo,
		true,
	); err != nil {
		return &ClientError{errors.Wrap(err, "error initializing result holder struct").Error()}
	}

	for fqn, value := range fqnToValue {
		targetFields, ok := remoteFeatureMap[fqn]
		if !ok {
			// If not a has-one remote feature, e.g. user.account.balance
			fieldIndices, ok := namespaceMemo.ResolvedFieldNameToIndices[fqn]
			if !ok {
				// For forward compatibility, i.e. when clients add
				// more fields to their dataclasses in chalkpy, we want
				// to default to not erring when trying to deserialize
				// a new field that does not yet exist in the Go struct.
				// Eventually we might consider exposing a flag.
				continue
			}
			targetFields = make([]reflect.Value, len(fieldIndices))
			for i, fieldIdx := range fieldIndices {
				targetFields[i] = structValue.Field(fieldIdx)
			}
		}

		for _, field := range targetFields {
			if value == nil {
				if field.Type().Kind() == reflect.Map && field.IsNil() {
					field.Set(reflect.MakeMap(field.Type()))
				}
				continue
			}
			if err := setFeatureSingle(field, fqn, value, allMemo); err != nil {
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
					return &ClientError{Message: errors.Wrapf(err, "error unmarshaling feature '%s' into the struct '%s'", fqn, structName).Error()}
				}
			}
		}

	}
	return nil
}

func validateOnlineQueryResultHolder(resultHolder any) error {
	value := reflect.ValueOf(resultHolder)
	kind := value.Type().Kind()
	if kind != reflect.Pointer {
		return &ClientError{Message: fmt.Sprintf("argument should be a pointer, got '%s' instead", kind.String())}
	}

	kindPointedTo := value.Elem().Kind()
	if kindPointedTo != reflect.Struct {
		return &ClientError{Message: fmt.Sprintf("argument should be pointer to a struct, got a pointer to a '%s' instead", kindPointedTo.String())}
	}
	return nil
}

func UnmarshalOnlineQueryResponse(response *commonv1.OnlineQueryResponse, resultHolder any) error {
	if err := validateOnlineQueryResultHolder(resultHolder); err != nil {
		return err
	}
	fqnToValue := map[Fqn]any{}
	for _, featureResult := range response.GetData().GetResults() {
		convertedValue, err := convertIfHasManyMap(featureResult.Value.AsInterface())
		if err != nil {
			return errors.Wrapf(err, "error converting has-many value for feature '%s'", featureResult.Field)
		}
		fqnToValue[featureResult.Field] = convertedValue
	}
	res := UnmarshalInto(resultHolder, fqnToValue, nil)
	if res == (*ClientError)(nil) {
		// TODO: Return `error` from `UnmarshalInto` [CHA-4153]
		return nil
	}
	return res
}

func UnmarshalOnlineQueryBulkResponse(response *commonv1.OnlineQueryBulkResponse, resultHolders any) error {
	scalars, err := internal.ConvertBytesToTable(response.GetScalarsData())
	if err != nil {
		return errors.Wrap(err, "error deserializing scalars table")
	}
	return unmarshalTableInto(scalars, resultHolders)
}

func ConvertTableToRows(table arrow.Table) ([]map[string]any, error) {
	features, _, err := internal.ExtractFeaturesFromTable(table, false)
	return features, err
}
