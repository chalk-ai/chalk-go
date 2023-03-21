package chalk

import (
	"fmt"
	"reflect"
)

type Feature struct {
	Fqn string
}

func unwrapFeature(t any) *Feature {
	ptrInDisguiseToFeature := reflect.ValueOf(t)
	if ptrInDisguiseToFeature.Kind() == reflect.Ptr {
		return (*Feature)(ptrInDisguiseToFeature.UnsafePointer())
	}

	typePointedTo := reflect.ValueOf(t).Elem().Type()
	panic(fmt.Sprintf("unsupported type: %s", typePointedTo))
}

func UnwrapFeature(t any) *Feature {
	return unwrapFeature(t)
}
