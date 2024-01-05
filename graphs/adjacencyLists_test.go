package graphs

import "testing"

func TestMinimalVertexCover(t *testing.T) {
	a := NewAL()
	mvc := a.MinimalVertexCover()
	if len(mvc) != 0 {
		t.Errorf("ERROR: For 'empty' expected 0 len, got %d len", len(mvc))
	}

	a = NewAL()
	n1 := NewVertex("A", 4)
	a.AddNode(&n1)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 0 {
		t.Errorf("ERROR: For 'one node' expected 0 len, got %d len", len(mvc))
	}

	a = NewAL()
	n1 = NewVertex("B", 4)
	n2 := NewVertex("C", 4)
	a.AddNode(&n1)
	a.AddNode(&n2)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 0 {
		t.Errorf("ERROR: For 'two nodes' expected 0 len, got %d len", len(mvc))
	}

	a = NewAL()
	n1 = NewVertex("B", 4)
	n2 = NewVertex("C", 4)
	a.AddEdge(&n1, &n2)
	mvc = a.MinimalVertexCover()
	if len(mvc) != 1 {
		t.Errorf("ERROR: For 'one edge' expected 1 len, got %d len", len(mvc))
	}

	// a = NewAL()
	// n1 = NewVertex("D", 4)
	// n2 = NewVertex("E", 4)
	// n3 := NewVertex("F", 4)
	// a.AddEdge(&n1, &n2)
	// a.AddEdge(&n2, &n3)
	// mvc = a.MinimalVertexCover()
	// if len(mvc) != 1 {
	// 	t.Errorf("ERROR: For 'one vertex, two whiskers' expected 1 len, got %d len", len(mvc))
	// }
}

// func TestE(t *testing.T) {
// 	testCases := []struct {
// 		n        int
// 		expected int64
// 	}{
// 		{1, 2},
// 		{2, 1},
// 		{3, 2},
// 		{4, 1},
// 		{5, 1},
// 		{6, 4},
// 		{7, 1},
// 		{8, 1},
// 		{9, 6},
// 	}

// 	for _, testCase := range testCases {
// 		answer := E(testCase.n)
// 		if answer != testCase.expected {
// 			t.Errorf("ERROR: For %d expected %d, got %d", testCase.n, testCase.expected, answer)
// 		}
// 	}
// }
