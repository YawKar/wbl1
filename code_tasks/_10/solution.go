package _10

func GroupByKeyFunc[K comparable, T any](f func(T) K, items ...T) map[K][]T {
	m := make(map[K][]T)
	for _, v := range items {
		key := f(v)
		var collected []T
		var found bool
		if collected, found = m[key]; !found {
			collected = make([]T, 0, 1)
		}
		collected = append(collected, v)
		m[key] = collected
	}
	return m
}

// groups: (-30.0; -20.0], (-20.0; -10.0], (-10.0, 0], [0; 10.0), [10.0, 20.0)
// all indices are shifted, so [0; 10.0) is 0th and (-10.0, 0] is 1st
func FloatsAbsGrouper(f float64) int {
	if f < 0 {
		return int(f/10.0) - 1
	} else {
		return int(f / 10.0)
	}
}

// get bounds for given key
// bounds may be of two kinds: (left, right], [left, right)
// if `leftInclusive` is `true` then it's the second type
func GroupKeyToBounds(key int) (left, right float64, leftInclusive bool) {
	left, right = float64(key*10), float64((key+1)*10)
	if key >= 0 {
		leftInclusive = true
	}
	return
}
