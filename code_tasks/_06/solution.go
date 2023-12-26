package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
}

// goroutine can be stopped using simple `return` statement
func StopUsingReturn() {
	go func() {
		return
	}()
}

// gracefully shutdown some goroutine using quit signal channel
func StopUsingQuitChan(d time.Duration) {
	quit := make(chan struct{})
	go func() {
		trash := true
		for {
			select {
			case <-quit:
				fmt.Println("stopped")
				return
			default:
				// doing some useful work
				trash = !trash
				fmt.Printf("doing trash: %t\n", trash)
			}
		}
	}()
	time.Sleep(d)
	quit <- struct{}{}
}

// gracefully shutdown some goroutine using context with cancellation
func StopUsingCancellableContext(d time.Duration) {
	ctx, cancel := context.WithCancel(context.Background())
	waitGLeave := make(chan struct{})

	go func() {
		trash := true
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped")
				waitGLeave <- struct{}{}
				return
			default:
				// doing some useful work
				trash = !trash
				fmt.Printf("doing trash: %t\n", trash)
			}
		}
	}()
	time.Sleep(d)
	cancel()
	<-waitGLeave
}

// shutdown goroutine using close on the channel which is listened by range on goroutine
func StopUsingRangeOnCh() {
	numbers := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		for v := range numbers {
			fmt.Println(v)
		}
		fmt.Println("finished")
		wg.Done()
	}()

	for i := 0; i < 10; i++ {
		numbers <- i
	}
	close(numbers)
	wg.Wait()
}

// goroutine may exit itself using Goexit func
func StopUsingGoexit() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("goroutine exited")
		time.Sleep(3 * time.Second)
		runtime.Goexit()
	}()
	wg.Wait()
}

// goroutine can temporarily stop its execution using Goshed
// although I couldn't make up with an idea of how to deterministically
// demonstrate it
func StopUsingGoshed() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1. about to drop the processor")
		runtime.Gosched()
		fmt.Println("1. returned back")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2. about to drop the processor")
		runtime.Gosched()
		fmt.Println("2. returned back")
	}()
	wg.Wait()
	/*
		Example output:
		2. about to drop the processor
		1. about to drop the processor
		2. returned back
		1. returned back
		Another one:
		2. about to drop the processor
		2. returned back
		1. about to drop the processor
		1. returned back
	*/
}

// goroutine can panic and then recover
// that way it won't corrupt the main goroutine while still will manage to stop its own execution
func StopViaPanic() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("hehe, boy")
	}()
	wg.Wait()
	fmt.Println("that's all")
}
