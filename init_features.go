package chalk

import (
	"fmt"
	"github.com/chalk-ai/chalk-go/internal"
	"github.com/cockroachdb/errors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func InitFeatures[T any](t *T) error {
	structValue := reflect.ValueOf(t).Elem()
	return initFeatures(structValue, "", make(map[string]bool))
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

		updatedFqn := cumulativeFqn + "." + resolvedName
		if cumulativeFqn == "" {
			updatedFqn = resolvedName
		}

		f := fm.Field
		if !f.CanSet() {
			continue
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
				bucketFqn := cumulativeFqn + "." + resolvedName + "__" + strconv.Itoa(seconds) + "__"
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

/* WarmUpUnmarshaller builds a memo to make unmarshalling efficient. This function should be called only once
 * at init time, instead of per query. If this function is not called, the first query will be slower, but
 * subsequent queries that unmarshals into the same structs will be faster because they will use the memo built
 * implicitly by the first query.
 *
 * This function takes in either an anonymous struct that contains all feature structs, or an individual
 * feature struct. It also recursively builds memos for all nested feature structs.
 *
 * Example usage:
 *  type User struct {
 *      Id *string
 *      Transactions *[]Transactions `has_many:"id,user_id"`
 *      Grade   *int `versioned:"default(2)"`
 *      GradeV1 *int `versioned:"true"`
 *      GradeV2 *int `versioned:"true"`
 *  }
 *  type Transactions struct {
 *      Id *string
 *      UserId *string
 *      Amount *float64
 *  }
 *  var Features struct {
 *      User *User
 *      Transactions *Transactions
 *  }
 *  func init() {
 *      if err := chalk.WarmUpUnmarshaller(&Features); err != nil {
 *          panic("error initializing unmarshalling")
 *      }
 *  }
 */
func WarmUpUnmarshaller[T any](featureStruct *T) error {
	elemType := reflect.TypeOf(featureStruct).Elem()
	if elemType.Kind() != reflect.Struct {
		return fmt.Errorf(
			"argument must be a pointer to a struct, found a pointer to `%s` instead",
			elemType.Kind(),
		)
	}
	return internal.PopulateAllNamespaceMemoNew(elemType, nil)
}

func pointerCheck(field reflect.Value) error {
	if field.Kind() != reflect.Ptr {
		return fmt.Errorf("expected a pointer type but found %s -- make sure the generated feature structs are unchanged, and that every field is of a pointer type except for Windowed feature types", field.Kind())
	}
	return nil
}

func SnakeCase(s string) string {
	return internal.LegacySnakeCase(s)
}
