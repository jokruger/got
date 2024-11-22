package got

import "iter"

// Partitions returns a sequence of partitions of the input items based on 'eq' comparator.
func Partitions[T any](items []T, eq func(a, b T) bool) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		sz := len(items)
		if sz == 0 {
			return
		}
		start := 0
		for i := 1; i < sz; i++ {
			if !eq(items[i], items[i-1]) {
				if !yield(items[start:i]) {
					return
				}
				start = i
			}
		}
		yield(items[start:])
	}
}
