package _13

import "golang.org/x/exp/constraints"

func GreedySwapPtrs[T any](a, b *T) {
	*a, *b = *b, *a
}

// It uses XOR to, firstly, take diff-mask betwenn `a` and `b` and place it in `a`,
// secondly, apply it to `b`, therefore, `b` = *original* `a` now,
// lastly, apply `b` (which is *original* `a` now) to diff-mask in `a`
// therefore, `a` is now `b`.
//
// In essence, it takes advantage of: a ⊕ b ⊕ a = a ⊕ a ⊕ b = 0 ⊕ b = b,
// where ⊕ is the XOR operator.
//
// I know that these pointers actually create themselves,
// so it's not a temporary-variables free solution.
// But I wanted to place my solution in a testable manner,
// that's why I wrapped this code in a function with pointers.
func SpecialIntegerSwap[T constraints.Integer](a, b *T) {
	*a = (*a | *b) & (^(*a & *b))
	*b = (*a | *b) & (^(*a & *b))
	*a = (*a | *b) & (^(*a & *b))
}
