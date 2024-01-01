package main

import (
	"fmt"
	"unsafe"
)

type t1 struct { // will have size of 8 bytes on 64 bit platform
	f int
}

type t2 struct { // will have size of 16 bytes on 64 bit platform
	f int
	g struct{} // due to alignment?
}

type t3 struct { // will still be 16 bytes sized
	f  int
	g  struct{}
	g1 struct{}
	g2 struct{}
	g3 struct{}
	g4 struct{}
	g5 struct{}
	g6 struct{}
}

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{})) // 0
	fmt.Println(unsafe.Sizeof(int32(3)))   // 4

	// 8 16 8 0
	fmt.Println(unsafe.Sizeof(t1{}), unsafe.Sizeof(t2{}), unsafe.Sizeof(t2{}.f), unsafe.Sizeof(t2{}.g))
	// 16
	fmt.Println(unsafe.Sizeof(t3{}))
}
