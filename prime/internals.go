package prime

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

import (
	"encoding/gob"
	"fmt"
	"math/bits"
	"os"

	"github.com/erikbryant/util-golang/system"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

	// Loaded from the gob file ...

	// wheel is the compressed storage for primes
	wheel = make([]uint8, maxPrime/30)

	// piCache is the cache of pre-computed counts
	piCache = []uint32{}

	// primeCount is the number of primes in wheel (this excludes {2, 3, 5})
	primeCount = 0

	// primeMax is the value of the largest prime in wheel
	primeMax = 0
)

const (
	// maxPrime is the highest value up to which we will search for primes
	maxPrime = 100*1000*1000 + 1000

	// gobName is the name of the gob file the primes are stored in
	gobName = "wheel.gob"

	// piStep is the size of the blocks in the prime counts cache
	piStep = 100 // A smaller step makes Index faster, but piCache larger
)

// init loads (or saves) the wheel
func init() {
	//primes := primesPkg.MakePrimes(maxPrime)
	//for _, prime := range primes {
	//	store(uint32(prime))
	//}
	//bake()
	//save()
	load()
}

// int2offset returns the bit/byte in the wheel that the int corresponds to
func int2offset(p int) (int, uint8, bool) {
	iByte := p / 30
	r := p % 30
	iBit := remainder2bit[r]
	return iByte, iBit, iBit > 0 || r == 1
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
		fmt.Printf("error reading primes gob: %v", err)
		panic(err)
	}

	err = decoder.Decode(&piCache)
	if err != nil {
		fmt.Printf("error reading primes gob: %v", err)
		panic(err)
	}

	err = decoder.Decode(&primeCount)
	if err != nil {
		fmt.Printf("error reading primes gob: %v", err)
		panic(err)
	}

	err = decoder.Decode(&primeMax)
	if err != nil {
		fmt.Printf("error reading primes gob: %v", err)
		panic(err)
	}
}

// store writes the prime to wheel (be sure to call bake() when done storing)
func store(p uint32) {
	if p <= 5 {
		return
	}
	iByte, iBit, ok := int2offset(int(p))
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
