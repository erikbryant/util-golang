package algebra

import "testing"

func TestNextSquare(t *testing.T) {
	testCases := []struct {
		c         uint
		expected  uint
		expected2 uint
	}{
		{0, 1, 1},
		{1, 4, 2},
		{9, 16, 4},
		{10, 16, 4},
		{15, 16, 4},
		{100, 121, 11},
		{101, 121, 11},
	}

	for _, testCase := range testCases {
		answer, answer2 := NextSquare(testCase.c)
		if answer != testCase.expected || answer2 != testCase.expected2 {
			t.Errorf("ERROR: For %d expected %d/%d, got %d/%d", testCase.c, testCase.expected, testCase.expected2, answer, answer2)
		}
	}
}
