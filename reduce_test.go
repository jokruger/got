package got

import (
	"slices"
	"testing"

	. "github.com/jokruger/got/basic"
	giter "github.com/jokruger/got/iter"
	gslices "github.com/jokruger/got/slices"
)

func TestReduce(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := gslices.Reduce(s, 0, Add)
		if r != 15 {
			t.Errorf("Expected 15, got %d", r)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		r := giter.Reduce(s, 0, Add)
		if r != 15 {
			t.Errorf("Expected 15, got %d", r)
		}
	})
}
