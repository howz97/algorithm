package search

import (
	"fmt"
	"github.com/howz97/algorithm/queue"
	"github.com/waiyva/binary-tree/btprinter"
)

func PreOrder(bt ITraversal, fn func(k Cmp, v T) bool) bool {
	if !fn(bt.Key(), bt.Val()) {
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

func InOrder(bt ITraversal, fn func(k Cmp, v T) bool) bool {
	if !bt.Left().IsNil() {
		if !InOrder(bt.Left(), fn) {
			return false
		}
	}
	if !fn(bt.Key(), bt.Val()) {
		return false
	}
	if !bt.Right().IsNil() {
		if !InOrder(bt.Right(), fn) {
			return false
		}
	}
	return true
}

func SufOrder(bt ITraversal, fn func(k Cmp, v T) bool) bool {
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
	return fn(bt.Key(), bt.Val())
}

func LevelOrder(bt ITraversal, fn func(k Cmp, v T) bool) {
	if bt.IsNil() {
		return
	}
	q := queue.NewLinkedQueue()
	q.PushBack(bt)
	for !q.IsEmpty() {
		bt = q.Front().(ITraversal)
		if !fn(bt.Key(), bt.Val()) {
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

func ReverseOrder(bt ITraversal, fn func(k Cmp, v T) bool) bool {
	if !bt.Right().IsNil() {
		if !ReverseOrder(bt.Right(), fn) {
			return false
		}
	}
	if !fn(bt.Key(), bt.Val()) {
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
	q := queue.NewLinkedQueue()
	q.PushBack(bt)
	for !q.IsEmpty() {
		bt = q.Front().(ITraversal)
		if bt.IsNil() {
			sli = append(sli, "#")
			continue
		}
		sli = append(sli, fmt.Sprint(bt.Val()))
		q.PushBack(bt.Left())
		q.PushBack(bt.Right())
	}
	btprinter.PrintTree(sli)
}
