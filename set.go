package got

import (
	"iter"
	"maps"
	"slices"
)

// Set is a set of elements of type T.
type Set[T comparable] map[T]interface{}

// NewSet creates a new set. The set can be initialized with other sets.
func NewSet[T comparable](sources ...Set[T]) Set[T] {
	set := make(Set[T])
	for _, s := range sources {
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

// NewSetFromArgs creates a new set from arguments.
func NewSetFromArgs[T comparable](es ...T) Set[T] {
	set := make(Set[T])
	for _, e := range es {
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

// AddIter adds elements to the set from an iterator.
func (s Set[T]) AddIter(is iter.Seq[T]) {
	for e := range is {
		s[e] = nil
	}
}

// Remove removes elements from the set.
func (s Set[T]) Remove(es ...T) {
	for _, e := range es {
		delete(s, e)
	}
}

// RemoveSet removes elements from the set present in another set.
func (s Set[T]) RemoveSet(other Set[T]) {
	for e := range other {
		delete(s, e)
	}
}

// RemoveSlice removes elements from the set present in a slice.
func (s Set[T]) RemoveSlice(es []T) {
	for _, e := range es {
		delete(s, e)
	}
}

// RemoveIter removes elements from the set present in a sequence.
func (s Set[T]) RemoveIter(is iter.Seq[T]) {
	for e := range is {
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

// ContainsAny returns true if the set contains all elements of other set.
func (s Set[T]) ContainsAllFromSet(other Set[T]) bool {
	for e := range other {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the set contains all elements of slice.
func (s Set[T]) ContainsAllFromSlice(es []T) bool {
	for _, e := range es {
		if !s.Contains(e) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the set contains all elements of sequence.
func (s Set[T]) ContainsAllFromIter(is iter.Seq[T]) bool {
	for e := range is {
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

// ContainsAnyFromSet returns true if the set contains any of the elements of other set.
func (s Set[T]) ContainsAnyFromSet(other Set[T]) bool {
	for e := range other {
		if s.Contains(e) {
			return true
		}
	}
	return false
}

// ContainsAnyFromSlice returns true if the set contains any of the elements of slice.
func (s Set[T]) ContainsAnyFromSlice(es []T) bool {
	for _, e := range es {
		if s.Contains(e) {
			return true
		}
	}
	return false
}

// ContainsAnyFromIter returns true if the set contains any of the elements of sequence.
func (s Set[T]) ContainsAnyFromIter(is iter.Seq[T]) bool {
	for e := range is {
		if s.Contains(e) {
			return true
		}
	}
	return false
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Clear removes all elements from the set.
func (s Set[T]) Clear() {
	clear(s)
}

// Clone returns a copy of the set.
func (s Set[T]) Clone() Set[T] {
	return NewSet(s)
}

// Equal returns true if the set is equal to another set.
func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	return s.ContainsAllFromSet(other)
}

// Slice returns the elements of the set as a slice.
func (s Set[T]) Slice() []T {
	return slices.Collect(maps.Keys(s))
}

// Iter returns an iterator over the elements of the set.
func (s Set[T]) Iter() iter.Seq[T] {
	return maps.Keys(s)
}

// IsEmpty returns true if the set is empty.
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// IsSubsetOf returns true if the set is a subset of another set.
func (s Set[T]) IsSubsetOf(other Set[T]) bool {
	return other.ContainsAllFromSet(s)
}

// IsSupersetOf returns true if the set is a superset of another set.
func (s Set[T]) IsSupersetOf(other Set[T]) bool {
	return s.ContainsAllFromSet(other)
}

// IsProperSubsetOf returns true if the set is a proper subset of another set.
func (s Set[T]) IsProperSubsetOf(other Set[T]) bool {
	return s.IsSubsetOf(other) && !s.Equal(other)
}

// IsProperSupersetOf returns true if the set is a proper superset of another set.
func (s Set[T]) IsProperSupersetOf(other Set[T]) bool {
	return s.IsSupersetOf(other) && !s.Equal(other)
}

// Diff returns the difference between the set and another set.
func (s Set[T]) Diff(other Set[T]) Set[T] {
	diff := NewSet(s)
	diff.RemoveSet(other)
	return diff
}

// SymmetricDiff returns the symmetric difference between the set and another set.
func (s Set[T]) SymmetricDiff(other Set[T]) Set[T] {
	t := s.Diff(other)
	t.AddSet(other.Diff(s))
	return t
}

// Union returns the union of the set and another set.
func (s Set[T]) Union(other Set[T]) Set[T] {
	union := NewSet(s)
	union.AddSet(other)
	return union
}
