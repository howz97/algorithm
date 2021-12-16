最小生成树
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

最短路径
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

符号图
```go
func TestSymbol(t *testing.T) {
	g, err := LoadGraph("./test_data/symbol_graph.yml")
	if err != nil {
		t.Fatal(err)
	}
	bfs := g.BFS(g.VetOf("姜文"))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁朝伟")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("宋慧乔")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("郎雄")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("周星驰")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁家辉")).Str(g.Symbol))

	//Output:
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->刘嘉玲, 刘嘉玲->《阿飞正传》, 《阿飞正传》->张学友, 张学友->《东邪西毒》, 《东邪西毒》->梁朝伟,
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《卧虎藏龙》, 《卧虎藏龙》->章子怡, 章子怡->《一代宗师》, 《一代宗师》->宋慧乔,
	//(distance=8): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《卧虎藏龙》, 《卧虎藏龙》->郎雄,
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->刘嘉玲, 刘嘉玲->《阿飞正传》, 《阿飞正传》->张曼玉, 张曼玉->《家有喜事》, 《家有喜事》->周星驰,
	//(distance=8): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《赌神2》, 《赌神2》->梁家辉,
}
```