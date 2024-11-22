package got

import "iter"

// Set is a set of elements of type T.
type Set[T comparable] map[T]interface{}

// NewSet creates a new set.
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
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

// Add adds elements to the set.
func (s Set[T]) Add(es ...T) {
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
