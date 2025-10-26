package algebra

import (
	"math/big"
	"slices"
	"testing"
)

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
		{1000 * 1000, 1002001, 1001},
	}

	for _, testCase := range testCases {
		answer, answer2 := NextSquare(testCase.c)
		if answer != testCase.expected || answer2 != testCase.expected2 {
			t.Errorf("ERROR: For %d expected %d/%d, got %d/%d", testCase.c, testCase.expected, testCase.expected2, answer, answer2)
		}
	}
}

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
		if !slices.Equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %d expected %v, got %v", testCase.n, testCase.expected, answer)
		}
	}
}

func TestFactors(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int
	}{
		{2, []int{2}},
		{3, []int{3}},
		{4, []int{2}},
		{5, []int{5}},
		{6, []int{2, 3}},
		{7, []int{7}},
		{8, []int{2}},
		{9, []int{3}},
		{10, []int{2, 5}},
		{11, []int{11}},
		{12, []int{2, 3}},
		{20, []int{2, 5}},
		{28, []int{2, 7}},
		{210, []int{2, 3, 5, 7}},
		{2310, []int{2, 3, 5, 7, 11}},
	}

	for _, testCase := range testCases {
		answer := Factors(testCase.n)
		if !slices.Equal(answer, testCase.expected) {
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

func TestMaxBigInt(t *testing.T) {
	testCases := []struct {
		a, b     *big.Int
		expected *big.Int
	}{
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(0), big.NewInt(1)},
		{big.NewInt(0), big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(3), big.NewInt(5), big.NewInt(5)},
		{big.NewInt(2), big.NewInt(4), big.NewInt(4)},
		{big.NewInt(2), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(3), big.NewInt(3)},
		{big.NewInt(9), big.NewInt(28), big.NewInt(28)},
		{big.NewInt(200), big.NewInt(100), big.NewInt(200)},
	}

	for _, testCase := range testCases {
		answer := MaxBigInt(testCase.a, testCase.b)
		if answer.Cmp(testCase.expected) != 0 {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}

func TestMinBigInt(t *testing.T) {
	testCases := []struct {
		a, b     *big.Int
		expected *big.Int
	}{
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(2), big.NewInt(1)},
		{big.NewInt(3), big.NewInt(5), big.NewInt(3)},
		{big.NewInt(2), big.NewInt(4), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(3), big.NewInt(2)},
		{big.NewInt(9), big.NewInt(28), big.NewInt(9)},
		{big.NewInt(200), big.NewInt(100), big.NewInt(100)},
	}

	for _, testCase := range testCases {
		answer := MinBigInt(testCase.a, testCase.b)
		if answer.Cmp(testCase.expected) != 0 {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}

func TestGCD(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{1, 2, 1},
		{3, 5, 1},
		{2, 4, 2},
		{9, 28, 1},
		{200, 100, 100},
	}

	for _, testCase := range testCases {
		answer := GCD(testCase.a, testCase.b)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}

func TestGCDBigInt(t *testing.T) {
	testCases := []struct {
		a, b     *big.Int
		expected *big.Int
	}{
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(0), big.NewInt(1)},
		{big.NewInt(0), big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(2), big.NewInt(1)},
		{big.NewInt(3), big.NewInt(5), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(4), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(3), big.NewInt(1)},
		{big.NewInt(9), big.NewInt(28), big.NewInt(1)},
		{big.NewInt(200), big.NewInt(100), big.NewInt(100)},
	}

	for _, testCase := range testCases {
		answer := GCDBigInt(testCase.a, testCase.b)
		if answer.Cmp(testCase.expected) != 0 {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}

func TestLCM(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 2, 2},
		{2, 2, 2},
		{2, 4, 4},
		{2, 3, 6},
	}

	for _, testCase := range testCases {
		answer := LCM(testCase.a, testCase.b)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}

func TestLCMBigInt(t *testing.T) {
	testCases := []struct {
		a, b     *big.Int
		expected *big.Int
	}{
		{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(0), big.NewInt(0)},
		{big.NewInt(0), big.NewInt(1), big.NewInt(0)},
		{big.NewInt(1), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(2), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(4), big.NewInt(4)},
		{big.NewInt(2), big.NewInt(3), big.NewInt(6)},
	}

	for _, testCase := range testCases {
		answer := LCMBigInt(testCase.a, testCase.b)
		if answer.Cmp(testCase.expected) != 0 {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.a, testCase.b, testCase.expected, answer)
		}
	}
}

func TestReduceFraction(t *testing.T) {
	testCases := []struct {
		a, b      int
		expectedN int
		expectedD int
	}{
		// Already reduced
		{1, 2, 1, 2},
		{2, 7, 2, 7},
		// Need reducing
		{12, 24, 1, 2},
		{49, 7, 7, 1},
		{2, 30, 1, 15},
	}

	for _, testCase := range testCases {
		answer, answer2 := ReduceFraction(testCase.a, testCase.b)
		if answer != testCase.expectedN || answer2 != testCase.expectedD {
			t.Errorf("ERROR: For %d/%d expected %d/%d, got %d/%d", testCase.a, testCase.b, testCase.expectedN, testCase.expectedD, answer, answer2)
		}
	}
}

func TestReduceFractionBigInt(t *testing.T) {
	testCases := []struct {
		a, b      *big.Int
		expectedN *big.Int
		expectedD *big.Int
	}{
		// Already reduced
		{big.NewInt(1), big.NewInt(2), big.NewInt(1), big.NewInt(2)},
		{big.NewInt(2), big.NewInt(7), big.NewInt(2), big.NewInt(7)},
		// Need reducing
		{big.NewInt(12), big.NewInt(24), big.NewInt(1), big.NewInt(2)},
		{big.NewInt(49), big.NewInt(7), big.NewInt(7), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(30), big.NewInt(1), big.NewInt(15)},
	}

	for _, testCase := range testCases {
		answer, answer2 := ReduceFractionBigInt(testCase.a, testCase.b)
		if answer.Cmp(testCase.expectedN) != 0 || answer2.Cmp(testCase.expectedD) != 0 {
			t.Errorf("ERROR: For %d/%d expected %d/%d, got %d/%d", testCase.a, testCase.b, testCase.expectedN, testCase.expectedD, answer, answer2)
		}
	}
}

func TestSumFraction(t *testing.T) {
	testCases := []struct {
		a, b      int
		c, d      int
		expectedN int
		expectedD int
	}{
		// Equal denominators
		{1, 2, 3, 2, 2, 1},
		{2, 7, 5, 7, 1, 1},
		// Differing denominators
		{1, 2, 1, 3, 5, 6},
		{1, 2, 3, 10, 4, 5},
		{2, 30, 2, 72, 17, 180},
	}

	for _, testCase := range testCases {
		answer, answer2 := SumFraction(testCase.a, testCase.b, testCase.c, testCase.d)
		if answer != testCase.expectedN || answer2 != testCase.expectedD {
			t.Errorf("ERROR: For %d/%d + %d/%d expected %d/%d, got %d/%d", testCase.a, testCase.b, testCase.c, testCase.d, testCase.expectedN, testCase.expectedD, answer, answer2)
		}
	}
}

func TestSumFractionBigInt(t *testing.T) {
	testCases := []struct {
		a, b      *big.Int
		c, d      *big.Int
		expectedN *big.Int
		expectedD *big.Int
	}{
		// Equal denominators
		{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(2), big.NewInt(2), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(7), big.NewInt(5), big.NewInt(7), big.NewInt(1), big.NewInt(1)},
		// Differing denominators
		{big.NewInt(1), big.NewInt(2), big.NewInt(1), big.NewInt(3), big.NewInt(5), big.NewInt(6)},
		{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(10), big.NewInt(4), big.NewInt(5)},
		{big.NewInt(2), big.NewInt(30), big.NewInt(2), big.NewInt(72), big.NewInt(17), big.NewInt(180)},
	}

	for _, testCase := range testCases {
		answer, answer2 := SumFractionBigInt(testCase.a, testCase.b, testCase.c, testCase.d)
		if answer.Cmp(testCase.expectedN) != 0 || answer2.Cmp(testCase.expectedD) != 0 {
			t.Errorf("ERROR: For %d/%d + %d/%d expected %d/%d, got %d/%d", testCase.a, testCase.b, testCase.c, testCase.d, testCase.expectedN, testCase.expectedD, answer, answer2)
		}
	}
}

func TestMulFraction(t *testing.T) {
	testCases := []struct {
		a, b      int
		c, d      int
		expectedN int
		expectedD int
	}{
		{3, 5, 0, 2, 0, 1},
		{2, 2, 3, 3, 1, 1},
		{1, 2, 3, 2, 3, 4},
		{5, 7, 2, 8, 5, 28},
		{6, 9, 1, 3, 2, 9},
		{3, 4, 3, 12, 3, 16},
		{2, 30, 2, 72, 1, 540},
	}

	for _, testCase := range testCases {
		answer, answer2 := MulFraction(testCase.a, testCase.b, testCase.c, testCase.d)
		if answer != testCase.expectedN || answer2 != testCase.expectedD {
			t.Errorf("ERROR: For %d/%d * %d/%d expected %d/%d, got %d/%d", testCase.a, testCase.b, testCase.c, testCase.d, testCase.expectedN, testCase.expectedD, answer, answer2)
		}
	}
}

func TestMulFractionBigInt(t *testing.T) {
	testCases := []struct {
		a, b      *big.Int
		c, d      *big.Int
		expectedN *big.Int
		expectedD *big.Int
	}{
		{big.NewInt(3), big.NewInt(5), big.NewInt(0), big.NewInt(2), big.NewInt(0), big.NewInt(1)},
		{big.NewInt(2), big.NewInt(2), big.NewInt(3), big.NewInt(3), big.NewInt(1), big.NewInt(1)},
		{big.NewInt(1), big.NewInt(2), big.NewInt(3), big.NewInt(2), big.NewInt(3), big.NewInt(4)},
		{big.NewInt(5), big.NewInt(7), big.NewInt(2), big.NewInt(8), big.NewInt(5), big.NewInt(28)},
		{big.NewInt(6), big.NewInt(9), big.NewInt(1), big.NewInt(3), big.NewInt(2), big.NewInt(9)},
		{big.NewInt(3), big.NewInt(4), big.NewInt(3), big.NewInt(12), big.NewInt(3), big.NewInt(16)},
		{big.NewInt(2), big.NewInt(30), big.NewInt(2), big.NewInt(72), big.NewInt(1), big.NewInt(540)},
	}

	for _, testCase := range testCases {
		answer, answer2 := MulFractionBigInt(testCase.a, testCase.b, testCase.c, testCase.d)
		if answer.Cmp(testCase.expectedN) != 0 || answer2.Cmp(testCase.expectedD) != 0 {
			t.Errorf("ERROR: For %d/%d * %d/%d expected %d/%d, got %d/%d", testCase.a, testCase.b, testCase.c, testCase.d, testCase.expectedN, testCase.expectedD, answer, answer2)
		}
	}
}

func TestIsInt(t *testing.T) {
	testCases := []struct {
		c        float64
		expected bool
	}{
		{-3.000, true},
		{-1.0, true},
		{0.0, true},
		{0.01, false},
		{1.0, true},
		{3.1415, false},
		{9999.0, true},
	}

	for _, testCase := range testCases {
		answer := IsInt(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %f expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIsSquare(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{-4, false},
		{0, true},
		{1, true},
		{2, false},
		{4, true},
		{1000000, true},
	}

	for _, testCase := range testCases {
		answer := IsSquare(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIsCube(t *testing.T) {
	testCases := []struct {
		c        int
		expected bool
	}{
		{-3, false},
		{-1, true},
		{0, true},
		{1, true},
		{2, false},
		{8, true},
		{125, true},
		{27000, true},
	}

	for _, testCase := range testCases {
		answer := IsCube(testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.c, testCase.expected, answer)
		}
	}
}

func TestIntToDigits(t *testing.T) {
	testCases := []struct {
		n        int
		expected []int8
	}{
		{56003, []int8{5, 6, 0, 0, 3}},
		{2, []int8{2}},
		{23, []int8{2, 3}},
		{1230, []int8{1, 2, 3, 0}},
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
		n        []int8
		expected int
	}{
		{[]int8{5, 6, 0, 0, 3}, 56003},
		{[]int8{2}, 2},
		{[]int8{2, 3}, 23},
		{[]int8{1, 2, 3, 0}, 1230},
		{[]int8{0}, 0},
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
		answer := Harshad(testCase.c)
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
		{0, 0}, // Totient() has had a lot of bugs. Use lots of test cases!
		{1, 1},
		{2, 1},
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
		{100, 40},
	}

	for _, testCase := range testCases {
		answer := Totient(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
		}
	}
}

func TestTotients(t *testing.T) {
	testCases := []struct {
		n        int
		expected int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
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
		{100, 40},
	}

	totients := Totients(100)

	for _, testCase := range testCases {
		answer := totients[testCase.n]
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
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

func TestKSmooth(t *testing.T) {
	testCases := []struct {
		n        int
		k        int
		expected bool
	}{
		// Out of bounds
		{0, 0, false},
		{0, 1, false},
		{1, 0, false},
		{1, 1, false},

		// Bad base
		{12, 4, false},

		// Valid data
		{1, 2, true},
		{5, 5, true},
		{5, 7, true},
		{700, 3, false},
		{700, 5, false},
		{700, 7, true},
	}

	for _, testCase := range testCases {
		answer := KSmooth(testCase.n, testCase.k)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d %d expected %t, got %t", testCase.n, testCase.k, testCase.expected, answer)
		}
	}
}

func TestHamming(t *testing.T) {
	testCases := []struct {
		n        int
		expected bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
		{7, false},
		{8, true},
		{9, true},
		{10, true},
		{11, false},
		{12, true},
		{13, false},
		{14, false},
		{15, true},
		{16, true},
		{18, true},
		{20, true},
		{24, true},
		{25, true},
		{27, true},
		{30, true},
		{32, true},
		{36, true},
		{40, true},
		{45, true},
		{48, true},
		{50, true},
		{54, true},
		{60, true},
		{64, true},
		{72, true},
		{75, true},
		{80, true},
		{81, true},
		{90, true},
		{96, true},
		{99, false},
		{100, true},
		{101, false},
	}

	for _, testCase := range testCases {
		answer := Hamming(testCase.n)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d expected %t, got %t", testCase.n, testCase.expected, answer)
		}
	}
}

func TestHammings(t *testing.T) {
	testCases := []int{
		1,
		2,
		3,
		4,
		5,
		6,
		8,
		9,
		10,
		12,
		15,
		16,
		18,
		20,
		24,
		25,
		27,
		30,
		32,
		36,
		40,
		45,
		48,
		50,
		54,
		60,
		64,
		72,
		75,
		80,
		81,
		90,
		96,
		100,
	}

	hammings := Hammings(100)

	expected := 34
	answer := len(hammings)
	if answer != expected {
		t.Errorf("ERROR: For len(hammings) expected %d, got %d", answer, expected)
	}

	for i, expected := range testCases {
		answer := hammings[i]
		if answer != expected {
			t.Errorf("ERROR: For hammings[%d] expected %d, got %d", i, expected, answer)
		}
	}
}

func TestPowerMod(t *testing.T) {
	testCases := []struct {
		b        int
		e        int
		m        int
		expected int
	}{
		{10, 1, 10, 0},

		// Primes 2,5 return 0
		{10, 1000 * 1000 * 1000, 2, 0},
		{10, 1000 * 1000 * 1000, 5, 0},

		{10, 1000 * 1000 * 1000, 3, 1},
		{10, 1000 * 1000 * 1000, 41, 1},
	}

	for _, testCase := range testCases {
		answer := PowerMod(testCase.b, testCase.e, testCase.m)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d^%d mod %d expected %d, got %d", testCase.b, testCase.e, testCase.m, testCase.expected, answer)
		}
	}
}

func TestCanPythagoreanTriples(t *testing.T) {
	testCases := []struct {
		c        int
		expected [][3]int
	}{
		{0, [][3]int{}},
		{1, [][3]int{}},
		{5, [][3]int{
			{3, 4, 5},
		}},
		{65, [][3]int{
			{3, 4, 5},
			{5, 12, 13},
			{8, 15, 17},
			{7, 24, 25},
			{20, 21, 29},
			{12, 35, 37},
			{9, 40, 41},
			{28, 45, 53},
			{11, 60, 61},
			{16, 63, 65},
			{33, 56, 65},
		}},
	}

	for _, testCase := range testCases {
		answer := PythagoreanTriples(testCase.c)
		if !slices.Equal(answer, testCase.expected) {
			t.Errorf("ERROR: For %d expected %v, got %v", testCase.c, testCase.expected, answer)
		}
	}
}
