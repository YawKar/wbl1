package _11

import (
	"slices"
	"testing"
)

func TestSubsetIsUnique(t *testing.T) {
	targetElement := 42
	a := make([]int, 5)
	for i := range a {
		a[i] = targetElement
	}
	b := make([]int, len(a))
	copy(b, a)

	result := MkSubset(a, b)
	if len(result) != 1 {
		t.Fatalf("subset doesn't only have 1 element: len(result) = %d", len(result))
	} else if result[0] != targetElement {
		t.Fatalf("subset has unique element, however it is invalid: %d != %d", result[0], targetElement)
	}
}

func TestWithEmptySlices(t *testing.T) {
	result := MkSubset[int](nil, nil)
	if result == nil {
		t.Fatalf("result slice is nil")
	}
	if len(result) != 0 {
		t.Fatalf("result slice's len is not 0: %d != 0", len(result))
	}
}

func TestIntersectedSlices(t *testing.T) {
	startIxs := []int{1, 200, 1000, 100000}
	lengths := make([]int, 10) // must be divisible by 4
	for i := 0; i < len(lengths); i++ {
		lengths[i] = 4 << (i + 1)
		if lengths[i] < 0 || lengths[i]%4 != 0 {
			t.Fatalf("test length is incorrect: %d", lengths[i])
		}
	}

	for _, startIx := range startIxs {
		for _, totalLen := range lengths {
			wholeSlice := make([]int, totalLen)
			for i := 0; i < totalLen; i++ {
				wholeSlice[i] = startIx + i
			}
			firstQuarter := totalLen / 4
			subSlicesLen := firstQuarter * 3 // 3/4 of the totalLen
			intersectionLen := totalLen / 2  // exactly 2/4th and 3/4th parts
			sliceA := wholeSlice[:subSlicesLen]
			sliceB := wholeSlice[firstQuarter:]
			result := MkSubset(sliceA, sliceB)
			if len(result) != intersectionLen {
				t.Fatalf("(startIx = %d; totalLen = %d) length of result != valid intersection length: %d != %d", startIx, totalLen, len(result), intersectionLen)
			}
			slices.Sort(result)
			for i := 0; i < len(result); i++ {
				valid := startIx + firstQuarter + i
				if result[i] != valid {
					t.Errorf("(startIx = %d; totalLen = %d) result[i] != valid value: %d != %d", startIx, totalLen, result[i], valid)
				}
			}
		}
	}
}

func TestNonIntersectedSlices(t *testing.T) {
	startIxs := []int{1, 1000, 10000, 100000}
	lengths := make([]int, 10) // must be divisible by 4
	for i := 0; i < len(lengths); i++ {
		lengths[i] = 4 << (i + 1)
		if lengths[i] < 0 || lengths[i]%4 != 0 {
			t.Fatalf("test length is incorrect: %d", lengths[i])
		}
	}

	for _, startIx := range startIxs {
		for _, totalLen := range lengths {
			wholeSlice := make([]int, totalLen)
			for i := 0; i < totalLen; i++ {
				wholeSlice[i] = startIx + i
			}
			half := totalLen / 2
			sliceA := wholeSlice[:half]
			sliceB := wholeSlice[half:]
			result := MkSubset(sliceA, sliceB)
			if len(result) != 0 {
				t.Fatalf("(startIx = %d; totalLen = %d) length of result != 0: %d != 0", startIx, totalLen, len(result))
			}
		}
	}
}
