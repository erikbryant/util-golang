package graphs

import "testing"

func TestHasNeighbor(t *testing.T) {
	v1 := NewVertex("ABC", 123)
	v2 := NewVertex("DEF", 345)
	v1.AddNeighbor(&v2)
	answer := v1.HasNeighbor(v2)
	if !answer {
		t.Errorf("ERROR: For '1 neighbor - v1/v2' expected true, got %t", answer)
	}
	answer = v2.HasNeighbor(v1)
	if answer {
		t.Errorf("ERROR: For '1 neighbor - v2/v1' expected false, got %t", answer)
	}
}

func TestFirstNeighbor(t *testing.T) {
	v1 := NewVertex("ABC", 123)
	answer := v1.FirstNeighbor()
	if answer != nil {
		t.Errorf("ERROR: For 'first neighbor' expected nil, got %s", answer.ID())
	}

	v1 = NewVertex("ABC", 123)
	v2 := NewVertex("DEF", 345)
	v1.AddNeighbor(&v2)
	answer = v1.FirstNeighbor()
	if answer.ID() != v2.ID() {
		t.Errorf("ERROR: For 'first neighbor' expected %s, got %s", v1.ID(), answer.ID())
	}
}

func TestRemoveNeighbor(t *testing.T) {
	v1 := NewVertex("ABC", 123)
	v2 := NewVertex("DEF", 345)
	v1.AddNeighbor(&v2)
	v1.RemoveNeighbor(v2)
	if len(v1.Neighbors()) != 0 {
		t.Errorf("ERROR: For 'remove neighbor' expected len 0, got len %d", len(v1.Neighbors()))
	}
}
