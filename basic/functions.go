package basic

import (
	. "github.com/jokruger/got/ifs"
)

// Range returns a slice of integers from 0 to n-1.
func Range[T Integer](n T) []T {
	if n <= 0 {
		return nil
	}
	is := make([]T, n)
	for i := T(0); i < n; i++ {
		is[i] = i
	}
	return is
}

// Choose returns the first non-zero value from the list of values.
func Choose[T comparable](values ...T) T {
	var zero T
	for _, v := range values {
		if v != zero {
			return v
		}
	}
	return zero
}

// Compare returns -1 if a < b, 0 if a == b, and 1 if a > b.
func Compare[T Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// CompareTo returns a function that compares a value to another value.
func CompareTo[T Ordered](b T) func(T) int {
	return func(a T) int {
		return Compare(a, b)
	}
}
