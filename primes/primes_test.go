package primes

import (
	"testing"
)

func init() {
	Load("../primes.gob")
}

func TestPi(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 2},
		{5, 3},
		{6, 3},
		{7, 4},
		{8, 4},
		{9, 4},
		{10, 4},
		{100, 25},
		{1000, 168},
		{10 * 1000, 1229},
		{100 * 1000, 9592},
		{1000 * 1000, 78498},
		{10 * 1000 * 1000, 664579},
		{100 * 1000 * 1000, 5761455},

		// This will trigger the panic, if you want to test that
		// {1000*1000*1000,50847534},
	}

	for _, testCase := range testCases {
		answer := Pi(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestPrime(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, false},
		{7, true},
		{101, true},
	}

	for _, testCase := range testCases {
		answer := Prime(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}
