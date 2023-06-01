package main

import (
	"encoding/binary"
	"fmt"
	"math"
)

func main() {
	// f1, _ := firstConceal("aaa")
	// f2, _ := firstConceal("aaaaaa")
	// f3, _ := firstConceal("")
	// f4, _ := firstConceal("aaaaaaa")
	// _, e := firstConceal("á")
	// fmt.Println(f1, f2)
	// fmt.Println(firstExtract(f1) == "aaa")
	// fmt.Println(firstExtract(f2) == "aaaaaa")
	// fmt.Println(firstExtract(f3) == "")
	// fmt.Println(firstExtract(f4) == "aaaaaa")
	// fmt.Println(e)

	f1 := secondConceal("aaa")
	// f2 := conceal("aaaaaa")
	fmt.Println(secondExtract(f1) == "aaa")
	// fmt.Println(extract(f2) == "aaaaaa")
	// fmt.Println(extract(f3) == "")
	// fmt.Println(extract(f4) == "aaaaaa")
	// fmt.Println(e)
}

func conceal(s string) float64 {
	bs := []byte(s)

	if len(bs) > 6 {
		return 0
	}

	nan := []byte{0b11111111, 0b11111000}
	nan[1] ^= byte(len(bs)) // XOR to codify length in the last 3 bits of 2nd bit

	for i := 0; i < 6-len(bs); i++ {
		nan = append(nan, 0)
	}

	nan = append(nan, bs...) // append byte to nan byte slice

	fmt.Println(binary.BigEndian.Uint64(nan))
	return math.Float64frombits(binary.BigEndian.Uint64(nan)) // Convert nan to bits and use to create float
}

func extract(f float64) string {
	bits := math.Float64bits(f)                // Return uint64 from float64
	b := make([]byte, 8)                       // Create a byte slice
	b = binary.BigEndian.AppendUint64(b, bits) // Use big endian notation to append uint64 into the byte slice
	l := int(b[1] & 0b111)                     // Use bitwise AND to get the length encoded in the last 3 bits
	return string(b[len(b)-l:])                // Return the string encoded in the fraction
}
