package algebra

import (
	"math"
	"math/big"

	"github.com/erikbryant/util-golang/primes"
)

// NextSquare returns the next highest square and its square root
func NextSquare(n uint) (uint, uint) {
	r := math.Sqrt(float64(n))
	root := uint(r)

	// If n is a square we can save a mul
	if float64(root) == r {
		next := n + root<<1 + 1
		return next, root + 1
	}

	next := root*root + root<<1 + 1
	return next, root + 1
}

type convergentSeries func(int) int64

// E returns the nth number (1-based) in the convergent series of e
func E(n int) int64 {
	// e.g., [2; 1,2,1, 1,4,1, 1,6,1, ... ,1,2k,1, ...]
	if n == 1 {
		return int64(2)
	}
	if n%3 == 0 {
		return int64(2 * n / 3)
	}
	return int64(1)
}

// Sqrt2 returns the nth number (1-based) in the convergent series of root 2
func Sqrt2(n int) int64 {
	// e.g., [2;(2)]

	if n == 1 {
		return int64(1)
	}
	return int64(2)
}

// Convergent returns the nth convergence of whichever series you pass in a function for
func Convergent(n int, fn convergentSeries) (*big.Int, *big.Int) {
	numerator := big.NewInt(fn(n))
	denominator := big.NewInt(1)

	for n > 1 {
		// Invert
		denominator, numerator = numerator, denominator

		// Add e(n-1)
		product := big.NewInt(fn(n - 1))
		product.Mul(product, denominator)
		numerator.Add(numerator, product)

		n--
	}

	return numerator, denominator
}

// Divisors returns a sorted list of all positive divisors of n
func Divisors(n int) []int {
	// Everything is divisible by 1
	d := []int{1}

	root := int(math.Sqrt(float64(n)))

	// Degenerate cases
	if root <= 1 {
		if n < 0 {
			return []int{}
		}
		if n == 0 || n == 1 {
			return d
		}
		d = append(d, n)
		return d
	}

	// Find the lower divisors
	for i := 2; i < root; i++ {
		if n%i == 0 {
			d = append(d, i)
		}
	}

	// Check for the special case of n being a perfect square
	var start int
	if root*root == n {
		d = append(d, root)
		start = len(d) - 2
	} else {
		if n%root == 0 {
			d = append(d, root)
		}
		start = len(d) - 1
	}

	// Add the upper divisors (the inverses of the lower divisors)
	for i := start; i >= 0; i-- {
		d = append(d, n/d[i])
	}

	return d
}

// Factors returns a sorted list of the unique prime factors of n (excluding n)
func Factors(n int) []int {
	if primes.Prime(n) {
		return []int{}
	}

	f := []int{}

	for i := 0; primes.PackedPrimes[i] <= n; i++ {
		if n%primes.PackedPrimes[i] == 0 {
			f = append(f, primes.PackedPrimes[i])
			n /= primes.PackedPrimes[i]
		}
	}

	return f
}

// FactorsCounted returns counts of how many times each factor divides into n
func FactorsCounted(n int) map[int]int {
	factors := make(map[int]int)

	// Find all the 2 factors, since they are quick
	for (n & 0x01) == 0 {
		factors[2]++
		n = n >> 1
		if n == 1 {
			return factors
		}
	}

	root := int(math.Sqrt(float64(n)))
	for i := 1; primes.PackedPrimes[i] <= root; i++ {
		p := primes.PackedPrimes[i]
		for n%p == 0 {
			factors[p]++
			n = n / p
			if n == 1 {
				return factors
			}
		}
	}

	// We did not find any factors for 'n',
	// so it must be prime.
	factors[n]++

	return factors
}

// MaxBigInt returns the largest of a or b
func MaxBigInt(a, b *big.Int) *big.Int {
	switch a.Cmp(b) {
	case -1:
		return b
	case 0:
		return a
	case 1:
		return a
	}

	return b
}

// MinBigInt returns the smallest of a or b
func MinBigInt(a, b *big.Int) *big.Int {
	switch a.Cmp(b) {
	case -1:
		return a
	case 0:
		return b
	case 1:
		return b
	}

	return a
}

// GCD returns the greatest common divisor of a and b
func GCD(a, b int) int {
	// https://en.wikipedia.org/wiki/Greatest_common_divisor

	if a == 0 && b == 0 {
		return 0
	}

	for a > 0 && b > 0 {
		a, b = min(a, b), max(a, b)%min(a, b)
	}

	return max(a, b)
}

// GCDBigInt returns the greatest common divisor of a and b
func GCDBigInt(a, b *big.Int) *big.Int {
	// https://en.wikipedia.org/wiki/Greatest_common_divisor

	zero := new(big.Int)

	if a.Cmp(zero) == 0 && b.Cmp(zero) == 0 {
		return zero
	}

	a2 := new(big.Int)
	b2 := new(big.Int)
	a2.Set(a)
	b2.Set(b)

	for a2.Cmp(zero) == 1 && b2.Cmp(zero) == 1 {
		z := new(big.Int)
		a2, b2 = MinBigInt(a2, b2), z.Mod(MaxBigInt(a2, b2), MinBigInt(a2, b2))
	}

	return MaxBigInt(a2, b2)
}

// LCM returns the least common multiple of a and b
func LCM(a, b int) int {
	// https://en.wikipedia.org/wiki/Least_common_multiple

	if a == 0 && b == 0 {
		return 0
	}

	return (a / GCD(a, b)) * b
}

// LCMBigInt returns the least common multiple of a and b
func LCMBigInt(a, b *big.Int) *big.Int {
	// https://en.wikipedia.org/wiki/Least_common_multiple

	zero := new(big.Int)

	if a.Cmp(zero) == 0 && b.Cmp(zero) == 0 {
		return zero
	}

	temp := new(big.Int)
	temp.Div(a, GCDBigInt(a, b))
	temp.Mul(temp, b)

	return temp
}

// ReduceFraction returns the lowest that n and d reduce to
func ReduceFraction(n, d int) (int, int) {
	gcd := GCD(n, d)
	return n / gcd, d / gcd
}

// ReduceFractionBigInt returns the lowest that n and d reduce to
func ReduceFractionBigInt(n, d *big.Int) (*big.Int, *big.Int) {
	gcd := GCDBigInt(n, d)
	n2 := new(big.Int)
	d2 := new(big.Int)
	n2.Set(n)
	d2.Set(d)
	return n2.Div(n, gcd), d2.Div(d, gcd)
}

// SumFraction returns the sum of the two fractions, still in fraction form
func SumFraction(n1, d1, n2, d2 int) (int, int) {
	lcm := LCM(d1, d2)
	n1Scalar := lcm / d1
	n2Scalar := lcm / d2
	return ReduceFraction((n1*n1Scalar)+(n2*n2Scalar), lcm)
}

// SumFractionBigInt returns the sum of the two fractions, still in fraction form
func SumFractionBigInt(n1, d1, n2, d2 *big.Int) (*big.Int, *big.Int) {
	lcm := LCMBigInt(d1, d2)
	temp1 := new(big.Int)
	temp2 := new(big.Int)

	n1Scalar := temp1.Div(lcm, d1)
	n2Scalar := temp2.Div(lcm, d2)

	temp1.Mul(n1, n1Scalar)
	temp2.Mul(n2, n2Scalar)

	return ReduceFractionBigInt(temp1.Add(temp1, temp2), lcm)
}

// MulFraction returns the product of the two fractions, still in fraction form
func MulFraction(n1, d1, n2, d2 int) (int, int) {
	n1, d1 = ReduceFraction(n1, d1)
	n2, d2 = ReduceFraction(n2, d2)
	a, b := ReduceFraction(n1*n2, d1*d2)
	return a, b
}

// MulFractionBigInt returns the product of the two fractions, still in fraction form
func MulFractionBigInt(n1, d1, n2, d2 *big.Int) (*big.Int, *big.Int) {
	tempN := new(big.Int)
	tempD := new(big.Int)

	tempN.Mul(n1, n2)
	tempD.Mul(d1, d2)

	return ReduceFractionBigInt(tempN, tempD)
}

// IsInt returns true if n is an integer
func IsInt(n float64) bool {
	return n == float64(int(n))
}

// IsSquare returns true if n is a square number
func IsSquare(n int) bool {
	root := math.Sqrt(float64(n))
	return IsInt(root)
}

// IsCube returns true if n is a cube number
func IsCube(n int) bool {
	root := math.Cbrt(float64(n))
	return IsInt(root)
}

// EqualIntSlice returns true if the two slices have identical contents
func EqualIntSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

// Reverse reverses the order of the elements in a slice
func Reverse(digits []int) []int {
	rev := make([]int, 0)

	for i := len(digits) - 1; i >= 0; i-- {
		rev = append(rev, digits[i])
	}

	return rev
}

// IntToDigits converts an int into a slice of its component digits
func IntToDigits(n int) []int {
	digits := make([]int, 0)

	for n > 0 {
		d := n % 10
		digits = append(digits, d)
		n = n / 10
	}

	return Reverse(digits)
}

// DigitsToInt converts a slice of digits to an int
func DigitsToInt(digits []int) int {
	number := 0

	for i := 0; i < len(digits); i++ {
		number += digits[i] * int(math.Pow(10.0, float64(len(digits)-1-i)))
	}

	return number
}

// DigitSum returns the sum of the digits in the number
func DigitSum(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

// Harshad returns true if n is divisible by the sum of its digits
func Harshad(n int) bool {
	return n%DigitSum(n) == 0
}

// Totient returns how many numbers k are relatively prime to n
func Totient(n int) int {
	// ... where  1 <= k < n. Relatively prime means that they have
	// no common divisors (other than 1). 1 is considered relatively
	// prime to all other numbers.
	//
	// From https://www.doc.ic.ac.uk/~mrh/330tutor/ch05s02.html
	//
	// The general formula to compute φ(n) is the following:
	//
	// If the prime factorisation of n is given by n =p1e1*...*pnen, then
	// φ(n) = n *(1 - 1/p1)* ... (1 - 1/pn).
	//
	// For example:
	//
	// 9 = 32, φ(9) = 9* (1-1/3) = 6
	//
	// 4 =22, φ(4) = 4* (1-1/2) = 2
	//
	// 15 = 3*5, φ(15) = 15* (1-1/3)*(1-1/5) = 15*(2/3)*(4/5) =8

	if primes.Prime(n) {
		return n - 1
	}

	factors := Factors(n)
	count := n

	for _, f := range factors {
		count /= f
		count *= f - 1
	}

	return count
}

// SquareFree returns true if no square of a prime divides n
func SquareFree(n int) bool {
	for _, prime := range primes.PackedPrimes {
		if prime > int(math.Sqrt(float64(n))) {
			break
		}

		if n%(prime*prime) == 0 {
			return false
		}
	}

	return true
}

// PascalTriangle returns a triangle of the max depth specified
func PascalTriangle(max int) [][]int {
	// We build the triangle left-justified. A cell is the sum of the cell above it
	// and the cell above and to the left.
	//
	//	1: 1
	//	2: 1 1
	//	3: 1 2 1
	//	4: 1 3 3 1
	//	5: 1 4 6 4 1

	rows := [][]int{}
	var row []int

	// Create the empty rows
	for i := 0; i < max; i++ {
		row = make([]int, i+1)
		rows = append(rows, row)
	}

	for i := 0; i < max; i++ {
		for j := range rows[i] {
			if j == 0 || j == len(rows[i])-1 {
				rows[i][j] = 1
				continue
			}
			rows[i][j] = rows[i-1][j] + rows[i-1][j-1]
		}
	}

	return rows
}
