package got

import "iter"

// Any returns true if any element in the slice satisfies the predicate, false otherwise.
func Any[T any](s []T, f func(T) bool) bool {
	for _, e := range s {
		if f(e) {
			return true
		}
	}
	return false
}

// AnySeq returns true if any element in the sequence satisfies the predicate, false otherwise.
func AnySeq[T any](s iter.Seq[T], f func(T) bool) bool {
	for e := range s {
		if f(e) {
			return true
		}
	}
	return false
}
