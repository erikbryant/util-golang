package main

import (
	"fmt"
	"github.com/erikbryant/util-golang/graphs/adjLists"
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"os"
)

func plotGraph(graph *adjLists.AdjLists) {
	title := fmt.Sprintf("Magnets")
	serial := graph.Serialize(title)
	fmt.Println(serial)
}

func makeGraph(row, col int) *adjLists.AdjLists {
	nodes := map[string]*vertexes.Vertexes{}
	graph := adjLists.NewAL()

	// Create the nodes
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			coord := fmt.Sprintf("%dx%d", r, c)
			node := vertexes.NewVertex(coord, 0)
			nodes[coord] = node
			graph.AddNode(node)
		}
	}

	// Connect the nodes into a grid
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			coord := fmt.Sprintf("%dx%d", r, c)
			node := nodes[coord]
			for _, dir := range [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}} {
				rDir, cDir := dir[0], dir[1]
				coord := fmt.Sprintf("%dx%d", r+rDir, c+cDir)
				node.AddNeighbor(nodes[coord])
			}
		}
	}

	return &graph
}

func main() {
	_, _ = fmt.Fprintf(os.Stderr, "Welcome to magnets research facility #848!\n")

	graph := makeGraph(5, 5)
	fmt.Fprintf(os.Stderr, "Symmetric: %t, Bipartite: %t, Complete: %t, Chromatic#: %d\n", graph.IsSymmetric(), graph.IsBipartite(), graph.IsComplete(), graph.ChromaticNumber())
	plotGraph(graph)
}
