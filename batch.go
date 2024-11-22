package got

// GoInBatches runs a function in batches of tasks.
func GoInBatches[T any](tasks []T, batchSize int, f func([]T) error) error {
	for len(tasks) > 0 {
		i := min(len(tasks), batchSize)
		if err := f(tasks[:i]); err != nil {
			return err
		}
		tasks = tasks[i:]
	}
	return nil
}
