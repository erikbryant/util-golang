package util

import (
	"math/big"
	"testing"
)

func TestE(t *testing.T) {
	testCases := []struct {
		n        int
		expected int64
	}{
		{1, 2},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 1},
		{6, 4},
		{7, 1},
		{8, 1},
		{9, 6},
	}

	for _, testCase := range testCases {
		answer := E(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestConvergentE(t *testing.T) {
	testCases := []struct {
		n         int
		expectedN int64
		expectedD int64
	}{
		{1, 2, 1},
		{2, 3, 1},
		{3, 8, 3},
		{4, 11, 4},
		{5, 19, 7},
		{6, 87, 32},
		{7, 106, 39},
		{8, 193, 71},
		{9, 1264, 465},
		{10, 1457, 536},
	}

	for _, testCase := range testCases {
		expectedN := big.NewInt(testCase.expectedN)
		expectedD := big.NewInt(testCase.expectedD)
		answerN, answerD := Convergent(testCase.n, E)
		if answerN.Cmp(expectedN) != 0 || answerD.Cmp(expectedD) != 0 {
			t.Errorf("ERROR: For %d expected %d/%d, got %d/%d", testCase.n, testCase.expectedN, testCase.expectedD, answerN, answerD)
		}
	}
}

func TestConvergentSqrt2(t *testing.T) {
	testCases := []struct {
		n         int
		expectedN int64
		expectedD int64
	}{
		{1, 1, 1},
		{2, 3, 2},
		{3, 7, 5},
		{4, 17, 12},
		{5, 41, 29},
		{6, 99, 70},
		{7, 239, 169},
		{8, 577, 408},
	}

	for _, testCase := range testCases {
		expectedN := big.NewInt(testCase.expectedN)
		expectedD := big.NewInt(testCase.expectedD)
		answerN, answerD := Convergent(testCase.n, Sqrt2)
		if answerN.Cmp(expectedN) != 0 || answerD.Cmp(expectedD) != 0 {
			t.Errorf("ERROR: For %d expected %d/%d, got %d/%d", testCase.n, testCase.expectedN, testCase.expectedD, answerN, answerD)
		}
	}
}

func TestDivisors(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{0, []int{1}},
		{1, []int{1}},
		{2, []int{1, 2}},
		{3, []int{1, 3}},
		{4, []int{1, 2, 4}},
		{5, []int{1, 5}},
		{6, []int{1, 2, 3, 6}},
		{7, []int{1, 7}},
		{8, []int{1, 2, 4, 8}},
		{9, []int{1, 3, 9}},
		{10, []int{1, 2, 5, 10}},
		{11, []int{1, 11}},
		{12, []int{1, 2, 3, 4, 6, 12}},
		{20, []int{1, 2, 4, 5, 10, 20}},
		{28, []int{1, 2, 4, 7, 14, 28}},
		{100, []int{1, 2, 4, 5, 10, 20, 25, 50, 100}},
		{210, []int{1, 2, 3, 5, 6, 7, 10, 14, 15, 21, 30, 35, 42, 70, 105, 210}},
		{2310, []int{1, 2, 3, 5, 6, 7, 10, 11, 14, 15, 21, 22, 30, 33, 35, 42, 55, 66, 70, 77, 105, 110, 154, 165, 210, 231, 330, 385, 462, 770, 1155, 2310}},
	}

	for _, testCase := range testCases {
		answer := Divisors(testCase.n)
		if !Equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %d expected %v, got %v", testCase.n, testCase.expected, answer)
		}
	}
}

func TestFactors(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{2, []int{}},
		{3, []int{}},
		{4, []int{2}},
		{5, []int{}},
		{6, []int{2, 3}},
		{7, []int{}},
		{8, []int{2}},
		{9, []int{3}},
		{10, []int{2, 5}},
		{11, []int{}},
		{12, []int{2, 3}},
		{20, []int{2, 5}},
		{28, []int{2, 7}},
		{210, []int{2, 3, 5, 7}},
		{2310, []int{2, 3, 5, 7, 11}},
	}

	for _, testCase := range testCases {
		answer := Factors(testCase.n)
		if !Equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %d expected %v, got %v", testCase.n, testCase.expected, answer)
		}
	}
}

func TestFactorsCounted(t *testing.T) {
	testCases := []struct {
		n        int
		expected map[int]int
	}{
		{2, map[int]int{2: 1}},
		{3, map[int]int{3: 1}},
		{4, map[int]int{2: 2}},
		{5, map[int]int{5: 1}},
		{6, map[int]int{2: 1, 3: 1}},
		{7, map[int]int{7: 1}},
		{8, map[int]int{2: 3}},
		{9, map[int]int{3: 2}},
		{10, map[int]int{2: 1, 5: 1}},
		{11, map[int]int{11: 1}},
		{12, map[int]int{2: 2, 3: 1}},
		{28, map[int]int{2: 2, 7: 1}},
		{210, map[int]int{2: 1, 3: 1, 5: 1, 7: 1}},
		{2310, map[int]int{2: 1, 3: 1, 5: 1, 7: 1, 11: 1}},
	}

	for _, testCase := range testCases {
		answer := FactorsCounted(testCase.n)
		if len(answer) != len(testCase.expected) {
			t.Errorf("ERROR: For %d expected len=%d, got len=%d %v", testCase.n, len(testCase.expected), len(answer), answer)
		}
		for key := range testCase.expected {
			if answer[key] != testCase.expected[key] {
				t.Errorf("ERROR: For %d expected %v, got %v", testCase.n, testCase.expected, answer)
			}
		}
	}
}

func TestIsPalindromeString(t *testing.T) {
	testCases := []struct {
		c        string
		expected bool
	}{
		{"", true},
		{"w", true},
		{"aba", true},
		{"aab", false},
		{"-22-", true},
	}

	for _, testCase := range testCases {
		answer := IsPalindromeString(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIsPalindromeInt(t *testing.T) {
	testCases := []struct {
		c        []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{6, 4, 4, 6}, true},
	}

	for _, testCase := range testCases {
		answer := IsPalindromeInt(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		n        []int
		expected []int
	}{
		{[]int{5, 6, 0, 0, 3}, []int{3, 0, 0, 6, 5}},
		{[]int{2}, []int{2}},
		{[]int{2, 3}, []int{3, 2}},
	}

	for _, testCase := range testCases {
		answer := Reverse(testCase.n)
		if len(answer) != len(testCase.expected) {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.n, testCase.expected, answer)
		}
		for i := 0; i < len(testCase.expected); i++ {
			if answer[i] != testCase.expected[i] {
				t.Errorf("ERROR: For %v expected %v, got %v", testCase.n, testCase.expected, answer)
			}
		}
	}
}

func TestIntToDigits(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{56003, []int{5, 6, 0, 0, 3}},
		{2, []int{2}},
		{23, []int{2, 3}},
		{1230, []int{1, 2, 3, 0}},
		// {0, []int{0}},  // Not implemented yet.
	}

	for _, testCase := range testCases {
		answer := IntToDigits(testCase.n)
		if len(answer) != len(testCase.expected) {
			t.Errorf("ERROR: For %v expected %v, got %v", testCase.n, testCase.expected, answer)
		}
		for i := 0; i < len(testCase.expected); i++ {
			if answer[i] != testCase.expected[i] {
				t.Errorf("ERROR: For %v expected %v, got %v", testCase.n, testCase.expected, answer)
			}
		}
	}
}

func TestDigitsToInt(t *testing.T) {
	testCases := []struct {
		n        []int
		expected int
	}{
		{[]int{5, 6, 0, 0, 3}, 56003},
		{[]int{2}, 2},
		{[]int{2, 3}, 23},
		{[]int{1, 2, 3, 0}, 1230},
		{[]int{0}, 0},
	}

	for _, testCase := range testCases {
		answer := DigitsToInt(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %v expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestDigitSum(t *testing.T) {
	testCases := []struct {
		c        int
		expected int
	}{
		{0, 0},
		{5, 5},
		{10, 1},
		{25, 7},
		{100000, 1},
		{100001, 2},
	}

	for _, testCase := range testCases {
		answer := DigitSum(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.c, testCase.expected, answer)
		}
	}
}

func TestHarshad(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{201, true},
		{2011, false},
		{100000, true},
		{100001, false},
	}

	for _, testCase := range testCases {
		answer := Harshad(testCase.c, DigitSum(testCase.c))
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestTriangular(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{0, true},
		{1, true},
		{9, false},
		{10, true},
		{1000, false},
		{17526, false},
		{1000 * 1000, false},
	}

	for _, testCase := range testCases {
		answer := Triangular(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestTotient(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{2, 1}, // Totient() has had a lot of bugs. Use lots of test cases!
		{3, 2},
		{4, 2},
		{5, 4},
		{6, 2},
		{7, 6},
		{8, 4},
		{9, 6},
		{10, 4},
		{11, 10},
		{12, 4},
		{13, 12},
		{14, 6},
		{15, 8},
		{16, 8},
		{17, 16},
		{18, 6},
		{19, 18},
		{20, 8},
		{21, 12},
		{22, 10},
		{23, 22},
		{24, 8},
		{25, 20},
		{26, 12},
		{27, 18},
		{28, 12},
		{29, 28},
		{30, 8},
		{80, 32},
		{81, 54},
		{82, 40},
		{83, 82},
		{84, 24},
		{85, 64},
		{86, 42},
		{87, 56},
		{88, 40},
		{89, 88},
		{90, 24},
		{91, 72},
		{92, 44},
		{93, 60},
		{94, 46},
		{95, 72},
		{96, 32},
		{97, 96},
		{98, 42},
		{99, 60},
	}

	for _, testCase := range testCases {
		answer := Totient(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		w1       string
		w2       string
		expected bool
	}{
		{"", "", true},
		{"ab", "ba", true},
		{"ab", "ab", true},
		{"ignore", "region", true},
		{"aaaa", "aaa", false},
		{"dog", "gad", false},
	}

	for _, testCase := range testCases {
		answer := IsAnagram(testCase.w1, testCase.w2)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s/%s expected %t, got %t", testCase.w1, testCase.w2, testCase.expected, answer)
		}
	}
}

func TestCryptoquip(t *testing.T) {
	testCases := []struct {
		w1       string
		w2       string
		expected bool
	}{
		{"", "", true},
		{"feel", "felt", false},
		{"keep", "pool", false},
		{"keep", "loot", true},
		{"keep", "kelp", false},
		{"keep", "toot", false},
		{"abcddeeffaa", "12344556611", true},
		{"aaaa", "aaa", false},
		{"dog", "gad", true},
	}

	for _, testCase := range testCases {
		_, answer := Cryptoquip(testCase.w1, testCase.w2)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %s/%s expected %t, got %t", testCase.w1, testCase.w2, testCase.expected, answer)
		}
	}
}

func TestSquareFree(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{0, true},
		{1, true},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{6, true},
		{7, true},
		{8, false},
		{9, false},
		{10, true},
		{25, false},
		{100000, false},
		{100001, true},
	}

	for _, testCase := range testCases {
		answer := SquareFree(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestPascalTriangle(t *testing.T) {
	testCases := []struct {
		row      int
		col      int
		expected int
	}{
		{0, 0, 1},
		{1, 0, 1},
		{1, 1, 1},
		{2, 0, 1},
		{2, 1, 2},
		{2, 2, 1},
		{3, 0, 1},
		{3, 1, 3},
		{3, 2, 3},
		{3, 3, 1},
	}

	triangle := PascalTriangle(4)

	for _, testCase := range testCases {
		answer := triangle[testCase.row][testCase.col]
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d %d expected %d, got %d", testCase.row, testCase.col, testCase.expected, answer)
		}
	}
}
