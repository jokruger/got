package got

import (
	"slices"
	"testing"
)

func TestAny(t *testing.T) {
	t.Run("Any ints", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = false; want true", s)
		}
		if Any(s, GreaterThan(3)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})

	t.Run("AnySeq ints", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !AnySeq(slices.Values(s), GreaterThan(1)) {
			t.Errorf("Any(%v, f) = false; want true", s)
		}
		if AnySeq(slices.Values(s), GreaterThan(3)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})
}
