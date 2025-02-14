package tests

import (
	"testing"

	"github.com/jokruger/got/basicutil"
)

func TestChoose(t *testing.T) {
	t.Run("Choose", func(t *testing.T) {
		if r := basicutil.Choose(0, 1, 2, 3); r != 1 {
			t.Errorf("expected 1, got %d", r)
		}

		if r := basicutil.Choose(0, 0, 0, 0); r != 0 {
			t.Errorf("expected 0, got %d", r)
		}

		if r := basicutil.Choose(0, 0, 0, 0, 1); r != 1 {
			t.Errorf("expected 1, got %d", r)
		}
	})
}
