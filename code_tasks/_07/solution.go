package _07

import "sync"

type ConcMap[K comparable, V any] struct {
	m sync.RWMutex
	s map[K]V
}

func NewConcMap[K comparable, V any]() *ConcMap[K, V] {
	return &ConcMap[K, V]{
		m: sync.RWMutex{},
		s: make(map[K]V),
	}
}

func (cm *ConcMap[K, V]) Set(key K, val V) {
	cm.m.Lock()
	defer cm.m.Unlock()
	cm.s[key] = val
}

func (cm *ConcMap[K, V]) Apply(key K, f func(V) V) bool {
	cm.m.Lock()
	defer cm.m.Unlock()
	v, found := cm.s[key]
	if found {
		cm.s[key] = f(v)
	}
	return found
}

func (cm *ConcMap[K, V]) ApplyAndGetPrevious(key K, f func(V) V) (previous V, found bool) {
	cm.m.Lock()
	defer cm.m.Unlock()
	previous, found = cm.s[key]
	if found {
		cm.s[key] = f(previous)
	}
	return
}

func (cm *ConcMap[K, V]) SetAndGetPrevious(key K, val V) (previous V, found bool) {
	cm.m.Lock()
	defer cm.m.Unlock()
	previous, found = cm.s[key]
	cm.s[key] = val
	return
}

func (cm *ConcMap[K, V]) Get(key K) (value V, found bool) {
	cm.m.RLock()
	defer cm.m.RUnlock()
	value, found = cm.s[key]
	return
}

func (cm *ConcMap[K, V]) Has(key K) (found bool) {
	cm.m.RLock()
	defer cm.m.RUnlock()
	_, found = cm.s[key]
	return
}

func (cm *ConcMap[K, V]) Count() int {
	cm.m.RLock()
	defer cm.m.RUnlock()
	return len(cm.s)
}

func (cm *ConcMap[K, V]) IsEmpty() bool {
	cm.m.RLock()
	defer cm.m.RUnlock()
	return len(cm.s) == 0
}

func (cm *ConcMap[K, V]) Delete(key K) {
	cm.m.Lock()
	defer cm.m.Unlock()
	delete(cm.s, key)
}

func (cm *ConcMap[K, V]) DeleteAndGetPrevious(key K) (previous V, found bool) {
	cm.m.Lock()
	defer cm.m.Unlock()
	previous, found = cm.s[key]
	delete(cm.s, key)
	return
}
