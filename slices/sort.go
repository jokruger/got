package slices

import (
	"slices"
)

// Sort sorts a slice using a compare function.
func Sort[T any](s []T, compare func(T, T) int) {
	slices.SortFunc(s, compare)
}

// Sorted returns a sorted copy of a slice using a compare function.
func Sorted[T any](s []T, compare func(T, T) int) []T {
	res := make([]T, len(s))
	copy(res, s)
	slices.SortFunc(res, compare)
	return res
}
