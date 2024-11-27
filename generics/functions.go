package generics

// Ternary returns a if f() is true, otherwise b.
func Ternary[T any](a T, b T, f func() bool) T {
	if f() {
		return a
	}
	return b
}

// Flip returns a function that swaps the arguments of a binary function.
func Flip[T, R any](f func(T, T) R) func(T, T) R {
	return func(a, b T) R {
		return f(b, a)
	}
}
