package internal

import (
	"github.com/apache/arrow/go/v16/arrow"
	"github.com/cockroachdb/errors"
	"reflect"
	"sync"
)

type Codec func(structValue reflect.Value, arr arrow.Array, arrIdx int) error

var codecNoOp = func(structValue reflect.Value, arr arrow.Array, arrIdx int) error {
	return nil
}

type InitFeatureFunc func(structValue reflect.Value) (leafStructValue reflect.Value, err error)

var initFeatureNoOp InitFeatureFunc = func(structValue reflect.Value) (reflect.Value, error) { return reflect.Value{}, nil }

type GetValueFunc func(arr arrow.Array, arrIdx int) (reflect.Value, error)

type FqnMemo struct {
	Codec            Codec
	once             sync.Once
	generateCodecErr error
}

type AllFqnMemoT sync.Map

var AllFqnMemo = &AllFqnMemoT{}

func (m *AllFqnMemoT) LoadCodec(key string, generateCodec func() (Codec, error)) (Codec, error) {
	value, loaded := (*sync.Map)(m).Load(key)
	if !loaded {
		value, loaded = (*sync.Map)(m).LoadOrStore(key, &FqnMemo{})
	}
	fqnMemo := value.(*FqnMemo)
	fqnMemo.once.Do(func() {
		fqnMemo.Codec, fqnMemo.generateCodecErr = generateCodec()
	})
	if fqnMemo.generateCodecErr != nil {
		return nil, errors.Wrap(fqnMemo.generateCodecErr, "generating codec")
	}
	return fqnMemo.Codec, nil
}

func (m *AllFqnMemoT) Keys() []string {
	var keys []string
	(*sync.Map)(m).Range(func(key, _ any) bool {
		keys = append(keys, key.(string))
		return true
	})
	return keys
}
