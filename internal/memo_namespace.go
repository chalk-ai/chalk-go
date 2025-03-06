package internal

import (
	"github.com/cockroachdb/errors"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type NamespaceMemo struct {
	// Root and non-root FQN as keys
	ResolvedFieldNameToIndices map[string][]int
	// Non-root FQN as keys only
	StructFieldsSet map[string]bool
}

type AllNamespaceMemoT sync.Map

var AllNamespaceMemo = &AllNamespaceMemoT{}

type NamespaceMutex struct {
	mu   sync.RWMutex
	memo *NamespaceMemo
}

func NewNamespaceMemo() *NamespaceMemo {
	return &NamespaceMemo{
		ResolvedFieldNameToIndices: map[string][]int{},
		StructFieldsSet:            map[string]bool{},
	}
}

func (m *AllNamespaceMemoT) Load(key reflect.Type) (*NamespaceMemo, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if !ok {
		return nil, false
	}
	namespaceMutex := value.(*NamespaceMutex)
	namespaceMutex.mu.RLock()
	defer namespaceMutex.mu.RUnlock()
	return namespaceMutex.memo, true
}

func (m *AllNamespaceMemoT) LoadOrStoreLockedMutex(key reflect.Type, memo *NamespaceMemo) (*NamespaceMutex, bool) {
	namespaceMutex := &NamespaceMutex{
		mu:   sync.RWMutex{},
		memo: memo,
	}
	// The mutex should always be inserted in a locked state.
	// Otherwise, there is an albeit slim possibility where
	// concurrent readers of the memo retrieve and obtain an
	// RLock on the mutex before the writer obtains a write
	// lock.
	namespaceMutex.mu.Lock()
	v, loaded := (*sync.Map)(m).LoadOrStore(key, namespaceMutex)
	if loaded {
		namespaceMutex.mu.Unlock()
	}
	return v.(*NamespaceMutex), loaded
}

func (m *AllNamespaceMemoT) Keys() []reflect.Type {
	keys := []reflect.Type{}
	(*sync.Map)(m).Range(func(key, _ any) bool {
		keys = append(keys, key.(reflect.Type))
		return true
	})
	return keys
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
	allMemo := AllNamespaceMemo
	if typ.Kind() == reflect.Ptr {
		return PopulateAllNamespaceMemo(typ.Elem(), visited)
	} else if typ.Kind() == reflect.Struct && typ != reflect.TypeOf(time.Time{}) {
		if visited[typ] {
			return nil
		}
		visited[typ] = true

		nsMutex, loaded := allMemo.LoadOrStoreLockedMutex(typ, NewNamespaceMemo())
		if loaded {
			nsMutex.mu.RLock()
			//lint:ignore SA2001 Empty is fine because this just waits for the memo of the same type to finish populating
			nsMutex.mu.RUnlock()

			// Prevent infinite loops and processing the same struct more than once.
			return nil
		}
		defer nsMutex.mu.Unlock()

		structName := typ.Name()
		namespace := ChalkpySnakeCase(structName)
		nsMemo := nsMutex.memo
		for fieldIdx := 0; fieldIdx < typ.NumField(); fieldIdx++ {
			fm := typ.Field(fieldIdx)
			resolvedName, err := ResolveFeatureName(fm)
			if err != nil {
				return errors.Wrapf(err, "error resolving feature name: %s", fm.Name)
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
					return errors.Wrapf(
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
					return err
				}
			}

			if fm.Type.Kind() == reflect.Ptr && IsStruct(fm.Type.Elem()) && !IsTypeDataclass(fm.Type.Elem()) {
				nsMemo.StructFieldsSet[resolvedName] = true
			}
		}
	} else if typ.Kind() == reflect.Slice {
		return PopulateAllNamespaceMemo(typ.Elem(), visited)
	}
	return nil
}
