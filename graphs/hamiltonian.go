package graphs

import (
	"sort"
)

// traversePaths finds [all] paths that touch each vertex
func (a AdjacencyList) traversePaths(ch chan []*Vertex, terminal1, terminal2 *Vertex, stopOnFirstPath bool) {
	defer close(ch)

	var startNodes []*Vertex

	// If we have terminal node overrides, use those instead
	if terminal1 == nil {
		for _, node := range a.Nodes() {
			startNodes = append(startNodes, node)
		}
	} else {
		startNodes = append(startNodes, terminal1)
		if terminal2 != nil {
			startNodes = append(startNodes, terminal2)
		}
	}

	sort.Slice(startNodes, func(i, j int) bool {
		return startNodes[i].NeighborCount() < startNodes[j].NeighborCount()
	})

	// Look for paths from each possible starting node
	for _, node := range startNodes {

		todo := NewPath(a.EdgeCount())
		path := NewPath(a.NodeCount())

		// Initialize
		todo.Push(node, path.Len())

		for todo.Len() > 0 {
			next, depth := todo.Pop()

			// We have a new node to put in the path, but
			// it may go in waaaay back near the start
			for path.Len() > depth {
				path.Pop()
			}

			// if path.Contains(*next) {
			// 	// We have already visited this node
			// 	continue
			// }

			path.Push(next, -1)

			if path.Len() == a.NodeCount() {
				// We have a path!!!
				ch <- path.Get()
				if stopOnFirstPath {
					return
				}
				continue
			}

			for _, node := range next.NeighborsSorted() {
				if !path.Contains(*node) {
					todo.Push(node, path.Len())
				}
			}
		}
	}
}

// allPotentialPaths returns all combinations of vertex orderings (valid paths or not)
func (a AdjacencyList) allPotentialPaths(terminal1, terminal2 *Vertex) [][]*Vertex {
	allPaths := [][]*Vertex{}

	ch := make(chan []*Vertex)
	go a.traversePaths(ch, terminal1, terminal2, true)

	for {
		path, ok := <-ch
		if !ok {
			break
		}
		allPaths = append(allPaths, path)
	}

	return allPaths
}

// HamiltonianPaths returns slices, the traversal of which touch each vertex once
func (a AdjacencyList) HamiltonianPaths() [][]*Vertex {
	// https://en.wikipedia.org/wiki/Hamiltonian_path

	if !a.Connected() {
		// An unconnected graph cannot have a path
		// This also eliminates empty graphs
		return nil
	}

	// --- Vertex count >= 1 and they are connected ---

	if a.NodeCount() <= 2 {
		// All such graphs have a Hamiltonian path
		path := [][]*Vertex{}
		path = append(path, []*Vertex{})
		for _, node := range a.Nodes() {
			path[0] = append(path[0], node)
		}
		return path
	}

	whiskers := a.Whiskers()

	if len(whiskers) > 2 {
		// No possible path
		return nil
	}

	// --- Vertex count >= 3 and they are connected ---

	// Convert whisker map to something we can index into.
	terminals := []*Vertex{}
	for _, node := range whiskers {
		terminals = append(terminals, node)
	}

	// Pad so we are sure to have at least 2 elements.
	terminals = append(terminals, nil)
	terminals = append(terminals, nil)

	return a.allPotentialPaths(terminals[0], terminals[1])
}
