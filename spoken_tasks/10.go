package main

import "fmt"

func update(p *int) {
	b := 2
	p = &b // changes local copy of `p` pointer, therefore, doesn't affect outer pointer
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p) // copies `p` pointer
	fmt.Println(*p)
}
