package adjLists

import (
	"fmt"
	"github.com/erikbryant/util-golang/graphs/vertexes"
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

	n1 := vertexes.NewVertex("A", 4)
	n2 := vertexes.NewVertex("B", 4)
	a.AddEdge(n1, n2)

	answer = a.NodeCount()
	if answer != 2 {
		t.Errorf("ERROR: Expected len=2, got len=%d", answer)
	}

	answer = a.EdgeCount()
	if answer != 1 {
		t.Errorf("ERROR: Expected len=1, got len=%d", answer)
	}

	// Explicit addition of nodes

	n3 := vertexes.NewVertex("C", 4)

	a.AddNode(n3)
	a.AddEdge(n1, n3)

	answer = a.NodeCount()
	if answer != 3 {
		t.Errorf("ERROR: Expected len=3, got len=%d", answer)
	}

	answer = a.EdgeCount()
	if answer != 2 {
		t.Errorf("ERROR: Expected len=2, got len=%d", answer)
	}

	// New edge, no new nodes

	a.AddEdge(n2, n3)

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
	n1 := vertexes.NewVertex("A", 4)
	a.AddNode(n1)
	whiskers = a.Whiskers()
	if len(whiskers) != 0 {
		t.Errorf("ERROR: For 'one node' expected len=0, got len=%d", len(whiskers))
	}

	a = NewAL()
	n1 = vertexes.NewVertex("B", 4)
	n2 := vertexes.NewVertex("C", 4)
	a.AddNode(n1)
	a.AddNode(n2)
	whiskers = a.Whiskers()
	if len(whiskers) != 0 {
		t.Errorf("ERROR: For 'two nodes' expected len=0, got len=%d", len(whiskers))
	}

	a = NewAL()
	n1 = vertexes.NewVertex("B", 4)
	n2 = vertexes.NewVertex("C", 4)
	a.AddEdge(n1, n2)
	whiskers = a.Whiskers()
	if len(whiskers) != 1 {
		t.Errorf("ERROR: For 'one edge' expected len=1, got len=%d", len(whiskers))
	}

	a = NewAL()
	n1 = vertexes.NewVertex("D", 4)
	n2 = vertexes.NewVertex("E", 4)
	n3 := vertexes.NewVertex("F", 4)
	a.AddEdge(n1, n2)
	a.AddEdge(n2, n3)
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
	n1 := vertexes.NewVertex("A", 4)
	a.AddNode(n1)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 0 {
		t.Errorf("ERROR: For 'one node' expected len=0, got len=%d", len(mvc))
	}

	a = NewAL()
	n1 = vertexes.NewVertex("B", 4)
	n2 := vertexes.NewVertex("C", 4)
	a.AddNode(n1)
	a.AddNode(n2)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 0 {
		t.Errorf("ERROR: For 'two nodes' expected len=0, got len=%d", len(mvc))
	}

	a = NewAL()
	n1 = vertexes.NewVertex("B", 4)
	n2 = vertexes.NewVertex("C", 4)
	a.AddEdge(n1, n2)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 1 {
		t.Errorf("ERROR: For 'one edge' expected len=1, got len=%d", len(mvc))
	}

	a = NewAL()
	n1 = vertexes.NewVertex("D", 4)
	n2 = vertexes.NewVertex("E", 4)
	n3 := vertexes.NewVertex("F", 4)
	a.AddEdge(n1, n2)
	a.AddEdge(n2, n3)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 1 {
		t.Errorf("ERROR: For 'one vertex, two whiskers' expected len=1, got len=%d", len(mvc))
		for _, node := range mvc {
			fmt.Println(node.Name())
		}
	}
}

func TestMinimalVertexCoverOctahedral(t *testing.T) {
	// Test against the octahedral graph
	// https://reference.wolfram.com/language/ref/FindVertexCover.html
	// https://www.researchgate.net/figure/Octahedral-graph_fig3_365219344

	a := NewAL()
	nA := vertexes.NewVertex("A", 4)
	nB := vertexes.NewVertex("B", 4)
	nC := vertexes.NewVertex("C", 4)
	nD := vertexes.NewVertex("D", 4)
	nE := vertexes.NewVertex("E", 4)
	nF := vertexes.NewVertex("F", 4)

	a.AddEdge(nA, nB)
	a.AddEdge(nB, nC)
	a.AddEdge(nC, nA)

	a.AddEdge(nD, nE)
	a.AddEdge(nE, nF)
	a.AddEdge(nF, nD)

	a.AddEdge(nD, nA)
	a.AddEdge(nD, nB)

	a.AddEdge(nE, nA)
	a.AddEdge(nE, nC)

	a.AddEdge(nF, nB)
	a.AddEdge(nF, nC)

	mvc := a.MinimalVertexCover()
	if len(mvc) != 4 {
		t.Errorf("ERROR: For 'two nodes' expected len=4, got len=%d", len(mvc))
	}
}

func TestIsSymmetric(t *testing.T) {
	a := NewAL()

	answer := a.IsSymmetric()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	nA := vertexes.NewVertex("A", 4)
	nB := vertexes.NewVertex("B", 4)

	a.AddEdge(nA, nB)

	answer = a.IsSymmetric()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}
}

func TestIsBipartite(t *testing.T) {
	a := NewAL()

	// Empty
	answer := a.IsBipartite()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	// Two connected vertexes
	nA := vertexes.NewVertex("A", 4)
	nB := vertexes.NewVertex("B", 4)
	a.AddEdge(nA, nB)
	answer = a.IsBipartite()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	// Three vertexes connected in a V
	nC := vertexes.NewVertex("C", 4)
	a.AddEdge(nA, nC)
	answer = a.IsBipartite()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	// Three vertexes connected in a triangle
	a.AddEdge(nB, nC)
	answer = a.IsBipartite()
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}
}

func TestIsComplete(t *testing.T) {
	a := NewAL()

	// Empty
	answer := a.IsComplete()
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	// Two connected vertexes
	nA := vertexes.NewVertex("A", 4)
	nB := vertexes.NewVertex("B", 4)
	a.AddEdge(nA, nB)
	answer = a.IsComplete()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	// Three vertexes connected in a V
	nC := vertexes.NewVertex("C", 4)
	a.AddEdge(nA, nC)
	answer = a.IsComplete()
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	// Three vertexes connected in a triangle
	a.AddEdge(nB, nC)
	answer = a.IsComplete()
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}
}

func TestChromaticNumber(t *testing.T) {
	a := NewAL()

	// Empty
	answer := a.ChromaticNumber()
	if answer != 1 {
		t.Errorf("ERROR: Expected 1, got %d", answer)
	}

	// Two non-connected vertexes
	nA := vertexes.NewVertex("A", 4)
	nB := vertexes.NewVertex("B", 4)
	a.AddNode(nA)
	a.AddNode(nB)
	answer = a.ChromaticNumber()
	if answer != 1 {
		t.Errorf("ERROR: Expected 1, got %d", answer)
	}

	// Two connected vertexes
	a.AddEdge(nA, nB)
	answer = a.ChromaticNumber()
	if answer != 2 {
		t.Errorf("ERROR: Expected 2, got %d", answer)
	}

	// Three vertexes connected in a V
	nC := vertexes.NewVertex("C", 4)
	a.AddEdge(nA, nC)
	answer = a.ChromaticNumber()
	if answer != 2 {
		t.Errorf("ERROR: Expected 2, got %d", answer)
	}

	// Three vertexes connected in a triangle (graph is complete)
	a.AddEdge(nB, nC)
	answer = a.ChromaticNumber()
	if answer != 3 {
		t.Errorf("ERROR: Expected 3, got %d", answer)
	}

	// Three vertexes connected in a triangle and 1 whisker
	nD := vertexes.NewVertex("D", 4)
	a.AddEdge(nC, nD)
	answer = a.ChromaticNumber()
	if answer != 3 {
		t.Errorf("ERROR: Expected 3, got %d", answer)
	}

	// Four vertexes connected in a diamond
	a.AddEdge(nD, nA)
	answer = a.ChromaticNumber()
	if answer != 3 {
		t.Errorf("ERROR: Expected 3, got %d", answer)
	}

	// Four vertexes complete-connected
	a.AddEdge(nD, nB)
	answer = a.ChromaticNumber()
	if answer != 4 {
		t.Errorf("ERROR: Expected 4, got %d", answer)
	}

	// Four vertexes complete-connected with 1 whisker
	nE := vertexes.NewVertex("C", 4)
	a.AddEdge(nD, nE)
	answer = a.ChromaticNumber()
	if answer != 4 {
		t.Errorf("ERROR: Expected 4, got %d", answer)
	}
}
