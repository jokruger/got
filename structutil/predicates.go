package structutil

import (
	"github.com/jokruger/got"
)

// InRange returns a predicate that checks if a value is in a range (inclusive).
func InRange[T any, V got.Comparable[T]](min, max T) func(V) bool {
	return func(v V) bool {
		return v.Compare(min) >= 0 && v.Compare(max) <= 0
	}
}

// GreaterThan returns a predicate that checks if a value is greater than a given value.
func GreaterThan[T any, V got.Comparable[T]](v T) func(V) bool {
	return func(x V) bool {
		return x.Compare(v) > 0
	}
}

// LessThan returns a predicate that checks if a value is less than a given value.
func LessThan[T any, V got.Comparable[T]](v T) func(V) bool {
	return func(x V) bool {
		return x.Compare(v) < 0
	}
}

// GreaterOrEqualTo returns a predicate that checks if a value is greater than or equal to a given value.
func GreaterOrEqualTo[T any, V got.Comparable[T]](v T) func(V) bool {
	return func(x V) bool {
		return x.Compare(v) >= 0
	}
}

// LessOrEqualTo returns a predicate that checks if a value is less than or equal to a given value.
func LessOrEqualTo[T any, V got.Comparable[T]](v T) func(V) bool {
	return func(x V) bool {
		return x.Compare(v) <= 0
	}
}

// EqualTo returns a predicate that checks if a value is equal to a given value.
func EqualTo[T any, V got.Comparable[T]](v T) func(V) bool {
	return func(x V) bool {
		return x.Compare(v) == 0
	}
}

// Equal checks if two values are equal.
func Equal[T any, V got.Comparable[T]](a V, b T) bool {
	return a.Compare(b) == 0
}

// IsZero returns result of zero testing.
func IsZero[T got.ZeroCheckable](v T) bool {
	return v.IsZero()
}
