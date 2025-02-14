package tests

import (
	"testing"

	"github.com/jokruger/got"
)

func TestTernary(t *testing.T) {
	if r := got.Ternary(1, 2, func() bool { return true }); r != 1 {
		t.Errorf("expected 1, got %d", r)
	}

	if r := got.Ternary(1, 2, func() bool { return false }); r != 2 {
		t.Errorf("expected 2, got %d", r)
	}
}
