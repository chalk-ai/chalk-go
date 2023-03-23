package chalk

import (
	"fmt"
	"reflect"
	"strings"
)

type Feature struct {
	Fqn string
}

func unwrapFeature(t any) *Feature {
	reflectValue := reflect.ValueOf(t)
	if reflectValue.Kind() == reflect.Ptr {
		// Everything but windowed features
		return (*Feature)(reflectValue.UnsafePointer())
	} else if reflectValue.Kind() == reflect.Map {
		// Windowed feature
		windowedFeatureMap := reflectValue
		keys := windowedFeatureMap.MapKeys()
		if len(keys) == 0 {
			panic("exception occurred obtaining all buckets for windowed feature - no buckets found")
		}
		key := keys[0]
		ptrToPseudofeature := windowedFeatureMap.MapIndex(key)
		castedPtrToPseudofeature := (*Feature)(ptrToPseudofeature.UnsafePointer())

		baseWindowedFeatureFqn := strings.Split(castedPtrToPseudofeature.Fqn, "__")[0]
		baseWindowedFeature := Feature{Fqn: baseWindowedFeatureFqn}
		return &baseWindowedFeature

	}
	typePointedTo := reflect.ValueOf(t).Elem().Type()
	panic(fmt.Sprintf("unsupported type: %s", typePointedTo))
}

func UnwrapFeature(t any) *Feature {
	return unwrapFeature(t)
}
