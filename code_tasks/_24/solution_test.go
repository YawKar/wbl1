package _24

import (
	"math"
	"testing"
)

func TestMkPoint(t *testing.T) {
	p := Point{1, 2}
	p1 := MkPoint(1, 2)
	if p != p1 {
		t.Errorf("%v != %v", p, p1)
	}
}

func TestNewPoint(t *testing.T) {
	p := Point{1, 2}
	p1 := NewPoint(1, 2)
	if p != *p1 {
		t.Errorf("%v != %v", p, *p1)
	}
}

func TestEuclideanDistance(t *testing.T) {
	tcases := []struct {
		p1    Point
		p2    Point
		valid float64
	}{
		{Point{1, 1}, Point{0, 0}, math.Sqrt2},
		{Point{-1, -1}, Point{0, 0}, math.Sqrt2},
		{Point{2, 2}, Point{2, 0}, 2},
	}
	for i, tcase := range tcases {
		if res := tcase.p1.EuclideanDistance(&tcase.p2); res != tcase.valid {
			t.Errorf("#%d test: res != tcase.valid: %f != %f", i, res, tcase.valid)
		}
	}
}

func TestManhattanDistance(t *testing.T) {
	tcases := []struct {
		p1    Point
		p2    Point
		valid float64
	}{
		{Point{1, 1}, Point{0, 0}, 2},
		{Point{-1, -1}, Point{0, 0}, 2},
		{Point{2, 2}, Point{2, 0}, 2},
	}
	for i, tcase := range tcases {
		if res := tcase.p1.ManhattanDistance(&tcase.p2); res != tcase.valid {
			t.Errorf("#%d test: res != tcase.valid: %f != %f", i, res, tcase.valid)
		}
	}
}
