package triangles

const (
	// MaxTriangle is the largest triangular number we will consider
	MaxTriangle = 100000
)

var (
	// Triangles stores whether the position in the slice is triangular or not
	Triangles [MaxTriangle + 1]bool
)

// Triangle returns true if the number is triangular
func Triangle(number int) bool {
	return Triangles[number]
}

func triangles() {
	for i := 0; i < len(Triangles); i++ {
		Triangles[i] = false
	}
	i := 1
	for {
		t := (i * (i + 1)) >> 1
		if t >= len(Triangles) {
			break
		}
		Triangles[t] = true
		i++
	}
}

// Init populates the Triangles cache
func Init() {
	triangles()
}
