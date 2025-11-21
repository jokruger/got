package tests

import (
	"testing"

	"github.com/jokruger/got"
	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sliceutil"
	"github.com/jokruger/got/structutil"
)

type TestCompareStruct struct {
	id   int
	name string
}

func (t TestCompareStruct) ID() int      { return t.id }
func (t TestCompareStruct) Name() string { return t.name }

type IDStruct struct {
	id int
}

func (i IDStruct) Compare(j IDStruct) int { return i.id - j.id }

type TestCompareStruct2 struct {
	id IDStruct
}

func (t TestCompareStruct2) ID() IDStruct { return t.id }

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
			{id: 3, name: "Three"},
			{id: 1, name: "One"},
			{id: 2, name: "Two"},
		}
		sliceutil.Sort(ms, basicutil.CompareBy[TestCompareStruct](structutil.ID))
		if len(ms) != 3 || ms[0].id != 1 || ms[1].id != 2 || ms[2].id != 3 {
			t.Error("Sort failed")
		}
		sliceutil.Sort(ms, got.Flip(basicutil.CompareBy[TestCompareStruct](structutil.ID)))
		if len(ms) != 3 || ms[0].id != 3 || ms[1].id != 2 || ms[2].id != 1 {
			t.Error("Sort failed")
		}
	})

	t.Run("structs slice, two fields", func(t *testing.T) {
		ms := []TestCompareStruct{
			{id: 3, name: "D"},
			{id: 1, name: "A"},
			{id: 2, name: "B"},
			{id: 3, name: "C"},
		}
		sliceutil.Sort(ms, basicutil.CompareBy2[TestCompareStruct](structutil.ID, structutil.Name))
		if len(ms) != 4 || ms[0].id != 1 || ms[1].id != 2 || ms[2].id != 3 || ms[3].id != 3 {
			t.Error("Sort failed")
		}
		if ms[0].name != "A" || ms[1].name != "B" || ms[2].name != "C" || ms[3].name != "D" {
			t.Error("Sort failed")
		}
		sliceutil.Sort(ms, got.Flip(basicutil.CompareBy2[TestCompareStruct](structutil.ID, structutil.Name)))
		if len(ms) != 4 || ms[0].id != 3 || ms[1].id != 3 || ms[2].id != 2 || ms[3].id != 1 {
			t.Error("Sort failed")
		}
		if ms[0].name != "D" || ms[1].name != "C" || ms[2].name != "B" || ms[3].name != "A" {
			t.Error("Sort failed")
		}
	})

	t.Run("compare structs by complex field", func(t *testing.T) {
		ms := []TestCompareStruct2{
			{id: IDStruct{id: 3}},
			{id: IDStruct{id: 1}},
			{id: IDStruct{id: 2}},
		}

		sliceutil.Sort(ms, structutil.CompareBy[TestCompareStruct2](structutil.ID))
		if len(ms) != 3 || ms[0].id.id != 1 || ms[1].id.id != 2 || ms[2].id.id != 3 {
			t.Error("Sort failed")
		}
	})
}
