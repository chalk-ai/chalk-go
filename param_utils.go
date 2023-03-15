package chalk

import (
	"fmt"
	"reflect"
)

func unwrapFeature(t any) *Feature {
	ptrInDisguiseToFeature := reflect.ValueOf(t)
	if ptrInDisguiseToFeature.Kind() == reflect.Ptr {
		return (*Feature)(ptrInDisguiseToFeature.UnsafePointer())
	}

	typePointedTo := reflect.ValueOf(t).Elem().Type()
	panic(fmt.Sprintf("unsupported type: %s", typePointedTo))
}

func unwrapFeatureInterface(t any) *Feature {
	return unwrapFeature(t)
}
