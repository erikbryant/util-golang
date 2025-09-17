package figurate

import "testing"

func TestTriangular(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 6},
		{4, 10},
		{5, 15},
		{6, 21},
		{-1, 0},
		{-2, 1},
		{-3, 3},
		{-4, 6},
		{-5, 10},
		{-6, 15},
	}

	for _, testCase := range testCases {
		answer := Triangular(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

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

func TestPentagonal(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 5},
		{3, 12},
		{4, 22},
		{5, 35},
		{6, 51},
		{-1, 2},
		{-2, 7},
		{-3, 15},
		{-4, 26},
		{-5, 40},
		{-6, 57},
	}

	for _, testCase := range testCases {
		answer := Pentagonal(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
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
		{0, false},
		{1, true},
		{2, true},
		{5, true},
		{7, true},
		{12, true},
		{15, true},
		{22, true},
		{26, true},
		{40, true},
		{57, true},
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

func TestHexagonal(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 6},
		{3, 15},
		{4, 28},
		{5, 45},
		{6, 66},
		{-1, 3},
		{-2, 10},
		{-3, 21},
		{-4, 36},
		{-5, 55},
		{-6, 78},
	}

	for _, testCase := range testCases {
		answer := Hexagonal(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestIsHexagonal(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{-1, false},
		{0, false},
		{1, true},
		{3, true},
		{5, false},
		{6, true},
		{10, true},
		{12, false},
		{15, true},
		{21, true},
		{28, true},
		{36, true},
		{55, true},
		{78, true},
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
