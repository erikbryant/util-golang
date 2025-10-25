package primes

// https://primes.utm.edu/howmany.html
//
// pi(x) = approx # of primes <= x
//
//                          x                pi(x)
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
	"math"
	"os"

	"github.com/erikbryant/util-golang/system"
)

var (
	// Primes is a list of the first n prime numbers
	Primes []int
	// PrimesEnd is the index of the final value in the Primes slice
	PrimesEnd int
)

const (
	// maxPrime is the highest value up to which we will search for primes
	maxPrime = 100*1000*1000 + 1000
	gobName  = "primes.gob"
)

func init() {
	//primes := MakePrimes(maxPrime)
	//Save(primes)
	fileName := system.MyPath(gobName)
	Primes = Load(fileName)
	PrimesEnd = len(Primes) - 1
}

// Pi is the prime counting function, returning the number of primes below n
// https://en.wikipedia.org/wiki/Prime-counting_function
func Pi(n int) int {
	if n < Primes[0] {
		return 0
	}

	if n > Primes[PrimesEnd] {
		err := fmt.Errorf("pi(%d) exceeded max prime; did you call Init()", n)
		panic(err)
	}

	i := PackedIndex(n)
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

	root := int(math.Sqrt(float64(n)))

	if root > Primes[PrimesEnd] {
		err := fmt.Errorf("SlowPrime(%d) exceeded max prime; did you call Init()", n)
		panic(err)
	}

	// Check each potential divisor to see if number divides evenly (i.e., is not prime).
	for i := 0; Primes[i] <= root; i++ {
		if n%Primes[i] == 0 {
			return false
		}
	}

	return true
}

// Prime returns true if number is prime
func Prime(number int) bool {
	if number > Primes[PrimesEnd] {
		return SlowPrime(number)
	}
	return PackedIndex(number) >= 0
}

// PackedIndex returns the index in Primes of n, or -1 if not found
func PackedIndex(n int) int {
	if n <= 1 {
		return -1
	}

	upper := PrimesEnd
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
func MakePrimes(maxPrime int) []int {
	// Sieve of Eratosthenes
	// Original Python Code by David Eppstein, UC Irvine, 28 Feb 2002
	// http://code.activestate.com/recipes/117119/
	// Found on:
	// https://stackoverflow.com/questions/567222/simple-prime-number-generator-in-python

	// Maps composites to primes witnessing their compositeness.
	// This is memory efficient, as the sieve is not "run forward"
	// indefinitely, but only as long as required by the current
	// number being tested.

	primes := []int{}
	D := map[int][]int{}

	// The running integer that's checked for primeness
	for q := 2; ; q++ {
		_, ok := D[q]
		if !ok {
			if q > maxPrime {
				break
			}
			// q is a new prime.
			// Yield it and mark its first multiple that isn't
			// already marked in previous iterations
			primes = append(primes, q)
			D[q*q] = []int{q}
		} else {
			// q is composite. D[q] is the list of primes that
			// divide it. Since we've reached q, we no longer
			// need it in the map, but we'll mark the next
			// multiples of its witnesses to prepare for larger
			// numbers
			for _, p := range D[q] {
				_, ok := D[p+q]
				if !ok {
					D[p+q] = []int{}
				}
				D[p+q] = append(D[p+q], p)
				delete(D, q)
			}
		}
	}

	return primes
}

// Save writes an int slice to the gob file
func Save(primes []int) {
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
func Load(fName string) []int {
	file, err := os.Open(fName)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		panic(err)
	}
	defer file.Close()

	primes := []int{}
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&primes)
	if err != nil {
		fmt.Printf("error reading packedPrimes: %v", err)
		panic(err)
	}

	return primes
}
