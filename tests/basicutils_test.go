package tests

import (
	"regexp"
	"testing"

	"github.com/jokruger/got"
	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sliceutil"
)

func TestBasicArithmetic(t *testing.T) {
	if basicutil.Add(1, 2) != 3 {
		t.Error("Add(1, 2) != 3")
	}

	if basicutil.Sub(1, 2) != -1 {
		t.Error("Sub(1, 2) != -1")
	}

	if basicutil.Mul(3, 2) != 6 {
		t.Error("Mul(3, 2) != 6")
	}

	if basicutil.Div(6, 2) != 3 {
		t.Error("Div(6, 2) != 3")
	}
}

func TestBasicPredicates(t *testing.T) {
	t.Run("LessThan 2 OR GreaterThan 4", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := sliceutil.Filter(s, got.Or(basicutil.LessThan(2), basicutil.GreaterThan(4)))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != 1 || r[1] != 5 {
			t.Errorf("expected 1, 5, got %v", r)
		}
	})
	t.Run("GreaterOrEqualTo 2 AND LessOrEqualTo 4", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := sliceutil.Filter(s, got.And(basicutil.GreaterOrEqualTo(2), basicutil.LessOrEqualTo(4)))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 3 || r[2] != 4 {
			t.Errorf("expected 2, 3, 4, got %v", r)
		}
	})
	t.Run("Not EqualTo 3", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := sliceutil.Filter(s, got.Not(basicutil.EqualTo(3)))
		if len(r) != 4 {
			t.Errorf("expected 4, got %d", len(r))
		}
		if r[0] != 1 || r[1] != 2 || r[2] != 4 || r[3] != 5 {
			t.Errorf("expected 1, 2, 4, 5, got %v", r)
		}
	})

	t.Run("InRange", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := sliceutil.Filter(s, basicutil.InRange(2, 4))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 3 || r[2] != 4 {
			t.Errorf("expected 2, 3, 4, got %v", r)
		}
	})

	t.Run("String predicates", func(t *testing.T) {
		s := []string{"apple", "banana", "cherry", "date", "elderberry"}
		r := sliceutil.Filter(s, basicutil.ContainsSubstring("rr"))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "cherry" || r[1] != "elderberry" {
			t.Errorf("expected cherry, elderberry, got %v", r)
		}

		r = sliceutil.Filter(s, basicutil.InString("applepie"))
		if len(r) != 1 {
			t.Errorf("expected 1, got %d", len(r))
		}
		if r[0] != "apple" {
			t.Errorf("expected apple, got %v", r)
		}

		r = sliceutil.Filter(s, basicutil.StartsWithString("b"))
		if len(r) != 1 {
			t.Errorf("expected 1, got %d", len(r))
		}
		if r[0] != "banana" {
			t.Errorf("expected banana, got %v", r)
		}

		r = sliceutil.Filter(s, basicutil.EndsWithString("e"))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "apple" || r[1] != "date" {
			t.Errorf("expected apple, date, got %v", r)
		}

		r = sliceutil.Filter(s, basicutil.MatchesRegexp(".*rr.*"))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "cherry" || r[1] != "elderberry" {
			t.Errorf("expected cherry, elderberry, got %v", r)
		}

		rc := regexp.MustCompile(".*rr.*")
		r = sliceutil.Filter(s, basicutil.MatchesRegexpCompiled(rc))
		if len(r) != 2 {
			t.Errorf("expected 2, got %d", len(r))
		}
		if r[0] != "cherry" || r[1] != "elderberry" {
			t.Errorf("expected cherry, elderberry, got %v", r)
		}
	})

}
