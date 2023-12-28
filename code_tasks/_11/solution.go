package _11

// resulting slice is guaranteed to have unique elements (accurate to comparable)
func MkSubset[T comparable](a []T, b []T) []T {
	if a == nil || b == nil {
		return []T{}
	}
	seen := make(map[T]struct{})
	for _, v := range a {
		seen[v] = struct{}{}
	}
	out := make([]T, 0)
	for _, v := range b {
		if _, found := seen[v]; found {
			out = append(out, v)
			delete(seen, v)
		}
	}
	return out
}
