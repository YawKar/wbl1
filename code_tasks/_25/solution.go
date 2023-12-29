package _25

import (
	"runtime"
	"time"
)

func Sleep(d time.Duration) {
	deadline := time.Now().Add(d)
	for time.Now().Before(deadline) {
		runtime.Gosched()
	}
}
