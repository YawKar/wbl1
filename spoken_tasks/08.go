package main

func main() {
	var _ *struct{ name string } = new(struct{ name string }) // returns a pointer to newly allocated zero-valued object
	var _ map[int]int = make(map[int]int)                     // returns value of the same type that was provided + initializes it

	var _ *int = new(int) // can allocate int

	var _ *chan int = new(chan int)       // zero-valued chan (nil), unusable because uninitialized
	var _ *map[int]int = new(map[int]int) // zero-valued map (nil), unusable because uninitialized
	var _ *[]int = new([]int)             // zero-valued slice (nil), unusable because uninitialized

	// make may only be used to create chan, map or slice (and initialize them along the way)
	var _ chan int = make(chan int)       // usable initialized chan
	var _ map[int]int = make(map[int]int) // usable initialized map
	var _ []int = make([]int, 0)          // usable initialized slice
}
