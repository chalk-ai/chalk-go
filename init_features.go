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
	return newFeatureInitializer().initFeatures(structValue, "", make(map[string]bool), nil)
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

type featureInitializer struct {
	fieldsMap     map[string][]reflect.Value
	namespaceMemo internal.NamespaceMemo
}

func newFeatureInitializer() *featureInitializer {
	return &featureInitializer{
		fieldsMap:     map[string][]reflect.Value{},
		namespaceMemo: internal.NamespaceMemo{},
	}
}

// initFeatures is a recursive function that:
//
//  1. If not scoped:
//     Initializes all features in the struct that is passed in. Each feature is initialized
//     as a pointer to a Feature struct with the appropriate FQN.
//
//  2. If scoped:
//     Only the features that are in scope (stored in the form of a trie) are initialized.
//     In scope means that the feature is requested as an output and is returned in the
//     query response.
func (fi *featureInitializer) initFeatures(
	structValue reflect.Value,
	cumulativeFqn string,
	visited map[string]bool,
	scope *scopeTrie,
) error {
	isScoped := scope != nil
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
			_, inScope = scope.children[resolvedName]
		}

		if f.Kind() == reflect.Ptr && internal.IsTypeDataclass(f.Type().Elem()) && isScoped {
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

			if isScoped {
				if f.IsNil() {
					featureSet := reflect.New(f.Type().Elem())
					f.Set(featureSet)
				}
				newScope := scope.children[resolvedName]
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
				if isScoped {
					if _, bucketInScope := scope.children[updatedResolvedName]; !bucketInScope {
						continue
					}
					// Make map only if one of the bucket features need to be set
					// If no bucket features need to be set, map is nil.
					if f.IsNil() {
						f.Set(reflect.MakeMap(f.Type()))
					}
					fi.fieldsMap[bucketFqn] = append(fi.fieldsMap[bucketFqn], f)
				} else {
					if f.IsNil() {
						f.Set(reflect.MakeMap(f.Type()))
					}
					feature := Feature{Fqn: bucketFqn}
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

			if isScoped {
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

/*  buildNamespaceMemo populates a memo to make bulk-unmarshalling and has-many unmarshalling efficient.
 *  i.e. Don't need to do the same work for the same features class multiple times. Given:
 *  type User struct {
 *      Id *string
 *      Transactions *[]Transactions `has_many:"id,user_id"`
 *  }
 *  type Transactions struct {
 *      Id *string
 *      UserId *string
 *      Amount *float64
 *  }
 *  The namespace memo will be:
 *  {
 *      "User": {
 *          ResolvedFieldNameToIndex: {
 *              "id": 0,
 *              "user.id": 0,
 *              "transactions": 1,
 *              "user.transactions": 1,
 *          }
 *      },
 *      "Transactions": {
 *          ResolvedFieldNameToIndex: {
 *              "id": 0,
 *              "transactions.id": 0,
 *              "user_id": 1,
 *              "transactions.user_id": 1,
 *              "amount": 2,
 *              "transactions.amount": 2,
 *          }
 *      }
 *  }
 */
func (fi *featureInitializer) buildNamespaceMemo(typ reflect.Type) error {
	if typ.Kind() == reflect.Ptr {
		return fi.buildNamespaceMemo(typ.Elem())
	} else if typ.Kind() == reflect.Struct && typ != reflect.TypeOf(time.Time{}) {
		structName := typ.Name()
		namespace := internal.ChalkpySnakeCase(structName)
		for fieldIdx := 0; fieldIdx < typ.NumField(); fieldIdx++ {
			fm := typ.Field(fieldIdx)
			resolvedName, err := internal.ResolveFeatureName(fm)
			if err != nil {
				return errors.Wrapf(err, "error resolving feature name: %s", fm.Name)
			}

			if _, ok := fi.namespaceMemo[structName]; !ok {
				fi.namespaceMemo[structName] = &internal.NamespaceMemoItem{}
			}
			nsMemo := fi.namespaceMemo[structName]
			if nsMemo.ResolvedFieldNameToIndex == nil {
				nsMemo.ResolvedFieldNameToIndex = map[string]int{}
			}
			nsMemo.ResolvedFieldNameToIndex[resolvedName] = fieldIdx
			// Has-many features come back as a list of structs whose keys are namespaced FQNs.
			// Here we map those keys to their respective indices in the struct, so that we
			// don't have to do any string manipulation to deprefix the FQN when unmarshalling.
			nsMemo.ResolvedFieldNameToIndex[namespace+"."+resolvedName] = fieldIdx

			// Handle exploding windowed features
			if fm.Type.Kind() == reflect.Map {
				// Is a windowed feature
				intTags, err := internal.GetWindowBucketsSecondsFromStructTag(fm)
				if err != nil {
					return errors.Wrapf(
						err,
						"error getting window buckets for field '%s' in struct '%s'",
						fm.Name,
						structName,
					)
				}
				for _, tag := range intTags {
					bucketFqn := fmt.Sprintf("%s__%d__", resolvedName, tag)
					nsMemo.ResolvedFieldNameToIndex[bucketFqn] = fieldIdx
					nsMemo.ResolvedFieldNameToIndex[namespace+"."+bucketFqn] = fieldIdx
				}
			} else {
				if err := fi.buildNamespaceMemo(fm.Type); err != nil {
					return err
				}
			}
		}
	} else if typ.Kind() == reflect.Slice {
		return fi.buildNamespaceMemo(typ.Elem())
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
