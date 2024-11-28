package got

import (
	"slices"
	"testing"

	giter "github.com/jokruger/got/iter"
	gslices "github.com/jokruger/got/slices"
)

type TestExtractorsStruct struct {
	ID   int
	Name string
}

func (t TestExtractorsStruct) GetID() int {
	return t.ID
}

func (t TestExtractorsStruct) GetName() string {
	return t.Name
}

func TestExtractors(t *testing.T) {
	ms := []TestExtractorsStruct{
		{ID: 1, Name: "One"},
		{ID: 2, Name: "Two"},
		{ID: 3, Name: "Three"},
	}

	t.Run("extract ids from slice", func(t *testing.T) {
		ids := gslices.ExtractIDs(ms)
		if len(ids) != 3 {
			t.Fatalf("expected 3 ids, got %d", len(ids))
		}
		if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
			t.Fatalf("expected [1, 2, 3], got %v", ids)
		}

		ids = slices.Collect(gslices.ExtractIDsToSeq(ms))
		if len(ids) != 3 {
			t.Fatalf("expected 3 ids, got %d", len(ids))
		}
		if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
			t.Fatalf("expected [1, 2, 3], got %v", ids)
		}
	})

	t.Run("extract ids from iter", func(t *testing.T) {
		ids := giter.ExtractIDsToSlice(slices.Values(ms))
		if len(ids) != 3 {
			t.Fatalf("expected 3 ids, got %d", len(ids))
		}
		if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
			t.Fatalf("expected [1, 2, 3], got %v", ids)
		}

		ids = slices.Collect(giter.ExtractIDs(slices.Values(ms)))
		if len(ids) != 3 {
			t.Fatalf("expected 3 ids, got %d", len(ids))
		}
		if ids[0] != 1 || ids[1] != 2 || ids[2] != 3 {
			t.Fatalf("expected [1, 2, 3], got %v", ids)
		}
	})

	t.Run("extract names from slice", func(t *testing.T) {
		names := gslices.ExtractNames(ms)
		if len(names) != 3 {
			t.Fatalf("expected 3 names, got %d", len(names))
		}
		if names[0] != "One" || names[1] != "Two" || names[2] != "Three" {
			t.Fatalf("expected [\"One\", \"Two\", \"Three\"], got %v", names)
		}

		names = slices.Collect(gslices.ExtractNamesToSeq(ms))
		if len(names) != 3 {
			t.Fatalf("expected 3 names, got %d", len(names))
		}
		if names[0] != "One" || names[1] != "Two" || names[2] != "Three" {
			t.Fatalf("expected [\"One\", \"Two\", \"Three\"], got %v", names)
		}
	})

	t.Run("extract names from iter", func(t *testing.T) {
		names := giter.ExtractNamesToSlice(slices.Values(ms))
		if len(names) != 3 {
			t.Fatalf("expected 3 names, got %d", len(names))
		}
		if names[0] != "One" || names[1] != "Two" || names[2] != "Three" {
			t.Fatalf("expected [\"One\", \"Two\", \"Three\"], got %v", names)
		}

		names = slices.Collect(giter.ExtractNames(slices.Values(ms)))
		if len(names) != 3 {
			t.Fatalf("expected 3 names, got %d", len(names))
		}
		if names[0] != "One" || names[1] != "Two" || names[2] != "Three" {
			t.Fatalf("expected [\"One\", \"Two\", \"Three\"], got %v", names)
		}
	})
}
