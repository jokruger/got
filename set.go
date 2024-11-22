package got

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

// Add adds an element to the set.
func (s Set[T]) Add(e T) {
	s[e] = nil
}

// Remove removes an element from the set.
func (s Set[T]) Remove(e T) {
	delete(s, e)
}

// Contains returns true if the set contains the element.
func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}
