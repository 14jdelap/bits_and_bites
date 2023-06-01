package main

import (
	"fmt"
)

func main() {

	var limit uint64 = 1 << 30

	for i := uint64(0); i < limit; i++ {
		enc := encode(i)
		dec := decode(enc)
		if dec != i {
			fmt.Println("FAILED", enc, dec)
		}
	}
}

// encode converts a uint64 into a Protobuf Base 128 Varint in []byte format
// and little endian ordering.
// encode loops until the uint64 intput it 0. While it's greater than 0, it
// copies the first 7 bits of the uint64, shifts right by 7m and marks its
// first bit as 1 if there's a successive byte (i.e., if it's value is still
// greater than 0). When the loop finishes it returns the []byte result.
func encode(i uint64) []byte {
	result := make([]byte, 0, 10)
	for i > 0 {
		b := byte(i & 0b01111111)
		i >>= 7
		if i > 0 {
			b |= 0b10000000
		}
		result = append(result, b)
	}
	return result
}

// decode converts a Protobuf Base 128 varint into a uint64.
// It keeps the payload of each byte (i.e., it removes the signbit)
// and concatenates all payloads with little endian byte ordering
// by shifting each payload left by 7 and using a bitwise OR operator.
func decode(bs []byte) uint64 {
	var result uint64
	for i := 0; i < len(bs); i++ {
		result |= uint64(bs[i]&0b01111111) << (i * 7)
	}
	return result
}
