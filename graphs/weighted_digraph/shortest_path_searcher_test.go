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
	if s0.Size() != s1.Size() {
		return false
	}
	for {
		e0, ok := s0.Pop()
		if !ok {
			break
		}
		eg0 := e0.(*Edge)
		e1, _ := s1.Pop()
		eg1 := e1.(*Edge)
		if eg0.From() != eg1.From() || eg0.To() != eg1.To() {
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
