package got

import (
	"testing"
)

func TestPartitions(t *testing.T) {
	t.Run("PartitionToSeq", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5}
		tgi, fgi := PartitionToSeq(slice, func(e int) bool { return e%2 == 0 })
		tg := make([]int, 0)
		for e := range tgi {
			tg = append(tg, e)
		}
		fg := make([]int, 0)
		for e := range fgi {
			fg = append(fg, e)
		}
		if len(tg) != 2 || tg[0] != 2 || tg[1] != 4 {
			t.Errorf("PartitionToSeq failed, expected [2, 4], got %v", tg)
		}
		if len(fg) != 3 || fg[0] != 1 || fg[1] != 3 || fg[2] != 5 {
			t.Errorf("PartitionToSeq failed, expected [1, 3, 5], got %v", fg)
		}
	})

	t.Run("PartitionConsEq", func(t *testing.T) {
		s := []int{1, 1, 2, 2, 2, 3, 3, 3, 3}
		r := PartitionConsEq(s, Equal)
		if len(r) != 3 {
			t.Errorf("PartitionConsEq failed, expected 3, got %d", len(r))
		}
		if len(r[0]) != 2 || r[0][0] != 1 || r[0][1] != 1 {
			t.Errorf("PartitionConsEq failed, expected [1, 1], got %v", r[0])
		}
		if len(r[1]) != 3 || r[1][0] != 2 || r[1][1] != 2 || r[1][2] != 2 {
			t.Errorf("PartitionConsEq failed, expected [2, 2, 2], got %v", r[1])
		}
		if len(r[2]) != 4 || r[2][0] != 3 || r[2][1] != 3 || r[2][2] != 3 || r[2][3] != 3 {
			t.Errorf("PartitionConsEq failed, expected [3, 3, 3, 3], got %v", r[2])
		}
	})
}
