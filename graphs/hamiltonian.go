package graphs

import "slices"

func VisitAll(node Vertex, visited map[uint64]bool) {
	if visited[node.ID()] {
		return
	}

	// Visit this node
	visited[node.ID()] = true

	// Vist each neighbor
	for _, neighbor := range node.Neighbors() {
		VisitAll(*neighbor, visited)
	}
}

// Connected returns true if every vertex is reachable from every other vertex
func (a AdjacencyList) Connected() bool {
	// https://en.wikipedia.org/wiki/Connectivity_(graph_theory)

	if len(a.Nodes()) == 0 {
		// An empty graph has no connections
		return false
	}

	visited := map[uint64]bool{}

	VisitAll(*a.FirstNode(), visited)

	// If each vertices has been visited, the graph is connected
	for _, node := range a.Nodes() {
		if !visited[node.ID()] {
			return false
		}
	}

	return true
}

// traversePaths finds [all] paths that touch each vertex
func (a AdjacencyList) traversePaths(ch chan []*Vertex, terminal1, terminal2 *Vertex, stopOnFirstPath bool) {
	defer close(ch)

	startNodes := a.Nodes()

	// If we have terminal node overrides, use those instead
	if terminal1 != nil {
		startNodes = map[uint64]*Vertex{}
		startNodes[terminal1.ID()] = terminal1
		if terminal2 != nil {
			startNodes[terminal2.ID()] = terminal2
		}
	}

	keys := []uint64{}
	for key := range startNodes {
		keys = append(keys, key)
	}
	slices.Sort(keys)

	// Look for paths from each possible starting node
	for _, key := range keys {
		node := startNodes[key]

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
