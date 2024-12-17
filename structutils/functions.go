package structutils

import (
	"github.com/jokruger/got"
)

// Choose returns the first non-zero value from the list of values.
func Choose[T got.ZeroCheckable](values ...T) T {
	for _, v := range values {
		if !v.IsZero() {
			return v
		}
	}
	if len(values) == 0 {
		var t T
		return t
	}
	return values[len(values)-1]
}

// Compare returns the result of comparing two values.
func Compare[T any, V got.Comparable[T]](a V, b T) int {
	return a.Compare(b)
}

// CompareTo returns a function that compares a value to another value.
func CompareTo[T any, V got.Comparable[T]](b T) func(V) int {
	return func(a V) int {
		return a.Compare(b)
	}
}

// CompareBy returns a function that compares two values by a given function.
// The function must return a value that implements the CompareProvider interface.
func CompareBy[T any, V got.Comparable[V]](f func(T) V) func(T, T) int {
	return func(a, b T) int {
		return f(a).Compare(f(b))
	}
}
