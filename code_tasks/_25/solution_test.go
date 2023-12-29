package _25

import (
	"testing"
	"time"
)

func TestSleep(t *testing.T) {
	eps := 10 * time.Millisecond
	duration := time.Second

	start := time.Now()
	Sleep(duration)
	after := time.Now()
	delta := after.Sub(start)
	if duration-eps > delta {
		t.Errorf("too early, diff = %d", delta-(duration-eps))
	} else if delta > duration+eps {
		t.Errorf("too late, diff = %d", delta-(duration+eps))
	}
}
