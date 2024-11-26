package got

import "iter"

// Find returns the index and the element of the first element in the slice that satisfies the predicate f.
// If no element satisfies the predicate, Find returns -1 and the zero value of the element type.
func Find[T any](s []T, f func(T) bool) (int, T) {
	if len(s) > 0 {
		for i, e := range s {
			if f(e) {
				return i, e
			}
		}
	}
	var zero T
	return -1, zero
}

func FindSeq[T any](s iter.Seq[T], f func(T) bool) (int, T) {
	if s != nil {
		i := 0
		for e := range s {
			if f(e) {
				return i, e
			}
			i++
		}
	}
	var zero T
	return -1, zero
}
