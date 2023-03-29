package chalk

import (
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
)

func InitFeatures[T any](t *T) error {
	structValue := reflect.ValueOf(t).Elem()
	return internal.InitFeatureInternal(structValue, "", make(map[string]bool), nil)
}
