package iter

import (
	"iter"
	"slices"
)

// Map applies a function to each element of a sequence and returns a new sequence.
func Map[T any, U any](is iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for i := range is {
			if !yield(f(i)) {
				return
			}
		}
	}
}

// MapToSlice applies a function to each element of a sequence and returns a new slice.
func MapToSlice[T any, U any](is iter.Seq[T], f func(T) U) []U {
	return slices.Collect(Map(is, f))
}

// MapUnique applies a function to each element of a sequence and returns a new sequence with unique elements only.
func MapUnique[T any, U comparable](is iter.Seq[T], f func(T) U) iter.Seq[U] {
	s := make(map[U]interface{})
	return func(yield func(U) bool) {
		for i := range is {
			u := f(i)
			if _, exists := s[u]; !exists {
				if !yield(u) {
					return
				}
				s[u] = nil
			}
		}
	}
}

// MapUniqueToSlice applies a function to each element of a sequence and returns a new slice with unique elements only.
func MapUniqueToSlice[T any, U comparable](is iter.Seq[T], f func(T) U) []U {
	var r []U
	s := make(map[U]interface{})
	for i := range is {
		u := f(i)
		if _, exists := s[u]; !exists {
			r = append(r, u)
			s[u] = nil
		}
	}
	return r
}
