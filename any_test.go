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
		s = []int{}
		if Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = nil
		if Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})

	t.Run("AnySeq ints", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		if !AnySeq(s, GreaterThan(1)) {
			t.Errorf("AnySeq(%v, f) = false; want true", s)
		}
		if AnySeq(s, GreaterThan(3)) {
			t.Errorf("AnySeq(%v, f) = true; want false", s)
		}
		s = slices.Values([]int{})
		if AnySeq(s, GreaterThan(1)) {
			t.Errorf("AnySeq(%v, f) = true; want false", s)
		}
		s = nil
		if AnySeq(s, GreaterThan(1)) {
			t.Errorf("AnySeq(%v, f) = true; want false", s)
		}
	})
}
