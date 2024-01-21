package adjMatrixes

import (
	"github.com/erikbryant/util-golang/graphs/adjLists"
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"testing"
)

func TestMatrix(t *testing.T) {
	m := Matrix(4)

	answer := m.Size()
	if answer != 4 {
		t.Errorf("ERROR: Expected 4, got %d", answer)
	}
}

//func TestReachable(t *testing.T) {
//}

func TestMatrixFromAdjList(t *testing.T) {
	graph := adjLists.NewAL()
	nA := vertexes.NewVertex("A", 0)
	nB := vertexes.NewVertex("B", 0)
	nC := vertexes.NewVertex("C", 0)
	graph.AddEdge(nA, nB)
	graph.AddEdge(nB, nC)

	m := MatrixFromAdjList(&graph)

	testCases := []struct {
		r, c     int
		expected int
	}{
		// Major diagonal
		{0, 0, 0},
		{1, 1, 0},
		{2, 2, 0},

		// Weights
		{0, 1, 1},
		{0, 2, 2},
		{1, 0, 1},
		{1, 2, 1},
		{2, 0, 2},
		{2, 1, 1},
	}

	for _, testCase := range testCases {
		answer := m.GetValue(testCase.r, testCase.c)
		if answer != testCase.expected {
			t.Errorf("ERROR: For %d,%d expected %d, got %d", testCase.r, testCase.c, testCase.expected, answer)
		}
	}
}

func TestSize(t *testing.T) {
	m := Matrix(4)

	answer := m.Size()
	if answer != 4 {
		t.Errorf("ERROR: Expected 4, got %d", answer)
	}
}

func TestSetValueGetValue(t *testing.T) {
	m := Matrix(3)

	testCases := []struct {
		r, c int
		v    int
	}{
		{0, 0, 100},
		{1, 0, 10},
		{0, 1, 1},
		{1, 1, -1},
	}

	for _, testCase := range testCases {
		m.SetValue(testCase.r, testCase.c, testCase.v)
		answer := m.GetValue(testCase.r, testCase.c)
		if answer != testCase.v {
			t.Errorf("ERROR: For %d, %d expected %d, got %d", testCase.r, testCase.c, testCase.v, answer)
		}
	}
}

//func TestNodeFromIndex(t *testing.T) {
//}

//func TestIndexFromID(t *testing.T) {
//}

func TestTrunc(t *testing.T) {
	testCases := []struct {
		s        string
		l        int
		expected string
	}{
		{"", 4, ""},
		{"", 0, ""},
		{"blah", 0, ""},
		{"   this   is   a  test", 4, "this"},
		{"the", 5, "the"},
	}

	for _, testCase := range testCases {
		answer := trunc(testCase.s, testCase.l)
		if answer != testCase.expected {
			t.Errorf("ERROR: For '%s' expected '%s', got '%s'", testCase.s, testCase.expected, answer)
		}
	}
}
