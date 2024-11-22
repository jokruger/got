package got

import "iter"

// Batches returns an iterator that yields batches of tasks.
func Batches[T any](tasks []T, batchSize int) iter.Seq[[]T] {
	return func(yield func([]T) bool) {
		for len(tasks) > 0 {
			i := min(len(tasks), batchSize)
			if !yield(tasks[:i]) {
				return
			}
			tasks = tasks[i:]
		}
	}
}

// GoInBatches runs a function in batches of tasks.
func GoInBatches[T any](tasks []T, batchSize int, f func([]T) error) error {
	for batch := range Batches(tasks, batchSize) {
		if err := f(batch); err != nil {
			return err
		}
	}
	return nil
}
