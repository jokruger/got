package sliceutil

import "cmp"

// MinMax calculates the minimum and maximum values of a slice based on a given function.
func MinMax[T any, K cmp.Ordered](s []T, f func(T) K) (K, K, bool) {
	if len(s) == 0 {
		var zero K
		return zero, zero, false
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

	return min, max, true
}
