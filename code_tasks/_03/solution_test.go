package _03

import (
	"fmt"
	"testing"
	"time"
)

var (
	arrays  [][]int64
	outputs []int64
)

func initArrays() {
	arrays = [][]int64{
		{},
		{2, 4, 6, 8, 10},
	}
	{
		length := 100
		array := make([]int64, 0, length)
		for i := 0; i < 100; i++ {
			array = append(array, int64(i))
		}
		arrays = append(arrays, array)
	}
}

func initOutputs() {
	outputs = make([]int64, len(arrays))
	for i, arr := range arrays {
		for _, el := range arr {
			outputs[i] += el * el
		}
	}
}

func init() {
	initArrays()
	initOutputs()
}

func TestCalculateBlocking(t *testing.T) {
	for i := 0; i < len(arrays); i++ {
		i := i
		t.Run(fmt.Sprintf("%d-th test array", i), func(t *testing.T) {
			result := sumSquaresConcurrentlyBlock(arrays[i]...)
			if result != outputs[i] {
				t.Errorf("result != valid output: %d != %d", result, outputs[i])
				t.Logf("tested array: %v", arrays[i])
			}
		})
	}
}

func TestCalculateNonBlocking(t *testing.T) {
	timeoutDuration := 10 * time.Second
	for i := 0; i < len(arrays); i++ {
		i := i
		t.Run(fmt.Sprintf("%d-th test array", i), func(t *testing.T) {
			resCh := sumSquaresConcurrentlyNonBlock(arrays[i]...)
			timeout := time.After(timeoutDuration)
			for {
				select {
				case <-timeout:
					t.Errorf("timed-out during the calculation, timeout = %v", timeoutDuration)
					return
				case result := <-resCh:
					if result != outputs[i] {
						t.Errorf("result != valid output: %d != %d", result, outputs[i])
						t.Logf("tested array: %v", arrays[i])
					}
					return
				}
			}
		})
	}
}
