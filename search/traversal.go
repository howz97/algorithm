package search

import (
	"fmt"
	"github.com/howz97/algorithm/queue"
	"github.com/waiyva/binary-tree/btprinter"
)

func PreOrder(bt IBinaryTree, order *queue.LinkedQueue) {
	if bt == nil {
		return
	}
	order.PushBack(bt)
	PreOrder(bt.Left(), order)
	PreOrder(bt.Right(), order)
}

func InOrder(bt IBinaryTree, order *queue.LinkedQueue) {
	if bt == nil {
		return
	}
	InOrder(bt.Left(), order)
	order.PushBack(bt)
	InOrder(bt.Right(), order)
}

func SufOrder(bt IBinaryTree, order *queue.LinkedQueue) {
	if bt == nil {
		return
	}
	SufOrder(bt.Left(), order)
	SufOrder(bt.Right(), order)
	order.PushBack(bt)
}

func LevelOrder(bt IBinaryTree) *queue.LinkedQueue {
	if bt == nil {
		return nil
	}
	order := queue.NewLinkedQueue()
	q := queue.NewLinkedQueue()
	q.PushBack(bt)
	for !q.IsEmpty() {
		bt = q.Front().(IBinaryTree)
		order.PushBack(bt)
		if bt.Left() != nil {
			q.PushBack(bt.Left())
		}
		if bt.Right() != nil {
			q.PushBack(bt.Right())
		}
	}
	return order
}

func PrintBinaryTree(bt IBinaryTree) {
	order := LevelOrder(bt)
	var sli []string
	for !order.IsEmpty() {
		sli = append(sli, fmt.Sprint(order.Front()))
	}
	btprinter.PrintTreeLevelOrder(sli)
}
