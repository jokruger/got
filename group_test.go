package got

import "testing"

func TestGroup(t *testing.T) {
	t.Run("GroupBy", func(t *testing.T) {
		type X struct {
			A int
			B string
		}
		slice := []X{
			{1, "a"},
			{2, "b"},
			{1, "c"},
			{2, "d"},
		}
		m := GroupBy(slice, func(x X) (int, string) { return x.A, x.B })
		if len(m) != 2 {
			t.Errorf("len(m) = %d; want 2", len(m))
		}
		if len(m[1]) != 2 {
			t.Errorf("len(m[1]) = %d; want 2", len(m[1]))
		}
		if m[1][0] != "a" {
			t.Errorf("m[1][0] = %q; want %q", m[1][0], "a")
		}
		if m[1][1] != "c" {
			t.Errorf("m[1][1] = %q; want %q", m[1][1], "c")
		}
		if len(m[2]) != 2 {
			t.Errorf("len(m[2]) = %d; want 2", len(m[2]))
		}
		if m[2][0] != "b" {
			t.Errorf("m[2][0] = %q; want %q", m[2][0], "b")
		}
		if m[2][1] != "d" {
			t.Errorf("m[2][1] = %q; want %q", m[2][1], "d")
		}
	})
}
