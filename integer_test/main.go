package main

import (
	wDigraph "github.com/howz97/algorithm/graphs/weighted_digraph"
)

func main() {
	g, err := wDigraph.ImportEWD("/Users/zhanghao1/code/algorithm/integer_test/tinyEWD.txt")
	if err != nil {
		panic(err)
	}
	pathSearcher, err := g.GenSearcherDijkstra()
	/*
		pathSearcher, err := g.GenSearcherTopological()
		pathSearcher, err := g.GenSearcherBellmanFord()
	*/
	if err != nil {
		panic(err)
	}
	pathSearcher.PrintPath(0,7)
}
