package main

import (
	"log"
	"unsafe"
)

func main() {
	var message string
	littleEndian := getEndian()
	if littleEndian {
		message = "little endian"
	} else {
		message = "big endian"
	}
	log.Printf("Your computer uses %s byte ordering\n", message)

	var i int32 = 10
	if isAnyBitOn(i) {
		message = "has at least one bit on"
	} else {
		message = "has no bits on"
	}

	log.Printf("The integer %d %s\n", i, message)

	// showInt(100)
	showString("hello")
	// showFloat(100.1)
}

// isAnyBitOn returns true if at least one bit in the integer is on.
func isAnyBitOn(i int32) bool {
	iup := unsafe.Pointer(&i)
	for i := 0; i < 4; i++ {
		bitOn := *(*byte)(unsafe.Add(iup, i)) | 0b0

		if bitOn != 0 {
			return true
		}
	}
	return false
}
