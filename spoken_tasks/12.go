package main

import "fmt"

func main() {
	n := 0 // outer `n`
	if true {
		n := 1 // shadows the outer `n`, from now on - it is a distinct variable
		n++    // changes local (in brackets) `n`, not the outer one
	}
	fmt.Println(n) // prints `0`
}
