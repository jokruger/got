package got

import (
	"slices"
	"testing"
)

func TestReduce(t *testing.T) {
	t.Run("Reduce ints", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := Reduce(s, 0, Add)
		if r != 15 {
			t.Errorf("Expected 15, got %d", r)
		}
	})

	t.Run("ReduceSeq ints", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		r := ReduceSeq(s, 0, Add)
		if r != 15 {
			t.Errorf("Expected 15, got %d", r)
		}
	})
}
