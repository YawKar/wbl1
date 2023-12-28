package _17

import "golang.org/x/exp/constraints"

// slice should be sorted beforehand
func BinarySearch[T constraints.Ordered](target T, slice []T) (pos int, found bool) {
	l, r := 0, len(slice)
	for r-l > 1 {
		m := l + (r-l)/2
		if slice[m] <= target {
			l = m
		} else {
			r = m
		}
	}
	if slice[l] == target {
		return l, true
	}
	return 0, false
}
