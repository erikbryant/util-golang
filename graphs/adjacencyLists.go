package graphs

import (
	"fmt"

	"github.com/emicklei/dot"
)

// AdjacencyList implements an undirected graph
type AdjacencyList struct {
	nodes map[string]*Vertex
}

// NewAL returns a new, empty adjacdency list
func NewAL() AdjacencyList {
	return AdjacencyList{
		nodes: map[string]*Vertex{},
	}
}

// HasNode returns true if the node is already in the adjacency list
func (a *AdjacencyList) HasNode(node Vertex) bool {
	return a.nodes[node.ID()] != nil
}

// FirstNode returns a node from the map
func (a *AdjacencyList) FirstNode() *Vertex {
	for _, node := range a.Nodes() {
		return node
	}
	return nil
}

// AddNode adds a node to the adjacency list if not already present
func (a *AdjacencyList) AddNode(node *Vertex) {
	a.nodes[node.ID()] = node
}

// RemoveNode removes a node from the adjacency list
func (a *AdjacencyList) RemoveNode(node *Vertex) {
	for _, n := range a.Nodes() {
		if n.ID() == node.ID() {
			// Remove it from its neighbors
			for _, neighbor := range node.Neighbors() {
				neighbor.RemoveNeighbor(*node)
			}
			delete(a.nodes, n.ID())
			return
		}
	}
}

// Copy returns a copy of the AdjacencyList
func (a *AdjacencyList) Copy() AdjacencyList {
	newAL := NewAL()

	for _, node := range a.nodes {
		newAL.AddNode(node)
	}

	return newAL
}

// Nodes returns the map of nodes in the adjacency list
func (a *AdjacencyList) Nodes() map[string]*Vertex {
	return a.nodes
}

// NodeCount returns the number of nodes in the adjacency list
func (a *AdjacencyList) NodeCount() int {
	return len(a.Nodes())
}

// AddEdge adds an edge, adding the nodes if they are not already present
func (a *AdjacencyList) AddEdge(n1, n2 *Vertex) {
	a.AddNode(n1)
	a.AddNode(n2)
	n1.AddNeighbor(n2)
	n2.AddNeighbor(n1)
}

// EdgeCount returns the number of distinct edges in the adjacency list
func (a *AdjacencyList) EdgeCount() int {
	edges := 0

	for _, node := range a.Nodes() {
		edges += node.EdgeCount()
	}

	// This is undirected, so each edge is listed twice
	return edges / 2
}

// Genus returns the genus number of the adjacency list
func (a *AdjacencyList) Genus() int {
	return a.EdgeCount() - a.NodeCount() + 1
}

// ValueSum returns the sum of all node values
func (a *AdjacencyList) ValueSum() int {
	sum := 0
	for _, node := range a.Nodes() {
		sum += node.Value()
	}
	return sum
}

// ValueLowest returns the node with the lowest value
func (a *AdjacencyList) ValueLowest() *Vertex {
	min := a.FirstNode()

	for _, node := range a.Nodes() {
		if node.Value() < min.Value() {
			min = node
		}
	}

	return min
}

// Whiskers returns a map of vertices that have only one edge
func (a *AdjacencyList) Whiskers() map[string]*Vertex {
	whiskers := map[string]*Vertex{}

	for _, node := range a.Nodes() {
		if node.EdgeCount() == 1 {
			// If we have already recorded the node at the other end of
			// this edge, do not also add this node. We count only one
			// node per edge.
			if whiskers[node.FirstNeighbor().ID()] != nil {
				continue
			}
			whiskers[node.ID()] = node
		}
	}

	return whiskers
}

// NodeWithMostEdges returns the node with the highest edge count
func (a *AdjacencyList) NodeWithMostEdges() *Vertex {
	var maxEdgeNode *Vertex

	maxEdges := -1
	maxEdgeNode = nil

	for _, node := range a.Nodes() {
		if node.EdgeCount() > maxEdges {
			maxEdges = node.EdgeCount()
			maxEdgeNode = node
		}
	}

	return maxEdgeNode
}

// RemoveOrphans removes all vertices that have no edges
func (a *AdjacencyList) RemoveOrphans() {
	for _, node := range a.Nodes() {
		if node.EdgeCount() == 0 {
			a.RemoveNode(node)
		}
	}
}

// MinimalVertexCover returns vertices that make a minimal (not guaranteed to be minimum) vertex cover
func (a *AdjacencyList) MinimalVertexCover() map[string]*Vertex {
	aCopy := a.Copy()
	mvc := map[string]*Vertex{}

	for len(aCopy.Nodes()) > 0 {
		// Axiom: Nodes with whiskers must be in the MVC
		for _, node := range aCopy.Whiskers() {
			neighbor := node.FirstNeighbor()
			if neighbor != nil {
				// We might have already removed this neighbor B if, for instance,
				// A<->B<->C and we are at node C and have already processed node A.
				mvc[neighbor.ID()] = neighbor
				aCopy.RemoveNode(neighbor)
			}
			aCopy.RemoveNode(node)
		}

		// Axiom: Nodes without edges are not in the MVC
		aCopy.RemoveOrphans()

		if len(aCopy.Nodes()) == 0 {
			break
		}

		// Heuristic: Assume the most connected node is in the MVC
		node := aCopy.NodeWithMostEdges()
		aCopy.RemoveNode(node)
		mvc[node.ID()] = node
	}

	return mvc
}

// label returns a label for the given node
func label(node Vertex) string {
	return fmt.Sprintf("%s %d", node.Name(), node.Value())
}

// Serialize returns the adjacency list in graphviz format
func (a *AdjacencyList) Serialize(title string) string {
	g := dot.NewGraph(dot.Undirected)
	g.Label(title)

	// Create the nodes
	nodes := map[string]dot.Node{}
	for _, node := range a.Nodes() {
		l := label(*node)
		id := node.ID()
		n := g.Node(id)
		n.Label(l)
		nodes[id] = n
	}

	// Create the edges
	added := map[string]bool{}
	for _, node := range a.Nodes() {
		for _, neighbor := range node.Neighbors() {
			if added[node.ID()+neighbor.ID()] {
				// We have already added this edge
				continue
			}
			// Add the edge
			n1 := nodes[node.ID()]
			n2 := nodes[neighbor.ID()]
			g.Edge(n1, n2)
			// Record that we have already added this edge
			added[node.ID()+neighbor.ID()] = true
			added[neighbor.ID()+node.ID()] = true
		}
	}

	return g.String()
}
