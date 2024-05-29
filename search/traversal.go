package search

import (
	"fmt"

	"github.com/howz97/algorithm/basic/queue"
	"github.com/howz97/algorithm/basic/stack"
	"github.com/waiyva/binary-tree/btprinter"
)

type ITraversal interface {
	fmt.Stringer
	IsNil() bool
	Left() ITraversal
	Right() ITraversal
}

func PreOrder(bt ITraversal, fn func(ITraversal) bool) bool {
	if !fn(bt) {
		return false
	}
	if !bt.Left().IsNil() {
		if !PreOrder(bt.Left(), fn) {
			return false
		}
	}
	if !bt.Right().IsNil() {
		if !PreOrder(bt.Right(), fn) {
			return false
		}
	}
	return true
}

// PreOrderIter traverse nodes in pre-order non-recursively
func PreOrderIter(bt ITraversal, fn func(ITraversal) bool) {
	right := stack.New[ITraversal](0)
	right.Push(bt)
	for right.Size() > 0 {
		n := right.Pop()
		for !n.IsNil() {
			if !fn(n) {
				return
			}
			if !n.Right().IsNil() {
				right.Push(n.Right())
			}
			n = n.Left()
		}
	}
}

func InOrder(bt ITraversal, fn func(ITraversal) bool) bool {
	if !bt.Left().IsNil() {
		if !InOrder(bt.Left(), fn) {
			return false
		}
	}
	if !fn(bt) {
		return false
	}
	if !bt.Right().IsNil() {
		if !InOrder(bt.Right(), fn) {
			return false
		}
	}
	return true
}

func SufOrder(bt ITraversal, fn func(ITraversal) bool) bool {
	if !bt.Left().IsNil() {
		if !SufOrder(bt.Left(), fn) {
			return false
		}
	}
	if !bt.Right().IsNil() {
		if !SufOrder(bt.Right(), fn) {
			return false
		}
	}
	return fn(bt)
}

func LevelOrder(bt ITraversal, fn func(ITraversal) bool) {
	if bt.IsNil() {
		return
	}
	q := queue.NewLinkQ[ITraversal]()
	q.PushBack(bt)
	for q.Size() > 0 {
		bt = q.PopFront()
		if !fn(bt) {
			break
		}
		if !bt.Left().IsNil() {
			q.PushBack(bt.Left())
		}
		if !bt.Right().IsNil() {
			q.PushBack(bt.Right())
		}
	}
}

func ReverseOrder(bt ITraversal, fn func(ITraversal) bool) bool {
	if !bt.Right().IsNil() {
		if !ReverseOrder(bt.Right(), fn) {
			return false
		}
	}
	if !fn(bt) {
		return false
	}
	if !bt.Left().IsNil() {
		if !ReverseOrder(bt.Left(), fn) {
			return false
		}
	}
	return true
}

func PrintBinaryTree(bt ITraversal) {
	var sli []string
	q := queue.NewLinkQ[ITraversal]()
	q.PushBack(bt)
	for q.Size() > 0 {
		bt = q.PopFront()
		if bt.IsNil() {
			sli = append(sli, "#")
			continue
		}
		sli = append(sli, bt.String())
		q.PushBack(bt.Left())
		q.PushBack(bt.Right())
	}
	btprinter.PrintTree(sli)
}
