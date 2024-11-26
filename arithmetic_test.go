package got

import (
	"testing"
)

func TestArithmetic(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add(1, 2) != 3")
	}

	if Sub(1, 2) != -1 {
		t.Error("Sub(1, 2) != -1")
	}

	if Mul(3, 2) != 6 {
		t.Error("Mul(3, 2) != 6")
	}

	if Div(6, 2) != 3 {
		t.Error("Div(6, 2) != 3")
	}
}
