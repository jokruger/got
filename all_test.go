package got

import (
	"slices"
	"testing"

	. "github.com/jokruger/got/basic"
	giter "github.com/jokruger/got/iter"
	gslices "github.com/jokruger/got/slices"
)

func TestAll(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !gslices.All(s, GreaterThan(0)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		if gslices.All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = true; want false", s)
		}
		s = []int{}
		if !gslices.All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		s = nil
		if !gslices.All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		if !giter.All(s, GreaterThan(0)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		if giter.All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = true; want false", s)
		}
		s = slices.Values([]int{})
		if !giter.All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		s = nil
		if !giter.All(s, GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
	})
}
