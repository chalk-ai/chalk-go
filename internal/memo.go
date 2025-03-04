package internal

import (
	"github.com/apache/arrow/go/v16/arrow"
	"reflect"
	"sync"
)

type Codec func(structValue reflect.Value, arr arrow.Array, arrIdx int) error

var codecNoOp = func(structValue reflect.Value, arr arrow.Array, arrIdx int) error {
	return nil
}

type Memo[V any] struct {
	Object V

	once sync.Once
	err  error
}

type MemosT[K, V any] sync.Map

func (m *MemosT[K, V]) LoadOrStore(key K, generateObject func() (V, error)) (V, error) {
	value, loaded := (*sync.Map)(m).Load(key)
	if !loaded {
		value, loaded = (*sync.Map)(m).LoadOrStore(key, &Memo[V]{})
	}
	memo := value.(*Memo[V])
	memo.once.Do(func() {
		memo.Object, memo.err = generateObject()
	})
	return memo.Object, memo.err
}

var CodecMemo = &MemosT[string, Codec]{}

type NamespaceMemo struct {
	// Root and non-root FQN as keys
	ResolvedFieldNameToIndices map[string][]int
	// Non-root FQN as keys only
	StructFieldsSet map[string]bool
}

type NamespaceMemosT MemosT[reflect.Type, NamespaceMemo]

var AllNamespaceMemo = &NamespaceMemosT{}

func (m *NamespaceMemosT) Load(structType reflect.Type) (NamespaceMemo, error) {
	return (*MemosT[reflect.Type, NamespaceMemo])(m).LoadOrStore(structType, func() (NamespaceMemo, error) {
		return generateNamespaceMemo(structType, nil)
	})
}
