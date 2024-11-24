package got

import (
	"iter"
	"regexp"
	"strings"

	"github.com/jokruger/got/set"
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

// InSet returns a predicate that checks if a value is in a set.
func InSet[T comparable](s set.Set[T]) func(T) bool {
	return func(v T) bool {
		return s.Contains(v)
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

// InRange returns a predicate that checks if a value is in a range (inclusive).
func InRange[T Ordered](min, max T) func(T) bool {
	return func(v T) bool {
		return v >= min && v <= max
	}
}

// EqualTo returns a predicate that checks if a value is equal to a given value.
func EqualTo[T comparable](v T) func(T) bool {
	return func(x T) bool {
		return x == v
	}
}

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

// InString returns a predicate that checks if a substring is in a string.
func InString(s string) func(string) bool {
	return func(sub string) bool {
		return strings.Contains(s, sub)
	}
}

// ContainsSubstring returns a predicate that checks if a string contains a substring.
func ContainsSubstring(sub string) func(string) bool {
	return func(s string) bool {
		return strings.Contains(s, sub)
	}
}

// StartsWithString returns a predicate that checks if a string starts with a substring.
func StartsWithString(prefix string) func(string) bool {
	return func(s string) bool {
		return strings.HasPrefix(s, prefix)
	}
}

// EndsWithString returns a predicate that checks if a string ends with a substring.
func EndsWithString(suffix string) func(string) bool {
	return func(s string) bool {
		return strings.HasSuffix(s, suffix)
	}
}

// MatchesRegexp returns a predicate that checks if a string matches a regular expression.
func MatchesRegexp(re string) func(string) bool {
	return func(s string) bool {
		m, err := regexp.MatchString(re, s)
		if err != nil {
			panic(err)
		}
		return m
	}
}

// MatchesRegexpCompiled returns a predicate that checks if a string matches a compiled regular expression.
func MatchesRegexpCompiled(re *regexp.Regexp) func(string) bool {
	return func(s string) bool {
		return re.MatchString(s)
	}
}

// Equal returns a predicate that checks if two values are equal.
func Equal[T comparable](a, b T) bool {
	return a == b
}

// Zero returns true if a value is the zero value of its type.
func Zero[T comparable](v T) bool {
	var zero T
	return v == zero
}
