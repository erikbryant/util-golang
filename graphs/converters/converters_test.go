package converters

import (
	"github.com/erikbryant/util-golang/graphs/adjLists"
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"testing"
)

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
