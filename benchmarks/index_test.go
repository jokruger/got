package benchmarks

import (
	"testing"

	"github.com/jokruger/got/index"
)

func BenchmarkMapIndex(b *testing.B) {
	s := uint16(1)
	e := uint16(1000)
	idx := make(map[uint16]string, 1001)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for id := s; id <= e; id++ {
			idx[id] = "123"
		}
		for id := 0; id < 1100; id++ {
			_, _ = idx[uint16(id)]
		}
	}
}

func BenchmarkIndex(b *testing.B) {
	s := uint16(1)
	e := uint16(1000)
	idx := index.New[uint16, string](s, e)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for id := s; id <= e; id++ {
			idx.Set(id, "123")
		}
		for id := 0; id < 1100; id++ {
			_, _ = idx.Get(uint16(id))
		}
	}
}
