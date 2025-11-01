package primey

import (
	"testing"
)

func TestWheelStart(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{1, 1},
	}

	for _, testCase := range testCases {
		answer := wheelStart()
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestInt2Offset(t *testing.T) {
	testCases := []struct {
		n         int
		expected1 int
		expected2 uint8
		expected3 bool
		expected4 int
	}{
		{0, 0, 0, false, 0},
		{1, 0, 0, true, 1},
	}

	for _, testCase := range testCases {
		answer1, answer2, answer3, answer4 := int2offset(testCase.n)
		if answer1 != testCase.expected1 || answer2 != testCase.expected2 || answer3 != testCase.expected3 || answer4 != testCase.expected4 {
			t.Errorf("ERROR: For %d expected %d, %d, %t, %d, got %d, %d, %t, %d", testCase.n, testCase.expected1, testCase.expected2, testCase.expected3, testCase.expected4, answer1, answer2, answer3, answer4)
		}
	}
}

func TestOffset2Int(t *testing.T) {
	testCases := []struct {
		iByte    int
		iBit     uint8
		expected int
	}{
		{0, 0, 1},
		{0, 1, 7},
		{0, 2, 11},
		{0, 3, 13},
		{0, 4, 17},
		{0, 5, 19},
		{0, 6, 23},
		{0, 7, 29},

		{1, 0, 31},
		{2, 1, 67},
		{3, 2, 101},
		{4, 3, 133},
		{5, 4, 167},
		{6, 5, 199},
		{7, 6, 233},
		{8, 7, 269},
	}

	for _, testCase := range testCases {
		answer := offset2int(testCase.iByte, testCase.iBit)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.iByte, testCase.iBit, testCase.expected, answer)
		}
	}
}

func TestSetBit(t *testing.T) {
	// TODO
}

func TestBitIsSet(t *testing.T) {
	testCases := []struct {
		iByte    int
		iBit     uint8
		expected bool
	}{
		{0, 0, false},
		{0, 1, true},
		{0, 2, true},
		{0, 3, true},
		{0, 4, true},
		{0, 5, true},
		{0, 6, true},
		{0, 7, true},

		{1, 0, true},
		{1, 1, true},
		{1, 2, true},
		{1, 3, true},
		{1, 4, true},
		{1, 5, false},
		{1, 6, true},
		{1, 7, true},
	}

	for _, testCase := range testCases {
		answer := bitIsSet(testCase.iByte, testCase.iBit)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %t, got %t", testCase.iByte, testCase.iBit, testCase.expected, answer)
		}
	}
}

func TestIndex2Offset(t *testing.T) {
	testCases := []struct {
		i         int
		expected1 int
		expected2 uint8
	}{
		// Wheel does not store primes {2, 3, 5}
		{0, 0, 0}, // 2
		{1, 0, 0}, // 3
		{2, 0, 0}, // 5

		// wheel[0]
		{3, 0, 1}, // 7
		{4, 0, 2}, // 11
		{5, 0, 3}, // 13
		{6, 0, 4}, // 17
		{7, 0, 5}, // 19
		{8, 0, 6}, // 23
		{9, 0, 7}, // 29

		// wheel[1]
		{10, 1, 0}, // 31
		{11, 1, 1}, // 37
		{12, 1, 2}, // 41
		{13, 1, 3}, // 43
		{14, 1, 4}, // 47
		{15, 1, 6}, // 53
		{16, 1, 7}, // 59

		// Transition from wheel[3] to wheel[4]
		{29, 3, 6}, // 113
		{30, 4, 1}, // 127
		{31, 4, 2}, // 131
	}

	for _, testCase := range testCases {
		answer1, answer2 := index2offset(testCase.i)
		if answer1 != testCase.expected1 || answer2 != testCase.expected2 {
			t.Errorf("ERROR: For %d expected %d, %d, got %d, %d", testCase.i, testCase.expected1, testCase.expected2, answer1, answer2)
		}
	}
}

func TestOffset2Index(t *testing.T) {
	testCases := []struct {
		iByte    int
		iBit     uint8
		expected int
	}{
		// wheel[0]
		{0, 1, 3}, // 7
		{0, 2, 4}, // 11
		{0, 3, 5}, // 13
		{0, 4, 6}, // 17
		{0, 5, 7}, // 19
		{0, 6, 8}, // 23
		{0, 7, 9}, // 29

		// wheel[1]
		{1, 0, 10}, // 31
		{1, 1, 11}, // 37
		{1, 2, 12}, // 41
		{1, 3, 13}, // 43
		{1, 4, 14}, // 47
		{1, 6, 15}, // 53
		{1, 7, 16}, // 59

		// Transition from wheel[3] to wheel[4]
		{3, 6, 29}, // 113
		{4, 1, 30}, // 127
		{4, 2, 31}, // 131
	}

	for _, testCase := range testCases {
		answer := offset2index(testCase.iByte, testCase.iBit)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.iByte, testCase.iBit, testCase.expected, answer)
		}
	}
}
