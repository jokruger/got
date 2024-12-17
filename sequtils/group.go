package sequtils

import "iter"

// GroupBy groups the elements of the sequence s by the key returned by the function kv.
func GroupBy[T any, K comparable, V any](s iter.Seq[T], kv func(T) (K, V)) map[K][]V {
	m := make(map[K][]V)
	for e := range s {
		k, v := kv(e)
		m[k] = append(m[k], v)
	}
	return m
}
