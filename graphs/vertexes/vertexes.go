package vertexes

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// Vertex implements a graph Vertex
type Vertex struct {
	name        string
	value       int
	neighbor    []*Vertex
	neighborIDs map[uint64]*Vertex
	id          uint64
}

// makeID returns a unique ID
func makeID() uint64 {
	// Use micro time as the ID. That would almost be unique on its own
	// if it were not the case that we might have a flurry of vertex
	// objects created all at once, even in the same microsecond.
	//
	// For uniqueness, we really only care about the lower half.
	// Fill the upper half with random garbage to make the composite
	// value unique.
	//
	// We could use a string to make this simpler to create, but
	// uint64 is a much faster map index, making 'neighborIDs' about
	// 1.3x faster to access.

	lowHalf := 0x00000000ffffffff & uint64(time.Now().UnixMicro())
	topHalf := rand.Uint64() << 32
	return topHalf | lowHalf
}

// NewVertex returns a new Vertex
func NewVertex(name string, value int) *Vertex {
	return &Vertex{
		name:        name,
		value:       value,
		neighbor:    []*Vertex{},
		neighborIDs: map[uint64]*Vertex{},
		id:          makeID(),
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
	v.value++
}

// Decrement decrements the vertex value by 1
func (v *Vertex) Decrement() {
	v.value--
}

// SetValue sets the vertex value
func (v *Vertex) SetValue(value int) {
	v.value = value
}

// ID returns the id of the vertex
func (v *Vertex) ID() uint64 {
	return v.id
}

// SetID sets the id of the vertex (use with extreme caution)
func (v *Vertex) SetID(id uint64) {
	v.id = id
}

// HasNeighbor returns true if the given vertex is already a neighbor
func (v *Vertex) HasNeighbor(vertex Vertex) bool {
	return v.neighborIDs[vertex.ID()] != nil
}

// Neighbors returns a slice of all neighbors
func (v *Vertex) Neighbors() []*Vertex {
	return v.neighbor
}

// SortNeighbors returns a sorted slice of all neighbors
func (v *Vertex) SortNeighbors() {
	sort.Slice(v.neighbor, func(i, j int) bool {
		return v.neighbor[i].NeighborCount() < v.neighbor[j].NeighborCount()
	})
}

// FirstNeighbor returns the first neighbor
func (v *Vertex) FirstNeighbor() *Vertex {
	if len(v.neighbor) == 0 {
		return nil
	}
	return v.neighbor[0]
}

// AddNeighbor adds a neighbor vertex
func (v *Vertex) AddNeighbor(vertex *Vertex) {
	if vertex == nil {
		return
	}
	v.neighbor = append(v.neighbor, vertex)
	v.neighborIDs[vertex.ID()] = vertex
}

// RemoveNeighbor removes a neighbor vertex
func (v *Vertex) RemoveNeighbor(vertex Vertex) {
	delete(v.neighborIDs, vertex.ID())

	for i := 0; i < len(v.neighbor); i++ {
		if vertex.Equal(*v.neighbor[i]) {
			// Copy terminal cell to here, shrink slice by one
			v.neighbor[i] = v.neighbor[len(v.neighbor)-1]
			v.neighbor = v.neighbor[:len(v.neighbor)-1]
			break
		}
	}
}

// NeighborCount returns the number of edges (neighbors)
func (v *Vertex) NeighborCount() int {
	return len(v.neighbor)
}

// Degree returns the degree of the incoming edges (loops count as 2)
func (v *Vertex) Degree() int {
	degree := 0

	for _, vertex := range v.Neighbors() {
		if v.ID() == vertex.ID() {
			// We own *both* ends of this edge
			degree++
		}
		degree++
	}

	return degree
}

// Equal returns true if v and vertex represent the same vertex
func (v *Vertex) Equal(vertex Vertex) bool {
	return v.ID() == vertex.ID()
}

// Label returns a label for the given vertex
func (v *Vertex) Label() string {
	return fmt.Sprintf("%s %d", v.Name(), v.Value())
}
