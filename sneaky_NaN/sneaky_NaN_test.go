package main

import (
	"errors"
	"fmt"
	"math"
	"testing"
)

func TestFirstConceal(t *testing.T) {
	var tests = []struct {
		s string
		u uint64
		e error
	}{
		{"aaa", 18445443769693110272, nil},
		{"aaaaaa", 18446288194629624161, nil},
		{"", 18444492273895866368, nil},
		{"aaaaaaa", 18446288194629624161, nil},
		{"á", 0, errors.New("string argument should only have ASCII characters")},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("String :: %s", tt.s)
		t.Run(testname, func(t *testing.T) {
			f, _ := firstConceal(tt.s)
			bits := math.Float64bits(f)
			if bits != tt.u {
				t.Errorf("expected and received floats have different values :: %d, %d", tt.u, bits)
			}
		})
	}
}
