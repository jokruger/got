package got

import (
	"slices"
	"testing"

	. "github.com/jokruger/got/basic"
	giter "github.com/jokruger/got/iter"
	gslices "github.com/jokruger/got/slices"
)

func TestCount(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		r := gslices.Count(s, GreaterThan(0))
		if r != 3 {
			t.Errorf("Count(%v, f) = %d; want 3", s, r)
		}
		r = gslices.Count(s, GreaterThan(1))
		if r != 2 {
			t.Errorf("Count(%v, f) = %d; want 2", s, r)
		}

		s = []int{}
		r = gslices.Count(s, GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}

		s = nil
		r = gslices.Count(s, GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		r := giter.Count(s, GreaterThan(0))
		if r != 3 {
			t.Errorf("Count(%v, f) = %d; want 3", s, r)
		}
		r = giter.Count(s, GreaterThan(1))
		if r != 2 {
			t.Errorf("Count(%v, f) = %d; want 2", s, r)
		}

		s = slices.Values([]int{})
		r = giter.Count(s, GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}

		s = nil
		r = giter.Count(s, GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}
	})
}
