package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutils"
	"github.com/jokruger/got/sequtils"
	"github.com/jokruger/got/sliceutils"
)

func TestFind(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		i, e := sliceutils.Find(s, basicutils.EqualTo(3))
		if i != 2 || e != 3 {
			t.Errorf("Find() = %v, %v; want 2, 3", i, e)
		}
		i, e = sliceutils.Find(s, basicutils.EqualTo(6))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
		s = nil
		i, e = sliceutils.Find(s, basicutils.EqualTo(3))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		i, e := sequtils.Find(s, basicutils.EqualTo(3))
		if i != 2 || e != 3 {
			t.Errorf("Find() = %v, %v; want 2, 3", i, e)
		}
		i, e = sequtils.Find(s, basicutils.EqualTo(6))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
		s = nil
		i, e = sequtils.Find(s, basicutils.EqualTo(3))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
	})
}
