package primey

// https://primes.utm.edu/howmany.html
//
// π(x) = approx # of primes <= x
//
//                          x                 π(x)
// 1                       10                    4
// 2                      100                   25
// 3                    1,000                  168
// 4                   10,000                1,229
// 5                  100,000                9,592
// 6                1,000,000               78,498
// 7               10,000,000              664,579
// 8              100,000,000            5,761,455
// 9            1,000,000,000           50,847,534
// 10          10,000,000,000          455,052,511
// 11         100,000,000,000        4,118,054,813
// 12       1,000,000,000,000       37,607,912,018
// 13      10,000,000,000,000      346,065,536,839
// 14     100,000,000,000,000    3,204,941,750,802
// 15   1,000,000,000,000,000   29,844,570,422,669
// 16  10,000,000,000,000,000  279,238,341,033,925

import (
	"fmt"
	"log"
	"math"
)

func initPrimes() {
	if primeCount == 0 {
		load()
	}
}

// PrimeMax returns the largest prime in the list of primes
func PrimeMax() int {
	initPrimes()
	return primeMax
}

// Len returns the length of the list of primes
func Len() int {
	initPrimes()
	return primeCount
}

// Iter returns an iterator over all Primes
func Iter() func(func(int, int) bool) {
	initPrimes()
	return Iterr(0, Len()-1)
}

// Iterr returns an iterator over a range of index values
func Iterr(start, end int) func(func(int, int) bool) {
	initPrimes()

	if start < 0 || start >= Len() {
		err := fmt.Errorf("start index out of range 0 >= %d > %d ", start, Len())
		panic(err)
	}

	if end < 0 || end >= Len() {
		err := fmt.Errorf("end index out of range 0 >= %d > %d ", end, Len())
		panic(err)
	}

	return func(yield func(int, int) bool) {
		// Initialize the starting point
		ctx := newContext(start)

		// Yield the primes
		for i := start; i < end; i++ {
			prime := ctx.next()
			if !yield(i-start, prime) {
				return
			}
		}
	}
}

// Iterp returns an iterator over a range of prime numbers
func Iterp(start, end int) func(func(int, int) bool) {
	initPrimes()

	if start < 2 || start > PrimeMax() {
		err := fmt.Errorf("start index out of range 2 >= %d >= %d ", start, PrimeMax())
		panic(err)
	}

	if end < 2 || end > PrimeMax() {
		err := fmt.Errorf("end index out of range 2 >= %d >= %d ", end, PrimeMax())
		panic(err)
	}

	return Iterr(Index(start), Index(end))
}

// Nth returns the nth prime
func Nth(n int) int {
	initPrimes()
	ctx := newContext(n)
	return ctx.next()
}

// Index returns the index of the prime, or if p is not prime then the index below the next highest prime
func Index(p int) int {
	initPrimes()

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

	return primesBelow(iByte, iBit) - adjust
}

// Pi returns the number of primes below (and including) n
func Pi(n int) int {
	if n < 2 {
		return 0
	}
	return Index(n) + 1
}

// Prime returns true if p is a prime
func Prime(p int) bool {
	initPrimes()

	if p <= 5 {
		return p == 2 || p == 3 || p == 5
	}

	if p > PrimeMax() {
		return SlowPrime(p)
	}

	iByte, iBit, ok, _ := int2offset(p)
	return ok && bitIsSet(iByte, iBit)
}

// SlowPrime returns whether a number is prime or not, using a brute force search
func SlowPrime(n int) bool {
	initPrimes()

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

// MakePrimes finds and returns all primes <= limit
func MakePrimes(limit uint) []int32 {
	// Sieve of Eratosthenes
	// Original Python Code by David Eppstein, UC Irvine, 28 Feb 2002
	// http://code.activestate.com/recipes/117119/
	// Found on:
	// https://stackoverflow.com/questions/567222/simple-prime-number-generator-in-python

	// Maps composites to primes witnessing their compositeness.
	// This is memory efficient, as the sieve is not "run forward"
	// indefinitely, but only as long as required by the current
	// number being tested.

	if limit > 4294967296 {
		// We calculate q*q below; verify q*q will not overflow uint
		log.Fatal("limit > sqrt(2^64 - 1)! ", limit)
	}

	primes := []int32{}
	D := map[uint][]uint{}

	// The running integer that's checked for primeness
	for q := uint(2); ; q++ {
		_, ok := D[q]
		if !ok {
			if q > limit {
				break
			}
			// q is a new prime.
			// Yield it and mark its first multiple that isn't
			// already marked in previous iterations
			primes = append(primes, int32(q))
			D[q*q] = []uint{q}
		} else {
			// q is composite. D[q] is the list of primes that
			// divide it. Since we've reached q, we no longer
			// need it in the map, but we'll mark the next
			// multiples of its witnesses to prepare for larger
			// numbers
			for _, p := range D[q] {
				_, ok := D[p+q]
				if !ok {
					D[p+q] = []uint{}
				}
				D[p+q] = append(D[p+q], p)
			}
			delete(D, q)
		}
	}

	return primes
}
