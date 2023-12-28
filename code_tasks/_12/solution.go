package _12

// essentially, returns slice of unique elements from the given ones
func MkSet[T comparable](items ...T) []T {
	seen := make(map[T]struct{})
	for _, v := range items {
		seen[v] = struct{}{}
	}
	out := make([]T, 0, len(seen))
	for k := range seen {
		out = append(out, k)
	}
	return out
}
