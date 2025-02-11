package sliceutil

// GroupBy groups elements of slice by key K calculated from the slice elements.
func GroupBy[T any, K comparable](slice []T, key func(T) K) map[K][]T {
	m := make(map[K][]T)
	for _, e := range slice {
		k := key(e)
		m[k] = append(m[k], e)
	}
	return m
}

// GroupByKV groups values V by key K where V and K are calculated from the slice elements.
func GroupByKV[T any, K comparable, V any](slice []T, kv func(T) (K, V)) map[K][]V {
	m := make(map[K][]V)
	for _, e := range slice {
		k, v := kv(e)
		m[k] = append(m[k], v)
	}
	return m
}
