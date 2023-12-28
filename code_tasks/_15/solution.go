package _15

import (
	"errors"
	"unsafe"
)

func CorrectWithoutLeakageFunc(stringGen func(length int) string) string {
	v := stringGen(1 << 10)
	slice := make([]byte, 100)
	copy(slice, v)
	return string(slice)
}

func ProofWithoutLeakageCorrectness() error {
	var vDataPtr *byte
	slice := CorrectWithoutLeakageFunc(func(length int) string {
		byts := make([]byte, length)
		gened := string(byts)
		vDataPtr = unsafe.StringData(gened)
		return gened
	})
	sliceDataPtr := unsafe.StringData(slice)
	if vDataPtr == sliceDataPtr {
		return errors.New("pointers are the same, memory leak occured")
	}
	return nil
}
