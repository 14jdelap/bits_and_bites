package power_of_two

import "testing"

func TestIsPowerOfTwo(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{0, false},
		{1, true},
		{3, false},
		{4, true},
		{-1, false},
		{1 << 31, true},
	}

	for _, tc := range testCases {
		result := IsPowerOfTwo(tc.input)
		if result != tc.expected {
			t.Errorf("expected isPowerOfTwo(%d) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
