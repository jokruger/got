package got

import "testing"

func TestMap(t *testing.T) {
	t.Run("Map ints", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 1}
		r := Map(s, func(i int) int { return i * 2 })
		if len(r) != 6 {
			t.Errorf("Expected 6, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 || r[5] != 2 {
			t.Errorf("Expected [2, 4, 6, 8, 10, 2], got %v", r)
		}
	})

	t.Run("Map structs", func(t *testing.T) {
		type T struct {
			A int
			B string
		}
		s := []T{{1, "a"}, {2, "b"}, {3, "c"}}
		r1 := Map(s, func(i T) int { return i.A })
		r2 := Map(s, func(i T) string { return i.B })
		if len(r1) != 3 {
			t.Errorf("Expected 3, got %d", len(r1))
		}
		if r1[0] != 1 || r1[1] != 2 || r1[2] != 3 {
			t.Errorf("Expected [1, 2, 3], got %v", r1)
		}
		if len(r2) != 3 {
			t.Errorf("Expected 3, got %d", len(r2))
		}
		if r2[0] != "a" || r2[1] != "b" || r2[2] != "c" {
			t.Errorf("Expected [a, b, c], got %v", r2)
		}
	})

	t.Run("MapToSeq ints", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		res := []int{}
		for r := range MapToSeq(s, func(i int) int { return i * 2 }) {
			res = append(res, r)
		}
		if len(res) != 5 {
			t.Errorf("Expected 5, got %d", len(res))
		}
		if res[0] != 2 || res[1] != 4 || res[2] != 6 || res[3] != 8 || res[4] != 10 {
			t.Errorf("Expected [2, 4, 6, 8, 10], got %v", res)
		}
	})
}

func TestMapUnique(t *testing.T) {
	t.Run("MapUnique ints", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5, 1, 2, 3}
		r := MapUnique(s, func(i int) int { return i * 2 })
		if len(r) != 5 {
			t.Errorf("Expected 5, got %d", len(r))
		}
		if r[0] != 2 || r[1] != 4 || r[2] != 6 || r[3] != 8 || r[4] != 10 {
			t.Errorf("Expected [2, 4, 6, 8, 10], got %v", r)
		}
	})

	t.Run("MapUnique structs", func(t *testing.T) {
		type T struct {
			A int
			B string
		}
		s := []T{{1, "a"}, {2, "b"}, {3, "c"}, {1, "b"}, {2, "c"}}
		r1 := MapUnique(s, func(i T) int { return i.A })
		r2 := MapUnique(s, func(i T) string { return i.B })
		if len(r1) != 3 {
			t.Errorf("Expected 3, got %d", len(r1))
		}
		if r1[0] != 1 || r1[1] != 2 || r1[2] != 3 {
			t.Errorf("Expected [1, 2, 3], got %v", r1)
		}
		if len(r2) != 3 {
			t.Errorf("Expected 3, got %d", len(r2))
		}
		if r2[0] != "a" || r2[1] != "b" || r2[2] != "c" {
			t.Errorf("Expected [a, b, c], got %v", r2)
		}
	})
}
