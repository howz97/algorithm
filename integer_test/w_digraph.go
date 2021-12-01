package main

import "github.com/howz97/algorithm/graphs/wdigraph"

func main() {
	g, err := wdigraph.ImportEWD("../graphs/wdigraph/tinyEWD.txt")
	if err != nil {
		panic(err)
	}
	pathSearcher, err := g.GenSearcherDijkstra()
	//pathSearcher, err := g.GenSearcherTopological()
	//pathSearcher, err := g.GenSearcherBellmanFord()
	if err != nil {
		panic(err)
	}
	pathSearcher.PrintPath(0, 7)
}
