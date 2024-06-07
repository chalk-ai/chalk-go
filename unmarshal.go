package chalk

import (
	"errors"
	"fmt"
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strconv"
	"strings"
)

type fqnToFields map[string][]reflect.Value

var FieldNotFoundError = errors.New("field not found")

func (f fqnToFields) addField(fqn string, field reflect.Value) {
	if _, ok := f[fqn]; !ok {
		f[fqn] = []reflect.Value{}
	}
	f[fqn] = append(f[fqn], field)
}

func setFeatureSingle(field reflect.Value, fqn string, value any) error {
	if internal.IsDataclassPointer(field) {
		structValue := field.Elem()
		if slice, isSlice := value.([]any); isSlice {
			if len(slice) != structValue.NumField() {
				return fmt.Errorf(
					"error unmarshalling value for dataclass "+
						"feature %s: expected %d fields, got %s",
					fqn,
					structValue.NumField(),
					slice,
				)
			}
			for idx, memberValue := range slice {
				memberFieldMeta := structValue.Type().Field(idx)
				memberField := structValue.Field(idx)
				pythonName := SnakeCase(memberFieldMeta.Name)
				if memberField == (reflect.Value{}) {
					return fmt.Errorf(
						"error unmarshalling value for dataclass feature %s: "+
							"field %s not found in struct %s",
						fqn, pythonName, structValue.Type().Name(),
					)
				}
				memberFqn := fqn + "." + pythonName
				if err := setFeatureSingle(memberField, memberFqn, memberValue); err != nil {
					return fmt.Errorf(
						"error unmarshalling value '%s' "+
							"for dataclass feature '%s': %w",
						pythonName, fqn, err,
					)
				}
			}
		} else if mapz, isMap := value.(map[string]any); isMap {
			nameToField := make(map[string]reflect.Value)
			for i := 0; i < structValue.NumField(); i++ {
				nameToField[SnakeCase(structValue.Type().Field(i).Name)] = structValue.Field(i)
			}
			for k, v := range mapz {
				memberField, fieldOk := nameToField[k]
				if !fieldOk {
					return fmt.Errorf(
						"error unmarshalling value for dataclass feature %s: "+
							"field %s not found in struct %s",
						fqn, k, structValue.Type().Name(),
					)
				}
				if err := setFeatureSingle(memberField, fqn+"."+k, v); err != nil {
					return fmt.Errorf(
						"error unmarshalling value '%s' for dataclass feature '%s': %w",
						k, fqn, err,
					)
				}
			}
		} else {
			return fmt.Errorf(
				"error unmarshalling value for dataclass "+
					"feature %s: value is not an `any` slice",
				fqn,
			)
		}
	} else if field.Kind() == reflect.Map {
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
		reflectValue, err := internal.GetReflectValue(value, field.Type().Elem().Elem())
		if err != nil {
			return fmt.Errorf("error unmarshalling value for windowed bucket feature %s: %w", fqn, err)
		}
		field.SetMapIndex(tagValue, reflectValue)
	} else {
		reflectValue, err := internal.GetReflectValue(value, field.Type().Elem())
		if err != nil {
			return fmt.Errorf("error unmarshalling value for feature %s: %w", fqn, err)
		}
		field.Set(reflectValue)
	}
	return nil
}

func (f fqnToFields) setFeature(fqn string, value any) error {
	fields, ok := f[fqn]
	if !ok {
		return FieldNotFoundError
	}

	// Versioned features can have multiple fields with the same FQN.
	// We need to set the value for each field.
	for _, field := range fields {
		if err := setFeatureSingle(field, fqn, value); err != nil {
			return err
		}
	}
	return nil
}

func (result *OnlineQueryResult) unmarshal(resultHolder any) (returnErr *ClientError) {
	fqnToValue := map[Fqn]any{}
	for _, featureResult := range result.Data {
		fqnToValue[featureResult.Field] = featureResult.Value
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

func UnmarshalInto(resultHolder any, fqnToValue map[Fqn]any, expectedOutputs []string) (returnErr *ClientError) {
	fieldMap := make(fqnToFields)
	structValue := reflect.ValueOf(resultHolder).Elem()

	// Has a side effect: fieldMap will be populated.
	initErr := initFeatures(structValue, "", make(map[string]bool), fieldMap, nil)
	if initErr != nil {
		return &ClientError{Message: fmt.Errorf("exception occurred while initializing result holder: %w", initErr).Error()}
	}

	for fqn, value := range fqnToValue {
		if value == nil {
			// Some fields are optional, so we leave the field as nil
			// TODO: Add validation for optional fields
			continue
		}
		err := fieldMap.setFeature(fqn, value)
		if err != nil {
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
				return &ClientError{Message: fmt.Errorf("error unmarshaling feature '%s' into the struct '%s': %w", fqn, structName, err).Error()}
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
