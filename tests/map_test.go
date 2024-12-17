package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/sequtil"
	"github.com/jokruger/got/sliceutil"
)

func TestMap(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 1}

		r := sliceutil.Map(s, func(i int) int { return i * 2 })
		if len(r) != 6 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 || r[5] != 2 {
			t.Errorf("Expected [2, 4, 6, 8, 10, 2], got %v", r)
		}

		r = sliceutil.MapUnique(s, func(i int) int { return i * 2 })
		if len(r) != 5 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 {
			t.Errorf("Expected [2, 4, 6, 8, 10], got %v", r)
		}

		r = make([]int, 0)
		for i := range sliceutil.MapToSeq(s, func(i int) int { return i * 2 }) {
			r = append(r, i)
		}
		if len(r) != 6 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 || r[5] != 2 {
			t.Errorf("Expected [2, 4, 6, 8, 10, 2], got %v", r)
		}

		r = make([]int, 0)
		for i := range sliceutil.MapUniqueToSeq(s, func(i int) int { return i * 2 }) {
			r = append(r, i)
		}
		if len(r) != 5 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 {
			t.Errorf("Expected [2, 4, 6, 8, 10], got %v", r)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5, 1})
		r := make([]int, 0)
		for i := range sequtil.Map(s, func(i int) int { return i * 2 }) {
			r = append(r, i)
		}
		if len(r) != 6 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 || r[5] != 2 {
			t.Errorf("Expected [2, 4, 6, 8, 10, 2], got %v", r)
		}

		r = make([]int, 0)
		for i := range sequtil.MapUnique(s, func(i int) int { return i * 2 }) {
			r = append(r, i)
		}
		if len(r) != 5 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 {
			t.Errorf("Expected [2, 4, 6, 8, 10], got %v", r)
		}

		r = sequtil.MapToSlice(s, func(i int) int { return i * 2 })
		if len(r) != 6 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 || r[5] != 2 {
			t.Errorf("Expected [2, 4, 6, 8, 10, 2], got %v", r)
		}

		r = sequtil.MapUniqueToSlice(s, func(i int) int { return i * 2 })
		if len(r) != 5 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 {
			t.Errorf("Expected [2, 4, 6, 8, 10], got %v", r)
		}
	})
}
