package tests

import (
	"testing"

	"github.com/jokruger/got/basicutil"
	"github.com/jokruger/got/sliceutil"
)

func TestGoInChunks(t *testing.T) {
	t.Run("GoInChunks1", func(t *testing.T) {
		tasks := basicutil.Range(999)
		chunks := make([][]int, 0)
		res := make([]int, 0)

		if err := sliceutil.GoInChunks(tasks, 100, func(chunk []int) error {
			chunks = append(chunks, chunk)
			return nil
		}); err != nil {
			t.Errorf("error: %s", err.Error())
		}

		if len(chunks) != 10 {
			t.Errorf("expected 10 chunks, got %d", len(chunks))
		}

		if len(chunks[0]) != 100 {
			t.Errorf("expected 100 items in first chunk, got %d", len(chunks[0]))
		}

		if len(chunks[9]) != 99 {
			t.Errorf("expected 99 items in last chunk, got %d", len(chunks[9]))
		}

		for _, chunk := range chunks {
			res = append(res, chunk...)
		}

		if len(res) != 999 {
			t.Errorf("expected 1000 items in all chunks, got %d", len(res))
		}

		for i := 0; i < 999; i++ {
			if res[i] != i {
				t.Errorf("expected %d, got %d", i, res[i])
			}
		}
	})

	t.Run("GoInChunks2", func(t *testing.T) {
		tasks := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		sums := make([]int, 0)

		if err := sliceutil.GoInChunks(tasks, 3, func(chunk []int) error {
			sum := 0
			for _, t := range chunk {
				sum += t
			}
			sums = append(sums, sum)
			return nil
		}); err != nil {
			t.Errorf("error: %s", err.Error())
		}

		if len(sums) != 4 {
			t.Errorf("expected 4 sums, got %d", len(sums))
		}

		if sums[0] != 3 {
			t.Errorf("expected 3, got %d", sums[0])
		}

		if sums[1] != 12 {
			t.Errorf("expected 12, got %d", sums[1])
		}

		if sums[2] != 21 {
			t.Errorf("expected 21, got %d", sums[2])
		}

		if sums[3] != 9 {
			t.Errorf("expected 9, got %d", sums[3])
		}
	})
}
