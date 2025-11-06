package primey

import (
	"slices"
	"testing"
)

func TestIter(t *testing.T) {
	primes := make([]int, 7)
	for i, prime := range Iter() {
		if i > 4 {
			break
		}
		primes[i] = prime
	}
	if !slices.Equal(primes, []int{2, 3, 5, 7, 11, 0, 0}) {
		t.Error("Iter failed to regenerate simple test!", primes)
	}
}

func TestIterr(t *testing.T) {
	testCases := []struct {
		start    int
		end      int
		expected []int
	}{
		// Just within the primeCache range
		{1, 2, []int{3, 0, 0, 0, 0}},

		// Just within the wheel range
		{5, 8, []int{13, 17, 19, 0, 0}},

		// Spanning the primeCache across into the wheel
		{2, 7, []int{5, 7, 11, 13, 17}},
	}

	CacheResize(3)

	for _, testCase := range testCases {
		answer := make([]int, 5)
		for i, prime := range Iterr(testCase.start, testCase.end) {
			answer[i] = prime
		}
		if !slices.Equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %d:%d expected %v, got %v", testCase.start, testCase.end, testCase.expected, answer)
		}
	}
}

func TestNth(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 2},
		{1, 3},
		{2, 5},
		{3, 7},
		{4, 11},
		{5, 13},
		{6, 17},
		{7, 19},
		{8, 23},
		{9, 29},
		{10, 31},
		{11, 37},

		{29, 113},
		{30, 127},
		{31, 131},
		{32, 137},

		{89, 463},
		{97, 521},
		{121, 673},
	}

	for _, testCase := range testCases {
		answer := Nth(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestIndex(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 0},
		{2, 0},
		{3, 1},
		{4, 1},
		{5, 2},
		{6, 2},
		{7, 3},
		{8, 3},
		{9, 3},
		{10, 3},
		{11, 4},

		{29, 9},
		{30, 9},
		{31, 10},
		{32, 10},

		{89, 23},
		{97, 24},
		{121, 29},
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

func TestPrimeSlow(t *testing.T) {
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
		{100001030, false},
	}

	for _, testCase := range testCases {
		answer := PrimeSlow(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}

func TestCacheResize(t *testing.T) {
	testCases := []struct {
		n         int
		expected1 []uint32
		expected2 int
	}{
		{0, []uint32{2, 3, 5}, 7},
		{1, []uint32{2, 3, 5}, 7},
		{2, []uint32{2, 3, 5}, 7},
		{3, []uint32{2, 3, 5}, 7},
		{4, []uint32{2, 3, 5, 7}, 11},
		{5, []uint32{2, 3, 5, 7, 11}, 13},
		{3, []uint32{2, 3, 5}, 7},
	}

	for _, testCase := range testCases {
		CacheResize(testCase.n)
		if !slices.Equal(primeCache, testCase.expected1) {
			t.Errorf("ERROR: For %d expected %v, got %v", testCase.n, testCase.expected1, primeCache)
		}

		// Verify that the next prime (which will come from wheel) is the proper one in the sequence
		// That is, that primeCache and wheel are properly aligned
		answer := Nth(max(testCase.n, 3))
		if answer != testCase.expected2 {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected2, answer)
		}
	}
}
