package graphs

import (
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"runtime"
	"time"
)

var (
	RunningID int64
)

// traversePaths pushes to resultsCh [all] paths from startNode that touch each vertex
func traversePaths(a AdjList, resultsCh chan []*vertexes.Vertex, startNode *vertexes.Vertex, stopOnFirstPath bool, myID int64) {
	// Be concurrency safe in here!
	// * Pass in a COPY of the AdjacencyList. This routine may continue
	//   to live and access it even after the caller has torn down.
	// * Use locks around access to global values

	targetLen := a.NodeCount()

	todo := NewPath(a.EdgeCount())
	path := NewPath(a.NodeCount())

	// Initialize
	todo.Push(startNode, path.Len())

	for todo.Len() > 0 {
		if myID != RunningID {
			return
		}

		next, depth := todo.Pop()

		// We have a new node to put in the path, but
		// it may go in way back near the start
		for path.Len() > depth {
			path.PopAndTrack()
		}

		path.PushAndTrack(next, -1)

		if path.Len() == targetLen {
			// We found have a path!!!
			if myID != RunningID {
				return
			}
			resultsCh <- path.Get()
			if stopOnFirstPath {
				return
			}
			continue
		}

		for _, node := range next.Neighbors() {
			if !path.Contains(*node) {
				todo.Push(node, path.Len())
			}
		}
	}
}

// isEqualReverse returns true if path1 is the reverse of path2
func isEqualReverse(path1, path2 []*vertexes.Vertex) bool {
	if len(path1) != len(path2) {
		return false
	}

	for i := range path1 {
		if path1[i].ID() != path2[len(path2)-1-i].ID() {
			return false
		}
	}

	return true
}

// foundForward returns true if there is a path that is the reverse of newPath
func hasReverse(foundPaths [][]*vertexes.Vertex, newPath []*vertexes.Vertex) bool {
	for _, path := range foundPaths {
		if isEqualReverse(path, newPath) {
			return true
		}
	}
	return false
}

// paths returns all combinations of vertex orderings (valid paths or not)
func (a *AdjList) paths(terminals []*vertexes.Vertex, stopOnFirstPath bool, includeReverse bool) [][]*vertexes.Vertex {
	allPaths := [][]*vertexes.Vertex{}

	resultsCh := make(chan []*vertexes.Vertex, a.NodeCount()+1000) // How the go routines send us results

	runID := time.Now().UnixMicro()
	RunningID = runID

	// Sometimes some are already running
	goRoutinesAlreadyStarted := runtime.NumGoroutine()

	// Create one worker for each starting node
	if len(terminals) > 0 {
		for _, node := range terminals {
			tmpA := a.Copy()
			tmpNode := tmpA.nodes[node.ID()]
			go traversePaths(tmpA, resultsCh, tmpNode, stopOnFirstPath, runID)
		}
	} else {
		for _, node := range a.Nodes() {
			tmpA := a.Copy()
			tmpNode := tmpA.nodes[node.ID()]
			go traversePaths(tmpA, resultsCh, tmpNode, stopOnFirstPath, runID)
		}
	}

	// Collect results from the workers
	for {
		var path []*vertexes.Vertex

		path = nil
		for path == nil {
			select {
			case path = <-resultsCh:
			default:
			}
			if runtime.NumGoroutine() <= goRoutinesAlreadyStarted && path == nil {
				break
			}
			if path == nil {
				continue
			}
		}

		if path == nil {
			break
		}

		if !includeReverse && hasReverse(allPaths, path) {
			// We have already collected this path
			continue
		}

		allPaths = append(allPaths, path)

		if stopOnFirstPath {
			break
		}
	}

	RunningID = 0

	return allPaths
}

// HamiltonianPaths returns paths, the traversal of which touch each vertex once
func (a *AdjList) HamiltonianPaths(minLength int, stopOnFirstPath bool, includeReverse bool) [][]*vertexes.Vertex {
	// https://en.wikipedia.org/wiki/Hamiltonian_path

	if !a.Connected() {
		// An unconnected graph cannot have a path
		// This also eliminates empty graphs
		return nil
	}

	if a.NodeCount() < minLength {
		// Not a large enough graph to satisfy the caller
		return nil
	}

	// --- Vertex count >= 1, and they are connected ---

	if a.NodeCount() <= 2 {
		// All such graphs have a Hamiltonian path
		path := [][]*vertexes.Vertex{}
		path = append(path, []*vertexes.Vertex{})
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

	// --- Vertex count >= 3, whisker count <= 2, and graph is connected ---

	// Convert whisker map to a slice
	terminals := []*vertexes.Vertex{}
	for _, node := range whiskers {
		terminals = append(terminals, node)
	}

	// Try to improve the speed of the traversal
	for _, node := range a.Nodes() {
		node.SortNeighbors()
	}

	return a.paths(terminals, stopOnFirstPath, includeReverse)
}
