package primes

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
	"fmt"
	"math"
	"math/bits"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var (
	wheel         = make([]uint8, maxPrime/30)
	bit2remainder = []int{1, 7, 11, 13, 17, 19, 23, 29}
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
	primeCountWheel = 0
	primeMaxWheel   = 0

	// Prime counts (pi) cache
	piCache = []uint32{}
	piStep  = 100 // A smaller step makes IndexWheel faster, but piCache larger
)

type Context struct {
	iByte  int
	iBit   uint8
	valid  bool
	primes []int
	pIndex int
}

func int2offset(p int) (int, uint8, bool) {
	iByte := p / 30
	r := p % 30
	iBit := remainder2bit[r]
	return iByte, iBit, iBit > 0 || r == 1
}

func offset2int(iByte int, iBit uint8) int {
	return iByte*30 + bit2remainder[iBit]
}

func setBit(iByte int, iBit uint8) {
	wheel[iByte] |= 1 << iBit
}

func bitIsSet(iByte int, iBit uint8) bool {
	return wheel[iByte]&(1<<iBit) != 0
}

func StoreWheel(p int) {
	if p <= 5 {
		return
	}
	iByte, iBit, ok := int2offset(p)
	if !ok {
		fmt.Printf("%d is not prime! Not storing.\n", p)
		return
	}

	if bitIsSet(iByte, iBit) {
		fmt.Printf("%d is already stored! Not storing.\n", p)
		return
	}

	primeMaxWheel = max(primeMaxWheel, p)
	primeCountWheel++
	setBit(iByte, iBit)
}

func Bake() {
	piCache = make([]uint32, primeCountWheel/piStep)

	primeCount := uint32(0)
	for i := 0; i < len(wheel); i++ {
		primeCount += uint32(bits.OnesCount8(wheel[i]))
		k := i/piStep + 1
		piCache[k] = primeCount
	}

	p := message.NewPrinter(language.English)
	p.Printf("Wheel statistics\n")
	p.Printf("primeCountWheel = %16d\n", primeCountWheel)
	p.Printf("primeMaxWheel   = %16d\n", primeMaxWheel)
	p.Printf("Wheel size      = %16d bytes\n", len(wheel))
	p.Printf("sizeof(piCache) = %16d bytes\n", len(piCache)*4)
	totalSizeWheel := len(wheel) + len(piCache)*4
	p.Printf("total size      = %16d bytes\n", totalSizeWheel)
	p.Printf("primes/byte     = %16f\n", float64(primeCountWheel)/float64(totalSizeWheel))

	p.Printf("\nPrime statistics\n")
	p.Printf("primeCount      = %16d\n", len(Primes))
	p.Printf("primeMax        = %16d\n", Primes[End])
	p.Printf("Prime size      = %16d bytes\n", len(Primes)*4)
	totalSizePrimes := len(Primes) * 4
	p.Printf("total size      = %16d bytes\n", totalSizePrimes)
	p.Printf("primes/byte     = %16f\n", float64(len(Primes))/float64(totalSizePrimes))
}

func inc(ctx *Context) {
	ctx.iBit++
	if ctx.iBit >= 8 {
		ctx.iBit = 0
		ctx.iByte++
		if ctx.iByte >= len(wheel) {
			ctx.valid = false
		}
	}
}

func context() Context {
	// Context indicates the next prime to return
	return Context{0, 0, true, []int{2, 3, 5}, 0}
}

func next(ctx *Context) int {
	if ctx.pIndex < len(ctx.primes) {
		p := ctx.primes[ctx.pIndex]
		ctx.pIndex++
		return p
	}

	for ctx.valid {
		if bitIsSet(ctx.iByte, ctx.iBit) {
			p := offset2int(ctx.iByte, ctx.iBit)
			inc(ctx)
			return p
		}
		inc(ctx)
	}
	return 0
}

func IndexWheel(p int) int {
	if p <= 5 {
		return []int{-1, -1, 0, 1, -1, 2}[p]
	}

	adjusted := false

	// Each byte in the wheel represents 0-7 primes
	// Count bits on the way up to p

	iByte, iBit, ok := int2offset(p)
	if !ok || !bitIsSet(iByte, iBit) {
		// p is not a prime; find the next higher prime
		adjusted = true
		ctx := Context{
			iByte: iByte,
			iBit:  0,
			valid: true,
		}
		for ctx.valid {
			if bitIsSet(ctx.iByte, ctx.iBit) {
				nextP := offset2int(ctx.iByte, ctx.iBit)
				if nextP > p {
					p = nextP
					break
				}
			}
			inc(&ctx)
		}
		iByte, iBit, _ = int2offset(p)
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

func PiWheel(n int) int {
	if n < 2 {
		return 0
	}
	i := IndexWheel(n)
	if i < 0 {
		i = -i
	}
	return i + 1
}

// IterWheel returns an iterator over all Primes
func IterWheel() func(func(int, int) bool) {
	return IterrWheel(0, -1)
}

// IterrWheel returns an iterator over a range of Primes
func IterrWheel(start, end int) func(func(int, int) bool) {
	return func(yield func(int, int) bool) {
		if end < 0 {
			end = len(Primes)
		}

		// Initialize the sequence
		ctx := context()
		var prime int

		// Fastforward to the first prime to yield
		for k := 0; k <= start; k++ {
			prime = next(&ctx)
			if !valid(ctx) {
				return
			}
		}

		// Yield the primes
		for i := start; i < end; i++ {
			if !valid(ctx) || !yield(i-start, prime) {
				return
			}
			prime = next(&ctx)
		}
	}
}

func valid(ctx Context) bool {
	return ctx.valid
}

// SlowPrimeWheel returns whether a number is prime or not, using a brute force search
func SlowPrimeWheel(n int) bool {
	if n <= 1 {
		return false
	}

	root := int(math.Sqrt(float64(n)))

	// Check each potential divisor to see if number divides evenly (i.e., is not prime).
	for _, prime := range IterWheel() {
		if prime > root {
			return true
		}
		if n%prime == 0 {
			return false
		}
	}

	return true
}

func PrimeWheel(p int) bool {
	if p <= 5 {
		return p == 2 || p == 3 || p == 5
	}
	iByte, iBit, ok := int2offset(p)
	return ok && bitIsSet(iByte, iBit)
}
