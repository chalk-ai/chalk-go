package chalk

import (
	"fmt"
	"reflect"
)

func InitFeatures[T any](t *T) {
	structValue := reflect.ValueOf(t).Elem()
	initFeatures(structValue, "", make(map[string]bool))
}

func initFeatures(s reflect.Value, fqn string, visited map[string]bool) {
	if s.Kind() != reflect.Struct {
		panic(fmt.Sprintf("Feature initialization function argument must be a reflect.Value of the kind reflect.Struct, found %s instead", s.Kind().String()))
	}

	namespace := s.Type().Name()
	if fqn == "" {
		fqn = snakeCase(namespace)
	}

	if isVisited, ok := visited[namespace]; ok && isVisited {
		// This is not memoization. Simply a cycle checker while DFSing.
		return
	} else {
		visited[namespace] = true
	}

	for i := 0; i < s.NumField(); i++ {
		attributeName := s.Type().Field(i).Name
		updatedFqn := fqn + snakeCase(attributeName)

		f := s.Field(i)
		if !f.CanSet() || f.Kind() != reflect.Pointer {
			continue
		}

		if f.Type().Elem().Kind() == reflect.Struct {
			// Create new Feature Set instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport = new(CreditReport)
			//
			featureSet := reflect.New(f.Type().Elem())
			// Must use underlying value that's not wrapped by reflect,
			// otherwise we will cause weird bugs. See
			//     https://github.com/chalk-ai/chalk-go/pull/11
			underlyingPointerValue := featureSet.Interface()
			ptrInDisguiseToFeatureSet := reflect.NewAt(f.Type().Elem(), reflect.ValueOf(underlyingPointerValue).UnsafePointer())
			f.Set(ptrInDisguiseToFeatureSet)
			featureSetInDisguise := f.Elem()
			initFeatures(featureSetInDisguise, updatedFqn+".", visited)
		} else {
			// Create new Feature instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport.Id = (*string)(unsafe.Pointer(&Feature{"user.credit_report.id"}))
			//
			feature := Feature{Fqn: updatedFqn}
			ptrInDisguiseToFeature := reflect.NewAt(f.Type().Elem(), reflect.ValueOf(&feature).UnsafePointer())
			f.Set(ptrInDisguiseToFeature)
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
