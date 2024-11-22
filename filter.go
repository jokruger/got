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

// FilterSet returns a new slice containing only the elements of slice 'is' that are in set 's'.
func FilterSet[T comparable](is []T, s Set[T]) []T {
	return Filter(is, func(i T) bool {
		return s.Contains(i)
	})
}

// FilterSetIter returns a new slice containing only the elements of sequence 'is' that are in set 's'.
func FilterSetIter[T comparable](is iter.Seq[T], s Set[T]) []T {
	return FilterIter(is, func(i T) bool {
		return s.Contains(i)
	})
}