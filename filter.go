package got

import "iter"

// Filter returns a new slice containing only the elements of slice 'is' that satisfy the predicate f.
func Filter[T any](is []T, f func(T) bool) []T {
	var js []T
	for _, v := range is {
		if f(v) {
			js = append(js, v)
		}
	}
	return js
}

// FilterIter returns a new slice containing only the elements of sequence 'is' that satisfy the predicate f.
func FilterIter[T any](is iter.Seq[T], f func(T) bool) []T {
	var js []T
	for i := range is {
		if f(i) {
			js = append(js, i)
		}
	}
	return js
}
