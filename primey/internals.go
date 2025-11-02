package primey

// Prime Wheel storage
//
// Concept from https://stackoverflow.com/a/15907156, implementation is all mine.
//
// All primes >5 are of the form:
//
//   30x+k where x is an integer >=0 and k ∈ {1, 7, 11, 13, 17, 19, 23, 29}
//
// Thus, primes from 1 to 30 can be stored in a single byte (ignoring 2, 3, 5),
// then primes from 31 to 60 in a single byte, then from 61 to 90, and so on.
// Handle 2, 3, and 5 separately.
//
// To create a new wheel gob file:
//
//   for _, prime := range FindPrimes(maxPrime) {
//     store(uint32(prime))
//   }
//   bake()
//   save()

import (
	"math/bits"
)

var (
	// bit2remainder maps n%30 to the corresponding bit position in wheel
	bit2remainder = []int{1, 7, 11, 13, 17, 19, 23, 29}

	// remainder2bit is the inverse of bit2remainder
	remainder2bit = []uint8{
		1:  0,
		7:  1,
		11: 2,
		13: 3,
		17: 4,
		19: 5,
		23: 6,
		29: 7,
	}

	// ---------- The prime cache ----------

	primeCache = []uint32{2, 3, 5}

	// ---------- Loaded from the gob file ----------

	// wheel is the compressed storage for primes
	wheel = []uint8{}

	// piCache is the cache of pre-computed counts
	piCache = []uint32{}

	// primeCount is the number of primes in wheel (this excludes {2, 3, 5})
	primeCount = 0

	// primeMax is the value of the largest prime in wheel
	primeMax = 0
)

const (
	// gobName is the name of the gob file the primes are stored in
	gobName = "wheel.gob"

	// piStep is the size of the blocks in the prime counts cache
	piStep = 100 // A smaller step makes Index faster, but piCache larger
)

func init() {
	load()
}

// int2offset returns the bit/byte in the wheel that the int corresponds to
func int2offset(p int) (int, uint8, bool, int) {
	iByte := p / 30
	r := p % 30
	if r == 1 {
		return iByte, 0, true, 1
	}
	iBit := remainder2bit[r]
	return iByte, iBit, iBit > 0, r
}

// offset2int returns the number corresponding to bit/byte
func offset2int(iByte int, iBit uint8) int {
	return iByte*30 + bit2remainder[iBit]
}

// setBit sets the given bit/byte in wheel
func setBit(iByte int, iBit uint8) {
	wheel[iByte] |= 1 << iBit
}

// bitIsSet returns true if the given bit is set in the given byte of the wheel
func bitIsSet(iByte int, iBit uint8) bool {
	return wheel[iByte]&(1<<iBit) != 0
}

func index2offset(index int) (int, uint8) {
	if index <= 2 {
		// Wheel does not store {2, 3, 5}
		return 0, 0
	}
	index -= 3

	iByte := 0
	iBit := uint8(0)

	// Jump forward to the step that index is in
	for s := 0; ; s++ {
		if int(piCache[s+1]) > index {
			iByte = s * piStep
			index -= int(piCache[s])
			break
		}
	}

	// Walk forward to the byte that index is in
	for {
		oc := bits.OnesCount8(wheel[iByte])
		if oc > index {
			break
		}
		iByte += 1
		index -= oc
	}

	// Crawl forward to the bit that index represents
	for b := uint8(0); index >= 0 && b <= 7; b++ {
		if !bitIsSet(iByte, b) {
			continue
		}
		iBit = b
		index -= 1
	}

	return iByte, iBit
}

// offset2index returns the index of this prime (i.e., the number of primes below the input). Input must be prime and must be > 5.
func offset2index(iByte int, iBit uint8) int {
	// Count 2, 3, and 5
	primesBelowP := 3

	// Each byte in the wheel represents 0-8 primes
	// Count bits on the way up to p

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

	return primesBelowP
}
