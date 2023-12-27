package _08

func WithIthBitSet(v int64, i int) int64 {
	if i < 0 || i > 63 {
		return v
	}
	return v | (1 << i)
}

func WithIthBitUnset(v int64, i int) int64 {
	if i < 0 || i > 63 {
		return v
	}
	return v & (^(1 << i))
}

func SetIthBit(v *int64, i int) {
	if i < 0 || i > 63 {
		return
	}
	*v |= 1 << i
}

func UnsetIthBit(v *int64, i int) {
	if i < 0 || i > 63 {
		return
	}
	*v &= ^(1 << i)
}
