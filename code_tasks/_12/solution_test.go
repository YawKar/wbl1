package _12

import (
	"slices"
	"testing"
)

func TestSample(t *testing.T) {
	sample := []string{"cat", "cat", "dog", "cat", "tree"}
	valid := []string{"cat", "dog", "tree"}
	slices.Sort(valid)

	set := MkSet(sample...)
	slices.Sort(set)

	if !slices.Equal(set, valid) {
		t.Fatalf("\n%v\n!=\n%v", set, valid)
	}
}

func TestEmpty(t *testing.T) {
	set := MkSet[int](nil...)
	if set == nil {
		t.Fatalf("set is nil")
	}
	if len(set) != 0 {
		t.Fatalf("set should be empty if no items were provided but has length: %d", len(set))
	}
}

func TestRepeating(t *testing.T) {
	numbers := make([]int, 10)
	set := MkSet(numbers...)
	if len(set) != 1 {
		t.Fatalf("set doesn't contain 1 element: %d != 1", len(set))
	}
	if set[0] != 0 {
		t.Fatalf("set contain nonsense: %d != 0", set[0])
	}
}
