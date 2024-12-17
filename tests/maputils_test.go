package tests

import (
	"slices"
	"testing"

	"github.com/jokruger/got/maputil"
)

func TestMaps(t *testing.T) {
	t.Run("keys/values", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}

		ks := maputil.Keys(m)
		slices.Sort(ks)
		if len(ks) != 3 || ks[0] != 1 || ks[1] != 2 || ks[2] != 3 {
			t.Errorf("Keys() = %v; want [1 2 3]", ks)
		}

		vs := maputil.Values(m)
		slices.Sort(vs)
		if len(vs) != 3 || vs[0] != "a" || vs[1] != "b" || vs[2] != "c" {
			t.Errorf("Values() = %v; want [a b c]", vs)
		}
	})

	t.Run("keys/values to", func(t *testing.T) {
		m := map[int]string{1: "a", 2: "b", 3: "c"}
		ks := make([]int, 10)
		vs := make([]string, 10)

		ks = maputil.AppendKeys(m, ks[:0])
		slices.Sort(ks)
		if len(ks) != 3 || ks[0] != 1 || ks[1] != 2 || ks[2] != 3 {
			t.Errorf("Keys() = %v; want [1 2 3]", ks)
		}

		ks = maputil.AppendKeys(m, nil)
		slices.Sort(ks)
		if len(ks) != 3 || ks[0] != 1 || ks[1] != 2 || ks[2] != 3 {
			t.Errorf("Keys() = %v; want [1 2 3]", ks)
		}

		vs = maputil.AppendValues(m, vs[:0])
		slices.Sort(vs)
		if len(vs) != 3 || vs[0] != "a" || vs[1] != "b" || vs[2] != "c" {
			t.Errorf("Values() = %v; want [a b c]", vs)
		}

		vs = maputil.AppendValues(m, nil)
		slices.Sort(vs)
		if len(vs) != 3 || vs[0] != "a" || vs[1] != "b" || vs[2] != "c" {
			t.Errorf("Values() = %v; want [a b c]", vs)
		}
	})
}
