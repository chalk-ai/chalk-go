package chalk

import (
	"reflect"
)

func ConvertFeatures[T any](t *T, path string) {
	s := reflect.ValueOf(t).Elem()
	if s.Kind() == reflect.Struct {
		namespace := s.Type().Name()
		if path == "" {
			path = namespace
		}
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			attributeName := s.Type().Field(i).Name
			newPath := path + "." + attributeName
			newPathSnake := SnakeCase(newPath)
			if f.CanSet() {
				if f.Kind() == reflect.Pointer {
					feature := Feature{Fqn: newPathSnake}
					fCopy := reflect.New(reflect.TypeOf(&feature))
					fCopy.Elem().Set(reflect.ValueOf(&feature))
					fakePointerToOriginalType := reflect.NewAt(f.Type().Elem(), fCopy.Elem().UnsafePointer())
					f.Set(fakePointerToOriginalType)
				}
			}
		}
	}
}

func SnakeCase(s string) string {
	var b []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isASCIIUpper(c) {
			if i > 0 && s[i-1] != '.' {
				b = append(b, '_')
			}
			c += 'a' - 'A'
		}
		b = append(b, c)
	}
	return string(b)
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}
