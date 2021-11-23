带权重的有向图
```go
package main

import wDigraph "github.com/howz97/algorithm/graphs/weighted_digraph"

func main() {
	g, err := wDigraph.ImportEWD("../graphs/weighted_digraph/tinyEWD.txt")
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

/*
output：
PATH: 0->2->7 (distance 0.6000000000000001)
*/
```
