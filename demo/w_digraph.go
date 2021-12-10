package main

import (
	"github.com/howz97/algorithm/graphs"
)

func main() {
	g, err := graphs.LoadWDigraph("../graphs/wdigraph/w_digraph.yml")
	if err != nil {
		panic(err)
	}
	pathSearcher, err := g.SearcherDijkstra()
	//pathSearcher, err := g.SearcherTopological()
	//pathSearcher, err := g.SearcherBellmanFord()
	if err != nil {
		panic(err)
	}
	pathSearcher.PrintPath(0, 7)
}
