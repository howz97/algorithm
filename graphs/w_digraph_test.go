package graphs

import (
	"fmt"
	"github.com/howz97/algorithm/stack"
	"testing"
)

func TestEWD_Integer(t *testing.T) {
	g, err := LoadWDigraph("no_cycle.yml")
	if err != nil {
		t.Fatal(err)
	}
	var (
		spsDijkstra *Searcher
		spsTop      *Searcher
		spsBF       *Searcher
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
	num := int(g.NumVert())
	for src := 0; src < num; src++ {
		for dst := 0; dst < num; dst++ {
			p0 := spsDijkstra.GetPath(src, dst)
			p1 := spsTop.GetPath(src, dst)
			if !isPathEqual(&p0.stk.Stack, &p1.stk.Stack) {
				t.Errorf("path(%d->%d) not equal: \np0=%s, \np1=%s \n", src, dst, p0, p1)
			}

			p0 = spsDijkstra.GetPath(src, dst)
			p1 = spsBF.GetPath(src, dst)
			if !isPathEqual(&p0.stk.Stack, &p1.stk.Stack) {
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
	g, err := LoadWDigraph("w_digraph.yml")
	if err != nil {
		t.Fatal(err)
	}

	sps, _ := g.SearcherDijkstra()

	num := int(g.NumVert())
	for src := 0; src < num; src++ {
		for dst := 0; dst < num; dst++ {
			fmt.Println(sps.GetPath(src, dst).Str(nil))
		}
	}
}

func TestNewSPS_Topological(t *testing.T) {
	g, err := LoadWDigraph("no_cycle.yml")
	if err != nil {
		t.Fatal(err)
	}
	sps, _ := g.SearcherTopological()
	num := int(g.NumVert())
	for src := 0; src < num; src++ {
		for dst := 0; dst < num; dst++ {
			fmt.Println(sps.GetPath(src, dst).Str(nil))
		}
	}
}

func TestPathTree_Top(t *testing.T) {
	g, err := LoadWDigraph("no_cycle.yml")
	if err != nil {
		t.Fatal(err)
	}
	tree, err := g.ShortestPathTree(1, Topological)
	if err != nil {
		t.Fatal(err)
	}
	g.IterateVertDFS(1, func(dst int) bool {
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
	g, err := LoadWDigraph("negative_cycle.yml")
	if err != nil {
		t.Fatal(err)
	}
	_, err = g.SearcherBellmanFord()
	if err == nil {
		t.Fatal("negative cycle exist, error should be received")
	}
	fmt.Println(err)
}
