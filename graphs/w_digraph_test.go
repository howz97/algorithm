package graphs

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestDijkstra_Simple(t *testing.T) {
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

func TestTopological_Simple(t *testing.T) {
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

func TestBellmanFord_Simple(t *testing.T) {
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

const (
	vertLowerLimit = 100
	vertRange      = 900
)

func TestTopological(t *testing.T) {
	LoopTestSearcher(t, 2, 0.5, func(wd *WDigraph) (*Searcher, error) {
		return wd.SearcherTopological()
	})
}

func TestDijkstra(t *testing.T) {
	LoopTestSearcher(t, 10, 0.00001, func(wd *WDigraph) (*Searcher, error) {
		return wd.SearcherDijkstra()
	})
}

func TestBellmanFord(t *testing.T) {
	LoopTestSearcher(t, 10, 0.0005, func(wd *WDigraph) (*Searcher, error) {
		return wd.SearcherBellmanFord()
	})
}

func TestSearcher(t *testing.T) {
	LoopTestSearcher(t, 10, 0.0005, func(wd *WDigraph) (*Searcher, error) {
		return wd.Searcher()
	})
}

func LoopTestSearcher(t *testing.T, edgeLimit int, negativeEdge float64,
	fn func(*WDigraph) (*Searcher, error)) {
	for i := 0; i < 10; i++ {
		wd := RandWDigraph(edgeLimit, negativeEdge)
		sps, err := fn(wd)
		if err != nil {
			t.Log(err)
			continue
		}
		CheckSearcher(t, sps, wd)
	}
}

func RandWDigraph(edgeLimit int, negativeEdge float64) (wd *WDigraph) {
	wd = NewWDigraph(uint(vertLowerLimit + rand.Intn(vertRange)))
	nv := int(wd.NumVert())
	for from := 0; from < nv; from++ {
		ne := rand.Intn(edgeLimit)
		for j := 0; j < ne; j++ {
			to := rand.Intn(nv)
			if from == to {
				continue
			}
			w := (rand.Float64() - negativeEdge) * 10000
			wd.AddEdge(from, to, w)
		}
	}
	return
}

func ExampleWDigraph() {
	g, _ := LoadWDigraph("test_data/no_cycle.yml")
	searcher, _ := g.SearcherDijkstra()
	// searcher, _ = g.SearcherTopological()
	// searcher, _ = g.SearcherBellmanFord()
	fmt.Println(searcher.GetPath(1, 2).Str(nil))

	// Output:
	// (distance=1.02): 1->3, 3->7, 7->2,
}
