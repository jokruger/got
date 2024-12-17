package basicutil

import (
	"github.com/jokruger/got"
)

// Range returns a slice of integers from 0 to n-1.
func Range[T got.Integer](n T) []T {
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
func Compare[T got.Ordered](a, b T) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// CompareTo returns a function that compares a value to another value.
func CompareTo[T got.Ordered](b T) func(T) int {
	return func(a T) int {
		return Compare(a, b)
	}
}

// CompareBy returns a function that compares two values by a function which returns an Ordered value.
func CompareBy[T any, V got.Ordered](f func(T) V) func(T, T) int {
	return func(a, b T) int {
		va := f(a)
		vb := f(b)
		if va < vb {
			return -1
		}
		if va > vb {
			return 1
		}
		return 0
	}
}

// CompareBy2 returns a function that compares two values by two functions which return Ordered values.
func CompareBy2[T any, V got.Ordered, U got.Ordered](f func(T) V, g func(T) U) func(T, T) int {
	return func(a, b T) int {
		va := f(a)
		vb := f(b)
		if va < vb {
			return -1
		}
		if va > vb {
			return 1
		}
		ua := g(a)
		ub := g(b)
		if ua < ub {
			return -1
		}
		if ua > ub {
			return 1
		}
		return 0
	}
}
