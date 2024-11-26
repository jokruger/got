package got

// Ternary returns a if f() is true, otherwise b.
func Ternary[T any](a T, b T, f func() bool) T {
	if f() {
		return a
	}
	return b
}
