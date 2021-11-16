package leetcode

import (
	"fmt"
)

func preOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.Key, " ")
	preOrder(n.leftSon)
	preOrder(n.rightSon)
}

func inOrder(n *Node) {
	if n == nil {
		return
	}
	inOrder(n.leftSon)
	fmt.Print(n.Key, " ")
	inOrder(n.rightSon)
}

func sufOrder(n *Node) {
	if n == nil {
		return
	}
	sufOrder(n.leftSon)
	sufOrder(n.rightSon)
	fmt.Print(n.Key, " ")
}

func levelOrder(n *Node) {
	if n == nil {
		return
	}
	q := new(queue)
	q.pushBack(n)
	for !q.empty() {
		n = q.front()
		fmt.Print(n.Key, " ")
		if n.leftSon != nil {
			q.pushBack(n.leftSon)
		}
		if n.rightSon != nil {
			q.pushBack(n.rightSon)
		}
	}
}
