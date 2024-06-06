package chalk

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func InitFeatures[T any](t *T) error {
	structValue := reflect.ValueOf(t).Elem()
	return initFeatures(structValue, "", make(map[string]bool), nil)
}

// initFeatures is a recursive function that initializes all features
// in the struct that is passed in. Each feature is initialized as
// a pointer to a Feature struct with the appropriate FQN.
func initFeatures(structValue reflect.Value, fqn string, visited map[string]bool, fieldMap fqnToFields) error {
	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf("feature initialization function argument must be a reflect.Value of the kind reflect.Struct, found %s instead", structValue.Kind().String())
	}

	namespace := structValue.Type().Name()
	if fqn == "" && namespace != "" {
		fqn = SnakeCase(namespace) + "."
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

		attributeName := SnakeCase(fieldMeta.Name)
		nameOverride := fieldMeta.Tag.Get("name")
		if nameOverride != "" {
			attributeName = nameOverride
		}
		updatedFqn := fqn + attributeName

		if !f.CanSet() {
			continue
		}

		if f.Type().Elem().Kind() == reflect.Struct && f.Type().Elem() != reflect.TypeOf(time.Time{}) {
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
				fieldMap.addField(updatedFqn, f)
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
				if fieldMap != nil {
					fieldMap.addField(windowFqn, f)
				}
			}
			f.Set(newMap)
			if fieldMap != nil {
				fieldMap.addField(updatedFqn, f)
			}
		} else {
			// BASE CASE.
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return ptrErr
			}

			versioned := fieldMeta.Tag.Get("versioned")
			if versioned == "true" {
				parts := strings.Split(updatedFqn, "_")
				nameErr := fmt.Errorf("versioned feature must have a version suffix `VN` at the end of the attribute name, but found '%s' instead", fieldMeta.Name)
				if len(parts) == 1 {
					return nameErr
				}
				lastPart := parts[len(parts)-1]
				if !strings.HasPrefix(lastPart, "v") {
					return nameErr
				}
				version := lastPart[1:]
				baseFqn := strings.Join(parts[:len(parts)-1], "_")
				if version == "1" {
					updatedFqn = baseFqn
				} else {
					updatedFqn = baseFqn + "@" + version
				}
			} else if strings.HasPrefix(versioned, "default(") && strings.HasSuffix(versioned, ")") {
				version := versioned[len("default(") : len(versioned)-len(")")]
				_, convertErr := strconv.Atoi(version)
				if convertErr != nil {
					return fmt.Errorf("Expected struct tag `versioned:\"default(N)\"` where N is an integer, but found %s instead", versioned)
				}
				if version != "1" {
					updatedFqn = updatedFqn + "@" + version
				}
			} else if versioned != "" {
				return fmt.Errorf("Expected struct tag `versioned:\"true\"` or `versioned:\"default(N)\"` where N is an integer, but found '%s' instead", versioned)
			}

			if fieldMap != nil {
				fieldMap.addField(updatedFqn, f)
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

func SnakeCase(s string) string {
	return internal.ChalkpySnakeCase(s)
}
