package got

import (
	"slices"
	"testing"

	giter "github.com/jokruger/got/iter"
	gslices "github.com/jokruger/got/slices"
)

func TestFilter(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := gslices.Filter(s, func(i int) bool { return i%2 == 0 })
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}

		r = make([]int, 0)
		for i := range gslices.FilterToSeq(s, func(i int) bool { return i%2 == 0 }) {
			r = append(r, i)
		}
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		r := make([]int, 0)
		for i := range giter.Filter(s, func(i int) bool { return i%2 == 0 }) {
			r = append(r, i)
		}
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}

		r = giter.FilterToSlice(s, func(i int) bool { return i%2 == 0 })
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}
	})
}
