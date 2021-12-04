package integer

import (
	"github.com/howz97/algorithm/graphs"
	"github.com/howz97/algorithm/graphs/digraph"
	"github.com/howz97/algorithm/graphs/graph"
	"github.com/howz97/algorithm/sort"
	"github.com/howz97/algorithm/util"
	stdsort "sort"
	"testing"
)

func TestDFS_Graph(t *testing.T) {
	g := graph.New(9)
	var err error
	err = g.AddEdge(0, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(0, 3)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(1, 3)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(3, 6)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(4, 7)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(4, 8)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(2, 5)
	if err != nil {
		t.Fatal(err)
	}

	dfsResults := [][]int{
		0: {0, 1, 3, 6},
		1: {0, 1, 3, 6},
		2: {2, 5},
		3: {0, 1, 3, 6},
		4: {4, 7, 8},
		5: {2, 5},
		6: {0, 1, 3, 6},
		7: {4, 7, 8},
		8: {4, 7, 8},
	}
	checkDFSResults(t, g, dfsResults)
}

var digraph1 = [][]int{
	0: {3},
	1: {5},
	2: {1},
	3: {6, 7},
	4: {},
	5: {2, 8},
	6: {7},
	7: {},
	8: {7},
}

func TestDFS_Digraph(t *testing.T) {
	dg, err := digraph.NewBy2DSli(digraph1)
	if err != nil {
		t.Fatal(err)
	}
	dfsResults := [][]int{
		0: {0, 3, 6, 7},
		1: {1, 2, 5, 7, 8},
		2: {1, 2, 5, 7, 8},
		3: {3, 6, 7},
		4: {4},
		5: {1, 2, 5, 7, 8},
		6: {6, 7},
		7: {7},
		8: {7, 8},
	}
	checkDFSResults(t, dg, dfsResults)
}

func checkDFSResults(t *testing.T, g graphs.IGraph, dfsResults [][]int) {
	for src := range dfsResults {
		reach := graphs.ReachableSlice(g, src)
		sort.QuickSort(stdsort.IntSlice(reach))
		if !util.SliceEqual(reach, dfsResults[src]) {
			t.Errorf("v %d reach %v not equal %v", src, reach, dfsResults[src])
		}
	}
}

func TestRevDFS(t *testing.T) {
	g, err := digraph.NewBy2DSli(digraph1)
	if err != nil {
		t.Fatal(err)
	}
	order := graphs.VetRDFSOrder(g, 0).ToSlice()
	correct := []int{0, 3, 6, 7}
	if !util.SliceEqual(order, correct) {
		t.Errorf("rev dfs order %v not equal %v", order, correct)
	}
}
