package sliceutil

import (
	"iter"

	"github.com/jokruger/got/set"
)

// Map applies a function to each element of a slice and returns a new slice.
func Map[T any, U any](is []T, f func(T) U) []U {
	us := make([]U, len(is))
	for i, v := range is {
		us[i] = f(v)
	}
	return us
}

// MapToSeq applies a function to each element of a slice and returns a new sequence.
func MapToSeq[T any, U any](is []T, f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for _, v := range is {
			if !yield(f(v)) {
				return
			}
		}
	}
}

// MapUnique applies a function to each element of a slice and returns a new slice with unique elements.
func MapUnique[T any, U comparable](is []T, f func(T) U) []U {
	s := set.New[U]()
	r := make([]U, 0, len(is))
	for _, v := range is {
		u := f(v)
		if s.Add(u) {
			r = append(r, u)
		}
	}
	return r
}

// MapUniqueToSeq applies a function to each element of a slice and returns a new sequence with unique elements.
func MapUniqueToSeq[T any, U comparable](is []T, f func(T) U) iter.Seq[U] {
	s := set.New[U]()
	return func(yield func(U) bool) {
		for _, v := range is {
			u := f(v)
			if s.Add(u) {
				if !yield(u) {
					return
				}
			}
		}
	}
}
