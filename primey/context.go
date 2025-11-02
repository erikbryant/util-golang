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
	index int
	iByte int
	iBit  uint8
}

// newContext returns a new context, set to the given index
func newContext(start int) context {
	// ctx indicates the next prime to return
	ctx := context{
		index: start,
	}

	ctx.iByte, ctx.iBit = index2offset(start)

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

// inc increments ctx by one, trusting the caller to stay in bounds
func (ctx *context) inc() {
	ctx.iBit++
	if ctx.iBit > 7 {
		ctx.iByte++
		ctx.iBit = 0
	}
}

// next returns the next prime and advances ctx, trusting the caller to stay in bounds
func (ctx *context) next() int {
	for {
		if bitIsSet(ctx.iByte, ctx.iBit) {
			p := offset2int(ctx.iByte, ctx.iBit)
			ctx.inc()
			ctx.index++
			return p
		}
		ctx.inc()
	}
}
