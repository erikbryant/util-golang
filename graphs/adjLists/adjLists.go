package adjLists

import (
	"fmt"
	"github.com/erikbryant/util-golang/graphs/vertexes"

	"github.com/emicklei/dot"
)

// AdjLists implements an undirected graph
type AdjLists struct {
	nodes map[uint64]*vertexes.Vertexes
}

// NewAL returns a new, empty adjacency list
func NewAL() AdjLists {
	return AdjLists{
		nodes: map[uint64]*vertexes.Vertexes{},
	}
}

// HasNode returns true if the node is already in the adjacency list
func (a *AdjLists) HasNode(node vertexes.Vertexes) bool {
	return a.nodes[node.ID()] != nil
}

// FirstNode returns a node from the map
func (a *AdjLists) FirstNode() *vertexes.Vertexes {
	for _, node := range a.Nodes() {
		return node
	}
	return nil
}

// AddNode adds a node to the adjacency list if not already present
func (a *AdjLists) AddNode(node *vertexes.Vertexes) {
	a.nodes[node.ID()] = node
}

// RemoveNode removes a node from the adjacency list
func (a *AdjLists) RemoveNode(node vertexes.Vertexes) {
	for _, n := range a.Nodes() {
		if n.ID() == node.ID() {
			// Remove it from its neighbors
			for _, neighbor := range node.Neighbors() {
				neighbor.RemoveNeighbor(node)
			}
			delete(a.nodes, n.ID())
			return
		}
	}
}

// Copy returns a copy of the AdjacencyList
func (a *AdjLists) Copy() AdjLists {
	newAL := NewAL()

	// Create copies of each node
	for _, node := range a.Nodes() {
		n := vertexes.NewVertex(node.Name(), node.Value())
		n.SetID(node.ID())      // TODO: replace with a copy function?
		newAL.nodes[n.ID()] = n // TODO: This direct access is slightly sketchy
	}

	// Link the new nodes together
	for _, node := range a.Nodes() {
		n := newAL.nodes[node.ID()]
		for _, neighbor := range node.Neighbors() {
			n.AddNeighbor(newAL.nodes[neighbor.ID()])
		}
	}

	return newAL
}

// Nodes returns the map of nodes in the adjacency list
func (a *AdjLists) Nodes() map[uint64]*vertexes.Vertexes {
	return a.nodes
}

// NodeCount returns the number of nodes in the adjacency list
func (a *AdjLists) NodeCount() int {
	return len(a.Nodes())
}

// AddEdge adds an edge, adding the nodes if they are not already present
func (a *AdjLists) AddEdge(n1, n2 *vertexes.Vertexes) {
	a.AddNode(n1)
	a.AddNode(n2)
	n1.AddNeighbor(n2)
	n2.AddNeighbor(n1)
}

// EdgeCount returns the number of distinct edges in the adjacency list
func (a *AdjLists) EdgeCount() int {
	edges := 0

	for _, node := range a.Nodes() {
		edges += node.NeighborCount()
	}

	// This is undirected, so each edge is listed twice
	return edges / 2
}

// Genus returns the genus number of the adjacency list
func (a *AdjLists) Genus() int {
	return a.EdgeCount() - a.NodeCount() + 1
}

// ValueSum returns the sum of all node values
func (a *AdjLists) ValueSum() int {
	sum := 0
	for _, node := range a.Nodes() {
		sum += node.Value()
	}
	return sum
}

// ValueLowest returns the node with the lowest value
func (a *AdjLists) ValueLowest() *vertexes.Vertexes {
	minVal := a.FirstNode()

	for _, node := range a.Nodes() {
		if node.Value() < minVal.Value() {
			minVal = node
		}
	}

	return minVal
}

// Whiskers returns a map of vertices that have only one edge
func (a *AdjLists) Whiskers() map[uint64]*vertexes.Vertexes {
	whiskers := map[uint64]*vertexes.Vertexes{}

	for _, node := range a.Nodes() {
		if node.NeighborCount() == 1 {
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
func (a *AdjLists) NodeWithMostEdges() *vertexes.Vertexes {
	var maxEdgeNode *vertexes.Vertexes

	maxEdges := -1
	maxEdgeNode = nil

	for _, node := range a.Nodes() {
		if node.NeighborCount() > maxEdges {
			maxEdges = node.NeighborCount()
			maxEdgeNode = node
		}
	}

	return maxEdgeNode
}

// RemoveOrphans removes all vertices that have no edges
func (a *AdjLists) RemoveOrphans() {
	for _, node := range a.Nodes() {
		if node.NeighborCount() == 0 {
			a.RemoveNode(*node)
		}
	}
}

func visitAll(node vertexes.Vertexes, visited map[uint64]bool) {
	if visited[node.ID()] {
		return
	}

	// Visit this node
	visited[node.ID()] = true

	// Visit each neighbor
	for _, neighbor := range node.Neighbors() {
		visitAll(*neighbor, visited)
	}
}

// Connected returns true if every vertex is reachable from every other vertex
func (a *AdjLists) Connected() bool {
	// https://en.wikipedia.org/wiki/Connectivity_(graph_theory)

	if len(a.Nodes()) == 0 {
		// An empty graph has no connections
		return false
	}

	// If each vertex has been visited, the graph is connected
	visited := map[uint64]bool{}
	visitAll(*a.FirstNode(), visited)
	return len(visited) == len(a.Nodes())
}

// MinimalVertexCover returns vertices that make a minimal (not guaranteed to be minimum) vertex cover
func (a *AdjLists) MinimalVertexCover() map[uint64]*vertexes.Vertexes {
	aCopy := a.Copy()
	mvc := map[uint64]*vertexes.Vertexes{}

	for len(aCopy.Nodes()) > 0 {
		// Axiom: Nodes with whiskers must be in the MVC
		for _, node := range aCopy.Whiskers() {
			neighbor := node.FirstNeighbor()
			if neighbor != nil {
				// We might have already removed this neighbor B if, for instance,
				// A<->B<->C, and we are at node C and have already processed node A.
				mvc[neighbor.ID()] = neighbor
				aCopy.RemoveNode(*neighbor)
			}
			aCopy.RemoveNode(*node)
		}

		// Axiom: Nodes without edges are not in the MVC
		aCopy.RemoveOrphans()

		if len(aCopy.Nodes()) == 0 {
			break
		}

		// Heuristic: Assume the most connected node is in the MVC
		node := aCopy.NodeWithMostEdges()
		aCopy.RemoveNode(*node)
		mvc[node.ID()] = node
	}

	return mvc
}

// Serialize returns the adjacency list in graphviz format
func (a *AdjLists) Serialize(title string) string {
	g := dot.NewGraph(dot.Undirected)
	g.Attr("label", title)
	g.Attr("overlap", "false")

	// Create the nodes
	nodes := map[uint64]dot.Node{}
	for _, node := range a.Nodes() {
		l := node.Label()
		id := node.ID()
		n := g.Node(fmt.Sprintf("%d", id))
		n.Label(l)
		nodes[id] = n
	}

	// Create the edges
	added := map[string]bool{}
	for _, node := range a.Nodes() {
		for _, neighbor := range node.Neighbors() {
			edgeAB := fmt.Sprintf("%d:%d", node.ID(), neighbor.ID())
			edgeBA := fmt.Sprintf("%d:%d", neighbor.ID(), node.ID())
			if added[edgeAB] {
				// We have already added this edge
				continue
			}
			// Add the edge
			n1 := nodes[node.ID()]
			n2 := nodes[neighbor.ID()]
			g.Edge(n1, n2)
			// Record that we have already added this edge
			added[edgeAB] = true
			added[edgeBA] = true
		}
	}

	return g.String()
}
