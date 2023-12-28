package _15

import (
	"errors"
	"unsafe"
)

func StringLeakageFunc(stringGen func(length int, flag []byte) string) string {
	v := stringGen(1<<10, []byte("flag")) // allocates a string of 1024 bytes in content size
	return v[:100]                        // leakage of 1024 - 100 = 924 bytes that aren't used, but stay in memory
}

func ProofStringLeakageFunc() error {
	var vDataPtr *byte
	var flagUsed []byte
	var lengthUsed int
	slice := StringLeakageFunc(func(length int, flag []byte) string {
		flagUsed = flag
		lengthUsed = length
		byts := make([]byte, length)
		for i, ch := range flag {
			byts[len(byts)-len(flag)+i] = ch
		}
		gened := string(byts)
		vDataPtr = unsafe.StringData(gened)
		return gened
	})
	sliceDataPtr := unsafe.StringData(slice)
	ptrsAreSame := vDataPtr == sliceDataPtr
	if !ptrsAreSame {
		return errors.New("pointers to slice data and string data differ")
	}
	flagIsInPlace := true
	for i, ch := range flagUsed {
		flagIsInPlace = flagIsInPlace && ch == *(*byte)(unsafe.Add(unsafe.Pointer(sliceDataPtr), lengthUsed-len(flagUsed)+i))
	}
	if !flagIsInPlace {
		return errors.New("flag wasn't found in leaked part of string")
	}
	return nil
}
