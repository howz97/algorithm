package main

import (
	"fmt"
	"github.com/howz97/algorithm/binarytree/avltree"
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
	tree := avltree.New()
	for k, v := range pairs {
		tree.Insert(k, v)
	}
	for k, v := range pairs {
		if tree.Find(k) != v {
			panic("key-value not match")
		}
	}
	fmt.Printf("min:%s, max:%s\n", tree.FindMin(), tree.FindMax())
}
