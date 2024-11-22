package got

import "testing"

func TestSet(t *testing.T) {
	t.Run("Set of ints", func(t *testing.T) {
		s := NewSet[int]()
		if s.Contains(42) {
			t.Error("Set should not contain 42")
		}
		s.Add(42)
		if !s.Contains(42) {
			t.Error("Set should contain 42")
		}
		s.Remove(42)
		if s.Contains(42) {
			t.Error("Set should not contain 42")
		}
	})

	t.Run("Set of ints from slice", func(t *testing.T) {
		s := NewSetFromSlice([]int{1, 2, 3})
		if !s.Contains(1) {
			t.Error("Set should contain 1")
		}
		if !s.Contains(2) {
			t.Error("Set should contain 2")
		}
		if !s.Contains(3) {
			t.Error("Set should contain 3")
		}
		if s.Contains(4) {
			t.Error("Set should not contain 4")
		}
		if !s.ContainsAll(2, 3) {
			t.Error("Set should contain 2 and 3")
		}
		if s.ContainsAll(2, 4) {
			t.Error("Set should not contain 4")
		}
		if !s.ContainsAny(2, 4) {
			t.Error("Set should contain 2")
		}
		if s.ContainsAny(4, 5) {
			t.Error("Set should not contain 4 or 5")
		}

		s = NewSet[int]()
		s.Add(1, 2, 3)
		if !s.Contains(1) {
			t.Error("Set should contain 1")
		}
		if !s.Contains(2) {
			t.Error("Set should contain 2")
		}
		if !s.Contains(3) {
			t.Error("Set should contain 3")
		}
		if s.Contains(4) {
			t.Error("Set should not contain 4")
		}
	})
}
