package graphs

import (
	"math/rand"
	"slices"
	"time"
)

// Vertex implements a graph Vertex
type Vertex struct {
	name      string
	value     int
	neighbors map[uint64]*Vertex
	id        uint64
}

// makeID returns a unique ID
func makeID() uint64 {
	// Get the micro time. That would almost be unique on its own
	// if it were not the case that we might have a flurry of vertex
	// objects created all at once. For uniqueness we really only
	// care about the lower half.
	// Fill the upper half with random garbage to make the composite
	// value unique.
	// We could use a string to make this simpler to create, but
	// uint64 is a much faster map index, making 'neighbors' about
	// 2x faster to access.
	lowHalf := 0x0000000000ffffff & uint64(time.Now().UnixMicro())
	topHalf := 0xffffffffff000000 & (uint64(rand.Float64()*10000000) << 32)
	return topHalf | lowHalf
}

// NewVertex returns a new Vertex
func NewVertex(name string, value int) *Vertex {
	return &Vertex{
		name:      name,
		value:     value,
		neighbors: map[uint64]*Vertex{},
		id:        makeID(),
	}
}

// Name returns the vertex name
func (v Vertex) Name() string {
	return v.name
}

// SetName sets the vertex name
func (v *Vertex) SetName(name string) {
	v.name = name
}

// Value returns the vertex value
func (v Vertex) Value() int {
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
func (v Vertex) ID() uint64 {
	return v.id
}

// HasNeighbor returns true if the given vertex is already a neighbor
func (v Vertex) HasNeighbor(node Vertex) bool {
	return v.neighbors[node.ID()] != nil
}

// Neighbors returns a map of all neighbors
func (v Vertex) Neighbors() map[uint64]*Vertex {
	return v.neighbors
}

// NeighborsSorted returns a sorted slice of all neighbors
func (v Vertex) NeighborsSorted() []*Vertex {
	keys := []uint64{}
	for key := range v.neighbors {
		keys = append(keys, key)
	}

	slices.Sort(keys)

	neighbors := []*Vertex{}
	for _, key := range keys {
		neighbors = append(neighbors, v.neighbors[key])
	}

	return neighbors
}

// FirstNeighbor returns the first neighbor
func (v Vertex) FirstNeighbor() *Vertex {
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
func (v Vertex) NeighborCount() int {
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

// Equal returns true if v and node are the same vertex
func (v Vertex) Equal(node Vertex) bool {
	return v.ID() == node.ID()
}
