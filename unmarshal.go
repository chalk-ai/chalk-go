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

func getWindowedPseudofeatureMeta(fqn string, fieldMap fqnToFields) (*int, *reflect.Value, error) {
	sections := strings.Split(fqn, ".")
	lastSection := sections[len(sections)-1]

	lastSectionSplit := strings.Split(lastSection, "__")
	if len(lastSectionSplit) < 2 {
		return nil, nil, nil
	}
	secondsStr := lastSectionSplit[1]
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return nil, nil, nil
	}

	featureClassFqn := DesuffixFqn(fqn)
	baseFeatureFqn := featureClassFqn + "." + lastSectionSplit[0]
	baseFeatureFields, ok := fieldMap[baseFeatureFqn]
	if !ok {
		return nil, nil, nil
	}

	if len(baseFeatureFields) != 1 {
		return nil, nil, fmt.Errorf(
			"found more than one base feature field for windowed feature '%s', "+
				"likely because the windowed feature is versioned but we currently "+
				"do not support that",
			fqn,
		)
	}

	return &seconds, &baseFeatureFields[0], nil
}

func (t fqnToFields) setFeature(fqn string, value any) error {
	// Multiple versioned features can share the same fqn.
	// e.g. "Features.user.grade" and "Features.user.grade_v2" share
	//      the same fqn "user.grade@2" if the default version is 2.

	// Get the fields from the map
	// 1. Handle the different types of fields
	// 1a. WIndowed features
	//     - if the first field is a windowed feature, every field is a windowed feature.
	//     - if the first field is a windowed feautre, call setWindowedFeatureSingle on it.
	// 1b. Windowed features
	//     - if the first field is a

	if fields, ok := t[fqn]; ok {

	}
}

func (t fqnToFields) setFeatureSingle(fqn string, value any) error {
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
	} else if bucketDuration, baseFeatureField, windowedErr := getWindowedPseudofeatureMeta(fqn, t); (bucketDuration != nil && baseFeatureField != nil) || windowedErr != nil {
		if windowedErr != nil {
			return windowedErr
		}

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
