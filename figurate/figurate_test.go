package figurate

import "testing"

func TestIsTriangular(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{-1, false},
		{0, true},
		{1, true},
		{3, true},
		{5, false},
		{6, true},
		{10, true},
		{12, false},
		{99, false},
		{666, true},
		{667, false},
	}

	for _, testCase := range testCases {
		answer := IsTriangular(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIsPentagonal(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{-3, false},
		{-1, false},
		{0, true},
		{1, true},
		{5, true},
		{12, true},
		{22, true},
		{99, false},
		{1000, false},
		{1001, true},
	}

	for _, testCase := range testCases {
		answer := IsPentagonal(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIsHexagonal(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{-1, false},
		{0, true},
		{1, true},
		{5, false},
		{6, true},
		{12, false},
		{15, true},
		{28, true},
		{99, false},
		{946, true},
		{947, false},
	}

	for _, testCase := range testCases {
		answer := IsHexagonal(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}
