package graphs

import (
	"fmt"
	"math/rand"
	"time"
)

// Vertex implements a graph Vertex
type Vertex struct {
	name        string
	value       int
	neighbors   []*Vertex
	neighborIDs map[string]bool
	id          string
}

// NewVertex returns a new Vertex
func NewVertex(name string, value int) Vertex {
	return Vertex{
		name:        name,
		value:       value,
		neighbors:   []*Vertex{},
		neighborIDs: map[string]bool{},
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

// HasNeighbor returns true if the given vertex is already a neighbor
func (v *Vertex) HasNeighbor(node Vertex) bool {
	return v.neighborIDs[node.ID()]
}

// Neighbors returns a slice of all neighbors
func (v *Vertex) Neighbors() []*Vertex {
	return v.neighbors
}

// AddNeighbor adds a neighbor vertex unless it is already a neighbor
func (v *Vertex) AddNeighbor(node *Vertex) {
	if v.HasNeighbor(*node) {
		return
	}
	v.neighbors = append(v.neighbors, node)
	v.neighborIDs[node.ID()] = true
}

// ID returns the id of the vertex
func (v *Vertex) ID() string {
	return v.id
}

// EdgeCount returns the number of edges (neighbors)
func (v *Vertex) EdgeCount() int {
	return len(v.neighbors)
}
