package basicutils

import (
	"regexp"
	"strings"

	"github.com/jokruger/got"
)

// InRange returns a predicate that checks if a value is in a range (inclusive).
func InRange[T got.Ordered](min, max T) func(T) bool {
	return func(v T) bool {
		return v >= min && v <= max
	}
}

// GreaterThan returns a predicate that checks if a value is greater than a given value.
func GreaterThan[T got.Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x > v
	}
}

// LessThan returns a predicate that checks if a value is less than a given value.
func LessThan[T got.Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x < v
	}
}

// GreaterOrEqualTo returns a predicate that checks if a value is greater than or equal to a given value.
func GreaterOrEqualTo[T got.Ordered](v T) func(T) bool {
	return func(x T) bool {
		return x >= v
	}
}

// LessOrEqualTo returns a predicate that checks if a value is less than or equal to a given value.
func LessOrEqualTo[T got.Ordered](v T) func(T) bool {
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
