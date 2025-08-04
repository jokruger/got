package sliceutil

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

// FindPtr returns the pointer to the first element in the slice that satisfies the predicate f.
// If no element satisfies the predicate, FindPtr returns nil.
func FindPtr[T any](s []T, f func(*T) bool) *T {
	if len(s) > 0 {
		for i := range s {
			p := &s[i]
			if f(p) {
				return p
			}
		}
	}
	return nil
}
