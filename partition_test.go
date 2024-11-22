package got

import (
	"slices"
	"testing"
)

func TestPartitions(t *testing.T) {
	t.Run("Partitions ints", func(t *testing.T) {
		s := []int{1, 1, 2, 2, 2, 3, 4, 4, 4, 4}
		r := slices.Collect(Partitions(s, func(a, b int) bool { return a == b }))

		if len(r) != 4 {
			t.Errorf("Expected 4, got %d", len(r))
		}

		if len(r[0]) != 2 {
			t.Errorf("Expected 2, got %d", len(r[0]))
		}
		if r[0][0] != 1 || r[0][1] != 1 {
			t.Errorf("Expected [1, 1], got %v", r[0])
		}

		if len(r[1]) != 3 {
			t.Errorf("Expected 3, got %d", len(r[1]))
		}
		if r[1][0] != 2 || r[1][1] != 2 || r[1][2] != 2 {
			t.Errorf("Expected [2, 2, 2], got %v", r[1])
		}

		if len(r[2]) != 1 {
			t.Errorf("Expected 1, got %d", len(r[2]))
		}
		if r[2][0] != 3 {
			t.Errorf("Expected [3], got %v", r[2])
		}

		if len(r[3]) != 4 {
			t.Errorf("Expected 4, got %d", len(r[3]))
		}
		if r[3][0] != 4 || r[3][1] != 4 || r[3][2] != 4 || r[3][3] != 4 {
			t.Errorf("Expected [4, 4, 4, 4], got %v", r[3])
		}
	})
}
