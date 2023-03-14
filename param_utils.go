package chalk

import (
	"fmt"
	"reflect"
)

func unwrapFeature(t any) *Feature {
	if reflect.ValueOf(t).Kind() == reflect.Ptr {
		v := reflect.New(reflect.TypeOf(t))
		v.Elem().Set(reflect.ValueOf(t))
		return (*Feature)(v.Elem().UnsafePointer())
	}

	typePointedTo := reflect.ValueOf(t).Elem().Type()
	panic(fmt.Sprintf("unsupported type: %s", typePointedTo))
}

func unwrapFeatureInterface(t any) *Feature {
	return unwrapFeature(t)

	//switch typed := t.(type) {
	//case *any:
	//	return unwrapFeature(typed)
	//case *int8:
	//	return unwrapFeature(typed)
	//case *int16:
	//	return unwrapFeature(typed)
	//case *int32:
	//	return unwrapFeature(typed)
	//case *int64:
	//	return unwrapFeature(typed)
	//case *uint:
	//	return unwrapFeature(typed)
	//case *uint8:
	//	return unwrapFeature(typed)
	//case *uint16:
	//	return unwrapFeature(typed)
	//case *uint32:
	//	return unwrapFeature(typed)
	//case *uint64:
	//	return unwrapFeature(typed)
	//case *string:
	//	return unwrapFeature(typed)
	//case *float32:
	//	return unwrapFeature(typed)
	//case *float64:
	//	return unwrapFeature(typed)
	//case *bool:
	//	return unwrapFeature(typed)
	//case *time.Time:
	//	return unwrapFeature(typed)
	//default:
	// 	panic(fmt.Sprintf("unsupported type: %s", reflect.TypeOf(t)))
	//}
}
