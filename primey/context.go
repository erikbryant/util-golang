package primey

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
	iByteBit int // byte and bit offsets concatenated; bit's overflow/underflow automatically increments/decrements byte
	index    int
	end      int
}

// newContext returns a new context, set to the given index
func newContext(start int) context {
	// context indicates the next prime to return
	ctx := context{
		iByteBit: 0,
		index:    start,
	}

	iByte, iBit := index2offset(start)
	ctx.iByteBit = iByte<<3 + int(iBit)

	return ctx
}

// atStart returns true if ctx points to the start of the primes
func (ctx *context) atStart() bool {
	return ctx.index == 0
}

// atEnd returns true if ctx points to the end of the primes
func (ctx *context) atEnd() bool {
	return ctx.index == Len()
}

// dec decrements ctx by one, trusting the caller to stay in bounds
func (ctx *context) dec() {
	ctx.iByteBit--
}

// inc increments ctx by one, trusting the caller to stay in bounds
func (ctx *context) inc() {
	ctx.iByteBit++
}

// prev moves ctx to the previous prime and returns that prime, trusting the caller to stay in bounds
func (ctx *context) prev() int {
	// FIXME: This function has not been tested!
	if ctx.index < len(primeCache)-1 && !ctx.atStart() {
		ctx.index--
		return int(primeCache[ctx.index])
	}

	for !ctx.atStart() {
		iByte := ctx.iByteBit >> 3
		iBit := uint8(ctx.iByteBit & 0x07)
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

// next returns the next prime and advances ctx, trusting the caller to stay in bounds
func (ctx *context) next() int {
	// TODO: the commented-out code is for a context that spans both the
	// TODO: primeCache and wheel. We are not using that. Will we ever?

	//if ctx.index < len(primeCache) {
	//	p := primeCache[ctx.index]
	//	ctx.index++
	//	return p
	//}

	//if ctx.index == len(primeCache) {
	//	// Leaving primeCache and entering wheel; initialize the wheel index
	//	ctx.iByteBit = wheelStartByteBit
	//}

	for {
		iByte := ctx.iByteBit >> 3
		iBit := uint8(ctx.iByteBit & 0x07)
		if bitIsSet(iByte, iBit) {
			p := offset2int(iByte, iBit)
			ctx.inc()
			ctx.index++
			return p
		}
		ctx.inc()
	}
}
