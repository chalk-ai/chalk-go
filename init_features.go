package chalk

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/pkg/errors"
	"reflect"
	"regexp"
	"strings"
	"time"
)

func InitFeatures[T any](t *T) error {
	structValue := reflect.ValueOf(t).Elem()
	_, err := initFeatures(structValue, "", make(map[string]bool), "")
	return err
}

func initFeatureSingle(structValue reflect.Value, fqn string) ([]reflect.Value, error) {
	parts := strings.Split(fqn, ".")
	if len(parts) < 2 {
		return nil, fmt.Errorf("feature fqn should have at least two parts, found: '%s'", fqn)
	}
	return initFeatures(structValue, "", make(map[string]bool), strings.Join(parts[1:], "."))
}

// initFeatures is a recursive function that initializes all features
// in the struct that is passed in. Each feature is initialized as
// a pointer to a Feature struct with the appropriate FQN.
func initFeatures(
	structValue reflect.Value,
	cumulativeFqn string,
	visited map[string]bool,
	targetFqn string,
) ([]reflect.Value, error) {
	if structValue.Kind() != reflect.Struct {
		return nil, fmt.Errorf(
			"feature initialization function argument must be a reflect.Value"+
				" of the kind reflect.Struct, found %s instead",
			structValue.Kind().String(),
		)
	}

	namespace := structValue.Type().Name()
	if cumulativeFqn == "" && namespace != "" {
		cumulativeFqn = SnakeCase(namespace) + "."
	}

	if isVisited, ok := visited[namespace]; ok && isVisited {
		// Found a cycle. Just return.
		return nil, nil
	}
	visited[namespace] = true
	defer func() {
		visited[namespace] = false
	}()

	type fieldAndMeta struct {
		Field reflect.Value
		Meta  reflect.StructField
	}

	var fms []fieldAndMeta
	for i := 0; i < structValue.NumField(); i++ {
		fms = append(
			fms,
			fieldAndMeta{
				structValue.Field(i),
				structValue.Type().Field(i),
			},
		)
	}
	var resFields []reflect.Value
	for _, fm := range fms {
		resolvedName, err := internal.ResolveFeatureName(fm.Meta)
		if err != nil {
			return nil, errors.Wrapf(err, "error resolving feature name: %s", fm.Meta.Name)
		}
		updatedFqn := cumulativeFqn + resolvedName

		f := fm.Field
		if !f.CanSet() {
			continue
		}

		shouldNaivelySkip := targetFqn != "" && getFqnRoot(targetFqn) != resolvedName

		if f.Kind() == reflect.Ptr && internal.IsTypeDataclass(f.Type().Elem()) && targetFqn != "" {
			// If dataclasses are being initialized for purposes
			// of specifying query params, we want it to go to
			// the next block where we initialize it the same way
			// as a struct.
			//
			// If dataclass child fields are being selectively
			// initialized for purposes of deserialization of
			// query results, we simply return the field, as we
			// do in this block, and let the feature setter set
			// the value of the dataclass down the line.
			if shouldNaivelySkip {
				continue
			}
			return []reflect.Value{f}, nil
		} else if f.Type().Elem().Kind() == reflect.Struct &&
			f.Type().Elem() != reflect.TypeOf(time.Time{}) {
			// RECURSIVE CASE.
			// Create new Feature Set instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport = new(CreditReport)
			//
			if shouldNaivelySkip {
				continue
			}
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return nil, ptrErr
			}

			if targetFqn == "" {
				featureSet := reflect.New(f.Type().Elem())
				// TODO: This is an actual feature set we don't have to disguise it, just have to set it.
				ptrInDisguiseToFeatureSet := reflect.NewAt(f.Type().Elem(), featureSet.UnsafePointer())
				f.Set(ptrInDisguiseToFeatureSet)
				featureSetInDisguise := f.Elem()
				if _, err := initFeatures(featureSetInDisguise, updatedFqn+".", visited, targetFqn); err != nil {
					return nil, err
				}
			} else {
				if f.IsNil() {
					featureSet := reflect.New(f.Type().Elem())
					f.Set(featureSet)
				}
				parts := strings.Split(targetFqn, ".")
				if len(parts) < 2 {
					return nil, fmt.Errorf(
						"feature fqn should have at least two parts, found: '%s'",
						targetFqn,
					)
				}
				nextTargetFqn := strings.Join(parts[1:], ".")
				return initFeatures(f.Elem(), updatedFqn+".", visited, nextTargetFqn)
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

			if targetFqn != "" {
				fqnRoot := getFqnRoot(targetFqn)
				// The target fqn here could be
				//    - "some_other_feature_in_this_feature_class"
				//    - "this_windowed_feature__600__"
				//    - "other_windowed_feature__600__"
				// So let's extract "this_windowed_feature" from the target fqn
				// and compare it with the current feature's name.
				pattern, err := regexp.Compile(fmt.Sprintf(`^%s__\d+__$`, resolvedName))
				if err != nil {
					return nil, errors.Wrap(err, "error compiling regex to match windowed feature names")
				}
				if !pattern.Match([]byte(fqnRoot)) {
					// Not any bucket feature in this windowed feature class
					continue
				}
			}

			mapValueType := f.Type().Elem()
			if mapValueType.Kind() != reflect.Pointer {
				return nil, fmt.Errorf(
					"the map type for Windowed features should a pointer"+
						" as its value type, but found %s instead",
					mapValueType.Kind(),
				)
			}

			targetMap := f
			if targetFqn == "" {
				// Initializing all features, always
				// make a new map.
				targetMap = reflect.MakeMap(f.Type())
				f.Set(targetMap)
				windows := fm.Meta.Tag.Get("windows")
				for _, tag := range strings.Split(windows, ",") {
					seconds, parseErr := internal.ParseBucketDuration(tag)
					if parseErr != nil {
						return nil, fmt.Errorf("error parsing bucket duration: %s", parseErr)
					}
					windowFqn := updatedFqn + fmt.Sprintf("__%d__", seconds)
					if targetFqn == "" {
						feature := Feature{Fqn: windowFqn}
						ptrInDisguiseToFeature := reflect.NewAt(mapValueType.Elem(), reflect.ValueOf(&feature).UnsafePointer())
						targetMap.SetMapIndex(reflect.ValueOf(tag), ptrInDisguiseToFeature)
					}
				}
			} else {
				// Selectively initializing features,
				// use a new map only if the map is nil.
				if targetMap.IsNil() {
					targetMap = reflect.MakeMap(f.Type())
					f.Set(targetMap)
				}
				return []reflect.Value{f}, nil
			}
		} else {
			// BASE CASE.
			if shouldNaivelySkip {
				continue
			}
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return nil, ptrErr
			}

			if targetFqn != "" {
				resFields = append(resFields, f)
			} else {
				// Create new Feature instance and point to it.
				// The equivalent way of doing it without 'reflect':
				//
				//      Features.User.CreditReport.Id = (*string)(unsafe.Pointer(&Feature{"user.credit_report.id"}))
				//

				// Dataclass fields are not actually real features,
				// so when we are initializing the root Features struct,
				// we want to return the parent (real) feature FQN
				// instead of the fake FQN of the dataclass child field.
				if internal.IsDataclass(structValue) {
					updatedFqn = DesuffixFqn(updatedFqn)
				}

				feature := Feature{Fqn: updatedFqn}
				ptrInDisguiseToFeature := reflect.NewAt(f.Type().Elem(), reflect.ValueOf(&feature).UnsafePointer())
				f.Set(ptrInDisguiseToFeature)
			}
		}
	}
	return resFields, nil
}

func getFqnRoot(s string) string {
	return strings.Split(s, ".")[0]
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
