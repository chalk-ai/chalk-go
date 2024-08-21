package ptr

import "reflect"

// PtrOrNil returns a pointer copy of value if it's nonzero.
// Otherwise, returns nil pointer.
func PtrOrNil[T any](x T) *T {
	isZero := reflect.ValueOf(&x).Elem().IsZero()
	if isZero {
		return nil
	}

	return &x
}

func Ptr[T any](value T) *T {
	return &value
}
