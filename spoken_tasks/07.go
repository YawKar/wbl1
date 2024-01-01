package main

import "fmt"

func main() {
	// Answer: The order is uncertain
	m := make(map[int]int)
	m[0] = 1
	m[1] = 124
	m[2] = 281
	for k, v := range m {
		fmt.Println(k, v)
	}
	/* Possible outputs:
	0 1
	1 124
	2 281

	1 124
	2 281
	0 1

	2 281
	0 1
	1 124
	*/
}
