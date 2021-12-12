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
	fmt.Println(mst.String())
}
