package main

import (
	"fmt"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
)

func main() {
	avl := avltree.New()
	for i := 0; i < 20; i++ {
		avl.Put(search.Integer(i), i)
	}
	fmt.Printf("Size=%d Get(5)=%v \n", avl.Size(), avl.Get(search.Integer(5)))
	search.PrintBinaryTree(avl)

	for i := 0; i < 10; i++ {
		avl.Del(search.Integer(i))
	}
	fmt.Printf("Size=%d Get(5)=%v \n", avl.Size(), avl.Get(search.Integer(5)))
	search.PrintBinaryTree(avl)

	fmt.Println("traversal in order:")
	search.InOrder(avl, func(t search.ITraversal) bool {
		fmt.Printf("%v,", t.Val())
		return true
	})
}
