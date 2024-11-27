package iter

import "iter"

// Count returns the number of elements in the sequence that satisfy the predicate f.
func Count[T any](s iter.Seq[T], f func(T) bool) int {
	c := 0
	if s != nil {
		for e := range s {
			if f(e) {
				c++
			}
		}
	}
	return c
}
