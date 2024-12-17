package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sequtil"
	"github.com/jokruger/got/sliceutil"
)

func TestAll(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !sliceutil.All(s, basicutil.GreaterThan(0)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		if sliceutil.All(s, basicutil.GreaterThan(1)) {
			t.Errorf("All(%v, f) = true; want false", s)
		}
		s = []int{}
		if !sliceutil.All(s, basicutil.GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		s = nil
		if !sliceutil.All(s, basicutil.GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		if !sequtil.All(s, basicutil.GreaterThan(0)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		if sequtil.All(s, basicutil.GreaterThan(1)) {
			t.Errorf("All(%v, f) = true; want false", s)
		}
		s = slices.Values([]int{})
		if !sequtil.All(s, basicutil.GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		s = nil
		if !sequtil.All(s, basicutil.GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
	})
}
