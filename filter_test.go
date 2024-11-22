package got

import "testing"

func TestFilter(t *testing.T) {
	t.Run("Filter ints", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := Filter(s, func(i int) bool { return i%2 == 0 })
		if len(r) != 2 {
			t.Errorf("Expected 2, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 {
			t.Errorf("Expected [2, 4], got %v", r)
		}
	})
}
