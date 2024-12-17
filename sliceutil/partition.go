package sliceutil

import (
	"iter"

	"github.com/jokruger/got"
)

// Partition returns two slices, the first slice contains all elements that satisfy the predicate, and the second slice contains all elements that do not satisfy the predicate.
func Partition[T any](slice []T, predicate func(T) bool) (trueGroup []T, falseGroup []T) {
	trueGroup = make([]T, 0, len(slice))
	falseGroup = make([]T, 0, len(slice))
	for _, e := range slice {
		if predicate(e) {
			trueGroup = append(trueGroup, e)
		} else {
			falseGroup = append(falseGroup, e)
		}
	}
	return
}

// PartitionToSeq returns two sequences, the first sequence contains all elements that satisfy the predicate, and the second sequence contains all elements that do not satisfy the predicate.
func PartitionToSeq[T any](slice []T, predicate func(T) bool) (trueGroup iter.Seq[T], falseGroup iter.Seq[T]) {
	return FilterToSeq(slice, predicate), FilterToSeq(slice, got.Not(predicate))
}

// PartitionConsEq returns a slice of slices of consecutive equal elements.
func PartitionConsEq[T any](slice []T, eq func(a, b T) bool) [][]T {
	var result [][]T
	for r := range PartitionConsEqToSeq(slice, eq) {
		result = append(result, r)
	}
	return result
}

// PartitionConsEqToSeq returns a sequence of slices of consecutive equal elements.
func PartitionConsEqToSeq[T any](slice []T, eq func(a, b T) bool) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if len(slice) == 0 {
			return
		}
		start := 0
		for i := 1; i < len(slice); i++ {
			if !eq(slice[i], slice[i-1]) {
				if !yield(slice[start:i]) {
					return
				}
				start = i
			}
		}
		yield(slice[start:])
	}
}
