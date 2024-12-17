package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sequtil"
	"github.com/jokruger/got/sliceutil"
)

func TestAny(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !sliceutil.Any(s, basicutil.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = false; want true", s)
		}
		if sliceutil.Any(s, basicutil.GreaterThan(3)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = []int{}
		if sliceutil.Any(s, basicutil.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = nil
		if sliceutil.Any(s, basicutil.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		if !sequtil.Any(s, basicutil.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = false; want true", s)
		}
		if sequtil.Any(s, basicutil.GreaterThan(3)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = slices.Values([]int{})
		if sequtil.Any(s, basicutil.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = nil
		if sequtil.Any(s, basicutil.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})
}
