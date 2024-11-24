package got

import (
	"iter"
	"slices"

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

// MapSeqToSeq applies a function to each element of a sequence and returns a new sequence.
func MapSeqToSeq[T any, U any](is iter.Seq[T], f func(T) U) iter.Seq[U] {
	return func(yield func(U) bool) {
		for i := range is {
			if !yield(f(i)) {
				return
			}
		}
	}
}

// MapSeq applies a function to each element of a sequence and returns a new slice.
func MapSeq[T any, U any](is iter.Seq[T], f func(T) U) []U {
	return slices.Collect(MapSeqToSeq(is, f))
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

// MapSeqUniqueToSeq applies a function to each element of a sequence and returns a new sequence with unique elements.
func MapSeqUniqueToSeq[T any, U comparable](is iter.Seq[T], f func(T) U) iter.Seq[U] {
	s := set.New[U]()
	return func(yield func(U) bool) {
		for i := range is {
			u := f(i)
			if s.Add(u) {
				if !yield(u) {
					return
				}
			}
		}
	}
}

// MapSeqUnique applies a function to each element of a sequence and returns a new slice with unique elements.
func MapSeqUnique[T any, U comparable](is iter.Seq[T], f func(T) U) []U {
	return slices.Collect(MapSeqUniqueToSeq(is, f))
}
