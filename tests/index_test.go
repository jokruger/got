package tests

import (
	"testing"

	"github.com/jokruger/got/index"
)

func TestIndex(t *testing.T) {
	t.Run("uint16", func(t *testing.T) {
		a := 1
		b := 10
		c := 20
		d := 30
		idx := index.New[uint16, int](uint16(a), uint16(d))
		for i := a; i <= b; i++ {
			idx.Set(uint16(i), i)
		}
		for i := c; i <= d; i++ {
			idx.Set(uint16(i), i)
		}

		for i := -10; i < 40; i++ {
			if idx.Cap() != int(d-a+1) {
				t.Errorf("Expected index capacity to be %d, got %d", d-a+1, idx.Cap())
			}

			if idx.Len() != (b-a+1)+(d-c+1) {
				t.Errorf("Expected index length to be %d, got %d", (b-a+1)+(d-c+1), idx.Len())
			}

			if (i >= a && i <= b) || (i >= c && i <= d) {
				if !idx.Contains(uint16(i)) {
					t.Errorf("Expected index to contain %d, but it does not", i)
				}
				if v, ok := idx.Get(uint16(i)); !ok || v != i {
					t.Errorf("Expected index to return %d for %d, got %v (ok: %v)", i, i, v, ok)
				}
			} else {
				if idx.Contains(uint16(i)) {
					t.Errorf("Expected index to not contain %d, but it does", i)
				}
				if v, ok := idx.Get(uint16(i)); ok || v != 0 {
					t.Errorf("Expected index to return zero for %d, got %v (ok: %v)", i, v, ok)
				}
			}
		}
	})

	t.Run("int16", func(t *testing.T) {
		a := -5
		b := 10
		c := 20
		d := 30
		idx := index.New[int16, int](int16(a), int16(d))
		for i := a; i <= b; i++ {
			idx.Set(int16(i), i)
		}
		for i := c; i <= d; i++ {
			idx.Set(int16(i), i)
		}

		for i := -10; i < 40; i++ {
			if idx.Cap() != int(d-a+1) {
				t.Errorf("Expected index capacity to be %d, got %d", d-a+1, idx.Cap())
			}

			if idx.Len() != (b-a+1)+(d-c+1) {
				t.Errorf("Expected index length to be %d, got %d", (b-a+1)+(d-c+1), idx.Len())
			}

			if (i >= a && i <= b) || (i >= c && i <= d) {
				if !idx.Contains(int16(i)) {
					t.Errorf("Expected index to contain %d, but it does not", i)
				}
				if v, ok := idx.Get(int16(i)); !ok || v != i {
					t.Errorf("Expected index to return %d for %d, got %v (ok: %v)", i, i, v, ok)
				}
			} else {
				if idx.Contains(int16(i)) {
					t.Errorf("Expected index to not contain %d, but it does", i)
				}
				if v, ok := idx.Get(int16(i)); ok || v != 0 {
					t.Errorf("Expected index to return zero for %d, got %v (ok: %v)", i, v, ok)
				}
			}
		}
	})
}
