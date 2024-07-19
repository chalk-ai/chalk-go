package chalk

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"reflect"
	"strings"
	"time"
)

func InitFeatures[T any](t *T) error {
	structValue := reflect.ValueOf(t).Elem()
	initializer := NewFeatureInitializer(initializerModeAsFeature)
	return initializer.initFeatures(structValue, "", make(map[string]bool), nil)
}

type scopeTrie struct {
	Children map[string]*scopeTrie
}

func (s *scopeTrie) addStr(fqn string) {
	s.add(strings.Split(fqn, "."))
}

func (s *scopeTrie) add(fqnParts []string) {
	if len(fqnParts) == 0 {
		return
	}

	firstPart := fqnParts[0]
	if s.Children == nil {
		s.Children = map[string]*scopeTrie{}
	}

	kid := s.Children[firstPart]
	if kid == nil {
		s.Children[firstPart] = &scopeTrie{}
	}

	s.Children[firstPart].add(fqnParts[1:])
}

type featureInitializer struct {
	fieldsMap map[string][]reflect.Value
	isScoped  bool
}

type initializerMode string

var (
	// For creating `Feature` structs that contains
	// an FQN field. Used for specifying query params.
	initializerModeAsFeature = initializerMode("as_feature")

	// For unmarshalling into feature structs. We initialize
	// the fields, then another function takes care of setting
	// the fields to the correct value.
	initializerModeUnmarshal = initializerMode("unmarshal")
)

func NewFeatureInitializer(mode initializerMode) *featureInitializer {
	return &featureInitializer{
		fieldsMap: map[string][]reflect.Value{},
		isScoped:  mode == initializerModeUnmarshal,
	}
}

// initFeatures is a recursive function that:
//
//  1. If not scoped:
//     Initializes all features in the struct that is passed in. Each feature is initialized
//     as a pointer to a Feature struct with the appropriate FQN.
//
//  2. If is scoped:
//     Only the features that are in scope (stored in the form of a trie) are initialized.
//     In scope means that the feature is requested as an output and is returned in the
//     query response.
func (fi *featureInitializer) initFeatures(
	structValue reflect.Value,
	cumulativeFqn string,
	visited map[string]bool,
	scope *scopeTrie,
) error {
	if fi.isScoped && scope == nil {
		return errors.New("scope cannot be nil when initializing features in a scoped manner")
	}

	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf(
			"feature initialization function argument must be a reflect.Value"+
				" of the kind reflect.Struct, found %s instead",
			structValue.Kind().String(),
		)
	}

	namespace := structValue.Type().Name()
	if isVisited, ok := visited[namespace]; ok && isVisited {
		// Found a cycle. Just return.
		return nil
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
	for _, fm := range fms {
		resolvedName, err := internal.ResolveFeatureName(fm.Meta)
		if err != nil {
			return errors.Wrapf(err, "error resolving feature name: %s", fm.Meta.Name)
		}

		updatedFqn := fmt.Sprintf("%s.%s", cumulativeFqn, resolvedName)
		if cumulativeFqn == "" {
			updatedFqn = resolvedName
		}

		f := fm.Field
		if !f.CanSet() {
			continue
		}

		inScope := true
		if scope != nil {
			_, inScope = scope.Children[resolvedName]
		}

		if f.Kind() == reflect.Ptr && internal.IsTypeDataclass(f.Type().Elem()) && fi.isScoped {
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
			if !inScope {
				continue
			}
			fi.fieldsMap[updatedFqn] = append(fi.fieldsMap[updatedFqn], f)
		} else if f.Type().Elem().Kind() == reflect.Struct &&
			f.Type().Elem() != reflect.TypeOf(time.Time{}) {
			// RECURSIVE CASE.
			// Create new Feature Set instance and point to it.
			// The equivalent way of doing it without 'reflect':
			//
			//      Features.User.CreditReport = new(CreditReport)
			//
			if !inScope {
				continue
			}
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return ptrErr
			}

			if fi.isScoped {
				if f.IsNil() {
					featureSet := reflect.New(f.Type().Elem())
					f.Set(featureSet)
				}
				newScope := scope.Children[resolvedName]
				if newScope == nil {
					return errors.Newf("scope not found for feature '%s'", cumulativeFqn)
				}
				if err := fi.initFeatures(f.Elem(), updatedFqn, visited, newScope); err != nil {
					return err
				}
			} else {
				featureSet := reflect.New(f.Type().Elem())
				// TODO: This is an actual feature set we don't have to disguise it, just have to set it.
				ptrInDisguiseToFeatureSet := reflect.NewAt(f.Type().Elem(), featureSet.UnsafePointer())
				f.Set(ptrInDisguiseToFeatureSet)
				featureSetInDisguise := f.Elem()
				if err := fi.initFeatures(featureSetInDisguise, updatedFqn, visited, nil); err != nil {
					return err
				}
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

			// FIXME: Should be removeable
			//if fi.isScoped {
			//	// The target fqn here could be
			//	//    - "some_other_feature_in_this_feature_class"
			//	//    - "this_windowed_feature__600__"
			//	//    - "other_windowed_feature__600__"
			//	// So let's extract "this_windowed_feature" from the target fqn
			//	// and compare it with the current feature's name.
			//	pattern, err := regexp.Compile(fmt.Sprintf(`^%s__\d+__$`, resolvedName))
			//	if err != nil {
			//		return nil, errors.Wrap(err, "error compiling regex to match windowed feature names")
			//	}
			//	if !pattern.Match([]byte(getFqnRoot(targetFqn))) {
			//		// Not any bucket feature in this windowed feature class
			//		continue
			//	}
			//}

			mapValueType := f.Type().Elem()
			if mapValueType.Kind() != reflect.Pointer {
				return errors.Newf(
					"the map type for Windowed features should a pointer"+
						" as its value type, but found %s instead",
					mapValueType.Kind(),
				)
			}

			//if fi.isScoped {
			//	// Selectively initializing features,
			//	// use a new map only if the map is nil.
			//	if f.IsNil() {
			//		f.Set(reflect.MakeMap(f.Type()))
			//	}
			//	return []reflect.Value{f}, nil
			//} else {
			//	// Initializing all features, always
			//	// make a new map.
			//	f.Set(reflect.MakeMap(f.Type()))
			//	windows := fm.Meta.Tag.Get("windows")
			//	for _, tag := range strings.Split(windows, ",") {
			//		seconds, parseErr := internal.ParseBucketDuration(tag)
			//		if parseErr != nil {
			//			return nil, fmt.Errorf("error parsing bucket duration: %s", parseErr)
			//		}
			//		windowFqn := updatedFqn + fmt.Sprintf("__%d__", seconds)
			//		if targetFqn == "" {
			//			feature := Feature{Fqn: windowFqn}
			//			ptrInDisguiseToFeature := reflect.NewAt(mapValueType.Elem(), reflect.ValueOf(&feature).UnsafePointer())
			//			f.SetMapIndex(reflect.ValueOf(tag), ptrInDisguiseToFeature)
			//		}
			//	}
			//}

			windows := fm.Meta.Tag.Get("windows")
			for _, tag := range strings.Split(windows, ",") {
				seconds, err := internal.ParseBucketDuration(tag)
				if err != nil {
					return errors.Wrap(err, "error parsing bucket duration: %s")
				}
				updatedResolvedName := fmt.Sprintf("%s__%d__", resolvedName, seconds)
				windowFqn := fmt.Sprintf("%s.%s", cumulativeFqn, updatedResolvedName)
				if fi.isScoped {
					if _, bucketInScope := scope.Children[updatedResolvedName]; !bucketInScope {
						continue
					}
					// Make map only if one of the bucket features need to be set
					// If no bucket features need to be set, map is nil.
					if f.IsNil() {
						f.Set(reflect.MakeMap(f.Type()))
					}
					fi.fieldsMap[windowFqn] = append(fi.fieldsMap[windowFqn], f)
				} else {
					if f.IsNil() {
						f.Set(reflect.MakeMap(f.Type()))
					}
					feature := Feature{Fqn: windowFqn}
					ptrInDisguiseToFeature := reflect.NewAt(mapValueType.Elem(), reflect.ValueOf(&feature).UnsafePointer())
					f.SetMapIndex(reflect.ValueOf(tag), ptrInDisguiseToFeature)
				}

			}
		} else {
			// BASE CASE.
			if !inScope {
				continue
			}
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return ptrErr
			}

			if fi.isScoped {
				fi.fieldsMap[updatedFqn] = append(fi.fieldsMap[updatedFqn], f)
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
	return nil
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
