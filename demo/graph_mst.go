package main

import (
	"fmt"

	"github.com/howz97/algorithm/graphs"
)

func main() {
	g, err := graphs.LoadWGraph("../graphs/testdata/mst.yml")
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
