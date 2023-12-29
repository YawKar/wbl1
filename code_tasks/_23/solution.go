package _23

func RemoveByIndex[T any](ix int, s []T) []T {
	if ix < 0 || len(s) <= ix {
		out := make([]T, len(s))
		copy(out, s)
		return out
	}
	out := make([]T, len(s)-1)
	copy(out, s[:ix])
	copy(out[ix:], s[ix+1:])
	return out
}
