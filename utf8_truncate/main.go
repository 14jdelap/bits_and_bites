package main

import (
	"bufio"
	"fmt"
	"os"
)

// bytesRemainingInBits maps the integer of remaining bytes in a row to
// a hex that represents the start of a multi-byte UTF-8 it cannot include
var bytesRemainingInBits = map[int]byte{
	3: 0xF0, // 3 remaining bytes maps to start of 4 byte UTF-8
	2: 0xE0, // 2 remaining bytes maps to start of 3 byte UTF-8
	1: 0xC0, // 1 remaining byte maps to start of 2 byte UTF-8
}

func main() {
	file, err := os.Open("./files/cases")
	if err != nil {
		panic(err)
	}

	// Scanner is a wrapper around the io.Reader to read line-delimited files
	scanner := bufio.NewScanner(file)
	for {
		if !scanner.Scan() {
			// Return main when EOF (assumed only error from scanner.Scan())
			return
		}

		b := scanner.Bytes() // row of bytes
		truncateAtByte, content := int(b[0]), b[1:]
		result := make([]byte, 0, truncateAtByte)

		for i := 0; i <= truncateAtByte && i < len(content); i++ {
			// bytesToEnd is the number of bytes until the row is truncated or ends
			bytesToEnd := truncateAtByte - i
			if bytesToEnd > len(content)-i {
				bytesToEnd = len(content) - i
			}

			if bytesToEnd < 4 &&
				content[i]&bytesRemainingInBits[bytesToEnd] ==
					// Once the row's remaining bytes may leave junk at truncation (i.e. 3
					// or less, given UTF-8 goes up to 4 bytes), in each iteration check
					// that the current byte is not the start of a multi-byte UTF-8 that
					// exceeds the remaining bytes. If it exceeds, break to the next row.
					bytesRemainingInBits[bytesToEnd] {
				break
			}

			result = append(result, content[i])
		}
		fmt.Println(string(result))
	}
}
