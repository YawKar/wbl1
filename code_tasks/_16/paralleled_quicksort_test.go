package _16

import (
	"slices"
	"testing"
)

func TestParalleledReversed(t *testing.T) {
	maxLength2Pow := 10 // 2^maxLength2Pow - is the actual max length
	arrays := make([][]int, maxLength2Pow)
	sorted := make([][]int, maxLength2Pow)
	for i := 0; i < maxLength2Pow; i++ {
		length := 1 << i
		arrays[i] = make([]int, length)
		sorted[i] = make([]int, length)
		for j := 0; j < length; j++ {
			arrays[i][j] = length - j
			sorted[i][j] = j + 1
		}
	}

	for i := 0; i < maxLength2Pow; i++ {
		QuickSortParalleled(arrays[i])
		if !slices.Equal(arrays[i], sorted[i]) {
			t.Errorf("#%d array: slices are not equal after sorting:\n%v\n!=\n%v", i, arrays[i], sorted[i])
		}
	}
}

func TestParalleledEqualElems(t *testing.T) {
	array := make([]int, 100)
	QuickSortParalleled(array)
	for _, v := range array {
		if v != 0 {
			t.Errorf("non-zero element after sorting zero-only array: %d != 0", v)
		}
	}
}

func TestParalleledSmallAndEmpty(t *testing.T) {
	QuickSortParalleled[int](nil)

	array := make([]int, 0)
	QuickSortParalleled(array)
	if len(array) != 0 {
		t.Errorf("empty array is not empty after sorting: len(array) = %d", len(array))
	}
}
