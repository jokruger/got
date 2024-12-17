package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sequtil"
	"github.com/jokruger/got/sliceutil"
)

func TestCount(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		r := sliceutil.Count(s, basicutil.GreaterThan(0))
		if r != 3 {
			t.Errorf("Count(%v, f) = %d; want 3", s, r)
		}
		r = sliceutil.Count(s, basicutil.GreaterThan(1))
		if r != 2 {
			t.Errorf("Count(%v, f) = %d; want 2", s, r)
		}

		s = []int{}
		r = sliceutil.Count(s, basicutil.GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}

		s = nil
		r = sliceutil.Count(s, basicutil.GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		r := sequtil.Count(s, basicutil.GreaterThan(0))
		if r != 3 {
			t.Errorf("Count(%v, f) = %d; want 3", s, r)
		}
		r = sequtil.Count(s, basicutil.GreaterThan(1))
		if r != 2 {
			t.Errorf("Count(%v, f) = %d; want 2", s, r)
		}

		s = slices.Values([]int{})
		r = sequtil.Count(s, basicutil.GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}

		s = nil
		r = sequtil.Count(s, basicutil.GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}
	})
}
