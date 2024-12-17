package sequtil

import (
	"iter"

	"github.com/jokruger/got"
)

// Partition returns two sequences, the first sequence contains all elements that satisfy the predicate, and the second sequence contains all elements that do not satisfy the predicate.
func Partition[T any](s iter.Seq[T], predicate func(T) bool) (trueGroup iter.Seq[T], falseGroup iter.Seq[T]) {
	return Filter(s, predicate), Filter(s, got.Not(predicate))
}

// PartitionToSlice returns two slices, the first slice contains all elements that satisfy the predicate, and the second slice contains all elements that do not satisfy the predicate.
func PartitionToSlice[T any](s iter.Seq[T], predicate func(T) bool) (trueGroup []T, falseGroup []T) {
	trueGroup = make([]T, 0)
	falseGroup = make([]T, 0)
	for e := range s {
		if predicate(e) {
			trueGroup = append(trueGroup, e)
		} else {
			falseGroup = append(falseGroup, e)
		}
	}
	return
}

// PartitionConsEq returns a sequence of slices of consecutive equal elements.
func PartitionConsEq[T any](s iter.Seq[T], eq func(a, b T) bool) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		if s == nil {
			return
		}
		group := []T{}
		for e := range s {
			l := len(group)
			if l == 0 || eq(group[l-1], e) {
				group = append(group, e)
			} else {
				if !yield(group) {
					return
				}
				group = []T{e}
			}
		}
		if len(group) > 0 {
			yield(group)
		}
	}
}

// PartitionConsEqToSlice returns a slice of slices of consecutive equal elements.
func PartitionConsEqToSlice[T any](s iter.Seq[T], eq func(a, b T) bool) [][]T {
	var result [][]T
	for r := range PartitionConsEq(s, eq) {
		result = append(result, r)
	}
	return result
}
