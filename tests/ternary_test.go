package tests

import (
	"testing"

	"github.com/jokruger/got"
)

func TestTernary(t *testing.T) {
	if r := got.Ternary(1, 2, func() bool { return true }); r != 1 {
		t.Errorf("Expected 1, got %d", r)
	}

	if r := got.Ternary(1, 2, func() bool { return false }); r != 2 {
		t.Errorf("Expected 2, got %d", r)
	}
}
