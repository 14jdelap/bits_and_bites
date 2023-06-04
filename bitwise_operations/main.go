package main

import (
	"unsafe"
)

func main() {
	littleEndian := getEndian()
	isAnyBitOn(10, littleEndian)
}

// isAnyBitOn returns true if at least one bit in the integer is on.
func isAnyBitOn(i int32, littleEndian bool) bool {
	iup := unsafe.Pointer(&i)
	for i := 0; i < 4; i++ {
		bitOn := *(*byte)(unsafe.Add(iup, i)) | 0b00000000

		if bitOn != 0 {
			return true
		}
	}
	return false
}
