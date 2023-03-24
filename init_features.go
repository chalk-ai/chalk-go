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

// initFeatures is a recursive function that initializes all features
// in the struct that is passed in. Each feature is initialized as
// a pointer to a Feature struct with the appropriate FQN.
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

		isTimeField := f.Type().Elem().Kind() == reflect.Struct && f.Type().Elem().String() == "time.Time"

		attributeName := fieldMeta.Name
		updatedFqn := fqn + snakeCase(attributeName)

		if !f.CanSet() {
			continue
		}

		if f.Type().Elem().Kind() == reflect.Struct && !isTimeField {
			// RECURSIVE CASE.
			// Create new Feature Set instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport = new(CreditReport)
			//
			pointerCheck(f)
			featureSet := reflect.New(f.Type().Elem())
			ptrInDisguiseToFeatureSet := reflect.NewAt(f.Type().Elem(), featureSet.UnsafePointer())
			f.Set(ptrInDisguiseToFeatureSet)
			featureSetInDisguise := f.Elem()
			initFeatures(featureSetInDisguise, updatedFqn+".", visited, fieldMap)
		} else if f.Kind() == reflect.Map {
			// Creates a map of tag values to pointers to Features.
			// For example, if we have the tag "windows=6h,12h,1d",
			// then the map will be:
			//
			// 		map[string]*int64{
			// 			"6h": &Feature{Fqn: "user.clicks__21600__"},
			// 			"12h": &Feature{Fqn: "user.clicks__43200__"},
			// 			"1d": &Feature{Fqn: "user.clicks__86400__"},
			// 		}
			//
			// Notice that while the values is typed as *int64, it is
			// actually a pointer to a Feature struct. See BASE CASE
			// section below.
			mapValueType := f.Type().Elem()
			if mapValueType.Kind() != reflect.Pointer {
				panic(fmt.Sprintf("the map type for Windowed features should a pointer as its value type, but found %s instead", mapValueType.Kind()))
			}
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
					ptrInDisguiseToFeature := reflect.NewAt(mapValueType.Elem(), reflect.ValueOf(&feature).UnsafePointer())
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
			// BASE CASE.
			// Create new Feature instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport.Id = (*string)(unsafe.Pointer(&Feature{"user.credit_report.id"}))
			//
			pointerCheck(f)
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

func pointerCheck(field reflect.Value) {
	if field.Kind() != reflect.Ptr {
		panic(fmt.Sprintf("expected a pointer type but found %s -- make sure the generated feature structs are unchanged, and that every field is of a pointer type except for Windowed feature types", field.Kind()))
	}
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
