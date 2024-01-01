package main

import "fmt"

func main() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a") // creates totally new slice with new buffer, connection with the outer slice is lost
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice) // [b b a]
	}(slice)
	fmt.Println(slice) // [a a]
}
