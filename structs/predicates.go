package structs

import (
	. "github.com/jokruger/got/ifs"
)

// InRange returns a predicate that checks if a value is in a range (inclusive).
func InRange[T CompareProvider](min, max T) func(T) bool {
	return func(v T) bool {
		return v.Compare(min) >= 0 && v.Compare(max) <= 0
	}
}

/*
// GreaterThan returns a predicate that checks if a value is greater than a given value.
func GreaterThan[T Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x > v
	}
}

// LessThan returns a predicate that checks if a value is less than a given value.
func LessThan[T Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x < v
	}
}

// GreaterOrEqualTo returns a predicate that checks if a value is greater than or equal to a given value.
func GreaterOrEqualTo[T Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x >= v
	}
}

// LessOrEqualTo returns a predicate that checks if a value is less than or equal to a given value.
func LessOrEqualTo[T Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x <= v
	}
}

// EqualTo returns a predicate that checks if a value is equal to a given value.
func EqualTo[T comparable](v T) func(T) bool {
	return func(x T) bool {
		return x == v
	}
}

// Equal returns a predicate that checks if two values are equal.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Zero returns true if a value is the zero value of its type.
func IsZero[T comparable](v T) bool {
	var zero T
	return v == zero
}
*/
