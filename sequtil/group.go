package sequtil

import "iter"

// GroupBy groups elements of sequence by key K calculated from the sequence elements.
func GroupBy[T any, K comparable](s iter.Seq[T], key func(T) K) map[K][]T {
	m := make(map[K][]T)
	for e := range s {
		k := key(e)
		m[k] = append(m[k], e)
	}
	return m
}

// GroupByKV groups values V by key K where V and K are calculated from the sequence elements.
func GroupByKV[T any, K comparable, V any](s iter.Seq[T], kv func(T) (K, V)) map[K][]V {
	m := make(map[K][]V)
	for e := range s {
		k, v := kv(e)
		m[k] = append(m[k], v)
	}
	return m
}
