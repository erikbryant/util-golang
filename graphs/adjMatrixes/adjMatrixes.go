package adjMatrixes

import (
	"fmt"
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"slices"
	"strings"
)

type AdjMatrix struct {
	nodes        map[uint64]*vertexes.Vertexes
	nodesOrdered []uint64
	matrix       [][]int
}

func Matrix() *AdjMatrix {
	return &AdjMatrix{
		nodes: map[uint64]*vertexes.Vertexes{},
	}
}

// AddNode adds the given node to the matrix (be sure to call ComputeDistances after)
func (m *AdjMatrix) AddNode(id uint64, node *vertexes.Vertexes) {
	m.nodes[id] = node
	m.nodesOrdered = append(m.nodesOrdered, node.ID())
	slices.Sort(m.nodesOrdered)
}

// NodeCount returns the number of vertexes in the matrix
func (m *AdjMatrix) NodeCount() int {
	return len(m.nodesOrdered)
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

// initMatrix allocates space for the matrix
func (m *AdjMatrix) initMatrix() {
	size := m.NodeCount()

	if len(m.matrix) != size {
		// Initialize major diagonal with 0, -1 elsewhere
		m.matrix = make([][]int, size)
		for i := range m.matrix {
			m.matrix[i] = make([]int, size)
			for j := 0; j < size; j++ {
				m.matrix[i][j] = -1
			}
			m.matrix[i][i] = 0
		}
	}

}

// ComputeDistances sets the matrix distance values for each vertex pair
func (m *AdjMatrix) ComputeDistances() {
	// Make sure the matrix is allocated
	m.initMatrix()

	// Populate the distance from each vertex to each other vertex
	for r := range m.nodesOrdered {
		m.reachable(r)
	}
}

// GetValue returns the value at that row/col
func (m *AdjMatrix) GetValue(r, c int) int {
	// Make sure the matrix is allocated
	m.initMatrix()

	return m.matrix[r][c]
}

// SetValue returns the value at that row/col
func (m *AdjMatrix) SetValue(r, c, value int) {
	// Make sure the matrix is allocated
	m.initMatrix()

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

// Diameter returns the length of the path between the two most distance vertexes
func (m *AdjMatrix) Diameter() int {
	// https://en.wikipedia.org/wiki/Distance_(graph_theory)

	if m.NodeCount() == 0 {
		// Undefined for no nodes
		return -1
	}

	if m.NodeCount() == 1 {
		// The distance between a vertex and itself is zero
		return 0
	}

	size := m.NodeCount()
	maxDistance := -1
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			maxDistance = max(m.GetValue(r, c), maxDistance)
		}
	}

	return maxDistance
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
	s += fmt.Sprintf("%dx%d\n", m.NodeCount(), m.NodeCount())

	// Column headers
	s += "     "
	for c := 0; c < m.NodeCount(); c++ {
		label := "???"
		node := m.NodeFromIndex(c)
		if node != nil {
			label = trunc(node.Label(), 4)
		}
		s += fmt.Sprintf("%4s ", label)
	}
	s += "\n"

	// Body of the matrix
	for r := 0; r < m.NodeCount(); r++ {
		label := "???"
		node := m.NodeFromIndex(r)
		if node != nil {
			label = trunc(node.Label(), 4)
		}
		s += fmt.Sprintf("%4s ", label)
		for c := 0; c < m.NodeCount(); c++ {
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
