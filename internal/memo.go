package internal

import (
	"sync"
)

type Memo[V any] struct {
	Object *V

	once sync.Once
	err  error
}

type MemosT[K, V any] sync.Map

func (m *MemosT[K, V]) LoadOrStore(key K, generateObject func() (*V, error)) (*V, error) {
	value, loaded := (*sync.Map)(m).Load(key)
	if !loaded {
		value, _ = (*sync.Map)(m).LoadOrStore(key, &Memo[V]{})
	}
	memo := value.(*Memo[V])
	memo.once.Do(func() {
		memo.Object, memo.err = generateObject()
	})
	return memo.Object, memo.err
}
