package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{})) // 0
	fmt.Println(unsafe.Sizeof(int32(3)))   // 4
}
