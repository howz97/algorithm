Minimum spanning tree
```go
func Example() {
	g, err := LoadWGraph("testdata/mst.yml")
	if err != nil {
		panic(err)
	}
	mst := g.Prim()
	//mst := g.LazyPrim()
	//mst := g.Kruskal()
	fmt.Println(mst.String())

	// possible output:
	// 0 : 2 7
	// 1 : 7
	// 2 : 0 3 6
	// 3 : 2
	// 4 : 5
	// 5 : 7 4
	// 6 : 2
	// 7 : 0 1 5
}
```

Shortest path
```go
func ExampleWDigraph() {
	g, _ := LoadWDigraph("testdata/no_cycle.yml")
	searcher, _ := g.SearcherDijkstra()
	// searcher, _ = g.SearcherTopological()
	// searcher, _ = g.SearcherBellmanFord()
	fmt.Println(searcher.GetPath(1, 2).Str(nil))

	// Output:
	// (distance=1.02): 1->3, 3->7, 7->2,
}
```

Symbol graph
```go
func ExampleSymbol() {
	g, _ := LoadGraph("./testdata/symbol_graph.yml")
	bfs := g.BFS(g.VetOf("姜文"))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁朝伟")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("宋慧乔")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("郎雄")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("周星驰")).Str(g.Symbol))
	fmt.Println(bfs.ShortestPathTo(g.VetOf("梁家辉")).Str(g.Symbol))

	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->刘嘉玲, 刘嘉玲->《阿飞正传》, 《阿飞正传》->张学友, 张学友->《东邪西毒》, 《东邪西毒》->梁朝伟,
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《卧虎藏龙》, 《卧虎藏龙》->章子怡, 章子怡->《一代宗师》, 《一代宗师》->宋慧乔,
	//(distance=8): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《卧虎藏龙》, 《卧虎藏龙》->郎雄,
	//(distance=12): 姜文->《让子弹飞》, 《让子弹飞》->刘嘉玲, 刘嘉玲->《阿飞正传》, 《阿飞正传》->张曼玉, 张曼玉->《家有喜事》, 《家有喜事》->周星驰,
	//(distance=8): 姜文->《让子弹飞》, 《让子弹飞》->周润发, 周润发->《赌神2》, 《赌神2》->梁家辉,
}
```