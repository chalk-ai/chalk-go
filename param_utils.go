package chalk

import (
	"time"
	"unsafe"
)

func unwrapFeature[T any](t *T) *Feature {
	return (*Feature)(unsafe.Pointer(t))
}

func unwrapFeatureInterface(t any) *Feature {
	switch typed := t.(type) {
	case *int:
		return unwrapFeature(typed)
	case *int8:
		return unwrapFeature(typed)
	case *int16:
		return unwrapFeature(typed)
	case *int32:
		return unwrapFeature(typed)
	case *int64:
		return unwrapFeature(typed)
	case *uint:
		return unwrapFeature(typed)
	case *uint8:
		return unwrapFeature(typed)
	case *uint16:
		return unwrapFeature(typed)
	case *uint32:
		return unwrapFeature(typed)
	case *uint64:
		return unwrapFeature(typed)
	case *string:
		return unwrapFeature(typed)
	case *float32:
		return unwrapFeature(typed)
	case *float64:
		return unwrapFeature(typed)
	case *bool:
		return unwrapFeature(typed)
	case *time.Time:
		return unwrapFeature(typed)
	default:
		panic("unsupported type")
	}
}
