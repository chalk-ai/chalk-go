package chalk

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strings"
)

func InitFeatures[T any](t *T) error {
	structValue := reflect.ValueOf(t).Elem()
	return initFeatures(structValue, "", make(map[string]bool), nil)
}

// initFeatures is a recursive function that initializes all features
// in the struct that is passed in. Each feature is initialized as
// a pointer to a Feature struct with the appropriate FQN.
func initFeatures(structValue reflect.Value, fqn string, visited map[string]bool, fieldMap fqnToField) error {
	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf("feature initialization function argument must be a reflect.Value of the kind reflect.Struct, found %s instead", structValue.Kind().String())
	}

	namespace := structValue.Type().Name()
	if fqn == "" && namespace != "" {
		fqn = snakeCase(namespace) + "."
	}

	if isVisited, ok := visited[namespace]; ok && isVisited {
		// This is not memoization. Simply a cycle checker while DFSing.
		return nil
	} else {
		visited[namespace] = true
	}

	for i := 0; i < structValue.NumField(); i++ {
		f := structValue.Field(i)
		fieldMeta := structValue.Type().Field(i)

		isTimeField := f.Type().Elem().Kind() == reflect.Struct && f.Type().Elem().String() == "time.Time"

		attributeName := snakeCase(fieldMeta.Name)
		nameOverride := fieldMeta.Tag.Get("name")
		if nameOverride != "" {
			attributeName = nameOverride
		}
		updatedFqn := fqn + attributeName

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
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return ptrErr
			}
			featureSet := reflect.New(f.Type().Elem())
			ptrInDisguiseToFeatureSet := reflect.NewAt(f.Type().Elem(), featureSet.UnsafePointer())
			f.Set(ptrInDisguiseToFeatureSet)
			featureSetInDisguise := f.Elem()
			initErr := initFeatures(featureSetInDisguise, updatedFqn+".", visited, fieldMap)
			if initErr != nil {
				return initErr
			}
			if fieldMap != nil {
				fieldMap[updatedFqn] = f
			}
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
				return fmt.Errorf("the map type for Windowed features should a pointer as its value type, but found %s instead", mapValueType.Kind())
			}
			newMap := reflect.MakeMap(f.Type())
			windows := fieldMeta.Tag.Get("windows")
			for _, tag := range strings.Split(windows, ",") {
				seconds, parseErr := internal.ParseBucketDuration(tag)
				if parseErr != nil {
					return fmt.Errorf("error parsing bucket duration: %s", parseErr)
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
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return ptrErr
			}
			if fieldMap != nil {
				fieldMap[updatedFqn] = f
			} else {
				// Create new Feature instance and point to it.
				// The equivalent way of doing it without 'reflect':
				//
				//      Features.User.CreditReport.Id = (*string)(unsafe.Pointer(&Feature{"user.credit_report.id"}))
				//

				// Dataclass fields are not actually real features,
				// so when we are initializing the root Features struct (fieldMap == nil),
				// we want to return the parent (real) feature FQN
				// instead of the fake FQN of the dataclass field.
				if internal.IsDataclass(structValue) {
					updatedFqn = DesuffixFqn(updatedFqn)
				}

				feature := Feature{Fqn: updatedFqn}
				ptrInDisguiseToFeature := reflect.NewAt(f.Type().Elem(), reflect.ValueOf(&feature).UnsafePointer())
				f.Set(ptrInDisguiseToFeature)
			}
		}
	}

	visited[namespace] = false
	return nil
}

func pointerCheck(field reflect.Value) error {
	if field.Kind() != reflect.Ptr {
		return fmt.Errorf("expected a pointer type but found %s -- make sure the generated feature structs are unchanged, and that every field is of a pointer type except for Windowed feature types", field.Kind())
	}
	return nil
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
		} else if isASCIIDigit(c) && i > 0 && isASCIILower(s[i-1]) {
			b = append(b, '_')
		}
		b = append(b, c)
	}
	return string(b)
}

func isASCIILower(c byte) bool {
	return 'a' <= c && c <= 'z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}
