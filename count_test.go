package got

import (
	"slices"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("Count ints", func(t *testing.T) {
		s := []int{1, 2, 3}
		r := Count(s, GreaterThan(0))
		if r != 3 {
			t.Errorf("Count(%v, f) = %d; want 3", s, r)
		}
		r = Count(s, GreaterThan(1))
		if r != 2 {
			t.Errorf("Count(%v, f) = %d; want 2", s, r)
		}

		s = []int{}
		r = Count(s, GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}

		s = nil
		r = Count(s, GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}
	})

	t.Run("CountSeq ints", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		r := CountSeq(s, GreaterThan(0))
		if r != 3 {
			t.Errorf("CountSeq(%v, f) = %d; want 3", s, r)
		}
		r = CountSeq(s, GreaterThan(1))
		if r != 2 {
			t.Errorf("CountSeq(%v, f) = %d; want 2", s, r)
		}

		s = slices.Values([]int{})
		r = CountSeq(s, GreaterThan(1))
		if r != 0 {
			t.Errorf("CountSeq(%v, f) = %d; want 0", s, r)
		}

		s = nil
		r = CountSeq(s, GreaterThan(1))
		if r != 0 {
			t.Errorf("CountSeq(%v, f) = %d; want 0", s, r)
		}
	})
}
