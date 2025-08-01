package sliceutil

import "cmp"

// MinMax finds minimum and maximum elements in a slice.
func MinMax[T any](s []T, compare func(T, T) int) (T, T) {
	if len(s) == 0 {
		var zero T
		return zero, zero
	}

	min := s[0]
	max := min

	for _, v := range s[1:] {
		if compare(v, min) < 0 {
			min = v
		}
		if compare(v, max) > 0 {
			max = v
		}
	}

	return min, max
}

// MapMinMax applies a function to each element of the slice and finds the minimum and maximum values based on the results of that function.
func MapMinMax[T any, K cmp.Ordered](s []T, f func(T) K) (K, K) {
	if len(s) == 0 {
		var zero K
		return zero, zero
	}

	min := f(s[0])
	max := min

	for _, v := range s[1:] {
		val := f(v)
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}

	return min, max
}
