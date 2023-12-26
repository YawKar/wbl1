package _07

import (
	"sync"
	"testing"
)

func TestConcurrentMapCounter(t *testing.T) {
	testKey := "testKey"
	numG := 10000
	additionPerG := 100
	totalAddition := numG * additionPerG

	cmap := NewConcMap[string, int]()
	cmap.Set(testKey, 0)

	wg := sync.WaitGroup{}
	for i := 0; i < numG; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cmap.Apply(testKey, func(v int) int { return v + additionPerG })
		}()
	}
	wg.Wait()
	if v, ok := cmap.Get(testKey); !ok {
		t.Errorf("testKey wasn't found in conc map: %s", testKey)
	} else if v != totalAddition {
		t.Errorf("resulting sum doesn't add up to the total addition: %d != %d", v, totalAddition)
	}
}

func TestNewlyCreatedCmap(t *testing.T) {
	cmap := NewConcMap[struct{}, struct{}]()
	if count := cmap.Count(); count != 0 {
		t.Errorf("new conc map should be empty but its size is %d", count)
	}
	if isEmpty := cmap.IsEmpty(); !isEmpty {
		t.Errorf("new conc map should be empty but shows that it is not")
	}
}

func TestCmapCountMethod(t *testing.T) {
	cmap := NewConcMap[int, int]()

	// inserting
	for i := 1; i <= 10; i++ {
		cmap.Set(i, i)
		if count := cmap.Count(); count != i {
			t.Fatalf("conc map shows invalid number of elements in it: %d != %d", count, i)
		}
	}

	// deleting
	for i := 10; i > 0; i-- {
		cmap.Delete(i)
		if count := cmap.Count(); count != i-1 {
			t.Fatalf("conc map shows invalid number of elements in it: %d != %d", count, i-1)
		}
	}
}

func TestHas(t *testing.T) {
	cmap := NewConcMap[string, any]()

	key1 := "Nanami"
	if cmap.Has(key1) {
		t.Fatalf("conc map shows that it has a key which wasn't inserted: %s", key1)
	}
	cmap.Set(key1, struct{}{})
	if !cmap.Has(key1) {
		t.Fatalf("conc map shows that it hasn't the key that was inserted: %s", key1)
	}

	key2 := "Sukuna"
	if cmap.Has(key2) {
		t.Fatalf("conc map shows that it has a key which wasn't inserted: %s", key2)
	}
	cmap.Set(key2, struct{}{})
	if !cmap.Has(key2) {
		t.Fatalf("conc map shows that it hasn't the key that was inserted: %s", key2)
	}
}
