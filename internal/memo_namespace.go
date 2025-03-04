package internal

import (
	"github.com/cockroachdb/errors"
	"reflect"
	"strconv"
	"time"
)

type NamespaceMemo struct {
	// Root and non-root FQN as keys
	ResolvedFieldNameToIndices map[string][]int
	// Non-root FQN as keys only
	StructFieldsSet map[string]bool
}

type NamespaceMemosT MemosT[reflect.Type, NamespaceMemo]

var NamespaceMemos = &NamespaceMemosT{}

func (m *NamespaceMemosT) LoadOrStore(structType reflect.Type) (*NamespaceMemo, error) {
	return (*MemosT[reflect.Type, NamespaceMemo])(m).LoadOrStore(structType, func() (*NamespaceMemo, error) {
		return generateNamespaceMemo(structType, nil)
	})
}

func generateNamespaceMemo(typ reflect.Type, visited map[reflect.Type]bool) (*NamespaceMemo, error) {
	structName := typ.Name()
	namespace := ChalkpySnakeCase(structName)
	nsMemo := NamespaceMemo{
		ResolvedFieldNameToIndices: map[string][]int{},
		StructFieldsSet:            map[string]bool{},
	}
	for fieldIdx := 0; fieldIdx < typ.NumField(); fieldIdx++ {
		fm := typ.Field(fieldIdx)
		resolvedName, err := ResolveFeatureName(fm)
		if err != nil {
			return nil, errors.Wrapf(err, "error resolving feature name: %s", fm.Name)
		}
		nsMemo.ResolvedFieldNameToIndices[resolvedName] = append(nsMemo.ResolvedFieldNameToIndices[resolvedName], fieldIdx)
		// Has-many features come back as a list of structs whose keys are namespaced FQNs.
		// Here we map those keys to their respective indices in the struct, so that we
		// don't have to do any string manipulation to deprefix the FQN when unmarshalling.
		rootFqn := namespace + "." + resolvedName
		nsMemo.ResolvedFieldNameToIndices[rootFqn] = append(nsMemo.ResolvedFieldNameToIndices[rootFqn], fieldIdx)

		// Handle exploding windowed features
		if fm.Type.Kind() == reflect.Map {
			// Is a windowed feature
			intTags, err := GetWindowBucketsSecondsFromStructTag(fm)
			if err != nil {
				return nil, errors.Wrapf(
					err,
					"error getting window buckets for field '%s' in struct '%s'",
					fm.Name,
					structName,
				)
			}
			for _, tag := range intTags {
				bucketFqn := resolvedName + "__" + strconv.Itoa(tag) + "__"
				nsMemo.ResolvedFieldNameToIndices[bucketFqn] = append(nsMemo.ResolvedFieldNameToIndices[bucketFqn], fieldIdx)
				rootBucketFqn := namespace + "." + bucketFqn
				nsMemo.ResolvedFieldNameToIndices[rootBucketFqn] = append(nsMemo.ResolvedFieldNameToIndices[rootBucketFqn], fieldIdx)
			}
		} else {
			if err := PopulateAllNamespaceMemo(fm.Type, visited); err != nil {
				return nil, errors.Wrapf(err, "populating namespace memo for field '%s'", fm.Name)
			}
		}

		if fm.Type.Kind() == reflect.Ptr && IsStruct(fm.Type.Elem()) && !IsTypeDataclass(fm.Type.Elem()) {
			nsMemo.StructFieldsSet[resolvedName] = true
		}
	}

	return &nsMemo, nil
}

/*PopulateAllNamespaceMemo populates a memo to make bulk-unmarshalling and has-many unmarshalling efficient.
 *  i.e. Don't need to do the same work for the same features class multiple times. Given:
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
 *  The namespace memo will be:
 *  {
 *      "User": {
 *          ResolvedFieldNameToIndices: {
 *              "id": [0],
 *              "user.id": [0],
 *              "grade@2": [2, 4],
 *              "user.grade@2": [2, 4],
 *              "grade": [3],
 *              "user.grade": [3],
 *              "transactions": [1],
 *              "user.transactions": [1],
 *          }
 *      },
 *      "Transactions": {
 *          ResolvedFieldNameToIndices: {
 *              "id": [0],
 *              "transactions.id": [0],
 *              "user_id": [1],
 *              "transactions.user_id": [1],
 *              "amount": [2],
 *              "transactions.amount": [2],
 *          }
 *      }
 *  }
 */
func PopulateAllNamespaceMemo(typ reflect.Type, visited map[reflect.Type]bool) error {
	if visited == nil {
		visited = map[reflect.Type]bool{}
	}
	if typ.Kind() == reflect.Ptr {
		return PopulateAllNamespaceMemo(typ.Elem(), visited)
	} else if typ.Kind() == reflect.Struct && typ != reflect.TypeOf(time.Time{}) {
		if visited[typ] {
			return nil
		}
		visited[typ] = true
		if _, err := NamespaceMemos.LoadOrStore(typ); err != nil {
			return errors.Wrap(err, "load-or-storing memo")
		}
	} else if typ.Kind() == reflect.Slice {
		return PopulateAllNamespaceMemo(typ.Elem(), visited)
	}
	return nil
}
