package prime

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
	"math"
	"math/bits"
)

// Len returns the length of the list of primes
func Len() int {
	return primeCount
}

// Iter returns an iterator over all Primes
func Iter() func(func(int, int) bool) {
	return Iterr(0, Len()-1)
}

// Iterr returns an iterator over a range of Primes
func Iterr(start, end int) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		// Initialize the starting point
		ctx := newContext(start)

		if end >= start {
			// Yield the primes in forward order
			for i := start; i < end; i++ {
				prime := ctx.next()
				if ctx.atEnd() || !yield(i-start, prime) {
					return
				}
			}
		} else {
			// Yield the primes in reverse order
			for i := start; i > end; i-- {
				prime := ctx.prev()
				if ctx.atEnd() || !yield(start-i, prime) {
					return
				}
			}
		}
	}
}

// Index returns the index of the given number in the sorted list of primes
func Index(p int) int {
	if p <= 5 {
		return []int{-1, -1, 0, 1, -1, 2}[p]
	}

	adjusted := false

	// Each byte in the wheel represents 0-7 primes
	// Count bits on the way up to p

	iByte, iBit, ok := int2offset(p)
	if !ok || !bitIsSet(iByte, iBit) {
		// p is not a prime; find the next higher prime
		p, iByte, iBit = nextHigherPrime(p, iByte)
		adjusted = true
	}

	// Count 2, 3, and 5
	primesBelowP := 3

	// Count the primes in bytes below p
	primesBelowP += int(piCache[iByte/piStep])
	for b := iByte - iByte%piStep; b < iByte; b++ {
		primesBelowP += bits.OnesCount8(wheel[b])
	}

	// Count the primes in bits below p
	for m := uint8(0); m < iBit; m++ {
		if bitIsSet(iByte, m) {
			primesBelowP++
		}
	}

	if adjusted {
		return -(primesBelowP - 1)
	}

	return primesBelowP
}

// Pi returns the number of primes below (and including) n
func Pi(n int) int {
	if n < 2 {
		return 0
	}
	i := Index(n)
	if i < 0 {
		i = -i
	}
	return i + 1
}

// Prime returns true if p is a prime
func Prime(p int) bool {
	if p <= 5 {
		return p == 2 || p == 3 || p == 5
	}
	iByte, iBit, ok := int2offset(p)
	return ok && bitIsSet(iByte, iBit)
}

// SlowPrime returns whether a number is prime or not, using a brute force search
func SlowPrime(n int) bool {
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
