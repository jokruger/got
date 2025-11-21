package tests

import (
	"testing"

	"github.com/jokruger/got/sliceutil"
	"github.com/jokruger/got/structutil"
)

type TestGettersStruct struct {
	id   int
	name string
}

func (t TestGettersStruct) ID() int {
	return t.id
}

func (t TestGettersStruct) Name() string {
	return t.name
}

func TestGetters(t *testing.T) {
	ms := []TestGettersStruct{
		{id: 1, name: "One"},
		{id: 2, name: "Two"},
		{id: 3, name: "Three"},
	}

	t.Run("get id", func(t *testing.T) {
		ids := sliceutil.Map(ms, structutil.ID)
		if len(ids) != 3 {
			t.Fatalf("expected 3 ids, got %d", len(ids))
		}
		if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
			t.Fatalf("expected [1, 2, 3], got %v", ids)
		}
	})

	t.Run("get name", func(t *testing.T) {
		names := sliceutil.Map(ms, structutil.Name)
		if len(names) != 3 {
			t.Fatalf("expected 3 names, got %d", len(names))
		}
		if names[0] != "One" || names[1] != "Two" || names[2] != "Three" {
			t.Fatalf("expected [\"One\", \"Two\", \"Three\"], got %v", names)
		}
	})
}
