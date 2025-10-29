package primey

import (
	"slices"
	"testing"
)

func TestIndex(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, -1},
		{1, -1},
		{2, 0},
		{3, 1},
		{4, -1},
		{5, 2},
		{6, -2},
		{7, 3},
		{8, -3},
		{9, -3},
		{10, -3},
		{11, 4},
		{29, 9},
		{30, -9},
		{31, 10},
		{32, -10},
		{89, 23},
		{97, 24},
		{121, -29},
	}

	for _, testCase := range testCases {
		answer := Index(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
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

func TestIter(t *testing.T) {
	primes := []int{}
	for i, prime := range Iter() {
		if i > 4 {
			break
		}
		primes = append(primes, prime)
	}
	if !slices.Equal(primes, []int{2, 3, 5, 7, 11}) {
		t.Error("Iter failed to regenerate simple test!", primes)
	}
}

func TestIterr(t *testing.T) {
	primes := make([]int, 5)
	for i, prime := range Iterr(1, 5) {
		primes[i] = prime
	}
	if !slices.Equal(primes, []int{3, 5, 7, 11, 0}) {
		t.Error("Iterr failed to regenerate simple test!", primes)
	}
}

func TestSlowPrime(t *testing.T) {
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
		{9, false},
		{101, true},
		{PrimeMax() + 1, false},
		{100001029, true},
	}

	for _, testCase := range testCases {
		answer := SlowPrime(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}

func TestPrime(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{-1, false},
		{0, false},
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
