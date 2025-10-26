package primes

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
	"encoding/gob"
	"fmt"
	"log"
	"math"
	"os"

	"github.com/erikbryant/util-golang/system"
)

var (
	// Primes is a list of the first n prime numbers
	Primes []int32
	// End is the index of the final value in the Primes slice
	End int
)

const (
	// maxPrime is the highest value up to which we will search for primes
	maxPrime = 100*1000*1000 + 1000
	gobName  = "primes.gob"
)

// init loads the primes into memory
func init() {
	//primes := MakePrimes(maxPrime)
	//Save(primes)
	fileName := system.MyPath(gobName)
	Primes = Load(fileName)
	End = len(Primes) - 1
}

// Iter returns an iterator over all Primes
func Iter() func(func(int, int) bool) {
	return Iterr(0, -1)
}

// Iterr returns an iterator over a range of Primes
func Iterr(start, end int) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		if end < 0 {
			end = len(Primes)
		}
		for i, prime := range Primes[start:end] {
			if !yield(i, int(prime)) {
				return
			}
		}
	}
}

// boundsCheck panics if n is outside the range of Primes
func boundsCheck(n int32) {
	if n <= 0 || n > Primes[End] {
		err := fmt.Errorf("exceeded max prime; did you call Init() n = %d", n)
		panic(err)
	}
}

// Pi is the prime counting function, returning the number of primes <= n
// https://en.wikipedia.org/wiki/Prime-counting_function
func Pi(n int) int {
	if n < 2 {
		return 0
	}

	i := Index(n)
	if i < 0 {
		i *= -1
	}

	return i + 1
}

// SlowPrime returns whether a number is prime or not, using a brute force search
func SlowPrime(n int) bool {
	if n <= 1 {
		return false
	}

	root := int32(math.Sqrt(float64(n)))

	// Check each potential divisor to see if number divides evenly (i.e., is not prime).
	boundsCheck(root)
	for i := 0; Primes[i] <= root; i++ {
		if n%int(Primes[i]) == 0 {
			return false
		}
	}

	return true
}

// Prime returns true if n is prime
func Prime(n int) bool {
	if n > int(Primes[End]) {
		return SlowPrime(n)
	}
	return Index(n) >= 0
}

// Index returns the index in Primes of n, or negative of the next highest index if not found
func Index(N int) int {
	if N <= 1 {
		return -1
	}

	n := int32(N)
	upper := End
	lower := 0

	for upper > lower {
		mid := (upper + lower) >> 1

		if n > Primes[mid] {
			if n < Primes[mid+1] {
				if Primes[mid] != n {
					// n is not prime
					return -1 * mid
				}
				return mid
			}
			lower = mid + 1
		} else {
			if n == Primes[mid] {
				if Primes[mid] != n {
					// n is not prime
					return -1 * mid
				}
				return mid
			}
			if mid == 0 {
				if Primes[mid] != n {
					// n is not prime
					return -1 * mid
				}
				return mid
			}
			upper = mid - 1
		}

	}

	if Primes[upper] != n {
		// n is not prime
		return -1 * upper
	}

	return upper
}

// MakePrimes returns all primes <= maxPrime
func MakePrimes(maxPrime uint) []int32 {
	// Sieve of Eratosthenes
	// Original Python Code by David Eppstein, UC Irvine, 28 Feb 2002
	// http://code.activestate.com/recipes/117119/
	// Found on:
	// https://stackoverflow.com/questions/567222/simple-prime-number-generator-in-python

	// Maps composites to primes witnessing their compositeness.
	// This is memory efficient, as the sieve is not "run forward"
	// indefinitely, but only as long as required by the current
	// number being tested.

	if maxPrime > 4294967296 {
		// We calculate q*q below; verify q*q will not overflow uint
		log.Fatal("maxPrime > sqrt(2^64 - 1)! ", maxPrime)
	}

	primes := []int32{}
	D := map[uint][]uint{}

	// The running integer that's checked for primeness
	for q := uint(2); ; q++ {
		_, ok := D[q]
		if !ok {
			if q > maxPrime {
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

// Save writes an int slice to the gob file
func Save(primes []int32) {
	file, err := os.Create(gobName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		panic(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	encoder.Encode(primes)
}

// Load returns the contents of the gob file as an int slice
func Load(name string) []int32 {
	file, err := os.Open(name)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		panic(err)
	}
	defer file.Close()

	primes := []int32{}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&primes)
	if err != nil {
		fmt.Printf("error reading primes gob: %v", err)
		panic(err)
	}

	return primes
}
