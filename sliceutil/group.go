package sliceutil

// GroupBy groups the elements of the slice s by the key returned by the function kv.
func GroupBy[T any, K comparable, V any](slice []T, kv func(T) (K, V)) map[K][]V {
	m := make(map[K][]V)
	for _, e := range slice {
		k, v := kv(e)
		m[k] = append(m[k], v)
	}
	return m
}
