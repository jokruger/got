package index

import "fmt"

type Integer interface {
	~int8 | ~int16 | ~int32 | ~uint8 | ~uint16 | ~uint32
}

// Index provides a generic, ID-based lookup structure optimized for fast, contiguous access to values of type T by integer IDs.
type Index[ID Integer, T any] struct {
	minID  ID
	maxID  ID
	values []T
	exists []bool
	count  int
}

// New creates a new Index with a range from minID to maxID (inclusive).
func New[ID Integer, T any](minID, maxID ID) *Index[ID, T] {
	if minID > maxID {
		minID, maxID = maxID, minID
	}
	size := int(maxID - minID + 1)
	return &Index[ID, T]{
		minID:  minID,
		maxID:  maxID,
		values: make([]T, size),
		exists: make([]bool, size),
	}
}

// GetMinID returns the minimum ID in the index.
func (idx *Index[ID, T]) GetMinID() ID {
	return idx.minID
}

// GetMaxID returns the maximum ID in the index.
func (idx *Index[ID, T]) GetMaxID() ID {
	return idx.maxID
}

// Set stores a value for the given ID and marks it as valid.
func (idx *Index[ID, T]) Set(id ID, value T) error {
	if id < idx.minID || id > idx.maxID {
		return fmt.Errorf("ID %v out of range [%v, %v]", id, idx.minID, idx.maxID)
	}
	offset := int(id - idx.minID)
	if !idx.exists[offset] {
		idx.count++
	}
	idx.values[offset] = value
	idx.exists[offset] = true
	return nil
}

// Get retrieves the value associated with the ID. Returns (value, true) if set, otherwise (zero, false).
func (idx *Index[ID, T]) Get(id ID) (T, bool) {
	if id < idx.minID || id > idx.maxID {
		var zero T
		return zero, false
	}
	offset := int(id - idx.minID)
	return idx.values[offset], idx.exists[offset]
}

// GetPtr retrieves pointer to the value associated with the ID. Returns *value if set, otherwise nil.
func (idx *Index[ID, T]) GetPtr(id ID) *T {
	if id < idx.minID || id > idx.maxID {
		return nil
	}
	offset := int(id - idx.minID)
	if !idx.exists[offset] {
		return nil
	}
	return &idx.values[offset]
}

// Cap returns the total number of possible values (capacity).
func (idx *Index[ID, T]) Cap() int {
	return len(idx.values)
}

// Len returns the number of actually set elements.
func (idx *Index[ID, T]) Len() int {
	return idx.count
}

// Contains returns true if value is set for a given ID.
func (idx *Index[ID, T]) Contains(id ID) bool {
	if id < idx.minID || id > idx.maxID {
		return false
	}
	offset := int(id - idx.minID)
	return idx.exists[offset]
}
