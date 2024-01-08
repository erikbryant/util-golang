package graphs

import "testing"

func TestPush(t *testing.T) {
	testCases := []struct {
		v     *Vertex
		depth int
	}{
		{NewVertex("A", 5), 3},
		{NewVertex("B", 12), -1},
	}

	path := NewPath(10)
	for _, testCase := range testCases {
		path.Push(testCase.v, testCase.depth)
		v, depth := path.Pop()
		if v != testCase.v || depth != testCase.depth {
			t.Errorf("ERROR: For %v expected %v/%d, got %v/%d", testCase.v, testCase.v, testCase.depth, v, depth)
		}
	}
}

func TestPop(t *testing.T) {
	path := NewPath(10)

	// Empty stack
	l := path.Len()
	if l != 0 {
		t.Errorf("ERROR: Expected 0, got %d", l)
	}
	v, depth := path.Pop()
	if v != nil || depth != 0 {
		t.Errorf("ERROR: Expected nil/0, got %v/%d", v, depth)
	}

	// Can pop a node that was just pushed
	v1 := NewVertex("X", 999)
	path.Push(v1, 99)
	v, depth = path.Pop()
	if v != v1 || depth != 99 {
		t.Errorf("ERROR: Expected %v/%d, got %v/%d", v1, 99, v, depth)
	}
}

func TestGet(t *testing.T) {
	path := NewPath(10)

	v0 := NewVertex("X", 0)
	v1 := NewVertex("Y", 1)
	v2 := NewVertex("Z", 2)

	path.Push(v0, 0)
	path.Push(v1, 1)
	path.Push(v2, 2)

	p := path.Get()
	if len(p) != 3 {
		t.Errorf("ERROR: Expected 3, got %d", len(p))
	}
	if p[0] != v0 {
		t.Errorf("ERROR: Expected %v, got %v", v0, p[0])
	}
	if p[1] != v1 {
		t.Errorf("ERROR: Expected %v, got %v", v1, p[1])
	}
	if p[2] != v2 {
		t.Errorf("ERROR: Expected %v, got %v", v2, p[2])
	}
}

func TestContains(t *testing.T) {
	path := NewPath(10)

	v0 := NewVertex("X", 0)
	v1 := NewVertex("Y", 1)
	v2 := NewVertex("Z", 2)

	path.Push(v0, 0)
	path.Push(v1, 1)
	path.Push(v2, 2)

	answer := path.Contains(*v1)
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	path.Pop()
	path.Pop()
	answer = path.Contains(*v1)
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}
}

func TestReset(t *testing.T) {
	path := NewPath(10)

	v0 := NewVertex("X", 0)
	v1 := NewVertex("Y", 1)
	v2 := NewVertex("Z", 2)

	path.Push(v0, 0)
	path.Push(v1, 1)
	path.Push(v2, 2)

	path.Reset()
	answer := path.Contains(*v1)
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	l := path.Len()
	if l != 0 {
		t.Errorf("ERROR: Expected 0, got %d", l)
	}

	p := path.Get()
	l = len(p)
	if l != 0 {
		t.Errorf("ERROR: Expected 0, got %d", l)
	}
}

func TestIsPath(t *testing.T) {
	// Empty path
	path := []*Vertex{}
	answer := IsPath(path)
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	v0 := NewVertex("X", 0)
	v1 := NewVertex("Y", 1)
	v2 := NewVertex("Z", 2)

	// Not a path
	path = []*Vertex{v0, v1, v2}
	answer = IsPath(path)
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	// Is a path
	v0.AddNeighbor(v1)
	v1.AddNeighbor(v2)
	answer = IsPath(path)
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}
}

func TestIsCycle(t *testing.T) {
	// Empty path
	path := []*Vertex{}
	answer := IsPath(path)
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	v0 := NewVertex("X", 0)
	v1 := NewVertex("Y", 1)
	v2 := NewVertex("Z", 2)

	// Not a path
	path = []*Vertex{v0, v1, v2}
	answer = IsPath(path)
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	// Is a path
	v0.AddNeighbor(v1)
	v1.AddNeighbor(v2)
	answer = IsPath(path)
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	// Is a path & a cycle (by neighbors)
	v0.AddNeighbor(v1)
	v1.AddNeighbor(v2)
	v2.AddNeighbor(v0)
	answer = IsPath(path)
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	// Is a path & a cycle (by equal start/end nodes)
	path = append(path, v0)
	answer = IsPath(path)
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}
}
