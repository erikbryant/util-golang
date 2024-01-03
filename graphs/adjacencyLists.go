package graphs

import (
	"fmt"

	"github.com/emicklei/dot"
)

// AdjacencyList implements an undirected graph
type AdjacencyList struct {
	nodes   []*Vertex
	nodeIDs map[string]bool
}

// NewAL returns a new, empty adjacdency list
func NewAL() AdjacencyList {
	return AdjacencyList{
		nodes:   []*Vertex{},
		nodeIDs: map[string]bool{},
	}
}

// HasNode returns true if the node is already in the adjacency list
func (a *AdjacencyList) HasNode(node Vertex) bool {
	return a.nodeIDs[node.ID()]
}

// AddNode adds a node to the adjacency list if not already present
func (a *AdjacencyList) AddNode(node *Vertex) {
	if a.HasNode(*node) {
		return
	}
	a.nodes = append(a.nodes, node)
	a.nodeIDs[node.ID()] = true
}

// Nodes returns the slice of nodes in the adjacency list
func (a *AdjacencyList) Nodes() []*Vertex {
	return a.nodes
}

// NodeCount returns the number of nodes in the adjacency list
func (a *AdjacencyList) NodeCount() int {
	return len(a.Nodes())
}

// AddEdge adds the two nodes as an edge, adding the nodes if they are not already present
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
	if len(a.nodes) == 0 {
		return nil
	}

	min := a.nodes[0]

	for _, node := range a.Nodes() {
		if node.Value() < min.Value() {
			min = node
		}
	}

	return min
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
