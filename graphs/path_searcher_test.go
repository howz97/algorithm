package graphs

import (
	"fmt"
	"github.com/howz97/algorithm/stack"
	"testing"
)

func TestEWD_Integer(t *testing.T) {
	g, err := LoadWDigraph("no_cycle.yml", false)
	if err != nil {
		t.Fatal(err)
	}
	var (
		spsDijkstra *PathSearcher
		spsTop      *PathSearcher
		spsBF       *PathSearcher
	)
	spsDijkstra, err = g.SearcherDijkstra()
	if err != nil {
		t.Fatal(err)
	}
	spsTop, err = g.SearcherTopological()
	if err != nil {
		t.Fatal(err)
	}
	spsBF, err = g.SearcherBellmanFord()
	if err != nil {
		t.Fatal(err)
	}
	num := int(g.NumVertical())
	for src := 0; src < num; src++ {
		for dst := 0; dst < num; dst++ {
			p0 := spsDijkstra.Path(src, dst)
			p1 := spsTop.Path(src, dst)
			if !isPathEqual(p0.stk, p1.stk) {
				t.Errorf("path(%d->%d) not equal: \np0=%s, \np1=%s \n", src, dst, p0, p1)
			}

			p0 = spsDijkstra.Path(src, dst)
			p1 = spsBF.Path(src, dst)
			if !isPathEqual(p0.stk, p1.stk) {
				t.Errorf("path(%d->%d) not equal: \np0=%s, \np1=%s \n", src, dst, p0, p1)
			}
		}
	}
}

func isPathEqual(s0, s1 *stack.Stack) bool {
	if stack.SizeOf(s0) != stack.SizeOf(s1) {
		return false
	}
	if s0 == nil {
		return true
	}
	for s0.Size() > 0 {
		e0 := s0.Pop()
		e1 := s1.Pop()
		if e0 != e1 {
			return false
		}
	}
	return true
}

func TestNewSPS_Dijkstra(t *testing.T) {
	g, err := LoadWDigraph("w_digraph.yml", false)
	if err != nil {
		t.Fatal(err)
	}

	sps, _ := g.SearcherDijkstra()

	num := int(g.NumVertical())
	for src := 0; src < num; src++ {
		for dst := 0; dst < num; dst++ {
			fmt.Println(sps.Path(src, dst).String())
		}
	}
}

func TestNewSPS_Topological(t *testing.T) {
	g, err := LoadWDigraph("no_cycle.yml", false)
	if err != nil {
		t.Fatal(err)
	}
	sps, _ := g.SearcherTopological()
	num := int(g.NumVertical())
	for src := 0; src < num; src++ {
		for dst := 0; dst < num; dst++ {
			fmt.Println(sps.Path(src, dst).String())
		}
	}
}

func TestPathTree_Top(t *testing.T) {
	g, err := LoadWDigraph("no_cycle.yml", false)
	if err != nil {
		t.Fatal(err)
	}
	tree, err := g.NewShortestPathTree(1, Topological)
	if err != nil {
		t.Fatal(err)
	}
	g.IterateVetDFS(1, func(dst int) bool {
		if dst == 1 {
			return true
		}
		path := tree.PathTo(dst)
		if path == nil {
			t.Errorf("failed to find shortest path from %d->%d\n", 1, dst)
		} else {
			t.Logf("path %d->%d: %s\n", 1, dst, path)
		}
		return true
	})
}

func TestNewSPS_BellmanFord(t *testing.T) {
	g, err := LoadWDigraph("negative_cycle.yml", false)
	if err != nil {
		t.Fatal(err)
	}
	_, err = g.SearcherBellmanFord()
	if err == nil {
		t.Fatal("negative cycle exist, error should be received")
	}
	fmt.Println(err)
}
