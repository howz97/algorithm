package search

import (
	"github.com/howz97/algorithm/queue"
	"github.com/waiyva/binary-tree/btprinter"
)

func PreOrder(bt ITraversal, order *queue.LinkedQueue) {
	if bt == nil {
		return
	}
	order.PushBack(bt)
	PreOrder(bt.Left(), order)
	PreOrder(bt.Right(), order)
}

func InOrder(bt ITraversal, order *queue.LinkedQueue) {
	if bt == nil {
		return
	}
	InOrder(bt.Left(), order)
	order.PushBack(bt)
	InOrder(bt.Right(), order)
}

func SufOrder(bt ITraversal, order *queue.LinkedQueue) {
	if bt == nil {
		return
	}
	SufOrder(bt.Left(), order)
	SufOrder(bt.Right(), order)
	order.PushBack(bt)
}

func LevelOrder(bt ITraversal) *queue.LinkedQueue {
	if bt.IsNil() {
		return nil
	}
	order := queue.NewLinkedQueue()
	q := queue.NewLinkedQueue()
	q.PushBack(bt)
	for !q.IsEmpty() {
		bt = q.Front().(ITraversal)
		order.PushBack(bt)
		if bt.IsNil() {
			continue
		}
		q.PushBack(bt.Left())
		q.PushBack(bt.Right())
	}
	return order
}

func PrintBinaryTree(bt ITraversal) {
	order := LevelOrder(bt)
	var sli []string
	for !order.IsEmpty() {
		sli = append(sli, order.Front().(ITraversal).String())
	}
	btprinter.PrintTreeLevelOrder(sli)
}
