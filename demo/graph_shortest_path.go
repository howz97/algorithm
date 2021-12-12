package main

import (
	"fmt"
	"github.com/howz97/algorithm/graphs"
)

func main() {
	g, err := graphs.LoadWDigraph("../graphs/test_data/w_digraph.yml", false)
	if err != nil {
		panic(err)
	}
	searcher, err := g.SearcherDijkstra()
	//searcher, err := g.SearcherTopological()
	//searcher, err := g.SearcherBellmanFord()
	if err != nil {
		panic(err)
	}
	fmt.Println(searcher.Path(1, 2).String())
}
