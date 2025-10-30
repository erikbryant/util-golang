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
	iByte int
	iBit  int8
	index int
}

// newContext returns a new context, starting at the given position
func newContext(start int) context {
	// context indicates the next prime to return
	ctx := context{
		iByte: 0,
		iBit:  0,
		index: 0,
	}

	for ctx.index < start && !ctx.atEnd() {
		ctx.next()
	}

	return ctx
}

// atStart returns true if ctx points to the start of the primes
func (ctx *context) atStart() bool {
	return ctx.iByte == 0 && ctx.iBit == 0
}

// atEnd returns true if ctx points to the end of the primes
func (ctx *context) atEnd() bool {
	return ctx.iByte == len(wheel)-1 && ctx.iBit == 7
}

// dec decrements ctx by one
func (ctx *context) dec() {
	if ctx.atStart() {
		return
	}
	if ctx.iBit == 0 {
		ctx.iByte--
		ctx.iBit = 8
	}
	ctx.iBit--
}

// inc increments ctx by one
func (ctx *context) inc() {
	ctx.iBit++
	if ctx.iBit >= 8 {
		ctx.iByte++
		ctx.iBit = 0
	}
}

// prev moves ctx to the previous prime and returns that prime
func (ctx *context) prev() int {
	if ctx.index < 3 && !ctx.atStart() {
		ctx.index--
		return []int{2, 3, 5}[ctx.index]
	}

	for !ctx.atStart() {
		if bitIsSet(ctx.iByte, uint8(ctx.iBit)) {
			p := offset2int(ctx.iByte, uint8(ctx.iBit))
			ctx.dec()
			ctx.index--
			return p
		}
		ctx.dec()
	}

	return 0
}

// next moves ctx to the next prime and returns that prime
func (ctx *context) next() int {
	if ctx.index < 3 {
		p := []int{2, 3, 5}[ctx.index]
		ctx.index++
		return p
	}

	for !ctx.atEnd() {
		if bitIsSet(ctx.iByte, uint8(ctx.iBit)) {
			p := offset2int(ctx.iByte, uint8(ctx.iBit))
			ctx.inc()
			ctx.index++
			return p
		}
		ctx.inc()
	}

	return 0
}
