package chalk

import (
	"reflect"
)

func ConvertFeatures(t reflect.Value, path string, visited map[string]bool) {
	if visited == nil {
		visited = make(map[string]bool)
	}

	var s reflect.Value
	if t.Kind() == reflect.Struct {
		s = t
	} else {
		s = t.Elem()
	}

	namespace := s.Type().Name()
	if namespaceVisited, ok := visited[namespace]; ok && namespaceVisited {
		return
	} else {
		visited[namespace] = true
	}

	if s.Kind() == reflect.Struct {
		if path == "" {
			path = SnakeCase(namespace)
		}
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			attributeName := s.Type().Field(i).Name
			var newPath string
			if path == "" {
				newPath = SnakeCase(attributeName)
			} else {
				newPath = path + "." + SnakeCase(attributeName)
			}

			if f.CanSet() {
				if f.Kind() == reflect.Pointer {
					//if sliceContains([]string{"str", "int", "bool", "float64", "time.Time"}, f.Type().Elem().String()) {
					if f.Type().Elem().Kind() == reflect.Struct {
						// Should be has-ones
						newStructObj := reflect.New(f.Type().Elem())
						fakePointerToOriginalType := reflect.NewAt(f.Type().Elem(), reflect.ValueOf(&newStructObj).UnsafePointer())
						f.Set(fakePointerToOriginalType)
						ConvertFeatures(f.Elem(), newPath, visited)
					} else {
						feature := Feature{Fqn: newPath}
						fakePointerToOriginalType := reflect.NewAt(f.Type().Elem(), reflect.ValueOf(&feature).UnsafePointer())
						f.Set(fakePointerToOriginalType)
					}

					// OLD METHOD
					//fCopy := reflect.New(reflect.TypeOf(&feature))
					//fCopy.Elem().Set(reflect.ValueOf(&feature))
					//fakePointerToOriginalType := reflect.NewAt(f.Type().Elem(), fCopy.Elem().UnsafePointer())

				}
			}
		}
	}

	visited[namespace] = false
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

func sliceContains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
