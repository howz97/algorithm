package leetcode

import (
	"fmt"
	"math"
)

/*
	No.23
	https://leetcode-cn.com/problems/merge-k-sorted-lists/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

var (
	MaxNode = &ListNode{
		Val:  math.MaxInt64,
		Next: nil,
	}
)

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	if k == 1 {
		return lists[0]
	}
	loserTree := make([]int, k)
	winner := initLoserTree(loserTree, 1, lists)
	loserTree[0] = winner
	var result *ListNode
	var tail *ListNode
	for lists[loserTree[0]].Val != math.MaxInt64 {
		if result == nil {
			result = lists[loserTree[0]]
			tail = result
		} else {
			tail.Next = lists[loserTree[0]]
			tail = tail.Next
		}
		lists[loserTree[0]] = lists[loserTree[0]].Next
		adjust(loserTree, loserTree[0], lists)
	}
	return result
}

func initLoserTree(loserTree []int, root int, lists []*ListNode) int {
	if root >= len(loserTree) {
		return root - len(loserTree)
	}
	leftWin := initLoserTree(loserTree, 2*root, lists)
	rightWin := initLoserTree(loserTree, 2*root+1, lists)
	if lists[leftWin] == nil {
		lists[leftWin] = MaxNode
	}
	if lists[rightWin] == nil {
		lists[rightWin] = MaxNode
	}
	if lists[leftWin].Val < lists[rightWin].Val {
		loserTree[root] = rightWin
		return leftWin
	} else {
		loserTree[root] = leftWin
		return rightWin
	}
}

func adjust(loserTree []int, leafIdx int, lists []*ListNode) {
	parent := (leafIdx + len(loserTree)) >> 1
	winner := 0
	if lists[leafIdx] == nil {
		lists[leafIdx] = MaxNode
	}
	if lists[leafIdx].Val < lists[loserTree[parent]].Val {
		winner = leafIdx
	} else {
		winner = loserTree[parent]
		loserTree[parent] = leafIdx
	}
	for parent >>= 1; parent > 0; parent >>= 1 {
		if lists[winner].Val > lists[loserTree[parent]].Val {
			winner, loserTree[parent] = loserTree[parent], winner
		}
	}
	loserTree[0] = winner
}

func printList(list *ListNode) {
	for list != nil {
		fmt.Print(list.Val, " -> ")
		list = list.Next
	}
	fmt.Println("end of list")
}

func printAllLists(lists []*ListNode) {
	fmt.Println("----------------")
	for i := range lists {
		fmt.Printf("list %v : ", i)
		printList(lists[i])
	}
	fmt.Println("----------------")
}
