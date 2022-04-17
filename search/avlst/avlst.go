package avlst

import (
	"fmt"

	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/util"
	"golang.org/x/exp/constraints"
)

// AVL is a strictly balanced binary search tree
type AVL[Ord constraints.Ordered, T any] struct {
	root *Node[Ord, T]
	size uint
}

func New[Ord constraints.Ordered, T any]() *AVL[Ord, T] {
	return new(AVL[Ord, T])
}

func (avl *AVL[Ord, T]) Put(key Ord, val T) {
	exist := false
	avl.root, exist = avl.root.put(key, val)
	if !exist {
		avl.size++
	}
}

func (avl *AVL[Ord, T]) Get(key Ord) (T, bool) {
	n := avl.root.get(key)
	if n == nil {
		var v T
		return v, false
	}
	return n.value, true
}

func (avl *AVL[Ord, T]) GetMin() T {
	return avl.root.getMin().value
}

func (avl *AVL[Ord, T]) GetMax() T {
	return avl.root.getMax().value
}

func (avl *AVL[Ord, T]) Del(key Ord) {
	exist := false
	avl.root, exist = avl.root.del(key)
	if exist {
		avl.size--
	}
}

func (avl *AVL[Ord, T]) Size() uint {
	return avl.size
}

func (avl *AVL[Ord, T]) Print() {
	search.PrintBinaryTree(avl.root)
}

func (avl *AVL[Ord, T]) PreOrder(fn func(*Node[Ord, T]) bool) {
	search.PreOrder(avl.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

func (avl *AVL[Ord, T]) InOrder(fn func(*Node[Ord, T]) bool) {
	search.InOrder(avl.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

func (avl *AVL[Ord, T]) SufOrder(fn func(*Node[Ord, T]) bool) {
	search.SufOrder(avl.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

func (avl *AVL[Ord, T]) LevelOrder(fn func(*Node[Ord, T]) bool) {
	search.LevelOrder(avl.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

func (avl *AVL[Ord, T]) ReverseOrder(fn func(*Node[Ord, T]) bool) {
	search.ReverseOrder(avl.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

type Node[Ord constraints.Ordered, T any] struct {
	key   Ord
	value T
	h     int8
	left  *Node[Ord, T]
	right *Node[Ord, T]
}

func (n *Node[Ord, T]) put(k Ord, v T) (*Node[Ord, T], bool) {
	if n == nil {
		n = new(Node[Ord, T])
		n.key = k
		n.value = v
		return n, false
	}
	var exist bool
	if k < n.key {
		n.left, exist = n.left.put(k, v)
		if n.diff() > 1 {
			n = rotation(n)
		}
		n.updateHeight()
	} else if k > n.key {
		n.right, exist = n.right.put(k, v)
		if n.diff() < -1 {
			n = rotation(n)
		}
		n.updateHeight()
	} else {
		n.value = v
		exist = true
	}
	return n, exist
}

func (n *Node[Ord, T]) diff() int8 {
	return n.left.height() - n.right.height()
}

func (n *Node[Ord, T]) updateHeight() {
	n.h = util.Max(n.left.height(), n.right.height()) + 1
}

func rotation[Ord constraints.Ordered, T any](r *Node[Ord, T]) *Node[Ord, T] {
	diff := r.diff()
	switch true {
	case diff == 2:
		left := r.left
		if left.left.height() < left.right.height() {
			r.left = rightRotation(left)
		}
		r = leftRotation(r)
	case diff == -2:
		right := r.right
		if right.left.height() > right.right.height() {
			r.right = leftRotation(right)
		}
		r = rightRotation(r)
	default:
		panic(fmt.Sprintf("|diff| == |%v - %v| != 2", r.left.height(), r.right.height()))
	}
	return r
}

func leftRotation[Ord constraints.Ordered, T any](n *Node[Ord, T]) *Node[Ord, T] {
	replacer := n.left
	n.left = replacer.right
	replacer.right = n
	n.updateHeight()
	replacer.updateHeight()
	return replacer
}

func rightRotation[Ord constraints.Ordered, T any](n *Node[Ord, T]) *Node[Ord, T] {
	replacer := n.right
	n.right = replacer.left
	replacer.left = n
	n.updateHeight()
	replacer.updateHeight()
	return replacer
}

func (n *Node[Ord, T]) height() int8 {
	if n == nil {
		return -1
	}
	return n.h
}

func (n *Node[Ord, T]) get(k Ord) *Node[Ord, T] {
	if n == nil {
		return nil
	}
	if k < n.key {
		return n.left.get(k)
	} else if k > n.key {
		return n.right.get(k)
	} else {
		return n
	}
}

func (n *Node[Ord, T]) getMin() *Node[Ord, T] {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *Node[Ord, T]) getMax() *Node[Ord, T] {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *Node[Ord, T]) del(k Ord) (*Node[Ord, T], bool) {
	if n == nil {
		return nil, false
	}
	var exist bool
	if k < n.key {
		n.left, exist = n.left.del(k)
	} else if k > n.key {
		n.right, exist = n.right.del(k)
	} else {
		if n.left == nil {
			return n.right, true
		}
		if n.right == nil {
			return n.left, true
		}
		exist = true
		replacer := n.right.getMin()
		replacer.right, _ = n.right.del(replacer.key)
		replacer.left = n.left
		n = replacer
	}
	n.updateHeight()
	if !n.isBalance() {
		n = rotation(n)
	}
	return n, exist
}

func (n *Node[Ord, T]) Left() search.ITraversal {
	return n.left
}

func (n *Node[Ord, T]) Right() search.ITraversal {
	return n.right
}

func (n *Node[Ord, T]) IsNil() bool {
	return n == nil
}

func (n *Node[Ord, T]) Key() Ord {
	return n.key
}

func (n *Node[Ord, T]) String() string {
	return fmt.Sprint(n.value)
}

func (n *Node[Ord, T]) isBalance() bool {
	diff := n.diff()
	return diff > -2 && diff < 2
}
