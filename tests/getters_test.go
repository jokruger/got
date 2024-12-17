package tests

import (
	"testing"

	"github.com/jokruger/got/sliceutils"
	"github.com/jokruger/got/structutils"
)

type TestGettersStruct struct {
	ID   int
	Name string
}

func (t TestGettersStruct) GetID() int {
	return t.ID
}

func (t TestGettersStruct) GetName() string {
	return t.Name
}

func TestGetters(t *testing.T) {
	ms := []TestGettersStruct{
		{ID: 1, Name: "One"},
		{ID: 2, Name: "Two"},
		{ID: 3, Name: "Three"},
	}

	t.Run("get id", func(t *testing.T) {
		ids := sliceutils.Map(ms, structutils.GetID)
		if len(ids) != 3 {
			t.Fatalf("expected 3 ids, got %d", len(ids))
		}
		if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
			t.Fatalf("expected [1, 2, 3], got %v", ids)
		}
	})

	t.Run("get name", func(t *testing.T) {
		names := sliceutils.Map(ms, structutils.GetName)
		if len(names) != 3 {
			t.Fatalf("expected 3 names, got %d", len(names))
		}
		if names[0] != "One" || names[1] != "Two" || names[2] != "Three" {
			t.Fatalf("expected [\"One\", \"Two\", \"Three\"], got %v", names)
		}
	})
}
