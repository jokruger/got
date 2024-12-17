package maps

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func ValuesTo[K comparable, V any](m map[K]V, dest []V) []V {
	for _, value := range m {
		dest = append(dest, value)
	}
	return dest
}
