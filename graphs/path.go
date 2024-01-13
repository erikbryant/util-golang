package graphs

type Path struct {
	path    []*Vertex
	depth   []int
	index   int
	visited map[uint64]bool
}

func NewPath(maxLen int) Path {
	return Path{
		path:    make([]*Vertex, maxLen),
		depth:   make([]int, maxLen),
		index:   -1,
		visited: make(map[uint64]bool, maxLen),
	}
}

func (p *Path) PushAndTrack(node *Vertex, depth int) {
	p.Push(node, depth)
	p.visited[node.ID()] = true
}

func (p *Path) Push(node *Vertex, depth int) {
	p.index++
	p.path[p.index] = node
	p.depth[p.index] = depth
}

func (p *Path) PopAndTrack() (*Vertex, int) {
	node, depth := p.Pop()
	if node != nil {
		p.visited[node.ID()] = false
	}
	return node, depth
}

func (p *Path) Pop() (*Vertex, int) {
	if p.index < 0 {
		// The path is empty
		return nil, 0
	}

	node := p.path[p.index]
	depth := p.depth[p.index]
	p.index--

	return node, depth
}

func (p Path) Len() int {
	return p.index + 1
}

func (p Path) Get() []*Vertex {
	path := make([]*Vertex, p.index+1)
	copy(path, p.path)
	return path
}

func (p Path) Contains(node Vertex) bool {
	return p.visited[node.ID()]
}

func (p *Path) Reset() {
	*p = NewPath(len(p.path))
}

// IsPath returns true if the nodes form a path
func IsPath(path []*Vertex) bool {
	if len(path) == 0 {
		// An empty path is not a path
		return false
	}

	for i := 0; i < len(path)-1; i++ {
		if !path[i].HasNeighbor(*path[i+1]) {
			return false
		}
	}

	return true
}

// IsCycle returns true if the nodes form a cycle
func IsCycle(path []*Vertex) bool {
	if !IsPath(path) {
		// If it is not a path it cannot be a cycle
		return false
	}

	first := path[0]
	last := path[len(path)-1]

	// A cycle is a path that returns to the start vertex
	// or to a neighbor of the start vertex
	return first.Equal(*last) || first.HasNeighbor(*last)
}
