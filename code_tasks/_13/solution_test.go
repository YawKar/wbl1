package _13

import (
	"math"
	"testing"
)

func TestSwaps(t *testing.T) {
	intSamples := []struct{ a, b int }{
		{1, 2},
		{4, 5},
		{41, 50},
		{-1, -41},
		{math.MaxInt32, math.MinInt32},
	}
	stringSamples := []struct{ a, b string }{
		{"cat", "tree"},
		{"dog", "anmial"},
		{"who is", "Lila?"},
	}

	for _, ipair := range intSamples {
		a, b := ipair.a, ipair.b
		GreedySwapPtrs(&a, &b)
		if a != ipair.b {
			t.Errorf("a != ipair.b after swap: %d != %d", a, ipair.b)
		}
		if b != ipair.a {
			t.Errorf("b != ipair.a after swap: %d != %d", b, ipair.a)
		}
		if t.Failed() {
			t.FailNow()
		}
	}

	for _, spair := range stringSamples {
		a, b := spair.a, spair.b
		GreedySwapPtrs(&a, &b)
		if a != spair.b {
			t.Errorf("a != spair.b after swap: %s != %s", a, spair.b)
		}
		if b != spair.a {
			t.Errorf("b != spair.a after swap: %s != %s", b, spair.a)
		}
		if t.Failed() {
			t.FailNow()
		}
	}
}

func TestSpecialIntegerSwap(t *testing.T) {
	intSamples := []struct{ a, b int }{
		{1, 2},
		{4, 5},
		{41, 50},
		{-1, -41},
		{math.MaxInt32, math.MinInt32},
	}

	uintSamples := []struct{ a, b uint }{
		{1, 2},
		{4, 5},
		{41, 50},
		{math.MaxUint32, 0},
	}

	for _, ipair := range intSamples {
		a, b := ipair.a, ipair.b
		SpecialIntegerSwap(&a, &b)
		if a != ipair.b {
			t.Errorf("a != ipair.b after swap: %d != %d", a, ipair.b)
		}
		if b != ipair.a {
			t.Errorf("b != ipair.a after swap: %d != %d", b, ipair.a)
		}
		if t.Failed() {
			t.FailNow()
		}
	}

	for _, uipair := range uintSamples {
		a, b := uipair.a, uipair.b
		SpecialIntegerSwap(&a, &b)
		if a != uipair.b {
			t.Errorf("a != ipair.b after swap: %d != %d", a, uipair.b)
		}
		if b != uipair.a {
			t.Errorf("b != ipair.a after swap: %d != %d", b, uipair.a)
		}
		if t.Failed() {
			t.FailNow()
		}
	}
}
