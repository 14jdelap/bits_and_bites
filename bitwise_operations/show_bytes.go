package main

import (
	"encoding/hex"
	"log"
	"unsafe"
)

func showBytes(bytePointer unsafe.Pointer, size int) {
	bs := make([]byte, 0, size)
	for i := 0; i < int(size); i++ {
		bs = append(bs, *(*byte)(unsafe.Add(bytePointer, i)))
	}
	log.Printf("0x%v", hex.EncodeToString(bs), bs)
}

func showString(s string) {
	// A Go string under the hood uses the StringHeader struct (C doesn't).
	// This is why unsafe.Sizeof(&s) will return 16 rather than 5 bytes.
	// We could change the pointer to be a StringHeader pointer to access
	// its Len field using the compiler's knowledge of thr struct definition
	// by name and translating that to specific addresses.

	sp := unsafe.StringData(s)
	showBytes(unsafe.Pointer(sp), len([]byte(s)))
}

func showInt(i int) {
	showBytes(unsafe.Pointer(&i), int(unsafe.Sizeof(i)))
}

func showFloat(f float32) {
	showBytes(unsafe.Pointer(&f), int(unsafe.Sizeof(f)))
}
