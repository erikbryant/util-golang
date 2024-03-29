package main

// go fmt ./... && go vet ./... && go test
// time go run squareSumPath.go -cpuprofile cpu.prof | gvpack -u | dot -Tpng > test.png && open test.png
// go tool pprof cpu.prof
//
// The path count (ignoring reverses) should match this sequence
// https://oeis.org/A071983

import (
	"flag"
	"fmt"
	"github.com/erikbryant/util-golang/graphs/adjLists"
	"github.com/erikbryant/util-golang/graphs/vertexes"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/erikbryant/util-golang/algebra"
	"github.com/erikbryant/util-golang/system"
)

var (
	nodes      = []*vertexes.Vertexes{}
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
func connect(adj *adjLists.AdjLists, n int, addends []int) int {
	// Record this new node
	node := vertexes.NewVertex("", n)
	nodes = append(nodes, node)
	adj.AddNode(node)

	// Connect it to all (any?) numbers it sums with
	for _, addend := range addends {
		adj.AddEdge(nodes[n], nodes[addend])
	}

	return len(addends)
}

func listPaths(paths [][]*vertexes.Vertexes) {
	for _, path := range paths {
		fmt.Fprintf(os.Stderr, "  ")
		for _, node := range path {
			fmt.Fprintf(os.Stderr, "%d ", node.Value())
		}
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func plotGraph(lower, upper int, adj adjLists.AdjLists, paths [][]*vertexes.Vertexes) {
	title := fmt.Sprintf("%d..%d Connected: %t #Paths: %d", lower, upper, adj.Connected(), len(paths))
	serial := adj.Serialize(title)
	fmt.Println(serial)
}

func findPaths() {
	adj := adjLists.NewAL()
	var paths [][]*vertexes.Vertexes

	// Our numbers start at 1, put a placeholder in 0
	nodes = append(nodes, nil)

	lower := 1
	upper := 300

	for i := lower; i <= upper; i++ {
		addends := squareAddends(i)
		connections := connect(&adj, i, addends)
		symmetric := adj.IsSymmetric()
		bipartite := adj.IsBipartite()
		complete := adj.IsComplete()
		chromatic := adj.ChromaticNumber()
		paths = adj.HamiltonianPaths(2, true, false)
		fmt.Fprintf(os.Stderr, "Added: %6d   Connections: %3d   Paths: %6d   Symmetric: %5t  Bipartite: %5t  Complete: %5t  Chromatic#: %d  GoRoutines: %3d\n", i, connections, len(paths), symmetric, bipartite, complete, chromatic, runtime.NumGoroutine())
		// listPaths(paths)
	}

	// plotGraph(lower, upper, adj, paths)
}

func main() {
	fmt.Fprintf(os.Stderr, "\nWelcome to square sum path!\n\n")

	pid := system.InstallDebug()
	fmt.Fprintf(os.Stderr, "Debugging signal handler installd: $ kill -SIGUSR1 %d\n", pid)

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	findPaths()
}
