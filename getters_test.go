package got

import (
	"testing"

	gslices "github.com/jokruger/got/slices"
	gstructs "github.com/jokruger/got/structs"
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
		ids := gslices.Map(ms, gstructs.GetID)
		if len(ids) != 3 {
			t.Fatalf("expected 3 ids, got %d", len(ids))
		}
		if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
			t.Fatalf("expected [1, 2, 3], got %v", ids)
		}
	})

	t.Run("get name", func(t *testing.T) {
		names := gslices.Map(ms, gstructs.GetName)
		if len(names) != 3 {
			t.Fatalf("expected 3 names, got %d", len(names))
		}
		if names[0] != "One" || names[1] != "Two" || names[2] != "Three" {
			t.Fatalf("expected [\"One\", \"Two\", \"Three\"], got %v", names)
		}
	})
}
