package main

import (
	"fmt"

	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avlst"
)

func avltree() {
	avl := avlst.New[int, int]()
	for i := 0; i < 20; i++ {
		avl.Put(i, i)
	}
	v, ok := avl.Get(5)
	fmt.Printf("Size=%d Get(5)=(%v,%v) \n", avl.Size(), v, ok)
	search.PrintBinaryTree(avl)

	for i := 0; i < 10; i++ {
		avl.Del(i)
	}
	v, ok = avl.Get(5)
	fmt.Printf("Size=%d Get(5)=(%v,%v) \n", avl.Size(), v, ok)
	search.PrintBinaryTree(avl)

	fmt.Println("traversal in order:")
	search.InOrder(avl, func(t search.ITraversal) bool {
		fmt.Printf("%v,", t.String())
		return true
	})

	// Size=20 Get(5)=(5,true)
	//            7
	//           / \
	//          /   \
	//         /     \
	//        /       \
	//       /         \
	//      3          15
	//     / \         / \
	//    /   \       /   \
	//   1     5     11   17
	//  / \   / \   / \   / \
	// 0   2 4   6 /   \ 16 18
	//            9    13     \
	//           / \   / \    19
	//          8  10 12 14
	// Size=10 Get(5)=(0,false)
	//     15
	//     / \
	//    /   \
	//   11   17
	//  / \   / \
	// 10 13 16 18
	//    / \     \
	//   12 14    19
	// traversal in order:
	// 10,11,12,13,14,15,16,17,18,19,
}
