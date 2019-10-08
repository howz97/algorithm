package ewd

import (
	"fmt"
	"testing"
)

func TestNewSPS(t *testing.T) {
	g, err := ImportEWD("./tinyEWD.txt")
	if err != nil {
		t.Fatal(err)
	}
	sps := NewSPS(g, Dijkstra)
	for i := range g {
		printPath(sps, 0, i)
	}
}

func printPath(sps *ShortestPathSearcher, src, dst int) {
	p := sps.Path(src, dst)
	fmt.Print("PATH: ")
	for !p.IsEmpty() {
		e := p.Pop().(*Edge)
		fmt.Print("->", e.to)
	}
	fmt.Printf(" (distance %v)\n", sps.Distance(src, dst))
}
