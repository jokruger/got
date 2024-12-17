package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sequtil"
	"github.com/jokruger/got/sliceutil"
)

func TestReduce(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := sliceutil.Reduce(s, 0, basicutil.Add)
		if r != 15 {
			t.Errorf("Expected 15, got %d", r)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		r := sequtil.Reduce(s, 0, basicutil.Add)
		if r != 15 {
			t.Errorf("Expected 15, got %d", r)
		}
	})
}
