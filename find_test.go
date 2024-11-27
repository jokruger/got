package got

import (
	"slices"
	"testing"

	. "github.com/jokruger/got/basic"
	giter "github.com/jokruger/got/iter"
	gslices "github.com/jokruger/got/slices"
)

func TestFind(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		i, e := gslices.Find(s, EqualTo(3))
		if i != 2 || e != 3 {
			t.Errorf("Find() = %v, %v; want 2, 3", i, e)
		}
		i, e = gslices.Find(s, EqualTo(6))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
		s = nil
		i, e = gslices.Find(s, EqualTo(3))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		i, e := giter.Find(s, EqualTo(3))
		if i != 2 || e != 3 {
			t.Errorf("Find() = %v, %v; want 2, 3", i, e)
		}
		i, e = giter.Find(s, EqualTo(6))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
		s = nil
		i, e = giter.Find(s, EqualTo(3))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
	})
}
