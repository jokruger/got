package sliceutils

import (
	"iter"
	"slices"
)

// Chunks returns a slice of chunks of the input slice, each of given size at most.
func Chunks[T any](ts []T, size int) [][]T {
	return slices.Collect(ChunksToSeq(ts, size))
}

// ChunksToSeq returns a sequence of chunks of the input slice, each of given size at most.
func ChunksToSeq[T any](ts []T, size int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for len(ts) > 0 {
			i := min(len(ts), size)
			if !yield(ts[:i]) {
				return
			}
			ts = ts[i:]
		}
	}
}

// GoInChunks calls the given function for each chunk of the input slice, each of given size at most.
func GoInChunks[T any](ts []T, size int, f func([]T) error) error {
	for chunk := range ChunksToSeq(ts, size) {
		if err := f(chunk); err != nil {
			return err
		}
	}
	return nil
}
