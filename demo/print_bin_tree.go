package main

import (
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
)

var pairs = []string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
}

func main() {
	tree := avltree.New[float64, string]()
	for k, v := range pairs {
		tree.Put(float64(k)/100, v)
	}
	search.PrintBinaryTree(tree)
}
