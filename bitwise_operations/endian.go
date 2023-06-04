package main

import "unsafe"

// getEndian returns true if the machine uses little endian byte ordering.
func getEndian() bool {
	var i int32 = 1
	iup := unsafe.Pointer(&i)
	littleEndian := (*(*byte)(iup) | 0x0) != 0

	return littleEndian
}
