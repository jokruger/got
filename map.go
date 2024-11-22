package got

import "iter"

// Map applies a function to each element of a slice and returns a new slice.
func Map[T any, U any](is []T, f func(T) U) []U {
	us := make([]U, len(is))
	for i, v := range is {
		us[i] = f(v)
	}
	return us
}

// MapIter applies a function to each element of a sequence and returns a new slice.
func MapIter[T any, U any](is iter.Seq[T], f func(T) U) []U {
	var us []U
	for i := range is {
		us = append(us, f(i))
	}
	return us
}

// MapUnique applies a function to each element of a slice and returns a new slice with unique elements.
func MapUnique[T any, U comparable](is []T, f func(T) U) []U {
	m := make(map[U]interface{})
	r := make([]U, 0)
	for _, v := range is {
		u := f(v)
		if _, ok := m[u]; !ok {
			m[u] = nil
			r = append(r, u)
		}
	}
	return r
}

// MapUniqueIter applies a function to each element of a sequence and returns a new slice with unique elements.
func MapUniqueIter[T any, U comparable](is iter.Seq[T], f func(T) U) []U {
	m := make(map[U]interface{})
	r := make([]U, 0)
	for i := range is {
		u := f(i)
		if _, ok := m[u]; !ok {
			m[u] = nil
			r = append(r, u)
		}
	}
	return r
}
