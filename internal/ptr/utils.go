package ptr

import "reflect"

// OrNil returns a pointer copy of value if it's nonzero.
// Otherwise, returns nil pointer.
func OrNil[T any](x T) *T {
	isZero := reflect.ValueOf(&x).Elem().IsZero()
	if isZero {
		return nil
	}

	return &x
}

func OrZero[T any](t *T) T {
	if t != nil {
		return *t
	}
	var result T
	return result
}
