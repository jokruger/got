package maputil

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func AppendValues[K comparable, V any](m map[K]V, dest []V) []V {
	if cap(dest)-len(dest) >= len(m) {
		for _, val := range m {
			dest = append(dest, val)
		}
		return dest
	}

	newDest := make([]V, len(dest), len(dest)+len(m))
	copy(newDest, dest)
	for _, val := range m {
		newDest = append(newDest, val)
	}
	return newDest
}
