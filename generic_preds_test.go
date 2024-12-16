package got

import (
	"slices"
	"testing"

	. "github.com/jokruger/got/generics"
	"github.com/jokruger/got/set"
	gslices "github.com/jokruger/got/slices"
)

func TestPredicates(t *testing.T) {
	t.Run("InSet", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		f := set.NewFromElements(4, 3, 2)
		r := gslices.Filter(s, InSet(f))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 3 || r[2] != 4 {
			t.Errorf("expected 2, 3, 4, got %v", r)
		}
	})
	t.Run("Not InSet", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		f := set.NewFromElements(4, 3, 2)
		r := gslices.Filter(s, Not(InSet(f)))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != 1 || r[1] != 5 {
			t.Errorf("expected 1, 5, got %v", r)
		}
	})
	t.Run("InSlice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := gslices.Filter(s, InSlice([]int{4, 3, 2}))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 3 || r[2] != 4 {
			t.Errorf("expected 2, 3, 4, got %v", r)
		}
	})
	t.Run("Not InSlice", func(t *testing.T) {
		s := []string{"a", "b", "c", "d", "e"}
		r := gslices.Filter(s, Not(InSlice([]string{"d", "c", "b"})))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "a" || r[1] != "e" {
			t.Errorf("expected a, e, got %v", r)
		}
	})

	t.Run("InMap", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d"}
		r := gslices.Filter(s, InMap(m))
		if len(r) != 4 {
			t.Errorf("expected 4, got %d", len(r))
		}
		if r[0] != 1 || r[1] != 2 || r[2] != 3 || r[3] != 4 {
			t.Errorf("expected 1, 2, 3, 4, got %v", r)
		}
	})

	t.Run("InSeq", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := gslices.Filter(s, InSeq(slices.Values([]int{4, 3, 2})))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 3 || r[2] != 4 {
			t.Errorf("expected 2, 3, 4, got %v", r)
		}
	})
}
