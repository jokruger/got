package tests

import (
	"testing"

	"github.com/jokruger/got"
	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sliceutil"
	"github.com/jokruger/got/structutil"
)

type TestCompareStruct struct {
	ID   int
	Name string
}

func (t TestCompareStruct) GetID() int      { return t.ID }
func (t TestCompareStruct) GetName() string { return t.Name }

type IDStruct struct {
	ID int
}

func (i IDStruct) Compare(j IDStruct) int { return i.ID - j.ID }

type TestCompareStruct2 struct {
	ID IDStruct
}

func (t TestCompareStruct2) GetID() IDStruct { return t.ID }

func TestCompare(t *testing.T) {
	t.Run("ints slice", func(t *testing.T) {
		s := []int{3, 1, 2}
		sliceutil.Sort(s, basicutil.Compare)
		if len(s) != 3 || s[0] != 1 || s[1] != 2 || s[2] != 3 {
			t.Error("Sort failed")
		}
		sliceutil.Sort(s, got.Flip(basicutil.Compare[int]))
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
		sliceutil.Sort(ms, basicutil.CompareBy[TestCompareStruct](structutil.GetID))
		if len(ms) != 3 || ms[0].ID != 1 || ms[1].ID != 2 || ms[2].ID != 3 {
			t.Error("Sort failed")
		}
		sliceutil.Sort(ms, got.Flip(basicutil.CompareBy[TestCompareStruct](structutil.GetID)))
		if len(ms) != 3 || ms[0].ID != 3 || ms[1].ID != 2 || ms[2].ID != 1 {
			t.Error("Sort failed")
		}
	})

	t.Run("structs slice, two fields", func(t *testing.T) {
		ms := []TestCompareStruct{
			{ID: 3, Name: "D"},
			{ID: 1, Name: "A"},
			{ID: 2, Name: "B"},
			{ID: 3, Name: "C"},
		}
		sliceutil.Sort(ms, basicutil.CompareBy2[TestCompareStruct](structutil.GetID, structutil.GetName))
		if len(ms) != 4 || ms[0].ID != 1 || ms[1].ID != 2 || ms[2].ID != 3 || ms[3].ID != 3 {
			t.Error("Sort failed")
		}
		if ms[0].Name != "A" || ms[1].Name != "B" || ms[2].Name != "C" || ms[3].Name != "D" {
			t.Error("Sort failed")
		}
		sliceutil.Sort(ms, got.Flip(basicutil.CompareBy2[TestCompareStruct](structutil.GetID, structutil.GetName)))
		if len(ms) != 4 || ms[0].ID != 3 || ms[1].ID != 3 || ms[2].ID != 2 || ms[3].ID != 1 {
			t.Error("Sort failed")
		}
		if ms[0].Name != "D" || ms[1].Name != "C" || ms[2].Name != "B" || ms[3].Name != "A" {
			t.Error("Sort failed")
		}
	})

	t.Run("compare structs by complex field", func(t *testing.T) {
		ms := []TestCompareStruct2{
			{ID: IDStruct{ID: 3}},
			{ID: IDStruct{ID: 1}},
			{ID: IDStruct{ID: 2}},
		}

		sliceutil.Sort(ms, structutil.CompareBy[TestCompareStruct2](structutil.GetID))
		if len(ms) != 3 || ms[0].ID.ID != 1 || ms[1].ID.ID != 2 || ms[2].ID.ID != 3 {
			t.Error("Sort failed")
		}
	})
}
