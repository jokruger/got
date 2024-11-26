package got

import "iter"

// Any returns true if any element in the slice satisfies the predicate, false otherwise.
// If the slice is empty or nil, Any returns false.
func Any[T any](s []T, f func(T) bool) bool {
	if len(s) > 0 {
		for _, e := range s {
			if f(e) {
				return true
			}
		}
	}
	return false
}

// AnySeq returns true if any element in the sequence satisfies the predicate, false otherwise.
// If the sequence is nil, AnySeq returns false.
func AnySeq[T any](s iter.Seq[T], f func(T) bool) bool {
	if s != nil {
		for e := range s {
			if f(e) {
				return true
			}
		}
	}
	return false
}
