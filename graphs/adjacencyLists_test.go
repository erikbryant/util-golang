package graphs

import (
	"fmt"
	"testing"
)

func TestAddEdge(t *testing.T) {
	a := NewAL()

	// Empty adjacency list

	answer := a.NodeCount()
	if answer != 0 {
		t.Errorf("ERROR: Expected len=0, got len=%d", answer)
	}

	answer = a.EdgeCount()
	if answer != 0 {
		t.Errorf("ERROR: Expected len=0, got len=%d", answer)
	}

	// Implied addition of nodes

	n1 := NewVertex("A", 4)
	n2 := NewVertex("B", 4)
	a.AddEdge(&n1, &n2)

	answer = a.NodeCount()
	if answer != 2 {
		t.Errorf("ERROR: Expected len=2, got len=%d", answer)
	}

	answer = a.EdgeCount()
	if answer != 1 {
		t.Errorf("ERROR: Expected len=1, got len=%d", answer)
	}

	// Explicit addition of nodes

	n3 := NewVertex("C", 4)

	a.AddNode(&n3)
	a.AddEdge(&n1, &n3)

	answer = a.NodeCount()
	if answer != 3 {
		t.Errorf("ERROR: Expected len=3, got len=%d", answer)
	}

	answer = a.EdgeCount()
	if answer != 2 {
		t.Errorf("ERROR: Expected len=2, got len=%d", answer)
	}

	// New edge, no new nodes

	a.AddEdge(&n2, &n3)

	answer = a.NodeCount()
	if answer != 3 {
		t.Errorf("ERROR: Expected len=3, got len=%d", answer)
	}

	answer = a.EdgeCount()
	if answer != 3 {
		t.Errorf("ERROR: Expected len=3, got len=%d", answer)
	}
}

func TestWhiskers(t *testing.T) {
	a := NewAL()
	whiskers := a.Whiskers()
	if len(whiskers) != 0 {
		t.Errorf("ERROR: For 'empty' expected len=0, got len=%d", len(whiskers))
	}

	a = NewAL()
	n1 := NewVertex("A", 4)
	a.AddNode(&n1)
	whiskers = a.Whiskers()
	if len(whiskers) != 0 {
		t.Errorf("ERROR: For 'one node' expected len=0, got len=%d", len(whiskers))
	}

	a = NewAL()
	n1 = NewVertex("B", 4)
	n2 := NewVertex("C", 4)
	a.AddNode(&n1)
	a.AddNode(&n2)
	whiskers = a.Whiskers()
	if len(whiskers) != 0 {
		t.Errorf("ERROR: For 'two nodes' expected len=0, got len=%d", len(whiskers))
	}

	a = NewAL()
	n1 = NewVertex("B", 4)
	n2 = NewVertex("C", 4)
	a.AddEdge(&n1, &n2)
	whiskers = a.Whiskers()
	if len(whiskers) != 1 {
		t.Errorf("ERROR: For 'one edge' expected len=1, got len=%d", len(whiskers))
	}

	a = NewAL()
	n1 = NewVertex("D", 4)
	n2 = NewVertex("E", 4)
	n3 := NewVertex("F", 4)
	a.AddEdge(&n1, &n2)
	a.AddEdge(&n2, &n3)
	whiskers = a.Whiskers()
	if len(whiskers) != 2 {
		t.Errorf("ERROR: For 'one vertex, two whiskers' expected len=2, got len=%d", len(whiskers))
	}
}

func TestMinimalVertexCover(t *testing.T) {
	a := NewAL()
	mvc := a.MinimalVertexCover()
	if len(mvc) != 0 {
		t.Errorf("ERROR: For 'empty' expected len=0, got len=%d", len(mvc))
	}

	a = NewAL()
	n1 := NewVertex("A", 4)
	a.AddNode(&n1)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 0 {
		t.Errorf("ERROR: For 'one node' expected len=0, got len=%d", len(mvc))
	}

	a = NewAL()
	n1 = NewVertex("B", 4)
	n2 := NewVertex("C", 4)
	a.AddNode(&n1)
	a.AddNode(&n2)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 0 {
		t.Errorf("ERROR: For 'two nodes' expected len=0, got len=%d", len(mvc))
	}

	a = NewAL()
	n1 = NewVertex("B", 4)
	n2 = NewVertex("C", 4)
	a.AddEdge(&n1, &n2)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 1 {
		t.Errorf("ERROR: For 'one edge' expected len=1, got len=%d", len(mvc))
	}

	a = NewAL()
	n1 = NewVertex("D", 4)
	n2 = NewVertex("E", 4)
	n3 := NewVertex("F", 4)
	a.AddEdge(&n1, &n2)
	a.AddEdge(&n2, &n3)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 1 {
		t.Errorf("ERROR: For 'one vertex, two whiskers' expected len=1, got len=%d", len(mvc))
		for _, node := range mvc {
			fmt.Println(node.Name())
		}
	}
}
