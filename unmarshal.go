package chalk

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	commonv1 "github.com/chalk-ai/chalk-go/gen/chalk/common/v1"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/cockroachdb/errors"
	"reflect"
	"strings"
)

var FieldNotFoundError = errors.New("field not found")

func setFeatureSingle(field reflect.Value, fqn string, value any, nsMemo internal.NamespaceMemo) error {
	if field.Type().Kind() == reflect.Ptr {
		rVal, err := internal.GetReflectValue(&value, field.Type(), nsMemo)
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
		if err := internal.SetMapEntryValue(field, bucket, value, nsMemo); err != nil {
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
	fqnToValue := map[Fqn]any{}
	for _, featureResult := range result.Data {
		convertedValue, err := convertIfHasManyMap(featureResult.Value)
		if err != nil {
			return &ClientError{Message: errors.Wrapf(err, "error converting feature '%s' value", featureResult.Field).Error()}
		}
		fqnToValue[featureResult.Field] = convertedValue
	}
	return UnmarshalInto(resultHolder, fqnToValue, result.expectedOutputs)
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

	rows, scalarsErr := internal.ExtractFeaturesFromTable(table)
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

	memo := internal.NamespaceMemo{}
	if err := buildNamespaceMemo(memo, sliceElemType); err != nil {
		return errors.Wrap(err, "building namespace memo")
	}

	for _, row := range rows {
		res := reflect.New(sliceElemType)
		if err := innerUnmarshalInto(res.Interface(), row, nil, scope, memo); err != nil {
			return err
		}
		internal.SliceAppend(resultHolders, res.Elem())
	}

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
*/
func UnmarshalInto(resultHolder any, fqnToValue map[Fqn]any, expectedOutputs []string) (returnErr *ClientError) {
	memo := internal.NamespaceMemo{}
	if err := buildNamespaceMemo(memo, reflect.ValueOf(resultHolder).Elem().Type()); err != nil {
		return &ClientError{errors.Wrap(err, "error building namespace memo").Error()}
	}
	scope, err := buildScope(colls.Keys(fqnToValue))
	if err != nil {
		return &ClientError{
			errors.Wrap(err, "error building scope for initializing result holder struct").Error(),
		}
	}
	return innerUnmarshalInto(resultHolder, fqnToValue, expectedOutputs, scope, memo)
}
func innerUnmarshalInto(resultHolder any, fqnToValue map[Fqn]any, expectedOutputs []string, scope *scopeTrie, memo internal.NamespaceMemo) (returnErr *ClientError) {
	structValue := reflect.ValueOf(resultHolder).Elem()
	fieldMap := map[string][]reflect.Value{}

	structName := structValue.Type().Name()
	namespace := SnakeCase(structName)
	nsScope := scope.children[namespace]
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

	remoteFeatureMap := map[string][]reflect.Value{}
	if err := initRemoteFeatureMap(
		remoteFeatureMap,
		structValue,
		namespace,
		map[string]bool{},
		nsScope,
		memo,
		true,
	); err != nil {
		return &ClientError{errors.Wrap(err, "error initializing result holder struct").Error()}
	}

	nsMemo, ok := memo[structName]
	if !ok {
		return &ClientError{errors.Newf("namespace '%s' not found in memo", structName).Error()}
	}

	for fqn, value := range fqnToValue {
		if value == nil {
			// Some fields are optional, so we leave the field as nil
			// TODO: Add validation for optional fields
			continue
		}
		if _, shouldSkip := internal.SkipUnmarshalFqnRoots[getFqnRoot(fqn)]; shouldSkip {
			continue
		}

		targetFields, ok := remoteFeatureMap[fqn]
		if !ok {
			// If not a has-one remote feature, e.g. user.account.balance
			fieldIndices, ok := nsMemo.ResolvedFieldNameToIndices[fqn]
			if !ok {
				// For forward compatibility, i.e. when clients add
				// more fields to their dataclasses in chalkpy, we want
				// to default to not erring when trying to deserialize
				// a new field that does not yet exist in the Go struct.
				// Eventually we might consider exposing a flag.
				continue
			}
			for _, fieldIdx := range fieldIndices {
				targetFields = append(targetFields, structValue.Field(fieldIdx))
			}
		}

		for _, field := range targetFields {
			if _, ok := fieldMap[fqn]; !ok {
				fieldMap[fqn] = []reflect.Value{}
			}
			fieldMap[fqn] = append(fieldMap[fqn], field)
			if err := setFeatureSingle(field, fqn, value, memo); err != nil {
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
	for _, expectedOutput := range expectedOutputs {
		if fields, ok := fieldMap[expectedOutput]; ok {
			for _, field := range fields {
				if field.IsNil() {
					// TODO: Handle optional fields
					//return &ClientError{Message: fmt.Sprintf(
					//	"Unexpected error unmarshaling output feature '%s'. "+
					//		"Feature is still nil after unmarshaling",
					//	expectedOutput,
					//)}
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
	return internal.ExtractFeaturesFromTable(table)
}
