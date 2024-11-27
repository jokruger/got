package iter

import "iter"

// Filter returns a new sequence containing only the elements of sequence 's' that satisfy the predicate f.
func Filter[T any](s iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range s {
			if f(i) {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// FilterToSlice returns a new slice containing only the elements of sequence 's' that satisfy the predicate f.
func FilterToSlice[T any](s iter.Seq[T], f func(T) bool) []T {
	var res []T
	for i := range s {
		if f(i) {
			res = append(res, i)
		}
	}
	return res
}
