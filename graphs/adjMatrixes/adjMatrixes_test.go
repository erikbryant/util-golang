package adjMatrixes

import (
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"testing"
)

//func TestAddNode(t *testing.T) {
//}

func TestNodeCount(t *testing.T) {
	m := Matrix()

	answer := m.NodeCount()
	if answer != 0 {
		t.Errorf("ERROR: Expected 0, got %d", answer)
	}

	vA := vertexes.NewVertex("A", 1)
	m.AddNode(vA.ID(), vA)
	answer = m.NodeCount()
	if answer != 1 {
		t.Errorf("ERROR: Expected 1, got %d", answer)
	}

	vB := vertexes.NewVertex("B", 1)
	m.AddNode(vB.ID(), vB)
	answer = m.NodeCount()
	if answer != 2 {
		t.Errorf("ERROR: Expected 2, got %d", answer)
	}
}

//func TestReachable(t *testing.T) {
//}

//func TestInitMatrix(t *testing.T) {
//}

//func TestComputeDistance(t *testing.T) {
//}

func TestSetValueGetValue(t *testing.T) {
	m := Matrix()

	testCases := []struct {
		r, c int
		v    int
	}{
		{0, 0, 100},
		{1, 0, 10},
		{0, 1, 1},
		{1, 1, -1},
	}

	vA := vertexes.NewVertex("A", 1)
	m.AddNode(vA.ID(), vA)
	vB := vertexes.NewVertex("B", 1)
	m.AddNode(vB.ID(), vB)
	vC := vertexes.NewVertex("C", 1)
	m.AddNode(vC.ID(), vC)

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

func TestDiameter(t *testing.T) {
	m := Matrix()

	// No vertex
	m.ComputeDistances()
	answer := m.Diameter()
	if answer != -1 {
		t.Errorf("ERROR: Expected -1, got %d", answer)
	}

	// One vertex
	vA := vertexes.NewVertex("A", 1)
	m.AddNode(vA.ID(), vA)
	m.ComputeDistances()
	answer = m.Diameter()
	if answer != 0 {
		t.Errorf("ERROR: Expected 0, got %d", answer)
	}

	// Two vertexes, not connected
	vB := vertexes.NewVertex("B", 1)
	m.AddNode(vB.ID(), vB)
	m.ComputeDistances()
	answer = m.Diameter()
	if answer != 0 {
		t.Errorf("ERROR: Expected 0, got %d", answer)
	}

	// Two vertexes connected
	vA.AddNeighbor(vB)
	m.ComputeDistances()
	answer = m.Diameter()
	if answer != 1 {
		t.Errorf("ERROR: Expected 1, got %d", answer)
	}

	// Three vertexes connected
	vC := vertexes.NewVertex("C", 1)
	m.AddNode(vC.ID(), vC)
	vB.AddNeighbor(vC)
	m.ComputeDistances()
	answer = m.Diameter()
	if answer != 2 {
		t.Errorf("ERROR: Expected 2, got %d", answer)
	}

	// Four vertexes, 3 connected, 1 not
	vD := vertexes.NewVertex("D", 1)
	m.AddNode(vD.ID(), vD)
	m.ComputeDistances()
	answer = m.Diameter()
	if answer != 2 {
		t.Errorf("ERROR: Expected 2, got %d", answer)
	}
}

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
