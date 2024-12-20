package chalk

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/chalk-ai/chalk-go/internal/colls"
	"github.com/cockroachdb/errors"
	"reflect"
	"strings"
	"time"
)

func InitFeatures[T any](t *T) error {
	structValue := reflect.ValueOf(t).Elem()
	return initFeatures(structValue, "", make(map[string]bool))
}

type scopeTrie struct {
	children map[string]*scopeTrie
}

func (s *scopeTrie) addStr(fqn string) {
	s.add(strings.Split(fqn, "."))
}

func (s *scopeTrie) add(fqnParts []string) {
	if len(fqnParts) == 0 {
		return
	}
	firstPart := fqnParts[0]
	if s.children == nil {
		s.children = map[string]*scopeTrie{}
	}
	if _, found := s.children[firstPart]; !found {
		s.children[firstPart] = &scopeTrie{}
	}
	s.children[firstPart].add(fqnParts[1:])
}

func initRemoteFeatureMap(
	remoteFeatureMap map[string][]reflect.Value,
	structValue reflect.Value,
	cumulativeFqn string,
	visited map[string]bool,
	scope *scopeTrie,
	nsMemo internal.NamespaceMemo,
	scopeToJustStructs bool,
) error {
	if structValue.Kind() != reflect.Struct {
		return fmt.Errorf(
			"feature initialization function argument must be a reflect.Value"+
				" of the kind reflect.Struct, found %s instead",
			structValue.Kind().String(),
		)
	}

	structName := structValue.Type().Name()
	if isVisited, ok := visited[structName]; ok && isVisited {
		// Found a cycle. Just return.
		return nil
	}
	visited[structName] = true
	defer func() {
		visited[structName] = false
	}()

	memo := nsMemo[internal.ChalkpySnakeCase(structName)]

	var fieldNames []string
	if scopeToJustStructs {
		fieldNames = colls.Keys(memo.HasOneFieldsSet)
	} else {
		fieldNames = colls.Keys(scope.children)
	}

	for _, resolvedFieldName := range fieldNames {
		nextScope, inScope := scope.children[resolvedFieldName]
		if !inScope {
			continue
		}
		updatedFqn := fmt.Sprintf("%s.%s", cumulativeFqn, resolvedFieldName)
		fieldIndices, ok := memo.ResolvedFieldNameToIndices[resolvedFieldName]
		if !ok {
			// We arrive here when chalk-go receives a response that contains a feature
			// newly added to one of their has-one feature classes. They have not updated
			// their codegen'd structs yet, so we simply skip unmarshalling this new
			// feature to ensure forward compatibility.
			continue
		}

		for _, fieldIdx := range fieldIndices {
			f := structValue.Field(fieldIdx)

			if _, isStruct := memo.HasOneFieldsSet[resolvedFieldName]; isStruct {
				if !f.CanSet() {
					continue
				}
				if f.IsNil() {
					featureSet := reflect.New(f.Type().Elem())
					f.Set(featureSet)
				}
				if err := initRemoteFeatureMap(
					remoteFeatureMap,
					f.Elem(),
					updatedFqn,
					visited,
					nextScope,
					nsMemo,
					false,
				); err != nil {
					return err
				}
			} else {
				remoteFeatureMap[updatedFqn] = append(remoteFeatureMap[updatedFqn], f)
			}
		}
	}
	return nil
}

// initFeatures is a recursive function that initializes all features
// in the struct that is passed in. Each feature is initialized as a
// pointer to a Feature struct with the appropriate FQN.
func initFeatures(
	structValue reflect.Value,
	cumulativeFqn string,
	visited map[string]bool,
) error {
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
			return errors.Newf("field '%s' is not settable", fm.Meta.Name)
		}

		if f.Type().Elem().Kind() == reflect.Struct &&
			f.Type().Elem() != reflect.TypeOf(time.Time{}) {
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
			// TODO: This is an actual feature set we don't have to disguise it, just have to set it.
			ptrInDisguiseToFeatureSet := reflect.NewAt(f.Type().Elem(), featureSet.UnsafePointer())
			f.Set(ptrInDisguiseToFeatureSet)
			featureSetInDisguise := f.Elem()
			if err := initFeatures(featureSetInDisguise, updatedFqn, visited); err != nil {
				return err
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
				return errors.Newf(
					"the map type for Windowed features should a pointer"+
						" as its value type, but found %s instead",
					mapValueType.Kind(),
				)
			}

			windows := fm.Meta.Tag.Get("windows")
			for _, tag := range strings.Split(windows, ",") {
				seconds, err := internal.ParseBucketDuration(tag)
				if err != nil {
					return errors.Wrap(err, "error parsing bucket duration: %s")
				}
				updatedResolvedName := fmt.Sprintf("%s__%d__", resolvedName, seconds)
				bucketFqn := fmt.Sprintf("%s.%s", cumulativeFqn, updatedResolvedName)
				if f.IsNil() {
					f.Set(reflect.MakeMap(f.Type()))
				}
				feature := Feature{Fqn: bucketFqn}
				ptrInDisguiseToFeature := reflect.NewAt(mapValueType.Elem(), reflect.ValueOf(&feature).UnsafePointer())
				f.SetMapIndex(reflect.ValueOf(tag), ptrInDisguiseToFeature)
			}
		} else {
			// BASE CASE.
			if ptrErr := pointerCheck(f); ptrErr != nil {
				return ptrErr
			}

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
