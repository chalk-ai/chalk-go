package chalk

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type fqnToField map[string]reflect.Value

func (t fqnToField) setFeature(fqn string, value any) error {
	field, fieldFound := t[fqn]
	if !fieldFound {
		return errors.New("field not found")
	}

	if field.Type().Elem().Kind() == reflect.Int {
		switch typed := value.(type) {
		case float64:
			value = int(typed)
		}
	}
	if field.Type().Elem().String() == "time.Time" {
		stringValue := reflect.ValueOf(value).String()
		timeValue, timeErr := time.Parse(time.RFC3339, stringValue)
		if timeErr == nil {
			field.Set(reflect.ValueOf(&timeValue))
			return nil
		}

		dateValue, dateErr := time.Parse("2006-01-02", stringValue)
		if dateErr != nil {
			// Return original datetime parsing error
			return timeErr
		}
		field.Set(reflect.ValueOf(&dateValue))
	} else {
		copied := reflect.New(reflect.TypeOf(value))
		copied.Elem().Set(reflect.ValueOf(value))
		castedPointer := reflect.NewAt(field.Type().Elem(), copied.UnsafePointer())
		field.Set(castedPointer)
	}

	return nil
}

func (result *OnlineQueryResult) unmarshal(t any) *ClientError {
	featureMap := make(fqnToField)
	structValue := reflect.ValueOf(t).Elem()

	// Has a side effect: featureMap will be populated.
	initFeatures(structValue, "", make(map[string]bool), featureMap)

	for _, featureResult := range result.Data {
		fqn := featureResult.Field
		err := featureMap.setFeature(fqn, featureResult.Value)
		if err != nil {
			outputNamespace := "unknown namespace"
			sections := strings.Split(fqn, ".")
			if len(sections) > 0 {
				outputNamespace = sections[0]
			}
			detailedErr := fmt.Sprintf("Error unmarshaling feature '%s' into the struct '%s'. ", fqn, structValue.Type().String())
			detailedErr += fmt.Sprintf("Please check if you are passing a pointer to a struct that represents the output namespace '%s', ", outputNamespace)
			return &ClientError{Message: detailedErr}
		}
	}
	for _, expectedOutput := range result.expectedOutputs {
		if field, ok := featureMap[expectedOutput]; ok {
			if field.IsNil() {
				return &ClientError{Message: fmt.Sprintf("Unexpected error unmarshaling output feature '%s'. Feature is still nil after unmarshaling", expectedOutput)}
			}
		}
	}
	return nil
}
