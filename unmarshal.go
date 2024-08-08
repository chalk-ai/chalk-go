package chalk

import (
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"reflect"
	"strconv"
	"strings"
)

var FieldNotFoundError = errors.New("field not found")

func setFeatureSingle(field reflect.Value, fqn string, value any) error {
	if field.Type().Kind() == reflect.Ptr {
		rVal, err := internal.GetReflectValue(&value, field.Type())
		if err != nil {
			return errors.Wrapf(err, "error getting reflect value for feature '%s'", fqn)
		}
		field.Set(*rVal)
		return nil
	} else if field.Kind() == reflect.Map {
		// We are handling maps differently because they are typed as `map`
		// instead of a pointer to a `map` like all other types are.
		//
		// And handling it in setFeaturesSingleNew instead of in the recursive
		// GetReflectValue function checks out because we never encounter
		// maps in slices, other maps, or structs.
		sections := strings.Split(fqn, ".")
		lastSection := sections[len(sections)-1]
		lastSectionSplit := strings.Split(lastSection, "__")
		formatErr := fmt.Errorf(
			"error unmarshalling value for windowed bucket feature %s: "+
				"expected windowed bucket feature to have fqn of the format "+
				"`{fqn}__{bucket seconds}__` ",
			fqn,
		)
		if len(lastSectionSplit) < 2 {
			return formatErr
		}
		secondsStr := lastSectionSplit[1]
		seconds, err := strconv.Atoi(secondsStr)
		if err != nil {
			return formatErr
		}
		tagValue := reflect.ValueOf(internal.FormatBucketDuration(seconds))
		rVal, err := internal.GetReflectValue(value, field.Type().Elem().Elem())
		if err != nil {
			return errors.Wrapf(err, "error unmarshalling value for windowed bucket feature %s", fqn)
		}
		field.SetMapIndex(tagValue, internal.ReflectPtr(*rVal))
		return nil
	} else {
		return fmt.Errorf("expected a pointer type for feature '%s', found %s", fqn, field.Type().Kind())
	}
}

func getConvertedValue(value any) (any, error) {
	// Does processing such as converting has-many results from a map with "columns" and "values" keys
	// to a list of maps with the column names (namespace de-prefixed) as keys.
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
			cell := values[colIdx][rowIdx]
			colParts := strings.Split(colName, ".")
			fieldName := colParts[len(colParts)-1]
			newRow[fieldName] = cell
		}
		newValues[rowIdx] = newRow
	}
	return newValues, nil
}

func (result *OnlineQueryResult) unmarshal(resultHolder any) (returnErr *ClientError) {
	fqnToValue := map[Fqn]any{}
	for _, featureResult := range result.Data {
		convertedValue, err := getConvertedValue(featureResult.Value)
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

	for _, row := range rows {
		res := reflect.New(sliceElemType)
		if err := UnmarshalInto(res.Interface(), row, nil); err != nil {
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

func UnmarshalInto(resultHolder any, fqnToValue map[Fqn]any, expectedOutputs []string) (returnErr *ClientError) {
	structValue := reflect.ValueOf(resultHolder).Elem()

	fieldMap := map[string][]reflect.Value{}

	initializer := NewFeatureInitializer()
	scope, err := buildScope(lo.Keys(fqnToValue))
	if err != nil {
		return &ClientError{
			errors.Wrap(err, "error building scope for initializing result holder struct").Error(),
		}
	}

	namespace := SnakeCase(structValue.Type().Name())
	nsScope := scope.children[namespace]
	if nsScope == nil {
		return &ClientError{
			errors.Newf("Scope of fields to initialize not found for namespace '%s'", nsScope).Error(),
		}
	}

	if err := initializer.initFeatures(structValue, namespace, map[string]bool{}, nsScope); err != nil {
		return &ClientError{errors.Wrap(err, "error initializing result holder struct").Error()}
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
		targetFields, ok := initializer.fieldsMap[fqn]
		if !ok {
			return &ClientError{
				errors.Newf("error locating fields associated with feature '%s'", fqn).Error(),
			}
		}
		if err != nil {
			err = errors.Wrapf(
				err,
				"error initializing feature field '%s' in the struct '%s'",
				fqn,
				structValue.Type().String(),
			)
			return &ClientError{Message: err.Error()}
		}
		for _, field := range targetFields {
			if _, ok := fieldMap[fqn]; !ok {
				fieldMap[fqn] = []reflect.Value{}
			}
			fieldMap[fqn] = append(fieldMap[fqn], field)
			if err := setFeatureSingle(field, fqn, value); err != nil {
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
