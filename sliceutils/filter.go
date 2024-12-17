package sliceutils

import "iter"

// Filter returns a new slice containing only the elements of slice 's' that satisfy the predicate f.
func Filter[T any](s []T, f func(T) bool) []T {
	res := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// FilterInPlace removes all elements from slice 's' that do not satisfy the predicate f and returns the modified slice.
func FilterInPlace[T any](s []T, f func(T) bool) []T {
	i := 0
	for _, v := range s {
		if f(v) {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

// FilterToSeq returns a new sequence containing only the elements of slice 's' that satisfy the predicate f.
func FilterToSeq[T any](s []T, f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
