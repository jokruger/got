package got

import "iter"

// All returns true if all elements in s satisfy the predicate f.
// If s is empty or nil, All returns true.
func All[T any](s []T, f func(T) bool) bool {
	if len(s) > 0 {
		for _, e := range s {
			if !f(e) {
				return false
			}
		}
	}
	return true
}

// AllSeq returns true if all elements in s satisfy the predicate f.
// If s is nil, AllSeq returns true.
func AllSeq[T any](s iter.Seq[T], f func(T) bool) bool {
	if s != nil {
		for e := range s {
			if !f(e) {
				return false
			}
		}
	}
	return true
}
