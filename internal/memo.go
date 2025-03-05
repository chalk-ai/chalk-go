package internal

import (
	"sync"
)

type Memo[V any] struct {
	object *V
	err    error
	once   sync.Once
}

type MemosT[K, V any] sync.Map

func (m *MemosT[K, V]) LoadOrStore(key K, generateObject func() (*V, error)) (*V, error) {
	value, loaded := (*sync.Map)(m).Load(key)
	if !loaded {
		value, _ = (*sync.Map)(m).LoadOrStore(key, &Memo[V]{})
	}
	memo := value.(*Memo[V])
	memo.once.Do(func() {
		memo.object, memo.err = generateObject()
	})
	return memo.object, memo.err
}
