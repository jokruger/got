package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutils"
	"github.com/jokruger/got/sequtils"
	"github.com/jokruger/got/sliceutils"
)

func TestCount(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		r := sliceutils.Count(s, basicutils.GreaterThan(0))
		if r != 3 {
			t.Errorf("Count(%v, f) = %d; want 3", s, r)
		}
		r = sliceutils.Count(s, basicutils.GreaterThan(1))
		if r != 2 {
			t.Errorf("Count(%v, f) = %d; want 2", s, r)
		}

		s = []int{}
		r = sliceutils.Count(s, basicutils.GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}

		s = nil
		r = sliceutils.Count(s, basicutils.GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		r := sequtils.Count(s, basicutils.GreaterThan(0))
		if r != 3 {
			t.Errorf("Count(%v, f) = %d; want 3", s, r)
		}
		r = sequtils.Count(s, basicutils.GreaterThan(1))
		if r != 2 {
			t.Errorf("Count(%v, f) = %d; want 2", s, r)
		}

		s = slices.Values([]int{})
		r = sequtils.Count(s, basicutils.GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}

		s = nil
		r = sequtils.Count(s, basicutils.GreaterThan(1))
		if r != 0 {
			t.Errorf("Count(%v, f) = %d; want 0", s, r)
		}
	})
}
