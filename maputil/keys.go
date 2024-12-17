package maputil

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func AppendKeys[K comparable, V any](m map[K]V, dest []K) []K {
	if cap(dest)-len(dest) >= len(m) {
		for key := range m {
			dest = append(dest, key)
		}
		return dest
	}

	newDest := make([]K, len(dest), len(dest)+len(m))
	copy(newDest, dest)
	for key := range m {
		newDest = append(newDest, key)
	}
	return newDest
}
