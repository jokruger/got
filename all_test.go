package got

import (
	"slices"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("All ints", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !All(s, GreaterThan(0)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		if All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = true; want false", s)
		}
		s = []int{}
		if !All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		s = nil
		if !All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
	})

	t.Run("AllSeq ints", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		if !AllSeq(s, GreaterThan(0)) {
			t.Errorf("AllSeq(%v, f) = false; want true", s)
		}
		if AllSeq(s, GreaterThan(1)) {
			t.Errorf("AllSeq(%v, f) = true; want false", s)
		}
		s = slices.Values([]int{})
		if !AllSeq(s, GreaterThan(1)) {
			t.Errorf("AllSeq(%v, f) = false; want true", s)
		}
		s = nil
		if !AllSeq(s, GreaterThan(1)) {
			t.Errorf("AllSeq(%v, f) = false; want true", s)
		}
	})
}
