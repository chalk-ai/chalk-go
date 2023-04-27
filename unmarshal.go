package chalk

import (
	"errors"
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strings"
)

func (result *OnlineQueryResult) unmarshal(t any) (returnErr *ClientError) {
	fieldMap := make(internal.FqnToField)
	structValue := reflect.ValueOf(t).Elem()

	// Has a side effect: fieldMap will be populated.
	initErr := internal.InitFeatureInternal(structValue, "", make(map[string]bool), fieldMap)
	if initErr != nil {
		return &ClientError{Message: "exception occurred while initializing result holder"}
	}

	for _, featureResult := range result.Data {
		fqn := featureResult.Field
		err := fieldMap.SetFeature(fqn, featureResult.Value)
		if err != nil {
			structName := structValue.Type().String()
			outputNamespace := "unknown namespace"
			sections := strings.Split(fqn, ".")
			if len(sections) > 0 {
				outputNamespace = sections[0]
			}
			if errors.Is(err, internal.FieldNotFoundError) {
				fieldError := fmt.Sprintf("Error unmarshaling feature '%s' into the struct '%s'. ", fqn, structName)
				fieldError += fmt.Sprintf("First, check if you are passing a pointer to a struct that represents the output namespace '%s'. ", outputNamespace)
				fieldError += fmt.Sprintf("Also, make sure the feature name can be traced to a field in the struct '%s' and or its nested structs.", structName)
				return &ClientError{Message: fieldError}
			} else {
				return &ClientError{Message: fmt.Sprintf("Unknown error unmarshaling feature '%s' into the struct '%s'.", fqn, structName)}
			}

		}
	}
	for _, expectedOutput := range result.expectedOutputs {
		if field, ok := fieldMap[expectedOutput]; ok {
			if field.IsNil() {
				// TODO: Handle optional fields
				return &ClientError{Message: fmt.Sprintf("Unexpected error unmarshaling output feature '%s'. Feature is still nil after unmarshaling", expectedOutput)}
			}
		}
	}
	return nil
}
