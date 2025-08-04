package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sequtil"
	"github.com/jokruger/got/sliceutil"
)

func TestFind(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		i, e := sliceutil.Find(s, basicutil.EqualTo(3))
		if i != 2 || e != 3 {
			t.Errorf("Find() = %v, %v; want 2, 3", i, e)
		}
		i, e = sliceutil.Find(s, basicutil.EqualTo(6))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
		s = nil
		i, e = sliceutil.Find(s, basicutil.EqualTo(3))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		i, e := sequtil.Find(s, basicutil.EqualTo(3))
		if i != 2 || e != 3 {
			t.Errorf("Find() = %v, %v; want 2, 3", i, e)
		}
		i, e = sequtil.Find(s, basicutil.EqualTo(6))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
		s = nil
		i, e = sequtil.Find(s, basicutil.EqualTo(3))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
	})
}

func TestFindPtr(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		p := sliceutil.FindPtr(s, func(x *int) bool { return *x == 3 })
		if p == nil || *p != 3 {
			t.Errorf("Find() = %v; want &s[3]", p)
		}
	})
}
