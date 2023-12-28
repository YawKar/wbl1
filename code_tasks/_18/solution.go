package _18

import "sync/atomic"

type AtomicCounter struct {
	c atomic.Int64
}

func (ac *AtomicCounter) Inc(delta int64) {
	ac.c.Add(delta)
}

func (ac *AtomicCounter) Load() int64 {
	return ac.c.Load()
}
