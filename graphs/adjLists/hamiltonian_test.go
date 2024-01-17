package adjLists

import (
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"testing"
)

func TestConnected(t *testing.T) {
	a := NewAL()

	answer := a.Connected()
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	nA := vertexes.NewVertex("A", 4)
	a.AddNode(nA)
	answer = a.Connected()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	nB := vertexes.NewVertex("B", 4)
	a.AddNode(nB)
	answer = a.Connected()
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	a.AddEdge(nA, nB)
	answer = a.Connected()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}
}

func TestHamiltonianPath(t *testing.T) {
	a := NewAL()

	// Empty graph
	answer := a.HamiltonianPaths(0, true, true)
	if answer != nil {
		t.Errorf("ERROR: Expected nil, got %v", answer)
	}

	// One node
	nA := vertexes.NewVertex("A", 4)
	a.AddNode(nA)
	answer = a.HamiltonianPaths(0, true, true)
	if answer == nil {
		t.Errorf("ERROR: Expected path, got %v", answer)
	}

	// Not connected
	nB := vertexes.NewVertex("B", 4)
	a.AddNode(nB)
	answer = a.HamiltonianPaths(0, true, true)
	if answer != nil {
		t.Errorf("ERROR: Expected nil, got %v", answer)
	}

	// More than 2 whiskers
	nC := vertexes.NewVertex("C", 4)
	nD := vertexes.NewVertex("D", 4)
	a.AddEdge(nA, nB)
	a.AddEdge(nC, nB)
	a.AddEdge(nD, nB)
	answer = a.HamiltonianPaths(0, true, true)
	if answer != nil {
		t.Errorf("ERROR: Expected nil, got %v", answer)
	}
}
