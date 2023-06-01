package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
	"unicode"
)

const signbit string = "1"
const exponent string = "11111111111"

func firstConceal(s string) (float64, error) {
	fraction := "1"
	l := len(s)

	if l > 6 {
		fraction += fmt.Sprintf("%03s", strconv.FormatInt(6, 2))
		l = 6
	} else {
		fraction += fmt.Sprintf("%03s", strconv.FormatInt(int64(l), 2))
	}

	for i := 0; i < 6; i++ {
		if i < l {
			if isASCII(s[i]) {
				fraction += fmt.Sprintf("%08b", s[i])
			} else {
				return 0, errors.New("string argument should only have ASCII characters")
			}
		} else {
			fraction += "00000000"
		}
	}

	bitStr := signbit + exponent + fraction
	number, err := strconv.ParseUint(bitStr, 2, 64)

	if err != nil {
		return 0, errors.New("error when parsing string to float")
	}

	return math.Float64frombits(number), nil
}

func isASCII(b byte) bool {
	return b < unicode.MaxASCII
}

func firstExtract(f float64) string {
	bits := strconv.FormatUint(math.Float64bits(f), 2)
	fraction := bits[12:]
	fractionLength, _ := strconv.ParseInt(fraction[1:4], 2, 0)

	if fractionLength == 0 {
		return ""
	}

	byteArray := make([]byte, fractionLength)

	for i := 0; i < int(fractionLength); i++ {
		start, end := 4+(i*8), 4+((i+1)*8)
		asciiCode, _ := strconv.ParseInt(fraction[start:end], 2, 0)
		byteArray[i] = byte(asciiCode)
	}

	return string(byteArray)
}

func secondConceal(s string) float64 {
	byteFloat := []byte{0b11111111, 0b11111000}
	bs := []byte(s)
	bsLength := len(bs)

	// Use bytwise OR to set string length
	if bsLength > 6 {
		byteFloat[1] ^= byte(6)
		bsLength = 6
	} else {
		byteFloat[1] ^= byte(bsLength)
	}

	fmt.Println(byteFloat[1])

	// Append byte of value 0 for each spare bit
	for i := 0; i < (6 - bsLength); i++ {
		byteFloat = append(byteFloat, 0)
	}

	// Append string
	byteFloat = append(byteFloat, bs[:(bsLength+1)]...)

	// Create float64 from uint6
	fmt.Println(byteFloat, binary.BigEndian.Uint64(byteFloat))
	return math.Float64frombits(binary.BigEndian.Uint64(byteFloat))
}

func secondExtract(f float64) string {
	b := math.Float64bits(f) - 1
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, b)
	len := bs[1] & 0b111
	return string(bs[8-len : 8])
}
