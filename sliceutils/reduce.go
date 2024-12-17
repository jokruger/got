package sliceutils

// Reduce applies a function against an accumulator and each element in the slice (from left to right) to reduce it to a single value.
func Reduce[T any, A any](s []T, a A, f func(T, A) A) A {
	if len(s) > 0 {
		for _, e := range s {
			a = f(e, a)
		}
	}
	return a
}
