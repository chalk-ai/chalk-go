package chalk

import (
	"fmt"
	"reflect"
)

type FqnToField map[string]reflect.Value

func (t FqnToField) setFeature(fqn string, value any) error {

	if field, ok := t[fqn]; ok {
		if field.Type().Elem().Kind() == reflect.Int {
			switch typed := value.(type) {
			case float64:
				value = int(typed)
			}
		}
		p := reflect.New(reflect.TypeOf(value))
		p.Elem().Set(reflect.ValueOf(value))
		castedPointer := reflect.NewAt(field.Type().Elem(), p.UnsafePointer())
		field.Set(castedPointer)
		return nil
	}
	return fmt.Errorf("field not found for result with FQN '%s'", fqn)
}

func LoadFeaturesFromData[T any](t *T, data []FeatureResult) *ClientError {
	structValue := reflect.ValueOf(t).Elem()
	fqnToField := make(FqnToField)
	initFeatures(structValue, "", make(map[string]bool), fqnToField)

	for _, result := range data {
		err := fqnToField.setFeature(result.Field, result.Value)
		if err != nil {
			return &ClientError{Message: err.Error()}
		}
	}
	return nil
}
