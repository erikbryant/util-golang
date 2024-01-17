package main

import (
	"fmt"
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"os"

	"github.com/erikbryant/util-golang/graphs"
)

func populate() graphs.AdjList {
	n1 := vertexes.NewVertex("", 4)
	n2 := vertexes.NewVertex("", -2)
	n3 := vertexes.NewVertex("", -1)
	n4 := vertexes.NewVertex("", 0)

	a := graphs.NewAL()

	a.AddEdge(n1, n2)
	a.AddEdge(n2, n3)
	a.AddEdge(n3, n4)
	a.AddEdge(n4, n2)

	return a
}

func solvable(a graphs.AdjList) bool {
	return a.ValueSum() >= a.Genus()
}

func title(a graphs.AdjList) string {
	str := fmt.Sprintf("value: $%d genus: %d", a.ValueSum(), a.Genus())
	if solvable(a) {
		str += " -> solvable!"
	} else {
		str += " -> probably not solvable :("
	}
	return str
}

// solved returns true if there are no more negative values
func solved(a graphs.AdjList) bool {
	low := a.ValueLowest()
	return low.Value() >= 0
}

func printGraph(a graphs.AdjList) {
	s := a.Serialize(title(a))
	fmt.Println(s)
}

func requestAid(node *vertexes.Vertex) {
	for _, neighbor := range node.Neighbors() {
		node.Increment()
		neighbor.Decrement()
	}
}

func solve(a graphs.AdjList) {
	if !solvable(a) {
		return
	}

	for !solved(a) {
		lowest := a.ValueLowest()
		requestAid(lowest)
		lowest.SetName("*")
		printGraph(a)
		lowest.SetName("")
	}
}

func main() {
	fmt.Fprintf(os.Stderr, "Welcome to the dollar game!\n")

	a := populate()
	printGraph(a)
	solve(a)
}
