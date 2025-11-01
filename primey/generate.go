package primey

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/bits"
	"os"

	"github.com/erikbryant/util-golang/system"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// makeWheel generates and saves a new wheel
func makeWheel() {
	// maxPrime is the highest value up to which we will search for primes
	maxPrime := uint(100*1000*1000 + 1000)

	primes := FindPrimes(maxPrime)
	for _, p := range primes {
		store(uint32(p))
	}
	bake()
	save()
}

// store writes the prime to wheel (be sure to call bake() when done storing)
func store(p uint32) {
	if p <= 5 {
		return
	}
	iByte, iBit, ok, _ := int2offset(int(p))
	if !ok {
		fmt.Printf("%d is not prime! Not storing.\n", p)
		return
	}

	if bitIsSet(iByte, iBit) {
		fmt.Printf("%d is already stored! Not storing.\n", p)
		return
	}

	primeMax = max(primeMax, int(p))
	primeCount++
	setBit(iByte, iBit)
}

// bake computes the derived values for a new wheel
func bake() {
	piCache = make([]uint32, primeCount/piStep)

	primeCount := uint32(0)
	for i := 0; i < len(wheel); i++ {
		primeCount += uint32(bits.OnesCount8(wheel[i]))
		k := i/piStep + 1
		piCache[k] = primeCount
	}

	p := message.NewPrinter(language.English)
	p.Printf("Wheel statistics\n")
	p.Printf("primeCount      = %16d\n", primeCount)
	p.Printf("primeMax        = %16d\n", primeMax)
	p.Printf("Wheel size      = %16d bytes\n", len(wheel))
	p.Printf("sizeof(piCache) = %16d bytes\n", len(piCache)*4)
	totalSizeWheel := len(wheel) + len(piCache)*4
	p.Printf("total size      = %16d bytes\n", totalSizeWheel)
	p.Printf("primes/byte     = %16f\n", float64(primeCount)/float64(totalSizeWheel))
}

// FindPrimes finds and returns all primes <= limit
func FindPrimes(limit uint) []int32 {
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

// save writes the wheel and its derived values to the gob file
func save() {
	file, err := os.Create(system.MyPath(gobName))
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		panic(err)
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)

	_ = encoder.Encode(wheel)
	_ = encoder.Encode(piCache)
	_ = encoder.Encode(primeCount)
	_ = encoder.Encode(primeMax)
}

// load reads the contents of the gob file into wheel and its derived values
func load() {
	file, err := os.Open(system.MyPath(gobName))
	if err != nil {
		fmt.Printf("error opening file: %v", err)
		panic(err)
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)

	err = decoder.Decode(&wheel)
	if err != nil {
		fmt.Printf("error reading primes gob - wheel: %v", err)
		panic(err)
	}

	err = decoder.Decode(&piCache)
	if err != nil {
		fmt.Printf("error reading primes gob - piCache: %v", err)
		panic(err)
	}

	err = decoder.Decode(&primeCount)
	if err != nil {
		fmt.Printf("error reading primes gob - primeCount: %v", err)
		panic(err)
	}

	err = decoder.Decode(&primeMax)
	if err != nil {
		fmt.Printf("error reading primes gob - primeMax: %v", err)
		panic(err)
	}
}
