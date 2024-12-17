package got

import (
	"slices"
	"testing"

	gmaps "github.com/jokruger/got/maps"
)

func TestMaps(t *testing.T) {
	t.Run("keys/values", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}

		ks := gmaps.Keys(m)
		slices.Sort(ks)
		if len(ks) != 3 || ks[0] != 1 || ks[1] != 2 || ks[2] != 3 {
			t.Errorf("Keys() = %v; want [1 2 3]", ks)
		}

		vs := gmaps.Values(m)
		slices.Sort(vs)
		if len(vs) != 3 || vs[0] != "a" || vs[1] != "b" || vs[2] != "c" {
			t.Errorf("Values() = %v; want [a b c]", vs)
		}
	})

	t.Run("keys/values to", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}
		ks := make([]int, 10)
		vs := make([]string, 10)

		ks = gmaps.KeysTo(m, ks[:0])
		slices.Sort(ks)
		if len(ks) != 3 || ks[0] != 1 || ks[1] != 2 || ks[2] != 3 {
			t.Errorf("Keys() = %v; want [1 2 3]", ks)
		}

		vs = gmaps.ValuesTo(m, vs[:0])
		slices.Sort(vs)
		if len(vs) != 3 || vs[0] != "a" || vs[1] != "b" || vs[2] != "c" {
			t.Errorf("Values() = %v; want [a b c]", vs)
		}
	})
}
