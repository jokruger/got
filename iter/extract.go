package iter

import (
	"iter"

	. "github.com/jokruger/got/ifs"
)

// ExtractIDs extracts the IDs of the elements in the sequence.
func ExtractIDs[I any, T IDProvider[I]](s iter.Seq[T]) iter.Seq[I] {
	return func(yield func(I) bool) {
		for e := range s {
			if !yield(e.GetID()) {
				return
			}
		}
	}
}

// ExtractIDsToSlice extracts the IDs of the elements in the sequence and returns them as a slice.
func ExtractIDsToSlice[I any, T IDProvider[I]](s iter.Seq[T]) []I {
	var res []I
	for e := range ExtractIDs(s) {
		res = append(res, e)
	}
	return res
}

// ExtractNames extracts the names of the elements in the sequence.
func ExtractNames[T NameProvider](s iter.Seq[T]) iter.Seq[string] {
	return func(yield func(string) bool) {
		for e := range s {
			if !yield(e.GetName()) {
				return
			}
		}
	}
}

// ExtractNamesToSlice extracts the names of the elements in the sequence and returns them as a slice.
func ExtractNamesToSlice[T NameProvider](s iter.Seq[T]) []string {
	var res []string
	for e := range ExtractNames(s) {
		res = append(res, e)
	}
	return res
}
