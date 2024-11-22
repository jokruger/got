package got

import "testing"

func TestGoInBatches(t *testing.T) {
	t.Run("GoInBatches1", func(t *testing.T) {
		tasks := Range(999)
		batches := make([][]int, 0)
		res := make([]int, 0)

		if err := GoInBatches(tasks, 100, func(batch []int) error {
			batches = append(batches, batch)
			return nil
		}); err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		if len(batches) != 10 {
			t.Errorf("Expected 10 batches, got %d", len(batches))
		}

		if len(batches[0]) != 100 {
			t.Errorf("Expected 100 items in first batch, got %d", len(batches[0]))
		}

		if len(batches[9]) != 99 {
			t.Errorf("Expected 99 items in last batch, got %d", len(batches[9]))
		}

		for _, batch := range batches {
			res = append(res, batch...)
		}

		if len(res) != 999 {
			t.Errorf("Expected 1000 items in all batches, got %d", len(res))
		}

		for i := 0; i < 999; i++ {
			if res[i] != i {
				t.Errorf("Expected %d, got %d", i, res[i])
			}
		}
	})

	t.Run("GoInBatches2", func(t *testing.T) {
		tasks := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		sums := make([]int, 0)

		if err := GoInBatches(tasks, 3, func(batch []int) error {
			sum := 0
			for _, t := range batch {
				sum += t
			}
			sums = append(sums, sum)
			return nil
		}); err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		if len(sums) != 4 {
			t.Errorf("Expected 4 sums, got %d", len(sums))
		}

		if sums[0] != 3 {
			t.Errorf("Expected 3, got %d", sums[0])
		}

		if sums[1] != 12 {
			t.Errorf("Expected 12, got %d", sums[1])
		}

		if sums[2] != 21 {
			t.Errorf("Expected 21, got %d", sums[2])
		}

		if sums[3] != 9 {
			t.Errorf("Expected 9, got %d", sums[3])
		}
	})
}
