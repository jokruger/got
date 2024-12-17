package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/basicutils"
)

func TestRange(t *testing.T) {
	t.Run("Range(0)", func(t *testing.T) {
		r := basicutils.Range(0)
		if r != nil {
			t.Errorf("Range(0) = %v; want nil", r)
		}
	})
	t.Run("Range(1)", func(t *testing.T) {
		r := basicutils.Range(1)
		e := []int{0}
		if !slices.Equal(r, e) {
			t.Errorf("Range(1) = %v; want %v", r, e)
		}
	})
	t.Run("Range(5)", func(t *testing.T) {
		r := basicutils.Range(5)
		e := []int{0, 1, 2, 3, 4}
		if !slices.Equal(r, e) {
			t.Errorf("Range(5) = %v; want %v", r, e)
		}
	})
}
