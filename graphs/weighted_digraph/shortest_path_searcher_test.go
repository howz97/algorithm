package weighted_digraph

import (
	"fmt"
	"github.com/howz97/algorithm/stack"
	"testing"
)

func TestEWD_Integer(t *testing.T) {
	g, err := ImportEWD("./tinyEWD.txt")
	if err != nil {
		t.Fatal(err)
	}
	var (
		spsDijkstra *ShortestPathSearcher
		spsTop      *ShortestPathSearcher
		spsBF       *ShortestPathSearcher
	)
	spsDijkstra, err = g.GenSearcherDijkstra()
	if err != nil {
		t.Fatal(err)
	}
	spsTop, err = g.GenSearcherTopological()
	if err != nil {
		t.Fatal(err)
	}
	spsBF, err = g.GenSearcherBellmanFord()
	if err != nil {
		t.Fatal(err)
	}
	for src := range g {
		for dst := range g {
			p0 := spsDijkstra.Path(src, dst)
			p1 := spsTop.Path(src, dst)
			if !isPathEqual(p0, p1) {
				t.Fatal("path not equal")
			}

			p0 = spsDijkstra.Path(src, dst)
			p1 = spsBF.Path(src, dst)
			if !isPathEqual(p0, p1) {
				t.Fatal("path not equal")
			}
		}
	}
}

func isPathEqual(s0, s1 *stack.Stack) bool {
	for !s0.IsEmpty() {
		e0 := s0.Pop().(*Edge)
		e1 := s1.Pop().(*Edge)
		if e0.From() != e1.From() || e0.To() != e1.To() {
			return false
		}
	}
	return true
}

func TestNewSPS_Dijkstra(t *testing.T) {
	g, err := ImportEWD("./tinyEWD.txt")
	if err != nil {
		t.Fatal(err)
	}
	sps, _ := g.GenSearcherDijkstra()
	for src := range g {
		for dst := range g {
			sps.PrintPath(src, dst)
		}
	}
}

func TestNewSPS_Topological(t *testing.T) {
	g, err := ImportEWD("./tinyEWDAG.txt")
	if err != nil {
		t.Fatal(err)
	}
	sps, _ := g.GenSearcherTopological()
	for src := range g {
		for dst := range g {
			sps.PrintPath(src, dst)
		}
	}
}

func TestNewSPS_BellmanFord(t *testing.T) {
	g, err := ImportEWD("./tinyEWDnc.txt")
	if err != nil {
		t.Fatal(err)
	}
	sps, err := g.GenSearcherBellmanFord()
	if err != nil {
		fmt.Println(err)
		return
	}
	for src := range g {
		for dst := range g {
			sps.PrintPath(src, dst)
		}
	}
}
