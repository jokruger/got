package sequtil

import "iter"

// All returns true if all elements in s satisfy the predicate f.
// If s is nil or empty, All returns true.
func All[T any](s iter.Seq[T], f func(T) bool) bool {
	if s != nil {
		for e := range s {
			if !f(e) {
				return false
			}
		}
	}
	return true
}
