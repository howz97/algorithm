package main

import (
	"fmt"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
	"github.com/howz97/algorithm/util"
)

func main() {
	avl := avltree.New()
	for i := 0; i < 20; i++ {
		avl.Put(util.Int(i), i)
	}
	fmt.Printf("Size=%d Get(5)=%v \n", avl.Size(), avl.Get(util.Int(5)))
	search.PrintBinaryTree(avl)

	for i := 0; i < 10; i++ {
		avl.Del(util.Int(i))
	}
	fmt.Printf("Size=%d Get(5)=%v \n", avl.Size(), avl.Get(util.Int(5)))
	search.PrintBinaryTree(avl)

	fmt.Println("traversal in order:")
	search.InOrder(avl, func(t search.ITraversal) bool {
		fmt.Printf("%v,", t.Val())
		return true
	})
}
