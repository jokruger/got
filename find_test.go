package got

import (
	"slices"
	"testing"
)

func TestFind(t *testing.T) {
	t.Run("Find ints", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		i, e := Find(s, EqualTo(3))
		if i != 2 || e != 3 {
			t.Errorf("Find() = %v, %v; want 2, 3", i, e)
		}
		i, e = Find(s, EqualTo(6))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
		s = nil
		i, e = Find(s, EqualTo(3))
		if i != -1 || e != 0 {
			t.Errorf("Find() = %v, %v; want -1, 0", i, e)
		}
	})

	t.Run("FindSeq ints", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		i, e := FindSeq(s, EqualTo(3))
		if i != 2 || e != 3 {
			t.Errorf("FindSeq() = %v, %v; want 2, 3", i, e)
		}
		i, e = FindSeq(s, EqualTo(6))
		if i != -1 || e != 0 {
			t.Errorf("FindSeq() = %v, %v; want -1, 0", i, e)
		}
		s = nil
		i, e = FindSeq(s, EqualTo(3))
		if i != -1 || e != 0 {
			t.Errorf("FindSeq() = %v, %v; want -1, 0", i, e)
		}
	})
}
