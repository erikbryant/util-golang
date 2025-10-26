package algebra

import (
	"math"
	"math/big"
	"slices"
	"sort"

	"github.com/erikbryant/util-golang/common"
	"github.com/erikbryant/util-golang/primes"
)

// NextSquare returns the next highest square and its square root
func NextSquare[T common.Numbers](n T) (uint, uint) {
	r := math.Sqrt(float64(n))
	root := uint(r)

	// If n is a square we can save a mul
	if float64(root) == r {
		next := uint(n) + root<<1 + 1
		return next, root + 1
	}

	next := root*root + root<<1 + 1
	return next, root + 1
}

// E returns the nth number (1-based) in the convergent series of e
func E[T common.Integers](n T) int64 {
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
func Sqrt2[T common.Integers](n T) int64 {
	// e.g., [2;(2)]
	if n == 1 {
		return int64(1)
	}
	return int64(2)
}

type convergentSeries func(int) int64

// Convergent returns the nth convergence of whichever series you pass in a function for
func Convergent[T common.Integers](n T, fn convergentSeries) (*big.Int, *big.Int) {
	numerator := big.NewInt(fn(int(n)))
	denominator := big.NewInt(1)

	for ; n > 1; n-- {
		// Invert
		denominator, numerator = numerator, denominator

		// Add e(n-1)
		product := big.NewInt(fn(int(n) - 1))
		product.Mul(product, denominator)
		numerator.Add(numerator, product)
	}

	return numerator, denominator
}

// Divisors returns a sorted list of all positive divisors of n
func Divisors[T common.Integers](n T) []T {
	// Everything is divisible by 1
	d := []T{1}

	// Degenerate cases
	if n <= 3 {
		// We are cheating here. Zero actually has an infinite number
		// of divisors [1..infinity]. We are just going to return [1].
		if n <= 1 {
			return d
		}
		d = append(d, n)
		return d
	}

	root := T(math.Sqrt(float64(n)))

	// Find the lower divisors
	for i := T(2); i < root; i++ {
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

// Factors returns a sorted list of the unique prime factors of n
func Factors(n int) []int {
	if primes.Prime(n) {
		return []int{n}
	}

	f := []int{}

	for i := 0; int(primes.Primes[i]) <= n; i++ {
		if n%int(primes.Primes[i]) == 0 {
			f = append(f, int(primes.Primes[i]))
			n /= int(primes.Primes[i])
			for n%int(primes.Primes[i]) == 0 {
				n /= int(primes.Primes[i])
			}
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
	}
	if n == 1 {
		return factors
	}

	root := int(math.Sqrt(float64(n)))
	for i := 1; int(primes.Primes[i]) <= root; i++ {
		p := int(primes.Primes[i])
		for n%p == 0 {
			factors[p]++
			n = n / p
		}
		if n == 1 {
			return factors
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
func GCD[T common.Integers](a, b T) T {
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
func LCM[T common.Integers](a, b T) T {
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
func ReduceFraction[T common.Integers](n, d T) (T, T) {
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
func SumFraction[T common.Integers](n1, d1, n2, d2 T) (T, T) {
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
func MulFraction[T common.Integers](n1, d1, n2, d2 T) (T, T) {
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
func IsInt[T common.Floats](n T) bool {
	return n == T(int(n))
}

// IsSquare returns true if n is a square number
func IsSquare[T common.Integers](n T) bool {
	root := math.Sqrt(float64(n))
	return IsInt(root)
}

// IsCube returns true if n is a cube number
func IsCube[T common.Integers](n T) bool {
	root := math.Cbrt(float64(n))
	return IsInt(root)
}

// IntToDigits converts an integer into a slice of its component digits
func IntToDigits[T common.Integers](n T) []int8 {
	digits := make([]int8, 0)

	for n > 0 {
		d := n % 10
		digits = append(digits, int8(d))
		n = n / 10
	}

	slices.Reverse(digits)
	return digits
}

// DigitsToInt converts a slice of digits to an int
func DigitsToInt(digits []int8) int {
	number := 0

	for i := 0; i < len(digits); i++ {
		number += int(digits[i]) * int(math.Pow(10.0, float64(len(digits)-1-i)))
	}

	return number
}

// DigitSum returns the sum of the digits in the number
func DigitSum[T common.Integers](n T) T {
	sum := T(0)

	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return sum
}

// Harshad returns true if n is divisible by the sum of its digits
func Harshad[T common.Integers](n T) bool {
	return n%DigitSum(n) == 0
}

// Totient returns how many numbers k are relatively prime to n
func Totient(n int) int {
	// ... where  1 <= k < n. Relatively prime means that they have
	// no common divisors (other than 1). 1 is considered relatively
	// prime to all other numbers.
	//
	// From https://en.wikipedia.org/wiki/Euler%27s_totient_function
	//
	// Given the prime factors of n are p1, p2, ... pk:
	// φ(n) = n * (1 - 1/p1) * (1 - 1/p2) * ... (1 - 1/pk)
	//      = n * (p1-1)/p1 * (p2-1)/p2 * ... (pk-1)/pk
	//
	// For example:
	//
	// Prime factors of  4 =   {2},  φ(4) =  4 * (1-1/2) = 2
	// Prime factors of  9 =   {3},  φ(9) =  9 * (1-1/3) = 6
	// Prime factors of 12 = {2,3}, φ(12) = 12 * (1-1/2) * (1-1/3) = 12 * 1/2 * 2/3 = 4
	// Prime factors of 15 = {3,5}, φ(15) = 15 * (1-1/3) * (1-1/5) = 15 * 2/3 * 4/5 = 8

	if primes.Prime(n) {
		return n - 1
	}

	count := n

	for _, f := range Factors(n) {
		count /= f
		count *= f - 1
	}

	return count
}

// Totients returns a slice where a[x] =ɸ(x) for 0 <= x <= upper
func Totients(upper int) []int {
	totients := make([]int, upper+1)

	for i := range totients {
		totients[i] = i
	}

	// Sieve of Eratosthenes

	// Fast mode
	for _, prime := range primes.Iter() {
		if prime > upper {
			break
		}
		for y := prime; y <= upper; y += prime {
			totients[y] -= totients[y] / prime
		}
	}

	// If we ran out of pre-computed primes, switch to slow mode
	for x := int(primes.Primes[len(primes.Primes)-1]) + 1; x <= upper; x++ {
		if totients[x] == x {
			for y := x; y <= upper; y += x {
				totients[y] -= totients[y] / x
			}
		}
	}

	return totients
}

// SquareFree returns true if no square of a prime divides n
func SquareFree(n int) bool {
	for _, prime := range primes.Iter() {
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

// KSmooth returns true if n is a k-smooth number
func KSmooth(n, k int) bool {
	if n < 1 || k < 2 || !primes.Prime(k) {
		// Invalid input
		return false
	}

	for _, prime := range primes.Iter() {
		if prime > k {
			break
		}
		for n%prime == 0 {
			n /= prime
		}
	}

	return n == 1
}

// Hamming returns true if n is a Hamming number (a 5-smooth number)
func Hamming(n int) bool {
	return KSmooth(n, 5)
}

func minimum[T common.Integers](s []T) T {
	m := s[0]
	for _, i := range s {
		m = min(m, i)
	}
	return m
}

// KSmooths returns a sorted list of all k-smooth numbers <= n
func KSmooths(n, k int) []int {
	// https://rosettacode.org/wiki/Hamming_numbers#Go
	h := []int{1}
	nexts := []int{}
	indices := []int{}

	for i := 0; int(primes.Primes[i]) <= k; i++ {
		nexts = append(nexts, int(primes.Primes[i]))
		indices = append(indices, 0)
	}

	for m := 1; ; m++ {
		next := minimum(nexts)
		if next > n {
			break
		}
		h = append(h, next)
		for i := 0; i < len(nexts); i++ {
			if h[m] == nexts[i] {
				indices[i]++
				nexts[i] = int(primes.Primes[i]) * h[indices[i]]
			}
		}
	}

	return h
}

// Hammings returns a sorted list of all Hamming numbers <= n
func Hammings(n int) []int {
	return KSmooths(n, 5)
}

// PowerMod returns (base^exp)%mod
func PowerMod(base, exp, mod int) int {
	// https://rosettacode.org/wiki/Modular_exponentiation#Python
	x := 1

	for exp > 0 {
		if exp%2 != 0 {
			x = (base * x) % mod
		}
		base = (base * base) % mod
		exp /= 2
	}

	return x
}

// PythagoreanTriples returns Pythagorean Triple primitives [a,b,c] where c <= upper
func PythagoreanTriples(upper int) [][3]int {
	pt := [][3]int{}

	sortFunc := func(i int, j int) bool {
		a, _, c := 0, 1, 2
		// Primary sort on c
		if pt[i][c] < pt[j][c] {
			return true
		}
		if pt[i][c] > pt[j][c] {
			return false
		}
		// pt[i][c] == pt[j][c], secondary sort on a
		return pt[i][a] < pt[j][a]
	}

	var a, b, c int
	for m := 2; ; m++ {
		for n := 1; n < m; n++ {
			if GCD(m, n) != 1 {
				continue
			}
			if m&0x01 == n&0x01 {
				continue
			}
			a = m*m - n*n
			b = 2 * m * n
			c = m*m + n*n
			// These are not generated in sorter order. Keep generating
			// too far to ensure we have found all.
			if a >= upper {
				sort.Slice(pt, sortFunc)
				return pt
			}
			if c > upper {
				continue
			}
			a, b = min(a, b), max(a, b)
			pt = append(pt, [3]int{a, b, c})
		}
	}
}
