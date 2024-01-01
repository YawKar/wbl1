package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100       // sets the first element of the underlying buffer (affects outer buffer)
	v = append(v, b) // changes local `v`'s length (and buffer) ((and capacity))
	// now `v` is fully disconnected from the outer one
	v[0] = 2 // cannot be seen in main
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6) // passes `a` as a shallow copy (with shared buffer on the heap)
	fmt.Println(a)
}
