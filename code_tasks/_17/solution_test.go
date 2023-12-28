package _17

import (
	"math/rand"
	"testing"
)

const (
	arrayLen = 1 << 10
	randSeed = 42
)

var array []int

func init() {
	array = make([]int, arrayLen)
	for i := 0; i < arrayLen; i++ {
		array[i] = i
	}
}

func TestSample(t *testing.T) {
	rand := rand.New(rand.NewSource(randSeed))
	for i := 0; i < 100; i++ {
		target := rand.Intn(arrayLen)
		if pos, found := BinarySearch(target, array); !found {
			t.Errorf("%d wasn't found in the test array", target)
		} else if pos != target {
			t.Errorf("found position differ from valid: %d != %d", pos, target)
		}
	}
}

func TestOutOfBounds(t *testing.T) {
	rand := rand.New(rand.NewSource(randSeed))
	for i := 0; i < 100; i++ {
		target := rand.Intn(arrayLen) + arrayLen
		if rand.Float32() > 0.5 {
			target *= -1
		}
		if pos, found := BinarySearch(target, array); found {
			t.Errorf("%d was found on pos %d, although it didn't exist in the test array", target, pos)
		}
	}
}
