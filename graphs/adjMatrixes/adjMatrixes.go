package adjMatrixes

import (
	"fmt"
	"github.com/erikbryant/util-golang/graphs/adjLists"
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"slices"
	"strings"
)

type AdjMatrix struct {
	size         int
	nodes        map[uint64]*vertexes.Vertexes
	nodesOrdered []uint64
	matrix       [][]int
}

// Matrix returns a matrix of size x size with the diagonal set to zero and all other cells set to -1
func Matrix(size int) *AdjMatrix {
	// TODO: This is not used. Consider removing it.

	m := AdjMatrix{
		size: size,
	}

	m.nodes = make(map[uint64]*vertexes.Vertexes)
	m.nodesOrdered = make([]uint64, 0)

	// Initialize major diagonal with 0, -1 elsewhere
	m.matrix = make([][]int, m.size)
	for i := range m.matrix {
		m.matrix[i] = make([]int, m.size)
		for j := 0; j < m.size; j++ {
			m.matrix[i][j] = -1
		}
		m.matrix[i][i] = 0
	}

	return &m
}

// reachable sets the weights for each column (that is reachable) for row r
func (m *AdjMatrix) reachable(r int) {
	todo := map[uint64]bool{}
	visited := map[uint64]bool{}
	depth := 0

	node := m.NodeFromIndex(r)
	visited[node.ID()] = true
	for _, neighbor := range node.Neighbors() {
		todo[neighbor.ID()] = true
	}

	for len(todo) > 0 {
		next := map[uint64]bool{}

		for id := range todo {
			c := m.IndexFromID(id)
			m.SetValue(r, c, 1+depth)
			m.SetValue(c, r, 1+depth)
			visited[id] = true
			for _, neighbor := range m.nodes[id].Neighbors() {
				if visited[neighbor.ID()] {
					continue
				}
				next[neighbor.ID()] = true
			}
		}

		depth++
		todo = next
	}
}

// MatrixFromAdjList returns a new matrix representing al
func MatrixFromAdjList(al *adjLists.AdjLists) *AdjMatrix {
	m := Matrix(len(al.Nodes()))

	// Record a copy of the list of vertexes
	for id, node := range al.Nodes() {
		m.nodes[id] = node
		m.nodesOrdered = append(m.nodesOrdered, node.ID())
	}

	slices.Sort(m.nodesOrdered)

	// Populate the distance from each vertex to each other vertex
	for r := range m.nodesOrdered {
		m.reachable(r)
	}

	return m
}

// Size returns the width/height (they are identical) of the matrix
func (m *AdjMatrix) Size() int {
	return m.size
}

// GetValue returns the value at that row/col
func (m *AdjMatrix) GetValue(r, c int) int {
	return m.matrix[r][c]
}

// SetValue returns the value at that row/col
func (m *AdjMatrix) SetValue(r, c, value int) {
	m.matrix[r][c] = value
}

// NodeFromIndex returns the node for the given index or nil if not found
func (m *AdjMatrix) NodeFromIndex(i int) *vertexes.Vertexes {
	if i < 0 || i >= len(m.nodesOrdered) {
		return nil
	}
	return m.nodes[m.nodesOrdered[i]]
}

// IndexFromID returns the index of the vertex ID or -1 if not found
func (m *AdjMatrix) IndexFromID(id uint64) int {
	// TODO: Consider making this a map at matrix creation time
	for i, nodeID := range m.nodesOrdered {
		if nodeID == id {
			return i
		}
	}
	return -1
}

// trunc returns a [possibly] truncated string
func trunc(s string, l int) string {
	str := strings.TrimSpace(s)
	if len(str) > l {
		return str[:l]
	}
	return str
}

// Serialize returns a string representation of the matrix
func (m *AdjMatrix) Serialize() string {
	s := ""

	// Meta info
	s += fmt.Sprintf("%dx%d\n", m.size, m.size)

	// Column headers
	s += "     "
	for c := 0; c < m.size; c++ {
		label := "???"
		node := m.NodeFromIndex(c)
		if node != nil {
			label = trunc(node.Label(), 4)
		}
		s += fmt.Sprintf("%4s ", label)
	}
	s += "\n"

	// Body of the matrix
	for r := 0; r < m.size; r++ {
		label := "???"
		node := m.NodeFromIndex(r)
		if node != nil {
			label = trunc(node.Label(), 4)
		}
		s += fmt.Sprintf("%4s ", label)
		for c := 0; c < m.size; c++ {
			// -1 is a representation of no edge, not an actual value
			v := m.GetValue(r, c)
			if v == -1 {
				s += fmt.Sprintf("%4s ", "·")
			} else {
				s += fmt.Sprintf("%4d ", v)
			}
		}
		s += "\n"
	}

	return s
}
