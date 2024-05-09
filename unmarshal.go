package chalk

import (
	"errors"
	"fmt"
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
		dataclassValues, ok := value.([]any)
		if !ok {
			return fmt.Errorf("error unmarshalling value for dataclass feature %s: value is not an `any` slice", fqn)
		}
		if len(dataclassValues) != structValue.NumField() {
			return fmt.Errorf("error unmarshalling value for dataclass feature %s: expected %d fields, got %s", fqn, structValue.NumField(), dataclassValues)
		}
		for idx, memberValue := range dataclassValues {
			memberFieldMeta := structValue.Type().Field(idx)
			memberField := structValue.Field(idx)
			pythonName := snakeCase(memberFieldMeta.Name)
			if memberField == (reflect.Value{}) {
				return fmt.Errorf("error unmarshalling value for dataclass feature %s: field %s not found in struct %s", fqn, pythonName, structValue.Type().Name())
			}
			memberFqn := fqn + "." + pythonName
			if err := setFeatureSingle(memberField, memberFqn, memberValue); err != nil {
				return fmt.Errorf("error unmarshalling value '%s' for dataclass feature '%s': %w", pythonName, fqn, err)
			}
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
	fqnToValueMap := map[Fqn]any{}
	for _, featureResult := range result.Data {
		fqnToValueMap[featureResult.Field] = featureResult.Value
	}
	return UnmarshalInto(resultHolder, fqnToValueMap, result.expectedOutputs)
}

func UnmarshalInto(resultHolder any, fqnToValueMap map[Fqn]any, expectedOutputs []string) (returnErr *ClientError) {
	fieldMap := make(fqnToFields)
	structValue := reflect.ValueOf(resultHolder).Elem()

	// Has a side effect: fieldMap will be populated.
	initErr := initFeatures(structValue, "", make(map[string]bool), fieldMap)
	if initErr != nil {
		return &ClientError{Message: "exception occurred while initializing result holder"}
	}

	for fqn, value := range fqnToValueMap {
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
					//return &ClientError{Message: fmt.Sprintf("Unexpected error unmarshaling output feature '%s'. Feature is still nil after unmarshaling", expectedOutput)}
				}
			}
		}
	}
	return nil
}
