package got

import "iter"

func GroupBy[T any, K comparable, V any](slice []T, kv func(T) (K, V)) map[K][]V {
	m := make(map[K][]V)
	for _, e := range slice {
		k, v := kv(e)
		m[k] = append(m[k], v)
	}
	return m
}

func GroupSeqBy[T any, K comparable, V any](s iter.Seq[T], kv func(T) (K, V)) map[K][]V {
	m := make(map[K][]V)
	for e := range s {
		k, v := kv(e)
		m[k] = append(m[k], v)
	}
	return m
}
