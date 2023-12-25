package _03

import (
	"sync"
	"sync/atomic"
)

// Blocking function
func sumSquaresConcurrentlyBlock(numbers ...int64) int64 {
	wg := sync.WaitGroup{}
	wg.Add(len(numbers))
	sum := atomic.Int64{}
	for _, v := range numbers {
		go func(v int64, wg *sync.WaitGroup) {
			sum.Add(v*v)
			wg.Done()
		}(v, &wg)
	}
	wg.Wait()
	return sum.Load()
}

// Non-blocking function
func sumSquaresConcurrentlyNonBlock(numbers ...int64) <-chan int64 {
	result := make(chan int64)
	go func() {
		result <- sumSquaresConcurrentlyBlock(numbers...)
		close(result)
	}()
	return result
}
