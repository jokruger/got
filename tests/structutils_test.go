package tests

import (
	"testing"
	"time"

	"github.com/jokruger/got/sliceutil"
	"github.com/jokruger/got/structutil"
)

func TestStructs(t *testing.T) {
	t.Run("predicates", func(t *testing.T) {
		a := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
		b := time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)
		c := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		s := []time.Time{a, b, c}

		r := sliceutil.Map(s, structutil.InRange[time.Time, time.Time](a, b))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != true || r[1] != true || r[2] != false {
			t.Errorf("expected [true, true, false], got %v", r)
		}

		r = sliceutil.Map(s, structutil.GreaterThan[time.Time, time.Time](b))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != false || r[1] != false || r[2] != true {
			t.Errorf("expected [false, false, true], got %v", r)
		}

		r = sliceutil.Map(s, structutil.LessThan[time.Time, time.Time](b))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != true || r[1] != false || r[2] != false {
			t.Errorf("expected [true, false, false], got %v", r)
		}

		r = sliceutil.Map(s, structutil.GreaterOrEqualTo[time.Time, time.Time](b))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != false || r[1] != true || r[2] != true {
			t.Errorf("expected [false, true, true], got %v", r)
		}

		r = sliceutil.Map(s, structutil.LessOrEqualTo[time.Time, time.Time](b))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != true || r[1] != true || r[2] != false {
			t.Errorf("expected [true, true, false], got %v", r)
		}

		r = sliceutil.Map(s, structutil.EqualTo[time.Time, time.Time](b))
		if len(r) != 3 {
			t.Errorf("expected 3, got %d", len(r))
		}
		if r[0] != false || r[1] != true || r[2] != false {
			t.Errorf("expected [false, true, false], got %v", r)
		}

		if structutil.Equal[time.Time, time.Time](a, a) != true {
			t.Errorf("expected true, got false")
		}
		if structutil.Equal[time.Time, time.Time](a, b) != false {
			t.Errorf("expected false, got true")
		}

		s = append(s, time.Time{})
		r = sliceutil.Map(s, structutil.IsZero)
		if len(r) != 4 {
			t.Errorf("expected 4, got %d", len(r))
		}
		if r[0] != false || r[1] != false || r[2] != false || r[3] != true {
			t.Errorf("expected [false, false, false, true], got %v", r)
		}
	})
}
