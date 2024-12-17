package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutils"
	"github.com/jokruger/got/sequtils"
	"github.com/jokruger/got/sliceutils"
)

func TestAny(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !sliceutils.Any(s, basicutils.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = false; want true", s)
		}
		if sliceutils.Any(s, basicutils.GreaterThan(3)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = []int{}
		if sliceutils.Any(s, basicutils.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = nil
		if sliceutils.Any(s, basicutils.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		if !sequtils.Any(s, basicutils.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = false; want true", s)
		}
		if sequtils.Any(s, basicutils.GreaterThan(3)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = slices.Values([]int{})
		if sequtils.Any(s, basicutils.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
		s = nil
		if sequtils.Any(s, basicutils.GreaterThan(1)) {
			t.Errorf("Any(%v, f) = true; want false", s)
		}
	})
}
