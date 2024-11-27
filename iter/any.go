package iter

import "iter"

// Any returns true if any element in the sequence satisfies the predicate, false otherwise.
// If the sequence is nil or empty, Any returns false.
func Any[T any](s iter.Seq[T], f func(T) bool) bool {
	if s != nil {
		for e := range s {
			if f(e) {
				return true
			}
		}
	}
	return false
}
