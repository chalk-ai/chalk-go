package colls

func Keys[T comparable, U any](m map[T]U) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values creates an array of the map values.
func Values[K comparable, V any](in ...map[K]V) []V {
	size := 0
	for i := range in {
		size += len(in[i])
	}
	result := make([]V, 0, size)

	for i := range in {
		for k := range in[i] {
			result = append(result, in[i][k])
		}
	}

	return result
}
