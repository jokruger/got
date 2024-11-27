package generics

import (
	"iter"

	. "github.com/jokruger/got/ifs"
)

// Not returns a predicate that negates the result of the given predicate.
func Not[T any](f func(T) bool) func(T) bool {
	return func(v T) bool {
		return !f(v)
	}
}

// And returns a predicate that returns true if both predicates return true.
func And[T any](f, g func(T) bool) func(T) bool {
	return func(v T) bool {
		return f(v) && g(v)
	}
}

// Or returns a predicate that returns true if either predicate returns true.
func Or[T any](f, g func(T) bool) func(T) bool {
	return func(v T) bool {
		return f(v) || g(v)
	}
}

// InContainer returns a predicate that checks if a value is in container.
func InContainer[T any](s Container[T]) func(T) bool {
	return func(v T) bool {
		return s.Contains(v)
	}
}

// InSet returns a predicate that checks if a value is in a set.
func InSet[T comparable](s map[T]interface{}) func(T) bool {
	return func(v T) bool {
		_, ok := s[v]
		return ok
	}
}

// InMap returns a predicate that checks if a key is in a map.
func InMap[K comparable, V any](m map[K]V) func(K) bool {
	return func(k K) bool {
		_, ok := m[k]
		return ok
	}
}

// InSlice returns a predicate that checks if a value is in a slice.
func InSlice[T comparable](s []T) func(T) bool {
	return func(v T) bool {
		for _, e := range s {
			if e == v {
				return true
			}
		}
		return false
	}
}

// InSeq returns a predicate that checks if a value is in a sequence.
func InSeq[T comparable](s iter.Seq[T]) func(T) bool {
	return func(v T) bool {
		for i := range s {
			if i == v {
				return true
			}
		}
		return false
	}
}
