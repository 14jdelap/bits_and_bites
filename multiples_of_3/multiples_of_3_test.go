package multiples_of_three

import "testing"

func TestMultiplesOfThree(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{0, false},
		{1, false},
		{2, false},
		{3, true},
		{6, true},
		{9, true},
		{12, true},
		{15, true},
	}

	for _, tc := range testCases {
		result := IsMultipleOfThree(tc.input)
		if result != tc.expected {
			t.Errorf("expected IsMultiplesOfThree(%d) to be %v, but got %v", tc.input, tc.expected, result)
		}
	}
}
