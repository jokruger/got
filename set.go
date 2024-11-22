package got

import (
	"iter"
	"maps"
	"slices"
)

// Set is a set of elements of type T.
type Set[T comparable] map[T]interface{}

// NewSet creates a new set. The set can be initialized with other sets.
func NewSet[T comparable](src ...Set[T]) Set[T] {
	set := make(Set[T])
	for _, s := range src {
		set.AddSet(s)
	}
	return set
}

// NewSetFromSlice creates a new set from a slice.
func NewSetFromSlice[T comparable](s []T) Set[T] {
	set := make(Set[T])
	for _, e := range s {
		set[e] = nil
	}
	return set
}

// NewSetFromIter creates a new set from an iterator.
func NewSetFromIter[T comparable](is iter.Seq[T]) Set[T] {
	set := make(Set[T])
	for e := range is {
		set[e] = nil
	}
	return set
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Add adds elements to the set.
func (s Set[T]) Add(es ...T) {
	for _, e := range es {
		s[e] = nil
	}
}

// AddIter adds elements to the set from an iterator.
func (s Set[T]) AddIter(is iter.Seq[T]) {
	for e := range is {
		s[e] = nil
	}
}

// AddSet adds elements to the set from another set.
func (s Set[T]) AddSet(other Set[T]) {
	for e := range other {
		s[e] = nil
	}
}

// AddSlice adds elements to the set from a slice.
func (s Set[T]) AddSlice(es []T) {
	for _, e := range es {
		s[e] = nil
	}
}

// Remove removes elements from the set.
func (s Set[T]) Remove(es ...T) {
	for _, e := range es {
		delete(s, e)
	}
}

// Contains returns true if the set contains the element.
func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

// ContainsAll returns true if the set contains all the elements.
func (s Set[T]) ContainsAll(es ...T) bool {
	for _, e := range es {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the set contains any of the elements.
func (s Set[T]) ContainsAny(es ...T) bool {
	for _, e := range es {
		if s.Contains(e) {
			return true
		}
	}
	return false
}

// Slice returns the elements of the set as a slice.
func (s Set[T]) Slice() []T {
	return slices.Collect(maps.Keys(s))
}

// Iter returns an iterator over the elements of the set.
func (s Set[T]) Iter() iter.Seq[T] {
	return maps.Keys(s)
}
