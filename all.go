package got

import "iter"

func All[T any](s []T, f func(T) bool) bool {
	if len(s) > 0 {
		for _, e := range s {
			if !f(e) {
				return false
			}
		}
	}
	return true
}

func AllSeq[T any](s iter.Seq[T], f func(T) bool) bool {
	if s != nil {
		for e := range s {
			if !f(e) {
				return false
			}
		}
	}
	return true
}
