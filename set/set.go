package set

import (
	"iter"
	"maps"
	"slices"

	"github.com/jokruger/got/sequtil"
)

// Set is a set of elements.
type Set[T comparable] map[T]struct{}

// New creates a new set with the given options.
func New[T comparable](options ...Option) Set[T] {
	opts := Options{}
	for _, o := range options {
		o(&opts)
	}

	var s Set[T]

	if opts.capacity > 0 {
		s = make(Set[T], opts.capacity)
	} else {
		s = make(Set[T])
	}

	return s
}

// NewFromElements creates a new set with the given elements.
func NewFromElements[T comparable](elements ...T) Set[T] {
	s := make(Set[T])
	for _, e := range elements {
		s[e] = struct{}{}
	}
	return s
}

// NewFromSlice creates a new set with the elements of the slice.
func NewFromSlice[T comparable](source []T) Set[T] {
	s := make(Set[T])
	for _, e := range source {
		s[e] = struct{}{}
	}
	return s
}

// NewFromSeq creates a new set with the elements of the sequence.
func NewFromSeq[T comparable](source iter.Seq[T]) Set[T] {
	s := make(Set[T])
	for e := range source {
		s[e] = struct{}{}
	}
	return s
}

// NewFromSet creates a new set with the elements of the other set.
func NewFromSet[T comparable](source Set[T]) Set[T] {
	s := make(Set[T])
	for e := range source {
		s[e] = struct{}{}
	}
	return s
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// IsEmpty returns true if the set is empty.
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}

// Clear removes all elements from the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) Clear() Set[T] {
	clear(s)
	return s
}

// Clone returns a copy of the set.
func (s Set[T]) Clone() Set[T] {
	return NewFromSet(s)
}

// Equal returns true if the set is equal to another set.
func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for e := range s {
		if _, ok := other[e]; !ok {
			return false
		}
	}
	return true
}

// Add adds the element to the set. Returns true if the element was added. Returns false if the element was already in the set.
func (s Set[T]) Add(element T) bool {
	if _, ok := s[element]; ok {
		return false
	}
	s[element] = struct{}{}
	return true
}

// AddMany adds the elements to the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) AddMany(elements ...T) Set[T] {
	for _, e := range elements {
		s[e] = struct{}{}
	}
	return s
}

// AddSlice adds the elements of the slice to the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) AddSlice(source []T) Set[T] {
	for _, e := range source {
		s[e] = struct{}{}
	}
	return s
}

// AddSeq adds the elements of the sequence to the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) AddSeq(source iter.Seq[T]) Set[T] {
	for e := range source {
		s[e] = struct{}{}
	}
	return s
}

// AddSet adds the elements of the other set to the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) AddSet(source Set[T]) Set[T] {
	for e := range source {
		s[e] = struct{}{}
	}
	return s
}

// Remove removes the element from the set. Returns true if the element was removed. Returns false if the element was not in the set.
func (s Set[T]) Remove(element T) bool {
	if _, ok := s[element]; !ok {
		return false
	}
	delete(s, element)
	return true
}

// RemoveMany removes the elements from the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) RemoveMany(elements ...T) Set[T] {
	for _, e := range elements {
		delete(s, e)
	}
	return s
}

// RemoveSlice removes the elements of the slice from the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) RemoveSlice(source []T) Set[T] {
	for _, e := range source {
		delete(s, e)
	}
	return s
}

// RemoveSeq removes the elements of the sequence from the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) RemoveSeq(source iter.Seq[T]) Set[T] {
	for e := range source {
		delete(s, e)
	}
	return s
}

// RemoveSet removes the elements of the other set from the set. Modifies the set in place. Returns the set itself.
func (s Set[T]) RemoveSet(source Set[T]) Set[T] {
	for e := range source {
		delete(s, e)
	}
	return s
}

// Contains returns true if the set contains all elements.
func (s Set[T]) Contains(elements ...T) bool {
	for _, e := range elements {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// ContainsSlice returns true if the set contains all elements of the slice.
func (s Set[T]) ContainsSlice(source []T) bool {
	for _, e := range source {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// ContainsSeq returns true if the set contains all elements of the sequence.
func (s Set[T]) ContainsSeq(source iter.Seq[T]) bool {
	for e := range source {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// ContainsSet returns true if the set contains all elements of the other set.
func (s Set[T]) ContainsSet(other Set[T]) bool {
	for e := range other {
		if _, ok := s[e]; !ok {
			return false
		}
	}
	return true
}

// ContainsAny returns true if the set contains any of the elements.
func (s Set[T]) ContainsAny(elements ...T) bool {
	for _, e := range elements {
		if _, ok := s[e]; ok {
			return true
		}
	}
	return false
}

// ContainsAnyFromSlice returns true if the set contains any of the elements of the slice.
func (s Set[T]) ContainsAnyFromSlice(source []T) bool {
	for _, e := range source {
		if _, ok := s[e]; ok {
			return true
		}
	}
	return false
}

// ContainsAnyFromSeq returns true if the set contains any of the elements of the sequence.
func (s Set[T]) ContainsAnyFromSeq(source iter.Seq[T]) bool {
	for e := range source {
		if _, ok := s[e]; ok {
			return true
		}
	}
	return false
}

// ContainsAnyFromSet returns true if the set contains any of the elements of the other set.
func (s Set[T]) ContainsAnyFromSet(other Set[T]) bool {
	for e := range other {
		if _, ok := s[e]; ok {
			return true
		}
	}
	return false
}

// IsSubsetOf returns true if the set is a subset of the other set.
func (s Set[T]) IsSubsetOf(other Set[T]) bool {
	return other.ContainsSet(s)
}

// IsSupersetOf returns true if the set is a superset of the other set.
func (s Set[T]) IsSupersetOf(other Set[T]) bool {
	return s.ContainsSet(other)
}

// IsProperSubsetOf returns true if the set is a proper subset of the other set.
func (s Set[T]) IsProperSubsetOf(other Set[T]) bool {
	return s.IsSubsetOf(other) && !s.Equal(other)
}

// IsProperSupersetOf returns true if the set is a proper superset of the other set.
func (s Set[T]) IsProperSupersetOf(other Set[T]) bool {
	return s.IsSupersetOf(other) && !s.Equal(other)
}

// Diff returns the difference between the set and another set.
func (s Set[T]) Diff(other Set[T]) Set[T] {
	diff := NewFromSet(s)
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
	union := NewFromSet(s)
	union.AddSet(other)
	return union
}

func (s Set[T]) ToSlice() []T {
	return slices.Collect(maps.Keys(s))
}

func (s Set[T]) ToSeq() iter.Seq[T] {
	return maps.Keys(s)
}

// All returns true if all elements in set satisfy the predicate f.
// If set is empty, All returns true.
func (s Set[T]) All(f func(T) bool) bool {
	return sequtil.All(s.ToSeq(), f)
}

// Any returns true if any element in set satisfies the predicate, false otherwise.
// If set is empty, Any returns false.
func (s Set[T]) Any(f func(T) bool) bool {
	return sequtil.Any(s.ToSeq(), f)
}

// Count returns the number of elements in set that satisfy the predicate f.
func (s Set[T]) Count(f func(T) bool) int {
	return sequtil.Count(s.ToSeq(), f)
}

// Filter returns a new set containing only the elements that satisfy the predicate f.
func (s Set[T]) Filter(f func(T) bool) Set[T] {
	res := New[T]()
	for e := range s {
		if f(e) {
			res[e] = struct{}{}
		}
	}
	return res
}
