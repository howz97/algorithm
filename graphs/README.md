##有向图
####BFS
todo
####DFS
####可达性
####强联通分量(kosaraju)
####二分图
####拓扑排序
####有向环检测

##无向图
####子图
####无向环检测

##加权无向图-最小生成树
####Lazy-Prim
####Prim
####Kruskal

##加权有向图-最短路径
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/graphs"
)

func main() {
	g, err := graphs.LoadWDigraph("../graphs/test_data/w_digraph.yml")
	if err != nil {
		panic(err)
	}
	searcher, err := g.SearcherDijkstra()
	//searcher, err := g.SearcherTopological()
	//searcher, err := g.SearcherBellmanFord()
	if err != nil {
		panic(err)
	}
	fmt.Println(searcher.GetPath(1, 2).String())

	/*
		(distance=0): 1->3->6->2
	*/
}
```
####Dijkstra

####Topological

####Bellman-Ford