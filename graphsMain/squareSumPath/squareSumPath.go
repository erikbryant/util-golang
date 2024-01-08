package main

// go fmt ./... && go vet ./... && go test
// time go run squareSumPath.go -cpuprofile cpu.prof | gvpack -u | dot -Tpng > test.png && open test.png
// go tool pprof cpu.prof

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/graphs"
)

var (
	nodes      = []*graphs.Vertex{}
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func squareAddends(n int) []int {
	addends := []int{}

	for i := n - 1; i >= 1; i-- {
		if algebra.IsSquare(n + i) {
			addends = append(addends, i)
		}
	}

	return addends
}

// connect connects the given int to all addends in the graph
func connect(adj *graphs.AdjacencyList, n int, addends []int) {
	// Record this new node
	node := graphs.NewVertex("", n)
	nodes = append(nodes, node)
	adj.AddNode(node)

	// Connect it to all (any?) numbers it sums with
	for _, addend := range addends {
		adj.AddEdge(nodes[n], nodes[addend])
	}
}

func doit() {
	adj := graphs.NewAL()

	lower := 1
	upper := 66

	// Our numbers start at 1, put a placeholder in 0
	nodes = append(nodes, nil)

	for i := lower; i <= upper; i++ {
		addends := squareAddends(i)
		connect(&adj, i, addends)
		fmt.Fprintf(os.Stderr, "Added %6d: ", i)
		paths := adj.HamiltonianPaths()
		if paths == nil {
			fmt.Fprintf(os.Stderr, "\n")
		} else {
			fmt.Fprintf(os.Stderr, "Found a path!!\n")
		}
		// for _, path := range paths {
		// 	for _, node := range path {
		// 		fmt.Fprintf(os.Stderr, " %s->", node.Name())
		// 	}
		// 	fmt.Fprintf(os.Stderr, "\n\n")
		// }
	}

	// title := fmt.Sprintf("%d..%d Connected: %t Hamiltonian Path: %t", lower, upper, adj.Connected(), adj.HamiltonianPaths() != nil)
	// serial := adj.Serialize(title)
	// fmt.Println(serial)
}

func main() {
	fmt.Fprintf(os.Stderr, "\nWelcome to square sum path!\n\n")

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	doit()
}
