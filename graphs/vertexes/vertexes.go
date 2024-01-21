package vertexes

import (
	"fmt"
	"sort"
)

// Vertexes implements a graph vertex
type Vertexes struct {
	name        string
	value       int
	userValue   *interface{}
	neighbor    []*Vertexes
	neighborIDs map[uint64]*Vertexes
	id          uint64
}

var uniqueID uint64 = 0

// makeID returns a unique ID
func makeID() uint64 {
	uniqueID++
	return uniqueID
}

// NewVertex returns a new Vertexes
func NewVertex(name string, value int) *Vertexes {
	return &Vertexes{
		name:        name,
		value:       value,
		neighbor:    []*Vertexes{},
		neighborIDs: map[uint64]*Vertexes{},
		id:          makeID(),
	}
}

// Name returns the vertex name
func (v *Vertexes) Name() string {
	return v.name
}

// SetName sets the vertex name
func (v *Vertexes) SetName(name string) {
	v.name = name
}

// Value returns the vertex value
func (v *Vertexes) Value() int {
	return v.value
}

// Increment increments the vertex value by 1
func (v *Vertexes) Increment() {
	v.value++
}

// Decrement decrements the vertex value by 1
func (v *Vertexes) Decrement() {
	v.value--
}

// SetValue sets the vertex value
func (v *Vertexes) SetValue(value int) {
	v.value = value
}

// SetUV stores the arbitrary value sent by the caller
func (v *Vertexes) SetUV(uv *interface{}) {
	v.userValue = uv
}

// GetUV returns the arbitrary value that had been stored by the caller
func (v *Vertexes) GetUV() *interface{} {
	return v.userValue
}

// ID returns the id of the vertex
func (v *Vertexes) ID() uint64 {
	return v.id
}

// SetID sets the id of the vertex (use with extreme caution)
func (v *Vertexes) SetID(id uint64) {
	v.id = id
}

// HasNeighbor returns true if the given vertex is already a neighbor
func (v *Vertexes) HasNeighbor(vertex Vertexes) bool {
	return v.neighborIDs[vertex.ID()] != nil
}

// Neighbors returns a slice of all neighbors
func (v *Vertexes) Neighbors() []*Vertexes {
	return v.neighbor
}

// SortNeighbors returns a sorted slice of all neighbors
func (v *Vertexes) SortNeighbors() {
	sort.Slice(v.neighbor, func(i, j int) bool {
		return v.neighbor[i].NeighborCount() < v.neighbor[j].NeighborCount()
	})
}

// FirstNeighbor returns the first neighbor
func (v *Vertexes) FirstNeighbor() *Vertexes {
	if len(v.neighbor) == 0 {
		return nil
	}
	return v.neighbor[0]
}

// AddNeighbor adds a neighbor vertex
func (v *Vertexes) AddNeighbor(vertex *Vertexes) {
	if vertex == nil {
		return
	}
	v.neighbor = append(v.neighbor, vertex)
	v.neighborIDs[vertex.ID()] = vertex
}

// RemoveNeighbor removes a neighbor vertex
func (v *Vertexes) RemoveNeighbor(vertex Vertexes) {
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
func (v *Vertexes) NeighborCount() int {
	return len(v.neighbor)
}

// Degree returns the degree of the incoming edges (loops count as 2)
func (v *Vertexes) Degree() int {
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
func (v *Vertexes) Equal(vertex Vertexes) bool {
	return v.ID() == vertex.ID()
}

// derefUV returns the value of userValue or nil if userValue is nil
func (v *Vertexes) derefUV() interface{} {
	uv := v.GetUV()
	if uv == nil {
		return nil
	}
	return *uv
}

// Label returns a label for the given vertex
func (v *Vertexes) Label() string {
	return fmt.Sprintf("%s %d %v", v.Name(), v.Value(), v.derefUV())
}
