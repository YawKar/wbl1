package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) { // wait group is accepted as a value, therefore, loses connection with original wg
			fmt.Println(i)
			wg.Done()
		}(wg, i) // passed as value, should be a pointer
	}
	// eventually, panics with "all goroutines are asleep"
	wg.Wait() // will fall asleep forever
	fmt.Println("exit")
}
