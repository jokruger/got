package got

import "iter"

// Filter returns a new slice containing only the elements of slice 's' that satisfy the predicate f.
func Filter[T any](s []T, f func(T) bool) []T {
	res := make([]T, 0, len(s))
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

// FilterSeq returns a new slice containing only the elements of sequence 's' that satisfy the predicate f.
func FilterSeq[T any](s iter.Seq[T], f func(T) bool) []T {
	var res []T
	for i := range s {
		if f(i) {
			res = append(res, i)
		}
	}
	return res
}

// FilterToSeq returns a new sequence containing only the elements of slice 's' that satisfy the predicate f.
func FilterToSeq[T any](s []T, f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// FilterSeqToSeq returns a new sequence containing only the elements of sequence 's' that satisfy the predicate f.
func FilterSeqToSeq[T any](s iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range s {
			if f(i) {
				if !yield(i) {
					return
				}
			}
		}
	}
}
