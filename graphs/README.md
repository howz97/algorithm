Detect ring
```go
func ExampleDigraph_FindCycle() {
	// (0)-------->(2)
	//  | ^         ^
	//  |  \        |
	//  |   ------  |
	//  |         \ |
	//  v          \|
	// (1)-------->(3)
	g := NewDigraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(3, 0)
	g.AddEdge(3, 2)
	c := g.FindCycle()
	fmt.Println(c.Error())

	// Output: (distance=3): 0->1, 1->3, 3->0,
}
```

Topological sort

![figure1](https://github.com/howz97/algorithm/blob/master/graphs/testdata/no_cycle.jpg)

```go
func ExampleDigraph_Topological() {
	dg, err := LoadDigraph(`.\testdata\no_cycle.yml`)
	if err != nil {
		panic(err)
	}
	for _, vet := range dg.Topological().Drain() {
		fmt.Printf("%d->", vet)
	}

	// Output: 5->1->3->6->4->7->0->2->
}
```

Bipartite
```go
func ExampleDigraph_IsBipartite() {
	dg, err := LoadDigraph(`.\testdata\no_cycle.yml`)
	if err != nil {
		panic(err)
	}
	fmt.Println(dg.IsBipartite())

	// Output: false
}
```

Reachable
```go
func ExampleReachable() {
	dg, err := LoadDigraph(`.\testdata\no_cycle.yml`)
	if err != nil {
		panic(err)
	}
	reach := dg.Reachable()
	fmt.Println(reach.CanReach(5, 2))
	fmt.Println(reach.CanReach(2, 5))

	// Output:
	// true
	// false
}
```

BFS
```go
func ExampleBFS() {
	dg, err := LoadDigraph(`.\testdata\no_cycle.yml`)
	if err != nil {
		panic(err)
	}
	bfs := dg.BFS(1)
	fmt.Println(bfs.CanReach(5))
	fmt.Println(bfs.ShortestPathTo(2).Str(nil))

	// Output:
	// false
	// (distance=6): 1->3, 3->6, 6->2,
}
```

Strongly connected components
```go
func ExampleSCC() {
	g := NewDigraph(13)
	g.AddEdge(0, 1)
	g.AddEdge(0, 5)
	g.AddEdge(5, 4)
	g.AddEdge(4, 3)
	g.AddEdge(4, 2)
	g.AddEdge(3, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 0)
	g.AddEdge(6, 0)
	g.AddEdge(6, 4)
	g.AddEdge(6, 9)
	g.AddEdge(9, 10)
	g.AddEdge(10, 12)
	g.AddEdge(12, 9)
	g.AddEdge(9, 11)
	g.AddEdge(11, 12)
	g.AddEdge(11, 4)
	g.AddEdge(7, 6)
	g.AddEdge(7, 8)
	g.AddEdge(8, 7)
	g.AddEdge(8, 9)
	scc := g.SCC()
	fmt.Println("amount of strongly connected component:", scc.NumComponents())
	var vertices []int
	scc.IterateComponent(0, func(v int) bool {
		vertices = append(vertices, v)
		return true
	})
	sort.Shell(vertices)
	fmt.Println("vertices strongly connected with 0:", vertices)
	fmt.Println(scc.IsStronglyConn(0, 6))

	// Output:
	// amount of strongly connected component: 5
	// vertices strongly connected with 0: [0 2 3 4 5]
	// false
}
```

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