package internal

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cockroachdb/errors"
	"reflect"
	"sync"
)

type GetValueFunc func(arr arrow.Array, arrIdx int) (reflect.Value, error)

type Codec func(structValue reflect.Value, arr arrow.Array, arrIdx int) error

var codecNoOp = func(structValue reflect.Value, arr arrow.Array, arrIdx int) error {
	return nil
}

type FqnMemo struct {
	Codec            Codec
	once             sync.Once
	generateCodecErr error
}

type AllFqnMemoT sync.Map

var AllFqnMemo = &AllFqnMemoT{}

func (m *AllFqnMemoT) LoadMemo(key string, generateCodec func() (Codec, error)) (*FqnMemo, error) {
	value, loaded := (*sync.Map)(m).Load(key)
	if !loaded {
		value, loaded = (*sync.Map)(m).LoadOrStore(key, &FqnMemo{})
	}
	fqnMemo := value.(*FqnMemo)
	fqnMemo.once.Do(func() {
		codec, err := generateCodec()
		if err != nil {
			fqnMemo.generateCodecErr = err
		} else {
			fqnMemo.Codec = codec
		}
	})
	if fqnMemo.generateCodecErr != nil {
		return nil, errors.Wrap(fqnMemo.generateCodecErr, "generating codec")
	}
	return fqnMemo, nil
}

func (m *AllFqnMemoT) Keys() []string {
	var keys []string
	(*sync.Map)(m).Range(func(key, _ any) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}
