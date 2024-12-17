package maputils

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func KeysTo[K comparable, V any](m map[K]V, dest []K) []K {
	for key := range m {
		dest = append(dest, key)
	}
	return dest
}
