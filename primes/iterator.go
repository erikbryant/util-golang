package primes

// Edge-based indexing
//
//     item   0   1   2   3   4   5
//          +---+---+---+---+---+---+
//          | P | y | t | h | o | n |
//          +---+---+---+---+---+---+
// iterator 0   1   2   3   4   5   6
//
// Detailed description: https://softwareengineering.stackexchange.com/a/290580

var (
	littlePrimes = []int{2, 3, 5}
)

type Context struct {
	iByte int
	iBit  int8
	index int
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
			if ctx.atEnd() {
				return
			}
			prime = ctx.next()
		}

		// Yield the primes
		for i := start; i < end; i++ {
			if ctx.atEnd() || !yield(i-start, prime) {
				return
			}
			prime = ctx.next()
		}
	}
}

func (ctx *Context) atStart() bool {
	return ctx.iByte == 0 && ctx.iBit == 0
}

func (ctx *Context) atEnd() bool {
	return ctx.iByte == len(wheel)-1 && ctx.iBit == 7
}

func (ctx *Context) dec() {
	if ctx.atStart() {
		return
	}
	if ctx.iBit == 0 {
		ctx.iByte--
		ctx.iBit = 8
	}
	ctx.iBit--
}

func (ctx *Context) inc() {
	if ctx.atEnd() {
		return
	}
	ctx.iBit++
	if ctx.iBit >= 8 {
		ctx.iBit = 0
		ctx.iByte++
	}
}

func context() Context {
	// Context indicates the next prime to return
	return Context{
		iByte: 0,
		iBit:  0,
		index: 0,
	}
}

func (ctx *Context) prev() int {
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

func (ctx *Context) next() int {
	if ctx.index < 3 && !ctx.atEnd() {
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

func nextHigherPrime(p, iByte int) (int, int, uint8) {
	// We can't call Index, because Index called us! Find it by hand.

	ctx := Context{
		iByte: iByte,
		iBit:  0,
	}

	for !ctx.atEnd() {
		if bitIsSet(ctx.iByte, uint8(ctx.iBit)) {
			nextP := offset2int(ctx.iByte, uint8(ctx.iBit))
			if nextP > p {
				return p, ctx.iByte, uint8(ctx.iBit)
			}
		}
		ctx.inc()
	}

	// We ran off the end of the list
	return -1, -1, 0
}
