package vertexes

import (
	"testing"
)

func TestMakeID(t *testing.T) {
	// If we generate a large number of IDs are any
	// duplicates?
	ids := map[uint64]bool{}
	for i := 0; i < 1000000; i++ {
		id := makeID()
		if ids[id] {
			t.Errorf("ERROR: Got duplicateID %d", id)
		}
		ids[id] = true
	}
}

func TestValues(t *testing.T) {
	v1 := NewVertex("ABC", 123)
	answer := v1.Value()
	if answer != 123 {
		t.Errorf("ERROR: Expected 123, got %d", answer)
	}

	v1.Increment()
	answer = v1.Value()
	if answer != 124 {
		t.Errorf("ERROR: Expected 124, got %d", answer)
	}

	v1.SetValue(101)
	answer = v1.Value()
	if answer != 101 {
		t.Errorf("ERROR: Expected 101, got %d", answer)
	}

	v1.Decrement()
	answer = v1.Value()
	if answer != 100 {
		t.Errorf("ERROR: Expected 100, got %d", answer)
	}
}

func TestHasNeighbor(t *testing.T) {
	v1 := NewVertex("ABC", 123)
	v2 := NewVertex("DEF", 345)
	v1.AddNeighbor(v2)
	answer := v1.HasNeighbor(*v2)
	if !answer {
		t.Errorf("ERROR: For '1 neighbor - v1/v2' expected true, got %t", answer)
	}
	answer = v2.HasNeighbor(*v1)
	if answer {
		t.Errorf("ERROR: For '1 neighbor - v2/v1' expected false, got %t", answer)
	}
}

func TestFirstNeighbor(t *testing.T) {
	v1 := NewVertex("ABC", 123)
	answer := v1.FirstNeighbor()
	if answer != nil {
		t.Errorf("ERROR: For 'first neighbor' expected nil, got %d", answer.ID())
	}

	v1 = NewVertex("ABC", 123)
	v2 := NewVertex("DEF", 345)
	v1.AddNeighbor(v2)
	answer = v1.FirstNeighbor()
	if answer.ID() != v2.ID() {
		t.Errorf("ERROR: For 'first neighbor' expected %d, got %d", v1.ID(), answer.ID())
	}
}

func TestRemoveNeighbor(t *testing.T) {
	v1 := NewVertex("ABC", 123)
	v2 := NewVertex("DEF", 345)
	v1.AddNeighbor(v2)
	v1.RemoveNeighbor(*v2)
	if len(v1.Neighbors()) != 0 {
		t.Errorf("ERROR: For 'remove neighbor' expected len 0, got len %d", len(v1.Neighbors()))
	}
}

func TestEdgeCountDegree(t *testing.T) {
	v1 := NewVertex("A", 1)
	answer := v1.NeighborCount()
	if answer != 0 {
		t.Errorf("ERROR: Expected 0, got %d", answer)
	}
	answer = v1.Degree()
	if answer != 0 {
		t.Errorf("ERROR: Expected 0, got %d", answer)
	}

	v2 := NewVertex("B", 2)
	answer = v2.NeighborCount()
	if answer != 0 {
		t.Errorf("ERROR: Expected 0, got %d", answer)
	}
	answer = v2.Degree()
	if answer != 0 {
		t.Errorf("ERROR: Expected 0, got %d", answer)
	}

	v1.AddNeighbor(v2)
	answer = v1.NeighborCount()
	if answer != 1 {
		t.Errorf("ERROR: Expected 1, got %d", answer)
	}
	answer = v1.Degree()
	if answer != 1 {
		t.Errorf("ERROR: Expected 1, got %d", answer)
	}

	v1.AddNeighbor(v1)
	answer = v1.NeighborCount()
	if answer != 2 {
		t.Errorf("ERROR: Expected 2, got %d", answer)
	}
	answer = v1.Degree()
	if answer != 3 {
		t.Errorf("ERROR: Expected 3, got %d", answer)
	}
}

func TestEqual(t *testing.T) {
	v1 := NewVertex("ABC", 123)
	answer := v1.Equal(*v1)
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}

	v2 := NewVertex("ABC", 123)
	answer = v1.Equal(*v2)
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	v1.AddNeighbor(v2)
	vN := v1.FirstNeighbor()
	answer = v1.Equal(*vN)
	if answer != false {
		t.Errorf("ERROR: Expected false, got %t", answer)
	}

	v3 := NewVertex("ABC", 123)
	v3.AddNeighbor(v3)
	vN = v3.FirstNeighbor()
	answer = v3.Equal(*vN)
	if answer != true {
		t.Errorf("ERROR: Expected true, got %t", answer)
	}
}
