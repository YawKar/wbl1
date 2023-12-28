package _18

import (
	"sync"
	"testing"
)

func TestConcurrentlyAddition(t *testing.T) {
	numG := 10000
	deltaPerG := 100
	totalValidSum := int64(numG) * int64(deltaPerG)

	counter := AtomicCounter{}

	wg := sync.WaitGroup{}
	for i := 0; i < numG; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc(int64(deltaPerG))
		}()
	}
	wg.Wait()

	result := counter.Load()
	if result != totalValidSum {
		t.Fatalf("result differ from valid one: %d != %d", result, totalValidSum)
	}
}
