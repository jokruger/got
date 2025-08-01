package tests

import (
	"testing"

	"github.com/jokruger/got/sliceutil"
)

func TestMinMax(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		min, max := sliceutil.MinMax(s, func(a, b int) int { return a - b })
		if min != 1 || max != 5 {
			t.Errorf("expected (1, 5) got (%v, %v)", min, max)
		}
		min, max = sliceutil.MapMinMax(s, func(x int) int { return x })
		if min != 1 || max != 5 {
			t.Errorf("expected (1, 5) got (%v, %v)", min, max)
		}
	})
}
