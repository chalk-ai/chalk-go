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

func setFeatureSingle(field reflect.Value, fqn string, value any, allMemo *internal.NamespaceMemosT) error {
	if field.Type().Kind() == reflect.Ptr {
		rVal, err := internal.GetReflectValue(&value, field.Type(), allMemo)
		if err != nil {
			return errors.Wrapf(err, "getting reflect value for feature '%s'", fqn)
		}
		field.Set(*rVal)
		return nil
	} else if field.Kind() == reflect.Map {
		bucket, err := internal.GetBucketFromFqn(fqn)
		if err != nil {
			return errors.Wrapf(err, "extracting bucket value for feature '%s'", fqn)
		}
		if err := internal.SetMapEntryValue(field, bucket, value, allMemo); err != nil {
			return errors.Wrapf(err, "setting map entry value for feature '%s'", fqn)
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

func (result *OnlineQueryResult) unmarshal(resultHolder any) (returnErr error) {
	fqnToValue := make(map[Fqn]any, len(result.Data))
	for _, featureResult := range result.Data {
		convertedValue, err := convertIfHasManyMap(featureResult.Value)
		if err != nil {
			return errors.Wrapf(err, "converting feature '%s' value", featureResult.Field)
		}
		fqnToValue[featureResult.Field] = convertedValue
	}
	return UnmarshalInto(resultHolder, fqnToValue)
}

// UnmarshalTableInto unmarshals the given Arrow table into the given result holders.
// The result holders should be a pointer to a slice of structs.
//
// Usage:
//
//	func printNumRelatives(chalkClient chalk.Client) {
//		result, _ := chalkClient.OnlineQueryBulk(
//		    context.Background(),
//		    chalk.OnlineQueryParams{}.WithOutputs(
//			    Features.User.Relatives,
//		    ).WithInput(Features.User.Id, []int{1, 2}),
//		    nil
//		)
//
//		relatives := make([]Relative, 0)
//		result.UnmarshalInto(&relatives)
//
//		feature, _ := chalk.UnwrapFeature(Features.User.Relatives)
//		fmt.Println("Number of relatives for all users: ", len(result.GroupsTable[feature.Fqn]))
//
//	}
func UnmarshalTableInto(table arrow.Table, resultHolders any) error {
	return internal.UnmarshalTableInto(table, resultHolders)
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
		if err := UnmarshalInto(&fm, fqnToValue); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fm)
		}
	}

To ensure fast unmarshals, see `WarmUpUnmarshaller`.
*/
func UnmarshalInto(resultHolder any, fqnToValue map[Fqn]any) (returnErr error) {
	allMemo := internal.AllNamespaceMemo
	if err := internal.PopulateAllNamespaceMemo(reflect.ValueOf(resultHolder).Elem().Type(), nil); err != nil {
		return errors.Wrap(err, "building namespace memo")
	}
	scope, err := internal.BuildScope(colls.Keys(fqnToValue))
	if err != nil {
		return errors.Wrap(err, "building scope for initializing result holder struct")
	}

	holderValue := reflect.ValueOf(resultHolder)
	structValue := holderValue.Elem()
	structName := structValue.Type().Name()
	namespace := internal.ChalkpySnakeCase(structName)
	nsScope := scope.Children[namespace]

	if nsScope != nil {
		// Single namespace unmarshalling
		if nsScope == nil {
			return errors.Newf(
				"Attempted to unmarshal into the feature struct '%s', "+
					"but results are from these feature class(es) '%v'",
				structName,
				colls.Keys(scope.Children),
			)
		}

		nsMemo, err := allMemo.Load(holderValue.Elem().Type())
		if err != nil {
			return errors.Wrapf(err, "loading memo for namespace '%s'", structName)
		}

		return thinUnmarshalInto(
			holderValue.Elem(),
			fqnToValue,
			namespace,
			nsScope,
			&nsMemo,
			allMemo,
		)
	}

	// Multi-namespace unmarshalling
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldMeta := structValue.Type().Field(i)
		if field.Type().Kind() != reflect.Struct {
			return errors.Newf(
				"If attempting single namespace unmarshalling, please make sure you're unmarshalling into the correct struct. "+
					"Attempted single namespace unmarshalling into struct '%s', but results are from these namespaces: %v. "+
					"If attempting multi-namespace unmarshalling, please pass in a pointer to a struct whose fields are all "+
					"structs (not struct pointers) corresponding to the output namespaces. The problematic field is '%s' of type '%s'.",
				structName,
				colls.Keys(scope.Children),
				fieldMeta.Name,
				field.Type().Name(),
			)
		}
		fieldNamespace := internal.ChalkpySnakeCase(field.Type().Name())

		fieldNsScope := scope.Children[fieldNamespace]
		if fieldNsScope == nil {
			return errors.Newf(
				"Please make sure you're unmarshalling into the correct struct. Attempted single namespace "+
					"unmarshalling into struct '%s', and attempted multi-namespace unmarshalling into the field '%s' "+
					"of type '%s', but results are from these namespaces: %v",
				structName,
				fieldMeta.Name,
				field.Type().Name(),
				colls.Keys(scope.Children),
			)
		}

		fieldNsMemo, err := allMemo.Load(field.Type())
		if err != nil {
			return errors.Wrapf(
				err,
				"loading namespace memo for struct '%s' of field '%s'", field.Type().Name(), fieldMeta.Name,
			)
		}
		if err := thinUnmarshalInto(
			field,
			fqnToValue,
			fieldNamespace,
			fieldNsScope,
			&fieldNsMemo,
			allMemo,
		); err != nil {
			return errors.Wrapf(err, "unmarshalling field '%s': %w", fieldMeta.Name)
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
	namespaceScope *internal.InitScope,
	namespaceMemo *internal.NamespaceMemo,
	allMemo *internal.NamespaceMemosT,
) (returnErr error) {
	remoteFeatureMap := map[string][]reflect.Value{}
	if err := internal.InitRemoteFeatureMap(
		remoteFeatureMap,
		structValue,
		namespace,
		map[string]bool{},
		namespaceScope,
		allMemo,
		true,
	); err != nil {
		return errors.Wrap(err, "initializing result holder struct")
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
			for _, fieldIdx := range fieldIndices {
				targetFields = append(targetFields, structValue.Field(fieldIdx))
			}
		}

		for _, field := range targetFields {
			if value == nil {
				if field.Type().Kind() == reflect.Map && field.IsNil() {
					field.Set(reflect.MakeMap(field.Type()))
					continue
				}

				// TODO: Add validation for optional fields
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
					return errors.New(fieldError)
				} else {
					return errors.Wrapf(err, "unmarshaling feature '%s' into the struct '%s'", fqn, structName)
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
		return errors.Newf("argument should be a pointer, got '%s' instead", kind.String())
	}

	kindPointedTo := value.Elem().Kind()
	if kindPointedTo != reflect.Struct {
		return errors.Newf("argument should be pointer to a struct, got a pointer to a '%s' instead", kindPointedTo.String())
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
			return errors.Wrapf(err, "converting has-many value for feature '%s'", featureResult.Field)
		}
		fqnToValue[featureResult.Field] = convertedValue
	}
	return UnmarshalInto(resultHolder, fqnToValue)
}

func UnmarshalOnlineQueryBulkResponse(response *commonv1.OnlineQueryBulkResponse, resultHolders any) error {
	scalars, err := internal.ConvertBytesToTable(response.GetScalarsData())
	if err != nil {
		return errors.Wrap(err, "deserializing scalars table")
	}
	return internal.UnmarshalTableInto(scalars, resultHolders)
}

func ConvertTableToRows(table arrow.Table) ([]map[string]any, error) {
	features, _, err := internal.ExtractFeaturesFromTable(table, false)
	return features, err
}
