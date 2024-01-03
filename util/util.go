package util

import (
	"fmt"
	"math"
	"math/big"
	"os"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"syscall"

	"github.com/erikbryant/project-euler/golang/primes"
)

func init() {
	filepath := path.Join(MyPath(), "../primes.gob")
	primes.Load(filepath)
}

// MyPath returns the absolute path of the source file calling this
func MyPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Cannot get caller information")
	}
	return path.Dir(filename)
}

// CtrlT prints a debugging message when SIGUSR1 is sent to this process.
func CtrlT(str string, val *int, digits []int) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)

	fmt.Println("$ kill -SIGUSR1", os.Getpid())

	go func() {
		for {
			<-c
			fmt.Println("^T] ", str, *val, digits)
		}
	}()
}

type convergentSeries func(int) int64

// E returns the nth number (1-based) in the convergent series
// of the number e [2; 1,2,1, 1,4,1, 1,6,1, ... ,1,2k,1, ...]
func E(n int) int64 {
	if n == 1 {
		return int64(2)
	}
	if n%3 == 0 {
		return int64(2 * n / 3)
	}
	return int64(1)
}

// Sqrt2 returns the nth number (1-based) in the convergent series
// of the square root of 2: [2;(2)]
func Sqrt2(n int) int64 {
	if n == 1 {
		return int64(1)
	}
	return int64(2)
}

// Convergent returns the nth convegence of whichever series you pass in a function for.
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

// Factors returns a sorted list of the unique prime factors of n (excluding n).
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

// FactorsCounted returns a map of prime factors of n with counts
// of how many times each factor divides into n.
func FactorsCounted(n int) map[int]int {
	factors := make(map[int]int)

	// Find all of the 2 factors, since they are quick
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

// IsSquare returns true if f is a square
func IsSquare(n int) bool {
	root := math.Sqrt(float64(n))
	return root == math.Trunc(root)
}

// heapPermutation generates a permutation using Heap Algorithm
// https://www.geeksforgeeks.org/heaps-algorithm-for-generating-permutations/
func heapPermutation(digits []int, size int, c chan []int) {
	if size == 1 {
		var temp []int
		for i := 0; i < len(digits); i++ {
			temp = append(temp, digits[i])
		}
		c <- temp
		return
	}

	for i := 0; i < size; i++ {
		heapPermutation(digits, size-1, c)

		// if size is odd, swap first and last element
		// If size is even, swap ith and last element
		swap := 0
		if size%2 == 0 {
			swap = i
		}
		digits[swap], digits[size-1] = digits[size-1], digits[swap]
	}
}

// MakeDigits generates all permutations of the first n digits.
// For example:
//
//	n=2 [1 2] [2 1]
//	n=3 [1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]
func MakeDigits(n int, c chan []int) {
	defer close(c)

	var digits []int
	for i := 1; i <= n; i++ {
		digits = append(digits, i)
	}

	heapPermutation(digits, len(digits), c)
}

// IsPalindromeString returns true if the string is a palindrome
func IsPalindromeString(p string) bool {
	head := 0
	tail := len(p) - 1

	for head < tail {
		if p[head] != p[tail] {
			return false
		}
		head++
		tail--
	}

	return true
}

// IsPalindromeInt returns true if the digits of p are a palindrome
func IsPalindromeInt(p []int) bool {
	head := 0
	tail := len(p) - 1

	for head < tail {
		if p[head] != p[tail] {
			return false
		}
		head++
		tail--
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

// DigitSum returns the sum of the digits in the number.
func DigitSum(n int) (sum int) {
	for n > 0 {
		sum += n % 10
		n /= 10
	}

	return
}

// Harshad returns true if n is divisible by the sum of its digits.
func Harshad(n, sum int) bool {
	return n%sum == 0
}

// Triangular returns true if n is a trianglar number
func Triangular(n int) bool {
	// n is triangular if 8*n+1 is a square
	root := math.Sqrt(float64(n<<3 + 1))
	return root == math.Trunc(root)
}

// Totient returns how many numbers k are relatively prime to n where
// 1 <= k < n. Relatively prime means that they have no common divisors (other
// than 1). 1 is considered relatively prime to all other numbers.
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
func Totient(n int) int {
	if primes.Prime(n) {
		return n - 1
	}

	factors := Factors(n)
	count := n

	for _, f := range factors {
		count /= f
		count *= (f - 1)
	}

	return count
}

// Equal returns true if the two slices have identical contents
func Equal(a, b []int) bool {
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

// IsAnagram returns true if w1 and w2 are anagrams of each other
func IsAnagram(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}

	for _, c := range w1 {
		w2 = strings.Replace(w2, string(c), "", 1)
	}

	return w2 == ""
}

// Cryptoquip returns whether the two strings have the same relative
// arrangement of letters. For instance, KEEP and LOOT.
func Cryptoquip(w1, w2 string) (map[byte]byte, bool) {
	if len(w1) != len(w2) {
		return nil, false
	}

	substitutions := make(map[byte]byte)

	for i := 0; i < len(w1); i++ {
		if val, ok := substitutions[w1[i]]; ok {
			if val != w2[i] {
				return nil, false
			}
			continue
		}
		if val, ok := substitutions[w2[i]]; ok {
			if val != w1[i] {
				return nil, false
			}
			continue
		}
		substitutions[w1[i]] = w2[i]
		substitutions[w2[i]] = w1[i]
	}

	return substitutions, true
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
// We build the triangle left-justified. A cell is the sum of the cell above it
// and the cell above and to the left.
//
//	1: 1
//	2: 1 1
//	3: 1 2 1
//	4: 1 3 3 1
//	5: 1 4 6 4 1
func PascalTriangle(max int) [][]int {
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
