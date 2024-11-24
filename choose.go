package got

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
