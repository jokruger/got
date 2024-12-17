package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutils"
	"github.com/jokruger/got/sequtils"
	"github.com/jokruger/got/sliceutils"
)

func TestAll(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3}
		if !sliceutils.All(s, basicutils.GreaterThan(0)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		if sliceutils.All(s, basicutils.GreaterThan(1)) {
			t.Errorf("All(%v, f) = true; want false", s)
		}
		s = []int{}
		if !sliceutils.All(s, basicutils.GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		s = nil
		if !sliceutils.All(s, basicutils.GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3})
		if !sequtils.All(s, basicutils.GreaterThan(0)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		if sequtils.All(s, basicutils.GreaterThan(1)) {
			t.Errorf("All(%v, f) = true; want false", s)
		}
		s = slices.Values([]int{})
		if !sequtils.All(s, basicutils.GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
		s = nil
		if !sequtils.All(s, basicutils.GreaterThan(1)) {
			t.Errorf("All(%v, f) = false; want true", s)
		}
	})
}
