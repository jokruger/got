package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/sequtil"
	"github.com/jokruger/got/sliceutil"
)

func TestFilter(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := sliceutil.Filter(s, func(i int) bool { return i%2 == 0 })
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}

		r = make([]int, 0)
		for i := range sliceutil.FilterToSeq(s, func(i int) bool { return i%2 == 0 }) {
			r = append(r, i)
		}
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}
	})

	t.Run("ints slice in place", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		s = sliceutil.FilterInPlace(s, func(i int) bool { return i%2 == 0 })
		if len(s) != 2 {
			t.Errorf("Expected 2, got %d", len(s))
		}
		if s[0] != 2 || s[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", s)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		r := make([]int, 0)
		for i := range sequtil.Filter(s, func(i int) bool { return i%2 == 0 }) {
			r = append(r, i)
		}
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}

		r = sequtil.FilterToSlice(s, func(i int) bool { return i%2 == 0 })
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}
	})
}
