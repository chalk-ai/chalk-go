package internal

import (
	"github.com/apache/arrow/go/v16/arrow"
	"os"
	"reflect"
	"strconv"
	"sync"
)

type GetValueFunc func(arr arrow.Array, arrIdx int) (reflect.Value, error)

type Codec func(structValue reflect.Value, arr arrow.Array, arrIdx int) error

var codecNoOp = func(structValue reflect.Value, arr arrow.Array, arrIdx int) error {
	return nil
}

type FqnMemo struct {
	Codec Codec
}

type AllFqnMemoT sync.Map

var AllFqnMemo = &AllFqnMemoT{}

func init() {
	if chunkSizeStr := os.Getenv(tableReaderChunkSizeKey); chunkSizeStr != "" {
		if newChunkSize, err := strconv.Atoi(chunkSizeStr); err == nil {
			TableReaderChunkSize = newChunkSize
		}
	}
}

func NewFqnMemo() *FqnMemo {
	return &FqnMemo{}
}

type FqnMutex struct {
	mu   sync.RWMutex
	memo *FqnMemo
}

func (m *AllFqnMemoT) Load(key string) (*FqnMemo, bool) {
	value, ok := (*sync.Map)(m).Load(key)
	if !ok {
		return nil, false
	}
	fqnMutex := value.(*FqnMutex)
	fqnMutex.mu.RLock()
	defer fqnMutex.mu.RUnlock()
	return fqnMutex.memo, true
}

func (m *AllFqnMemoT) LoadOrStoreLockedMutex(key string, memo *FqnMemo) (*FqnMutex, bool) {
	newMutex := &FqnMutex{
		mu:   sync.RWMutex{},
		memo: memo,
	}
	// The mutex should always be inserted in a locked state.
	// Otherwise, there is an albeit slim possibility where
	// concurrent readers of the memo retrieve and obtain an
	// RLock on the mutex before the writer obtains a write
	// lock.
	newMutex.mu.Lock()
	v, loaded := (*sync.Map)(m).LoadOrStore(key, newMutex)
	if loaded {
		newMutex.mu.Unlock()

		existingMutex := v.(*FqnMutex)
		existingMutex.mu.RLock()
		defer existingMutex.mu.RUnlock()
	}
	return v.(*FqnMutex), loaded
}

func (m *AllFqnMemoT) Keys() []string {
	var keys []string
	(*sync.Map)(m).Range(func(key, _ any) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}
