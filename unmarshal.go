package chalk

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"maps"
	"reflect"
	"slices"
)

func (result *OnlineQueryResult) unmarshal(resultHolder any) (returnErr error) {
	fqnToValue := make(map[Fqn]any, len(result.Data))
	for _, featureResult := range result.Data {
		convertedValue, err := internal.ConvertIfHasManyMap(featureResult.Value)
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
	allMemo := internal.NamespaceMemos
	if err := internal.PopulateNamespaceMemos(reflect.ValueOf(resultHolder).Elem().Type(), nil); err != nil {
		return errors.Wrap(err, "building namespace memo")
	}
	scope, err := internal.BuildScope(slices.Collect(maps.Keys(fqnToValue)))
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
		nsMemo, err := allMemo.LoadOrStore(holderValue.Elem().Type())
		if err != nil {
			return errors.Wrapf(err, "loading memo for struct '%s'", structName)
		}
		return internal.ThinUnmarshalInto(
			holderValue.Elem(),
			fqnToValue,
			namespace,
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
			return errors.Newf(
				"If attempting single namespace unmarshalling, please make sure you're unmarshalling into the correct struct. "+
					"Attempted single namespace unmarshalling into struct '%s', but results are from these namespaces: %v. "+
					"If attempting multi-namespace unmarshalling, please pass in a pointer to a struct whose fields are all "+
					"structs (not struct pointers) corresponding to the output namespaces. The problematic field is '%s' of type '%s'.",
				structName,
				slices.Collect(maps.Keys(scope.Children)),
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
				slices.Collect(maps.Keys(scope.Children)),
			)
		}

		fieldNsMemo, err := allMemo.LoadOrStore(field.Type())
		if err != nil {
			return errors.Wrapf(
				err,
				"loading memo for struct '%s' of field '%s'",
				field.Type().Name(), fieldMeta.Name,
			)
		}
		if err := internal.ThinUnmarshalInto(
			field,
			fqnToValue,
			fieldNamespace,
			fieldNsScope,
			fieldNsMemo,
			allMemo,
		); err != nil {
			return errors.Wrapf(err, "unmarshalling field %q", fieldMeta.Name)
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

func ConvertTableToRows(table arrow.Table) ([]map[string]any, error) {
	features, _, err := internal.ExtractFeaturesFromTable(table, false)
	return features, err
}
