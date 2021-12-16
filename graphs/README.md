```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/graphs"
)

func main() {
	g, err := graphs.LoadWGraph("../graphs/test_data/mst.yml")
	if err != nil {
		panic(err)
	}
	mst := g.Prim()
	//mst := g.LazyPrim()
	//mst := g.Kruskal()
	fmt.Println(mst.String())

	// Output:
	// 0 : 7 2
	// 1 : 7
	// 2 : 3 0 6
	// 3 : 2
	// 4 : 5
	// 5 : 7 4
	// 6 : 2
	// 7 : 0 1 5
}
```

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
	fmt.Println(searcher.GetPath(1, 2).Str(nil))

	// Output:
	// (distance=1.21): 1->3, 3->6, 6->2,
}
```