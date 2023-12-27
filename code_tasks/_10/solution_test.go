package _10

import (
	"slices"
	"testing"
)

func TestSample(t *testing.T) {
	numbers := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	validMap := map[int][]float64{
		-3: {-25.4, -27.0, -21.0}, // (-30.0; -20.0]
		1:  {13.0, 19.0, 15.5},    // [10.0; 20.0)
		2:  {24.5},                // [20.0; 30.0)
		3:  {32.5},                // [30.0; 40.0)
	}

	grouped := GroupByKeyFunc(FloatsAbsGrouper, numbers...)
	if len(grouped) != len(validMap) {
		t.Errorf("length of grouped elems doesn't equal to the length of the valid map: %d != %d", len(grouped), len(validMap))
	}
	for key, floats := range grouped {
		if valid, found := validMap[key]; !found {
			t.Errorf("%d key wasn't found in valid map but exists in grouped", key)
		} else {
			slices.Sort(floats)
			slices.Sort(valid)
			if !slices.Equal(floats, valid) {
				t.Errorf("grouped elements under %d key are not the same as in valid map:\n%v\n!=\n%v", key, floats, valid)
			}
		}
	}
}

func TestFloatAbsGrouperFuncAndToBounds(t *testing.T) {
	/* Note:
	solution_test.go:27: -1.010000 ∈ (-10.000000; 0.000000], delta = -0.010000
	solution_test.go:27: -1.000000 ∈ (-10.000000; 0.000000], delta = 0.000000
	solution_test.go:27: -0.990000 ∈ (-10.000000; 0.000000], delta = 0.010000
	solution_test.go:27: -0.010000 ∈ (-10.000000; 0.000000], delta = -0.010000
	solution_test.go:24: 0.000000 ∈ [0.000000; 10.000000), delta = 0.000000
	solution_test.go:24: 0.010000 ∈ [0.000000; 10.000000), delta = 0.010000
	*/
	for _, eps := range []float64{1e-1, 1e-2, 1e-3, 1e-4, 1e-5, 1e-6, 1e-7, 1e-8, 1e-9} {
		epss := []float64{-eps, 0, eps}
		for f := -100.0; f <= 100.0; f += 1 {
			for _, dt := range epss {
				fVal := f + dt
				groupKey := FloatsAbsGrouper(fVal)
				left, right, leftInclusive := GroupKeyToBounds(groupKey)
				if leftInclusive && left <= fVal && fVal < right {
					t.Logf("%f ∈ [%f; %f), delta = %f", fVal, left, right, dt)
					continue
				} else if !leftInclusive && left < fVal && fVal <= right {
					t.Logf("%f ∈ (%f; %f], delta = %f", fVal, left, right, dt)
					continue
				} else {
					t.Fatalf("initial %f value is not in bounds after grouping", fVal)
				}
			}
		}
	}
}
