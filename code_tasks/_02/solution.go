package _02

import (
	"fmt"
	"io"
	"sync"
)

// Need mutex to be able to test output, otherwise output may (high probability) become nonsensical
func calculateSquareAndPrintOut(m *sync.Mutex, out io.Writer, array ...int64) {
	wg := sync.WaitGroup{}
	wg.Add(len(array))
	for _, v := range array {
		go func(v int64) {
			m.Lock()
			fmt.Fprintln(out, v*v)
			m.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()
}
