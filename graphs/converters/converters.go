package converters

import (
	"github.com/erikbryant/util-golang/graphs/adjLists"
	"github.com/erikbryant/util-golang/graphs/adjMatrixes"
)

// MatrixFromAdjList returns a new AdjMatrix representing the given AdjLists
func MatrixFromAdjList(al *adjLists.AdjLists) *adjMatrixes.AdjMatrix {
	m := adjMatrixes.Matrix()

	// Copy the list of vertexes into the matrix
	for id, node := range al.Nodes() {
		m.AddNode(id, node)
	}

	// Calculate the vertex distances
	m.ComputeDistances()

	return m
}
