package chalk

import (
	"errors"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strconv"
	"strings"
)

type fqnToField map[string]reflect.Value

var FieldNotFoundError = errors.New("field not found")

func getWindowedPseudofeatureMeta(fqn string, fieldMap fqnToField) (*int, *reflect.Value) {
	sections := strings.Split(fqn, ".")
	lastSection := sections[len(sections)-1]

	lastSectionSplit := strings.Split(lastSection, "__")
	if len(lastSectionSplit) < 2 {
		return nil, nil
	}
	secondsStr := lastSectionSplit[1]
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return nil, nil
	}

	featureClassFqn := DesuffixFqn(fqn)
	baseFeatureFqn := featureClassFqn + "." + lastSectionSplit[0]
	baseFeatureField, ok := fieldMap[baseFeatureFqn]
	if !ok {
		return nil, nil
	}

	return &seconds, &baseFeatureField
}

func (t fqnToField) setFeature(fqn string, value any) error {
	if field, fieldFound := t[fqn]; fieldFound && internal.IsDataclassPointer(field) {
		structValue := field.Elem()
		dataclassValues, ok := value.([]any)
		if !ok {
			return fmt.Errorf("error unmarshalling value for dataclass feature %s: value is not a slice", fqn)
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
			if err := t.setFeature(memberFqn, memberValue); err != nil {
				return fmt.Errorf("error unmarshalling value '%s' for dataclass feature '%s': %w", pythonName, fqn, err)
			}
		}
	} else if bucketDuration, baseFeatureField := getWindowedPseudofeatureMeta(fqn, t); bucketDuration != nil && baseFeatureField != nil {
		tagValue := reflect.ValueOf(internal.FormatBucketDuration(*bucketDuration))

		if baseFeatureField.Kind() != reflect.Map {
			return fmt.Errorf(fmt.Sprintf("exception setting windowed feature '%s'", fqn))
		}

		reflectValue, err := internal.GetReflectValue(value, baseFeatureField.Type().Elem().Elem())
		if err != nil {
			return fmt.Errorf("error unmarshalling value for windowed feature %s: %w", fqn, err)
		}

		baseFeatureField.SetMapIndex(tagValue, reflectValue)
	} else {
		field, fieldFound = t[fqn]
		if !fieldFound {
			return FieldNotFoundError
		}
		reflectValue, err := internal.GetReflectValue(value, field.Type().Elem())
		if err != nil {
			return fmt.Errorf("error unmarshalling value for feature %s: %w", fqn, err)
		}
		field.Set(reflectValue)
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
	fieldMap := make(fqnToField)
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
				return &ClientError{Message: fmt.Sprintf("Unknown error unmarshaling feature '%s' into the struct '%s'.", fqn, structName)}
			}

		}
	}
	for _, expectedOutput := range expectedOutputs {
		if field, ok := fieldMap[expectedOutput]; ok {
			if field.IsNil() {
				// TODO: Handle optional fields
				//return &ClientError{Message: fmt.Sprintf("Unexpected error unmarshaling output feature '%s'. Feature is still nil after unmarshaling", expectedOutput)}
			}
		}
	}
	return nil
}
