package chalk

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strings"
)

func InitFeatures[T any](t *T) {
	structValue := reflect.ValueOf(t).Elem()
	initFeatures(structValue, "", make(map[string]bool), nil)
}

func initFeatures(structValue reflect.Value, fqn string, visited map[string]bool, fieldMap fqnToField) {
	if structValue.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Feature initialization function argument must be a reflect.Value of the kind reflect.Struct, found %s instead", structValue.Kind().String()))
	}

	namespace := structValue.Type().Name()
	if fqn == "" && namespace != "" {
		fqn = snakeCase(namespace) + "."
	}

	if isVisited, ok := visited[namespace]; ok && isVisited {
		// This is not memoization. Simply a cycle checker while DFSing.
		return
	} else {
		visited[namespace] = true
	}

	for i := 0; i < structValue.NumField(); i++ {
		f := structValue.Field(i)
		fieldMeta := structValue.Type().Field(i)

		attributeName := fieldMeta.Name
		updatedFqn := fqn + snakeCase(attributeName)

		if !f.CanSet() {
			continue
		}

		if f.Type().Elem().Kind() == reflect.Struct && f.Type().Elem().String() != "time.Time" {
			// Create new Feature Set instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport = new(CreditReport)
			//
			featureSet := reflect.New(f.Type().Elem())
			ptrInDisguiseToFeatureSet := reflect.NewAt(f.Type().Elem(), featureSet.UnsafePointer())
			f.Set(ptrInDisguiseToFeatureSet)
			featureSetInDisguise := f.Elem()
			initFeatures(featureSetInDisguise, updatedFqn+".", visited, fieldMap)
		} else if f.Kind() == reflect.Map {
			newMap := reflect.MakeMap(f.Type())
			windows := fieldMeta.Tag.Get("windows")
			for _, tag := range strings.Split(windows, ",") {
				seconds, parseErr := internal.ParseBucketDuration(tag)
				if parseErr != nil {
					panic(fmt.Sprintf("error parsing bucket duration: %s", parseErr))
				}
				windowFqn := updatedFqn + fmt.Sprintf("__%d__", seconds)
				if fieldMap == nil {
					feature := Feature{Fqn: windowFqn}
					ptrInDisguiseToFeature := reflect.NewAt(f.Type().Elem().Elem(), reflect.ValueOf(&feature).UnsafePointer())
					newMap.SetMapIndex(reflect.ValueOf(tag), ptrInDisguiseToFeature)
				} else {
					nilPointer := reflect.New(f.Type().Elem()).Elem()
					nilPointer.Set(reflect.Zero(nilPointer.Type()))
					newMap.SetMapIndex(reflect.ValueOf(tag), nilPointer)
				}
			}
			f.Set(newMap)
			if fieldMap != nil {
				fieldMap[updatedFqn] = f
			}
		} else {
			// Create new Feature instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport.Id = (*string)(unsafe.Pointer(&Feature{"user.credit_report.id"}))
			//
			if fieldMap != nil {
				fieldMap[updatedFqn] = f
			} else {
				feature := Feature{Fqn: updatedFqn}
				ptrInDisguiseToFeature := reflect.NewAt(f.Type().Elem(), reflect.ValueOf(&feature).UnsafePointer())
				f.Set(ptrInDisguiseToFeature)
			}

		}
	}

	visited[namespace] = false
}

func snakeCase(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isASCIIUpper(c) {
			if i > 0 && s[i-1] != '.' {
				b = append(b, '_')
			}
			c += 'a' - 'A'
		}
		b = append(b, c)
	}
	return string(b)
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}
