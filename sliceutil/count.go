package sliceutil

// Count returns the number of elements in the slice that satisfy the predicate f.
func Count[T any](s []T, f func(T) bool) int {
	c := 0
	if len(s) > 0 {
		for _, e := range s {
			if f(e) {
				c++
			}
		}
	}
	return c
}
