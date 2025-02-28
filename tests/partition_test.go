package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/sequtil"
	"github.com/jokruger/got/sliceutil"
)

func TestPartitions(t *testing.T) {
	t.Run("Partition iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4})
		tg, fg := sequtil.Partition(s, func(i int) bool { return i%2 == 0 })
		ts := slices.Collect(tg)
		fs := slices.Collect(fg)
		if len(ts) != 2 || ts[0] != 2 || ts[1] != 4 {
			t.Errorf("expected [2 4], got %v", ts)
		}
		if len(fs) != 2 || fs[0] != 1 || fs[1] != 3 {
			t.Errorf("expected [1 3], got %v", fs)
		}

		ts, fs = sequtil.PartitionToSlice(s, func(i int) bool { return i%2 == 0 })
		if len(ts) != 2 || ts[0] != 2 || ts[1] != 4 {
			t.Errorf("expected [2 4], got %v", ts)
		}
		if len(fs) != 2 || fs[0] != 1 || fs[1] != 3 {
			t.Errorf("expected [1 3], got %v", fs)
		}
	})

	t.Run("PartitionConsEq iter", func(t *testing.T) {
		s := slices.Values([]int{1, 1, 2, 2, 3, 4, 4, 4})
		r := slices.Collect(sequtil.PartitionConsEq(s, func(a, b int) bool { return a == b }))
		if len(r) != 4 {
			t.Errorf("expected 4 groups, got %v", r)
		}
		if len(r[0]) != 2 || r[0][0] != 1 || r[0][1] != 1 {
			t.Errorf("expected [1 1], got %v", r[0])
		}
		if len(r[1]) != 2 || r[1][0] != 2 || r[1][1] != 2 {
			t.Errorf("expected [2 2], got %v", r[1])
		}
		if len(r[2]) != 1 || r[2][0] != 3 {
			t.Errorf("expected [3], got %v", r[2])
		}
		if len(r[3]) != 3 || r[3][0] != 4 || r[3][1] != 4 || r[3][2] != 4 {
			t.Errorf("expected [4 4 4], got %v", r[3])
		}

		r = sequtil.PartitionConsEqToSlice(s, func(a, b int) bool { return a == b })
		if len(r) != 4 {
			t.Errorf("expected 4 groups, got %v", r)
		}
		if len(r[0]) != 2 || r[0][0] != 1 || r[0][1] != 1 {
			t.Errorf("expected [1 1], got %v", r[0])
		}
		if len(r[1]) != 2 || r[1][0] != 2 || r[1][1] != 2 {
			t.Errorf("expected [2 2], got %v", r[1])
		}
		if len(r[2]) != 1 || r[2][0] != 3 {
			t.Errorf("expected [3], got %v", r[2])
		}
		if len(r[3]) != 3 || r[3][0] != 4 || r[3][1] != 4 || r[3][2] != 4 {
			t.Errorf("expected [4 4 4], got %v", r[3])
		}
	})

	t.Run("Partition slices", func(t *testing.T) {
		s := []int{1, 2, 3, 4}
		ts, fs := sliceutil.Partition(s, func(i int) bool { return i%2 == 0 })
		if len(ts) != 2 || ts[0] != 2 || ts[1] != 4 {
			t.Errorf("expected [2 4], got %v", ts)
		}
		if len(fs) != 2 || fs[0] != 1 || fs[1] != 3 {
			t.Errorf("expected [1 3], got %v", fs)
		}

		tg, fg := sliceutil.PartitionToSeq(s, func(i int) bool { return i%2 == 0 })
		ts = slices.Collect(tg)
		fs = slices.Collect(fg)
		if len(ts) != 2 || ts[0] != 2 || ts[1] != 4 {
			t.Errorf("expected [2 4], got %v", ts)
		}
		if len(fs) != 2 || fs[0] != 1 || fs[1] != 3 {
			t.Errorf("expected [1 3], got %v", fs)
		}
	})

	t.Run("PartitionConsEq slices", func(t *testing.T) {
		s := []int{1, 1, 2, 2, 3, 4, 4, 4}
		r := sliceutil.PartitionConsEq(s, func(a, b int) bool { return a == b })
		if len(r) != 4 {
			t.Errorf("expected 4 groups, got %v", r)
		}
		if len(r[0]) != 2 || r[0][0] != 1 || r[0][1] != 1 {
			t.Errorf("expected [1 1], got %v", r[0])
		}
		if len(r[1]) != 2 || r[1][0] != 2 || r[1][1] != 2 {
			t.Errorf("expected [2 2], got %v", r[1])
		}
		if len(r[2]) != 1 || r[2][0] != 3 {
			t.Errorf("expected [3], got %v", r[2])
		}
		if len(r[3]) != 3 || r[3][0] != 4 || r[3][1] != 4 || r[3][2] != 4 {
			t.Errorf("expected [4 4 4], got %v", r[3])
		}

		r = slices.Collect(sliceutil.PartitionConsEqToSeq(s, func(a, b int) bool { return a == b }))
		if len(r) != 4 {
			t.Errorf("expected 4 groups, got %v", r)
		}
		if len(r[0]) != 2 || r[0][0] != 1 || r[0][1] != 1 {
			t.Errorf("expected [1 1], got %v", r[0])
		}
		if len(r[1]) != 2 || r[1][0] != 2 || r[1][1] != 2 {
			t.Errorf("expected [2 2], got %v", r[1])
		}
		if len(r[2]) != 1 || r[2][0] != 3 {
			t.Errorf("expected [3], got %v", r[2])
		}
		if len(r[3]) != 3 || r[3][0] != 4 || r[3][1] != 4 || r[3][2] != 4 {
			t.Errorf("expected [4 4 4], got %v", r[3])
		}
	})
}
