// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"strconv"

	"github.com/howz97/algorithm/search"
)

func main() {
	fmt.Println("BinaryTree:")
	demoBinaryTreeTraversal()
	fmt.Println("\nAVL:")
	demoAvlTraversal()
	fmt.Println("\nRedBlack:")
	demoRedBlackTraversal()
}

func demoBinaryTreeTraversal() {
	bt := search.NewBinTree[int, string]()
	nodes := []int{50, 20, 10, 30, 40, 45, 60}
	for _, v := range nodes {
		bt.Put(v, strconv.Itoa(v))
	}
	search.PrintTree(bt.Root(), func(nd *search.BtNode[int, string]) string { return nd.Value() })

	printNode := func(nd *search.BtNode[int, string]) bool {
		fmt.Print("\t", nd.Value())
		return true
	}
	fmt.Print("Preorder:")
	search.Preorder(bt.Root(), printNode)
	fmt.Println()

	fmt.Print("PreorderRecur:")
	search.PreorderRecur(bt.Root(), printNode)
	fmt.Println()

	fmt.Print("Inorder:")
	search.Inorder(bt.Root(), printNode)
	fmt.Println()

	fmt.Print("Postorder:")
	search.Postorder(bt.Root(), printNode)
	fmt.Println()
}

func demoAvlTraversal() {
	avl := search.NewAVL[int, string]()
	nodes := []int{50, 20, 10, 30, 40, 45, 60}
	for _, v := range nodes {
		avl.Put(v, strconv.Itoa(v))
	}
	search.PrintTree(avl.Root(), func(nd *search.AvlNode[int, string]) string { return nd.Value() })

	printNode := func(nd *search.AvlNode[int, string]) bool {
		fmt.Print("\t", nd.Value())
		return true
	}
	fmt.Print("Preorder:")
	search.Preorder(avl.Root(), printNode)
	fmt.Println()

	fmt.Print("PreorderRecur:")
	search.PreorderRecur(avl.Root(), printNode)
	fmt.Println()

	fmt.Print("Inorder:")
	search.Inorder(avl.Root(), printNode)
	fmt.Println()

	fmt.Print("Postorder:")
	search.Postorder(avl.Root(), printNode)
	fmt.Println()
}

func demoRedBlackTraversal() {
	rbt := search.NewRBTree[int, string]()
	nodes := []int{50, 20, 10, 30, 40, 45, 60}
	for _, v := range nodes {
		rbt.Put(v, strconv.Itoa(v))
	}
	search.PrintTree(rbt.Root(), func(nd *search.RbNode[int, string]) string {
		var c string
		if nd.IsRed() {
			c = "r"
		} else {
			c = "b"
		}
		return fmt.Sprintf("%s(%v)", nd.Value(), c)
	})

	printNode := func(nd *search.RbNode[int, string]) bool {
		fmt.Print("\t", nd.Value())
		return true
	}
	fmt.Print("Preorder:")
	search.Preorder(rbt.Root(), printNode)
	fmt.Println()

	fmt.Print("PreorderRecur:")
	search.PreorderRecur(rbt.Root(), printNode)
	fmt.Println()

	fmt.Print("Inorder:")
	search.Inorder(rbt.Root(), printNode)
	fmt.Println()

	fmt.Print("Postorder:")
	search.Postorder(rbt.Root(), printNode)
	fmt.Println()
}
