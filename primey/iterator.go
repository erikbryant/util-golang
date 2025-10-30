package primey

import "fmt"

// Edge-based indexing
//
//     item   0   1   2   3   4   5
//          +---+---+---+---+---+---+
//          | P | y | t | h | o | n |
//          +---+---+---+---+---+---+
// iterator 0   1   2   3   4   5   6
//
// Detailed description: https://softwareengineering.stackexchange.com/a/290580

// context stores the position of the iterator in the list of primes
type context struct {
	iByteBit int // byte and bit offsets together; bit overflow/underflow automatically increments/decrements byte
	index    int
	end      int
}

// newContext returns a new context, starting at the given position
func newContext(start int) context {
	// context indicates the next prime to return
	ctx := context{
		iByteBit: 0,
		index:    0,
		end:      (len(wheel)-1)<<3 + 7,
	}

	if start > PrimeMax() {
		err := fmt.Errorf("index out of range %d > %d", start, PrimeMax())
		panic(err)
	}

	for ctx.index < start && !ctx.atEnd() {
		ctx.next()
	}

	return ctx
}

// atStart returns true if ctx points to the start of the primes
func (ctx *context) atStart() bool {
	return ctx.iByteBit == 0
}

// atEnd returns true if ctx points to the end of the primes
func (ctx *context) atEnd() bool {
	return ctx.iByteBit == ctx.end
}

// dec decrements ctx by one
func (ctx *context) dec() {
	ctx.iByteBit--
}

// inc increments ctx by one
func (ctx *context) inc() {
	ctx.iByteBit++
}

// prev moves ctx to the previous prime and returns that prime
func (ctx *context) prev() int {
	if ctx.index < 3 && !ctx.atStart() {
		ctx.index--
		return []int{2, 3, 5}[ctx.index]
	}

	for !ctx.atStart() {
		iByte := ctx.iByteBit >> 3
		iBit := uint8(ctx.iByteBit & 7)
		if bitIsSet(iByte, iBit) {
			p := offset2int(iByte, iBit)
			ctx.dec()
			ctx.index--
			return p
		}
		ctx.dec()
	}

	return 0
}

// next returns the next prime and increments cts, trusting the caller to stay in bounds
func (ctx *context) next() int {
	if ctx.index < 3 {
		p := []int{2, 3, 5}[ctx.index]
		ctx.index++
		return p
	}

	for {
		iByte := ctx.iByteBit >> 3
		iBit := uint8(ctx.iByteBit & 7)
		if bitIsSet(iByte, iBit) {
			p := offset2int(iByte, iBit)
			ctx.inc()
			ctx.index++
			return p
		}
		ctx.inc()
	}
}
