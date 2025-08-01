package tests

import (
	"testing"

	"github.com/jokruger/got/sliceutil"
)

func TestMinMax(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		min, max, ok := sliceutil.MinMax(s, func(x int) int { return x })
		if min != 1 || max != 5 || ok != true {
			t.Errorf("expected (1, 5, true) got (%v, %v, %v)", min, max, ok)
		}

		min, max, ok = sliceutil.MinMax(nil, func(x int) int { return x })
		if min != 0 || max != 0 || ok != false {
			t.Errorf("expected (0, 0, false) got (%v, %v, %v)", min, max, ok)
		}
	})
}
