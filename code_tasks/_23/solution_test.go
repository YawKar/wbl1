package _23

import "testing"

func TestEmpty(t *testing.T) {
	s := make([]int, 0)
	result := RemoveByIndex(0, s)
	if len(result) != 0 {
		t.Fatalf("resulting array's not empty: len(result) = %d", len(result))
	}
}

func TestNil(t *testing.T) {
	var s []int
	result := RemoveByIndex(0, s)
	if len(result) != 0 {
		t.Fatalf("resulting array's not empty: len(result) = %d", len(result))
	}
}

func TestSamples(t *testing.T) {
	for i := 1; i < 100; i++ {
		s := make([]int, i)
		for j := range s {
			s[j] = j
		}
		for j := range s {
			result := RemoveByIndex(j, s)
			if len(result) != len(s)-1 {
				t.Fatalf("length of the result slice is not valid: %d != %d", len(result), len(s)-1)
			}
			if j < len(result) && result[j] != j+1 {
				t.Fatalf("%d-th position after removing doesn't contain next element: %d != %d", j, result[j], j+1)
			}
		}
	}
}
