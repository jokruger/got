package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutils"
	"github.com/jokruger/got/sequtils"
	"github.com/jokruger/got/sliceutils"
)

func TestReduce(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}
		r := sliceutils.Reduce(s, 0, basicutils.Add)
		if r != 15 {
			t.Errorf("Expected 15, got %d", r)
		}
	})

	t.Run("ints iter", func(t *testing.T) {
		s := slices.Values([]int{1, 2, 3, 4, 5})
		r := sequtils.Reduce(s, 0, basicutils.Add)
		if r != 15 {
			t.Errorf("Expected 15, got %d", r)
		}
	})
}
