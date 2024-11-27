package structs

import (
	. "github.com/jokruger/got/ifs"
)

// Choose returns the first non-zero value from the list of values.
func Choose[T ZeroTestProvider](values ...T) T {
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
