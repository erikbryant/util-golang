package main

import (
	"fmt"
	"github.com/erikbryant/util-golang/graphs"
	"os"
)

func plotGraph(graph *graphs.AdjList) {
	title := fmt.Sprintf("Magnets")
	serial := graph.Serialize(title)
	fmt.Println(serial)
}

func makeGraph(row, col int) *graphs.AdjList {
	nodes := map[string]*graphs.Vertex{}
	graph := graphs.NewAL()

	// Create the nodes
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			coord := fmt.Sprintf("%dx%d", r, c)
			node := graphs.NewVertex(coord, 0)
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

	graph := makeGraph(2, 2)
	plotGraph(graph)
}
