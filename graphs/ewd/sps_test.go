package ewd

import (
	"fmt"
	"testing"
)

func TestNewSPS_Dijkstra(t *testing.T) {
	g, err := ImportEWD("./tinyEWD.txt")
	if err != nil {
		t.Fatal(err)
	}
	sps, _ := NewSPS(g, Dijkstra)
	for i := range g {
		printPath(sps, 0, i)
	}
}

func TestNewSPS_Topological(t *testing.T) {
	g, err := ImportEWD("./tinyEWDAG.txt")
	if err != nil {
		t.Fatal(err)
	}
	sps, _ := NewSPS(g, Topological)
	for i := range g {
		printPath(sps, 1, i)
	}
}

func TestNewSPS_BellmanFord(t *testing.T) {
	g, err := ImportEWD("./tinyEWD.txt")
	if err != nil {
		t.Fatal(err)
	}
	sps, _ := NewSPS(g, BellmanFord)
	for src := range g {
		for dst := range g {
			printPath(sps, src, dst)
		}
	}
}

func printPath(sps *ShortestPathSearcher, src, dst int) {
	p := sps.Path(src, dst)
	fmt.Print("PATH: ", src)
	for !p.IsEmpty() {
		e := p.Pop().(*Edge)
		fmt.Print("->", e.to)
	}
	fmt.Printf(" (distance %v)\n", sps.Distance(src, dst))
}
