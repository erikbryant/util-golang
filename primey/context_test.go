package primey

import "testing"

func TestNewContext(t *testing.T) {
	testCases := []struct {
		c         int
		expected1 int
		expected2 int
	}{
		{0, 0, 0},
		{1, 1, 0},
		{2, 2, 0},
		{3, 3, 1},
		{4, 4, 2},
	}

	for _, testCase := range testCases {
		ctx := newContext(testCase.c)
		answer1 := ctx.index
		answer2 := ctx.iByteBit
		if answer1 != testCase.expected1 || answer2 != testCase.expected2 {
			t.Errorf("ERROR: For %d expected %d, %d, got %d, %d", testCase.c, testCase.expected1, testCase.expected2, answer1, answer2)
		}
	}
}

func TestAtStart(t *testing.T) {
	ctx := newContext(0)
	if !ctx.atStart() {
		t.Errorf("ERROR: Expected to be at start %v", ctx)
	}

	ctx = newContext(1)
	if ctx.atStart() {
		t.Errorf("ERROR: Expected to NOT be at start %v", ctx)
	}

	ctx = newContext(2)
	if ctx.atStart() {
		t.Errorf("ERROR: Expected to NOT be at start %v", ctx)
	}

	ctx = newContext(3)
	if ctx.atStart() {
		t.Errorf("ERROR: Expected to NOT be at start %v", ctx)
	}

	ctx = newContext(1000)
	if ctx.atStart() {
		t.Errorf("ERROR: Expected to NOT be at start %v", ctx)
	}
}

func TestAtEnd(t *testing.T) {
	ctx := newContext(0)
	if ctx.atEnd() {
		t.Errorf("ERROR: Expected to NOT be at end %v", ctx)
	}

	ctx = newContext(1)
	if ctx.atEnd() {
		t.Errorf("ERROR: Expected to NOT be at end %v", ctx)
	}

	ctx = newContext(2)
	if ctx.atEnd() {
		t.Errorf("ERROR: Expected to NOT be at end %v", ctx)
	}

	ctx = newContext(3)
	if ctx.atEnd() {
		t.Errorf("ERROR: Expected to NOT be at end %v", ctx)
	}

	ctx = newContext(Len())
	if !ctx.atEnd() {
		t.Errorf("ERROR: Expected to be at end %v", ctx)
	}
}

func TestDec(t *testing.T) {
}

func TestInc(t *testing.T) {
}

func TestPrev(t *testing.T) {
}

func TestNext(t *testing.T) {
	ctx := newContext(0)
	expected := 0
	if ctx.index != expected {
		t.Errorf("ERROR: Expected %d got %d", expected, ctx.index)
	}

	ctx.next()
	expected = 1
	if ctx.index != expected {
		t.Errorf("ERROR: Expected %d got %d", expected, ctx.index)
	}

	ctx.next()
	expected = 2
	if ctx.index != expected {
		t.Errorf("ERROR: Expected %d got %d", expected, ctx.index)
	}

	ctx.next()
	expected = 3
	if ctx.index != expected {
		t.Errorf("ERROR: Expected %d got %d", expected, ctx.index)
	}

	ctx.next()
	expected = 4
	if ctx.index != expected {
		t.Errorf("ERROR: Expected %d got %d", expected, ctx.index)
	}

	ctx.next()
	expected = 5
	if ctx.index != expected {
		t.Errorf("ERROR: Expected %d got %d", expected, ctx.index)
	}
}
