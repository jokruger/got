package basic

// Range returns a slice of integers from 0 to n-1.
func Range(n int) []int {
	if n <= 0 {
		return nil
	}
	is := make([]int, n)
	for i := range is {
		is[i] = i
	}
	return is
}
