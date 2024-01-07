package graphs

import (
	"fmt"
	"math/rand"
	"time"
)

// Vertex implements a graph Vertex
type Vertex struct {
	name      string
	value     int
	neighbors map[string]*Vertex
	id        string
}

// NewVertex returns a new Vertex
func NewVertex(name string, value int) Vertex {
	return Vertex{
		name:      name,
		value:     value,
		neighbors: map[string]*Vertex{},
		// Generate an ID guaranteed to be unique
		id: fmt.Sprintf("%v:%v", time.Now().UnixMicro(), rand.Float64()),
	}
}

// Name returns the vertex name
func (v *Vertex) Name() string {
	return v.name
}

// SetName sets the vertex name
func (v *Vertex) SetName(name string) {
	v.name = name
}

// Value returns the vertex value
func (v *Vertex) Value() int {
	return v.value
}

// Increment increments the vertex value by 1
func (v *Vertex) Increment() {
	// Should this call SetValue()?
	v.value++
}

// Decrement decrements the vertex value by 1
func (v *Vertex) Decrement() {
	// Should this call SetValue()?
	v.value--
}

// SetValue sets the vertex value
func (v *Vertex) SetValue(value int) {
	v.value = value
}

// ID returns the id of the vertex
func (v *Vertex) ID() string {
	return v.id
}

// HasNeighbor returns true if the given vertex is already a neighbor
func (v *Vertex) HasNeighbor(node Vertex) bool {
	return v.neighbors[node.ID()] != nil
}

// Neighbors returns a map of all neighbors
func (v *Vertex) Neighbors() map[string]*Vertex {
	return v.neighbors
}

// FirstNeighbor returns the first neighbor
func (v *Vertex) FirstNeighbor() *Vertex {
	for _, node := range v.neighbors {
		return node
	}
	return nil
}

// AddNeighbor adds a neighbor vertex
func (v *Vertex) AddNeighbor(node *Vertex) {
	v.neighbors[node.ID()] = node
}

// RemoveNeighbor removes a neighbor vertex
func (v *Vertex) RemoveNeighbor(node Vertex) {
	delete(v.neighbors, node.ID())
}

// NeighborCount returns the number of edges (neighbors)
func (v *Vertex) NeighborCount() int {
	return len(v.neighbors)
}

// Degree returns the degree of the incoming edges (loops count as 2)
func (v *Vertex) Degree() int {
	degree := 0

	for _, node := range v.Neighbors() {
		if v.ID() == node.ID() {
			// We own *both* ends of this edge
			degree++
		}
		degree++
	}

	return degree
}
