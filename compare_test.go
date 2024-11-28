package got

import (
	"testing"

	gbasic "github.com/jokruger/got/basic"
	. "github.com/jokruger/got/generics"
	gslices "github.com/jokruger/got/slices"
	gstructs "github.com/jokruger/got/structs"
)

type TestCompareStruct struct {
	ID   int
	Name string
}

func (t TestCompareStruct) GetID() int {
	return t.ID
}

func TestCompare(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{3, 1, 2}
		gslices.Sort(s, gbasic.Compare)
		if len(s) != 3 || s[0] != 1 || s[1] != 2 || s[2] != 3 {
			t.Error("Sort failed")
		}
		gslices.Sort(s, Flip(gbasic.Compare[int]))
		if len(s) != 3 || s[0] != 3 || s[1] != 2 || s[2] != 1 {
			t.Error("Sort failed")
		}
	})

	t.Run("structs slice", func(t *testing.T) {
		ms := []TestCompareStruct{
			{ID: 3, Name: "Three"},
			{ID: 1, Name: "One"},
			{ID: 2, Name: "Two"},
		}
		gslices.Sort(ms, gstructs.CompareBy[TestCompareStruct](gstructs.GetID))
		if len(ms) != 3 || ms[0].ID != 1 || ms[1].ID != 2 || ms[2].ID != 3 {
			t.Error("Sort failed")
		}
		gslices.Sort(ms, Flip(gstructs.CompareBy[TestCompareStruct](gstructs.GetID)))
		if len(ms) != 3 || ms[0].ID != 3 || ms[1].ID != 2 || ms[2].ID != 1 {
			t.Error("Sort failed")
		}
	})
}
