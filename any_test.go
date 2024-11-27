package got

import (
	"slices"
	"testing"

	. "github.com/jokruger/got/basic"
	giter "github.com/jokruger/got/iter"
	gslices "github.com/jokruger/got/slices"
)

func TestAny(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !gslices.Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = false; want true", s)
		}
		if gslices.Any(s, GreaterThan(3)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = []int{}
		if gslices.Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = nil
		if gslices.Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		if !giter.Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = false; want true", s)
		}
		if giter.Any(s, GreaterThan(3)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = slices.Values([]int{})
		if giter.Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = nil
		if giter.Any(s, GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})
}
