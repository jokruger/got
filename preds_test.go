package got

import (
	"regexp"
	"slices"
	"testing"

	. "github.com/jokruger/got/preds"
	"github.com/jokruger/got/set"
	gslices "github.com/jokruger/got/slices"
)

func TestPredicates(t *testing.T) {
	t.Run("InSet", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		f := set.New(4, 3, 2)
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
		f := set.New(4, 3, 2)
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
	t.Run("String predicates", func(t *testing.T) {
		s := []string{"apple", "banana", "cherry", "date", "elderberry"}
		r := gslices.Filter(s, ContainsSubstring("rr"))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "cherry" || r[1] != "elderberry" {
			t.Errorf("expected cherry, elderberry, got %v", r)
		}

		r = gslices.Filter(s, InString("applepie"))
		if len(r) != 1 {
			t.Errorf("expected 1, got %d", len(r))
		}
		if r[0] != "apple" {
			t.Errorf("expected apple, got %v", r)
		}

		r = gslices.Filter(s, StartsWithString("b"))
		if len(r) != 1 {
			t.Errorf("expected 1, got %d", len(r))
		}
		if r[0] != "banana" {
			t.Errorf("expected banana, got %v", r)
		}

		r = gslices.Filter(s, EndsWithString("e"))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "apple" || r[1] != "date" {
			t.Errorf("expected apple, date, got %v", r)
		}

		r = gslices.Filter(s, MatchesRegexp(".*rr.*"))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "cherry" || r[1] != "elderberry" {
			t.Errorf("expected cherry, elderberry, got %v", r)
		}

		rc := regexp.MustCompile(".*rr.*")
		r = gslices.Filter(s, MatchesRegexpCompiled(rc))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "cherry" || r[1] != "elderberry" {
			t.Errorf("expected cherry, elderberry, got %v", r)
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