package slices

import (
	"iter"

	. "github.com/jokruger/got/ifs"
)

// ExtractIDs returns a slice of IDs from a slice of elements.
func ExtractIDs[I any, T IDProvider[I]](s []T) []I {
	res := make([]I, len(s))
	for i, e := range s {
		res[i] = e.GetID()
	}
	return res
}

// ExtractIDsToIter returns an iterator that yields IDs from a slice of elements.
func ExtractIDsToSeq[I any, T IDProvider[I]](s []T) iter.Seq[I] {
	return func(yield func(I) bool) {
		for _, e := range s {
			if !yield(e.GetID()) {
				return
			}
		}
	}
}

// ExtractNames returns a slice of names from a slice of elements.
func ExtractNames[T NameProvider](s []T) []string {
	res := make([]string, len(s))
	for i, e := range s {
		res[i] = e.GetName()
	}
	return res
}

// ExtractNamesToIter returns an iterator that yields names from a slice of elements.
func ExtractNamesToSeq[T NameProvider](s []T) iter.Seq[string] {
	return func(yield func(string) bool) {
		for _, e := range s {
			if !yield(e.GetName()) {
				return
			}
		}
	}
}
