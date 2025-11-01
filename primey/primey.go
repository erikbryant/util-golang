package primey

// https://primes.utm.edu/howmany.html
//
// π(x) = approx # of primes <= x
//
//        x                  π(x)
//       10                     4
//      100                    25
//    1,000                   168
//   10,000                 1,229
//  100,000                 9,592
//     10^6                78,498
//     10^7               664,579
//     10^8             5,761,455
//     10^9            50,847,534
//    10^10           455,052,511
//    10^11         4,118,054,813
//    10^12        37,607,912,018
//    10^13       346,065,536,839
//    10^14     3,204,941,750,802
//    10^15    29,844,570,422,669
//    10^16   279,238,341,033,925

import (
	"fmt"
	"math"
)

// PrimeMax returns the largest prime in the list of primes
func PrimeMax() int {
	return primeMax
}

// Len returns the length of the list of primes
func Len() int {
	return primeCount
}

// Iter returns an iterator over all primes
func Iter() func(func(int, int) bool) {
	return Iterr(0, Len()-1)
}

// Iterr returns an iterator over a range of index values
func Iterr(start, end int) func(func(int, int) bool) {
	if end <= start {
		return func(yield func(int, int) bool) {}
	}

	if start < 0 || start >= Len() {
		err := fmt.Errorf("start index out of range 0 >= %d > %d ", start, Len())
		panic(err)
	}

	if end < 0 || end >= Len() {
		err := fmt.Errorf("end index out of range 0 >= %d > %d ", end, Len())
		panic(err)
	}

	// Initialize the starting point
	ctx := newContext(start)

	return func(yield func(int, int) bool) {
		// Yield the primes
		for i := start; i < end; i++ {
			prime := ctx.next()
			if !yield(i-start, prime) {
				return
			}
		}
	}
}

// Nth returns the value of the nth prime
func Nth(n int) int {
	ctx := newContext(n)
	return ctx.next()
}

// Index returns the index of the prime, or if p is not prime then the index below the next highest prime
func Index(p int) int {
	if p <= 5 {
		return []int{0, 0, 0, 1, 1, 2}[p]
	}

	iByte, iBit, ok, r := int2offset(p)
	adjust := 0
	if !ok || !bitIsSet(iByte, iBit) {
		// p is not a prime; find the next higher iBit
		for i, remainder := range bit2remainder {
			iBit = uint8(i)
			if remainder > r {
				break
			}
		}
		adjust = 1
	}

	return offset2index(iByte, iBit) - adjust
}

// Pi returns the number of primes up to and including n
func Pi(n int) int {
	if n < 2 {
		return 0
	}
	return Index(n) + 1
}

// Prime returns true if p is a prime
func Prime(p int) bool {
	if p <= 5 {
		return p == 2 || p == 3 || p == 5
	}

	if p > PrimeMax() {
		return PrimeSlow(p)
	}

	iByte, iBit, ok, _ := int2offset(p)
	return ok && bitIsSet(iByte, iBit)
}

// PrimeSlow returns whether a number is prime or not, used for primes > PrimeMax()
func PrimeSlow(n int) bool {
	if n <= 1 {
		return false
	}

	root := int(math.Sqrt(float64(n)))

	// Check each potential divisor to see if number divides evenly (i.e., is not prime).
	for _, prime := range Iter() {
		if prime > root {
			return true
		}
		if n%prime == 0 {
			return false
		}
	}

	return true
}
