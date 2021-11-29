package graphs

import (
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

func TestDFS_Digraph(t *testing.T) {
	data := [][]int{
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
	dg, err := digraph.NewBy2DSli(data)
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

func checkDFSResults(t *testing.T, g IGraph, dfsResults [][]int) {
	for src := range dfsResults {
		reach := DFSReachable(g, src)
		sort.QuickSort(stdsort.IntSlice(reach))
		if !util.SliceEqual(reach, dfsResults[src]) {
			t.Fatalf("v %d reach %v not equal %v", src, reach, dfsResults[src])
		}
	}
}
