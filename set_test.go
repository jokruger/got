package got

import (
	"slices"
	"testing"

	"github.com/jokruger/got/set"
)

func TestSet(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		s := set.New(1, 2, 2, 3)
		if s.Len() != 3 {
			t.Error("Set should have length 3")
		}

		s = set.NewFromSlice([]int{1, 2, 2, 3})
		if s.Len() != 3 {
			t.Error("Set should have length 3")
		}

		sl1 := []int{1, 2, 2, 3}
		s = set.NewFromSeq(slices.Values(sl1))
		if s.Len() != 3 {
			t.Error("Set should have length 3")
		}

		s = set.NewFromSet(set.New(1, 2, 2, 3))
		if s.Len() != 3 {
			t.Error("Set should have length 3")
		}
	})

	t.Run("Set of ints", func(t *testing.T) {
		s := set.New[int]()
		if s.Contains(42) {
			t.Error("Set should not contain 42")
		}
		s.Add(42)
		if !s.Contains(42) {
			t.Error("Set should contain 42")
		}
		s.Remove(42)
		if s.Contains(42) {
			t.Error("Set should not contain 42")
		}
	})

	t.Run("Set of ints from slice", func(t *testing.T) {
		s := set.NewFromSlice([]int{1, 2, 3})
		if !s.Contains(1) {
			t.Error("Set should contain 1")
		}
		if !s.Contains(2) {
			t.Error("Set should contain 2")
		}
		if !s.Contains(3) {
			t.Error("Set should contain 3")
		}
		if s.Contains(4) {
			t.Error("Set should not contain 4")
		}
		if !s.Contains(2, 3) {
			t.Error("Set should contain 2 and 3")
		}
		if s.Contains(2, 4) {
			t.Error("Set should not contain 4")
		}
		if !s.ContainsAny(2, 4) {
			t.Error("Set should contain 2")
		}
		if s.ContainsAny(4, 5) {
			t.Error("Set should not contain 4 or 5")
		}

		s = set.New[int]()
		s.AddMany(1, 2, 3)
		if !s.Contains(1) {
			t.Error("Set should contain 1")
		}
		if !s.Contains(2) {
			t.Error("Set should contain 2")
		}
		if !s.Contains(3) {
			t.Error("Set should contain 3")
		}
		if s.Contains(4) {
			t.Error("Set should not contain 4")
		}
	})

	t.Run("Set of ints from other sets", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := set.New(2, 3, 4)
		s := set.NewFromSet(s1)
		s.AddSet(s2)
		if !s.Contains(1) {
			t.Error("Set should contain 1")
		}
		if !s.Contains(2) {
			t.Error("Set should contain 2")
		}
		if !s.Contains(3) {
			t.Error("Set should contain 3")
		}
		if !s.Contains(4) {
			t.Error("Set should contain 4")
		}
		if s.Contains(5) {
			t.Error("Set should not contain 5")
		}
	})

	t.Run("Set of ints from iterators", func(t *testing.T) {
		slice := []int{1, 2, 3}
		s := set.NewFromSeq(slices.Values(slice))
		if !s.Contains(1) {
			t.Error("Set should contain 1")
		}
		if !s.Contains(2) {
			t.Error("Set should contain 2")
		}
		if !s.Contains(3) {
			t.Error("Set should contain 3")
		}
		if s.Contains(4) {
			t.Error("Set should not contain 4")
		}
	})

	t.Run("Set of ints add", func(t *testing.T) {
		s := set.New[int]()
		s.AddMany(1, 2, 3)
		if !s.Contains(1) {
			t.Error("Set should contain 1")
		}
		if !s.Contains(2) {
			t.Error("Set should contain 2")
		}
		if !s.Contains(3) {
			t.Error("Set should contain 3")
		}
		if s.Contains(4) {
			t.Error("Set should not contain 4")
		}

		s.AddSet(set.New(4, 5, 6))
		if !s.Contains(4) {
			t.Error("Set should contain 4")
		}
		if !s.Contains(5) {
			t.Error("Set should contain 5")
		}
		if !s.Contains(6) {
			t.Error("Set should contain 6")
		}

		s.AddSlice([]int{7, 8, 9})
		if !s.Contains(7) {
			t.Error("Set should contain 7")
		}
		if !s.Contains(8) {
			t.Error("Set should contain 8")
		}
		if !s.Contains(9) {
			t.Error("Set should contain 9")
		}

		slice := []int{10, 11, 12}
		s.AddSeq(slices.Values(slice))
		if !s.Contains(10) {
			t.Error("Set should contain 10")
		}
		if !s.Contains(11) {
			t.Error("Set should contain 11")
		}
		if !s.Contains(12) {
			t.Error("Set should contain 12")
		}
	})

	t.Run("Set of ints remove", func(t *testing.T) {
		s := set.New(1, 2, 3, 4, 5, 6, 7, 9)
		s.RemoveMany(1, 2)
		if s.Contains(1) {
			t.Error("Set should not contain 1")
		}
		if s.Contains(2) {
			t.Error("Set should not contain 2")
		}

		s.RemoveSet(set.New(3, 4))
		if s.Contains(3) {
			t.Error("Set should not contain 3")
		}
		if s.Contains(4) {
			t.Error("Set should not contain 4")
		}

		s.RemoveSlice([]int{5, 6})
		if s.Contains(5) {
			t.Error("Set should not contain 5")
		}
		if s.Contains(6) {
			t.Error("Set should not contain 6")
		}

		slice := []int{7, 9}
		s.RemoveSeq(slices.Values(slice))
		if s.Contains(7) {
			t.Error("Set should not contain 7")
		}
		if s.Contains(9) {
			t.Error("Set should not contain 9")
		}
	})

	t.Run("Set of ints clear", func(t *testing.T) {
		s := set.New(1, 2, 3)
		if s.IsEmpty() {
			t.Error("Set should not be empty")
		}
		s.Clear()
		if !s.IsEmpty() {
			t.Error("Set should be empty")
		}
	})

	t.Run("Set of ints contains all", func(t *testing.T) {
		s := set.New(1, 2, 3)
		if !s.Contains(1, 2) {
			t.Error("Set should contain 1 and 2")
		}
		if s.Contains(1, 4) {
			t.Error("Set should not contain 4")
		}
		if !s.ContainsSet(set.New(1, 2)) {
			t.Error("Set should contain 1 and 2")
		}
		if s.ContainsSet(set.New(1, 4)) {
			t.Error("Set should not contain 4")
		}
		if !s.ContainsSlice([]int{1, 2}) {
			t.Error("Set should contain 1 and 2")
		}
		if s.ContainsSlice([]int{1, 4}) {
			t.Error("Set should not contain 4")
		}

		slice := []int{1, 2}
		if !s.ContainsSeq(slices.Values(slice)) {
			t.Error("Set should contain 1 and 2")
		}
		slice = []int{1, 4}
		if s.ContainsSeq(slices.Values(slice)) {
			t.Error("Set should not contain 4")
		}
	})

	t.Run("Set of ints contains any", func(t *testing.T) {
		s := set.New(1, 2, 3)
		if !s.ContainsAny(1, 4) {
			t.Error("Set should contain 1")
		}
		if s.ContainsAny(4, 5) {
			t.Error("Set should not contain 4 or 5")
		}
		if !s.ContainsAnyFromSet(set.New(1, 4)) {
			t.Error("Set should contain 1")
		}
		if s.ContainsAnyFromSet(set.New(4, 5)) {
			t.Error("Set should not contain 4 or 5")
		}
		if !s.ContainsAnyFromSlice([]int{1, 4}) {
			t.Error("Set should contain 1")
		}
		if s.ContainsAnyFromSlice([]int{4, 5}) {
			t.Error("Set should not contain 4 or 5")
		}

		slice := []int{1, 4}
		if !s.ContainsAnyFromSeq(slices.Values(slice)) {
			t.Error("Set should contain 1")
		}
		slice = []int{4, 5}
		if s.ContainsAnyFromSeq(slices.Values(slice)) {
			t.Error("Set should not contain 4 or 5")
		}
	})

	t.Run("Set of ints len", func(t *testing.T) {
		s := set.New(1, 2, 3)
		if s.Len() != 3 {
			t.Error("Set should have length 3")
		}
	})

	t.Run("Set of ints equal", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := set.New(1, 2, 3)
		if !s1.Equal(s2) {
			t.Error("Sets should be equal")
		}
		s3 := set.New(1, 2, 3, 4)
		if s1.Equal(s3) {
			t.Error("Sets should not be equal")
		}
	})

	t.Run("Set of ints clone", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := s1.Clone()
		if !s1.Equal(s2) {
			t.Error("Sets should be equal")
		}
	})

	t.Run("Set of ints slice", func(t *testing.T) {
		s := set.New(1, 2, 3)
		slice := s.ToSlice()
		if len(slice) != 3 {
			t.Error("Slice should have length 3")
		}
		slices.Sort(slice)
		if slice[0] != 1 {
			t.Error("Slice should contain 1")
		}
		if slice[1] != 2 {
			t.Error("Slice should contain 2")
		}
		if slice[2] != 3 {
			t.Error("Slice should contain 3")
		}
	})

	t.Run("Set of ints iter", func(t *testing.T) {
		s := set.New(1, 2, 3)
		for e := range s.ToSeq() {
			if !s.Contains(e) {
				t.Errorf("Set should contain %d", e)
			}
		}
	})

	t.Run("Set of ints is subset of", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := set.New(1, 2, 3, 4)
		if !s1.IsSubsetOf(s2) {
			t.Error("Set should be a subset of other set")
		}
		if s2.IsSubsetOf(s1) {
			t.Error("Set should not be a subset of other set")
		}
	})

	t.Run("Set of ints is superset of", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := set.New(1, 2, 3, 4)
		if s1.IsSupersetOf(s2) {
			t.Error("Set should not be a superset of other set")
		}
		if !s2.IsSupersetOf(s1) {
			t.Error("Set should be a superset of other set")
		}
	})

	t.Run("Set of ints is proper subset of", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := set.New(1, 2, 3, 4)
		if !s1.IsProperSubsetOf(s2) {
			t.Error("Set should be a proper subset of other set")
		}
		if s2.IsProperSubsetOf(s1) {
			t.Error("Set should not be a proper subset of other set")
		}
		s3 := set.New(1, 2, 3)
		if s1.IsProperSubsetOf(s3) {
			t.Error("Set should not be a proper subset of other set")
		}
	})

	t.Run("Set of ints is proper superset of", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := set.New(1, 2, 3, 4)
		if s1.IsProperSupersetOf(s2) {
			t.Error("Set should not be a proper superset of other set")
		}
		if !s2.IsProperSupersetOf(s1) {
			t.Error("Set should be a proper superset of other set")
		}
		s3 := set.New(1, 2, 3)
		if s1.IsProperSupersetOf(s3) {
			t.Error("Set should not be a proper superset of other set")
		}
	})

	t.Run("Set of ints diff", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := set.New(1, 2, 3, 4)
		diff1 := s1.Diff(s2)
		if diff1.Len() != 0 {
			t.Error("Diff should be empty")
		}
		diff2 := s2.Diff(s1)
		if diff2.Len() != 1 {
			t.Error("Diff should have length 1")
		}
		if !diff2.Contains(4) {
			t.Error("Diff should contain 4")
		}
	})

	t.Run("Set of ints symmetric diff", func(t *testing.T) {
		s1 := set.New(1, 2, 3, 4, 5)
		s2 := set.New(2, 4, 6)
		diff := s1.SymmetricDiff(s2)
		if diff.Len() != 4 {
			t.Error("Diff should have length 4")
		}
		if !diff.Contains(1) {
			t.Error("Diff should contain 1")
		}
		if !diff.Contains(3) {
			t.Error("Diff should contain 3")
		}
		if !diff.Contains(5) {
			t.Error("Diff should contain 5")
		}
		if !diff.Contains(6) {
			t.Error("Diff should contain 6")
		}
	})

	t.Run("Set of ints union", func(t *testing.T) {
		s1 := set.New(1, 2, 3)
		s2 := set.New(2, 3, 4)
		union := s1.Union(s2)
		if union.Len() != 4 {
			t.Error("Union should have length 4")
		}
		if !union.Contains(1) {
			t.Error("Union should contain 1")
		}
		if !union.Contains(2) {
			t.Error("Union should contain 2")
		}
		if !union.Contains(3) {
			t.Error("Union should contain 3")
		}
		if !union.Contains(4) {
			t.Error("Union should contain 4")
		}
	})
}
