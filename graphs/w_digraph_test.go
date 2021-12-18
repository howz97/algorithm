package graphs

import (
	"fmt"
	"testing"
)

func TestNewSPS_Dijkstra(t *testing.T) {
	g, err := LoadWDigraph(".\\test_data\\w_digraph.yml")
	if err != nil {
		t.Fatal(err)
	}
	sps, err := g.SearcherDijkstra()
	if err != nil {
		t.Fatal(err)
	}
	CheckSearcher(t, sps, g)
}

func TestNewSPS_Topological(t *testing.T) {
	g, err := LoadWDigraph(".\\test_data\\no_cycle.yml")
	if err != nil {
		t.Fatal(err)
	}
	sps, err := g.SearcherTopological()
	if err != nil {
		t.Fatal(err)
	}
	CheckSearcher(t, sps, g)
}

func TestBellmanFord(t *testing.T) {
	g, err := LoadWDigraph(".\\test_data\\w_digraph.yml")
	if err != nil {
		t.Fatal(err)
	}
	sps, err := g.SearcherBellmanFord()
	if err != nil {
		t.Fatal(err)
	}
	CheckSearcher(t, sps, g)
}

func TestNegativeCycle(t *testing.T) {
	g, err := LoadWDigraph(".\\test_data\\negative_cycle.yml")
	if err != nil {
		t.Fatal(err)
	}
	_, err = g.SearcherBellmanFord()
	if err == nil {
		t.Fatalf("negative cycle exist, error should be received")
	}
	fmt.Println("negative cycle detected:", err)
}

func CheckSearcher(t *testing.T, s *Searcher, dg *WDigraph) {
	for _, spt := range s.spt {
		CheckSPT(t, spt, dg)
	}
}

func CheckSPT(t *testing.T, tree *PathTree, dg *WDigraph) {
	dg.IterateWEdge(func(from int, to int, w float64) bool {
		if tree.distTo[from]+w < tree.distTo[to] {
			t.Errorf("edge %d->%d should belong to SPT: %v + %v < %v", from, to, tree.distTo[from], w, tree.distTo[to])
			return false
		}
		return true
	})
}
