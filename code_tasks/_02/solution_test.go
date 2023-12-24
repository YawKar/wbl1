package _02

import (
	"bytes"
	"fmt"
	"slices"
	"sort"
	"sync"
	"testing"
)

var arrays = [][]int64{
	{1},
	{2, 4, 6, 8, 10},
	{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
}

var outputs = make([][]int64, 0, len(arrays))

func init() {
	length := 100
	longTestCase := make([]int64, 0, length)
	for i := 0; i < length; i++ {
		longTestCase = append(longTestCase, int64(i)+15)
	}
	arrays = append(arrays, longTestCase)

	for i := 0; i < len(arrays); i++ {
		squared := make([]int64, 0, len(arrays[i]))
		for _, v := range arrays[i] {
			squared = append(squared, v*v)
		}
		sort.Slice(squared, func(i, j int) bool { return squared[i] < squared[j] })
		outputs = append(outputs, squared)
	}
}

func Test(t *testing.T) {
	for i := 0; i < len(arrays); i++ {
		i := i
		t.Run(fmt.Sprintf("%dth test array", i), func(t *testing.T) {
			buffer := bytes.NewBufferString("")
			m := sync.Mutex{}
			calculateSquareAndPrintOut(&m, buffer, arrays[i]...)

			scanned := make([]int64, 0, len(arrays[i]))
			curScanned := int64(0)
			_, err := fmt.Fscan(buffer, &curScanned)
			for err == nil {
				scanned = append(scanned, curScanned)
				_, err = fmt.Fscan(buffer, &curScanned)
			}
			sort.Slice(scanned, func(i, j int) bool { return scanned[i] < scanned[j] })
			if cmpResult := slices.Compare(scanned, outputs[i]); cmpResult != 0 {
				t.Logf("scanned: %v", scanned)
				t.Logf("outputs[%d]: %v", i, outputs[i])
				t.Fatalf("scanned numbers are different from valid outputs")
			}
		})
	}
}
