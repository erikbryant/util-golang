package matrices

import "testing"

func TestNew(t *testing.T) {
	testCases := []struct {
		r int
		c int
	}{
		{1, 1},
		{1, 9},
		{9, 1},
	}

	for _, testCase := range testCases {
		m := New[int](testCase.r, testCase.c)
		answerR := m.Rows()
		answerC := m.Cols()
		if answerR != testCase.r || answerC != testCase.c {
			t.Errorf("ERROR: Expected %dx%d, got %dx%d", testCase.r, testCase.c, answerR, answerC)
		}
	}
}
