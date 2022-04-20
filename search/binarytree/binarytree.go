package binarytree

import (
	"fmt"
	"time"

	"github.com/howz97/algorithm/search"
	"golang.org/x/exp/constraints"
)

// BinaryTree is a simple binary search tree that does not guarantee balance
type BinaryTree[Ord constraints.Ordered, T any] struct {
	root *Node[Ord, T]
	size uint
}

func New[Ord constraints.Ordered, T any]() *BinaryTree[Ord, T] {
	return new(BinaryTree[Ord, T])
}

func (st *BinaryTree[Ord, T]) Put(key Ord, val T) {
	exist := false
	st.root, exist = st.root.put(key, val)
	if !exist {
		st.size++
	}
}

func (st *BinaryTree[Ord, T]) Get(key Ord) (T, bool) {
	n := st.root.get(key)
	if n == nil {
		var v T
		return v, false
	}
	return n.value, true
}

func (st *BinaryTree[Ord, T]) GetMin() T {
	return st.root.getMin().value
}

func (st *BinaryTree[Ord, T]) GetMax() T {
	return st.root.getMax().value
}

func (st *BinaryTree[Ord, T]) Del(key Ord) {
	exist := false
	st.root, exist = st.root.del(key)
	if exist {
		st.size--
	}
}

func (st *BinaryTree[Ord, T]) Size() uint {
	return st.size
}

func (st *BinaryTree[Ord, T]) PreOrder(fn func(*Node[Ord, T]) bool) {
	search.PreOrderIter(st.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

func (st *BinaryTree[Ord, T]) InOrder(fn func(*Node[Ord, T]) bool) {
	search.InOrder(st.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

func (st *BinaryTree[Ord, T]) SufOrder(fn func(*Node[Ord, T]) bool) {
	search.SufOrder(st.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

func (st *BinaryTree[Ord, T]) LevelOrder(fn func(*Node[Ord, T]) bool) {
	search.LevelOrder(st.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

func (st *BinaryTree[Ord, T]) ReverseOrder(fn func(*Node[Ord, T]) bool) {
	search.ReverseOrder(st.root, func(trv search.ITraversal) bool {
		return fn(trv.(*Node[Ord, T]))
	})
}

type Node[Ord constraints.Ordered, T any] struct {
	value T
	key   Ord
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
	} else if k > n.key {
		n.right, exist = n.right.put(k, v)
	} else {
		n.value = v
		exist = true
	}
	return n, exist
}

func (n *Node[Ord, T]) get(k Ord) *Node[Ord, T] {
	if n == nil {
		return nil
	}
	if k < n.key {
		n = n.left.get(k)
	} else if k > n.key {
		n = n.right.get(k)
	}
	return n
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
		var replacer *Node[Ord, T]
		if time.Now().UnixNano()&1 == 1 { // to make it randomly
			replacer = n.left.getMax()
			n.left, _ = n.left.del(replacer.key)
		} else {
			replacer = n.right.getMin()
			n.right, _ = n.right.del(replacer.key)
		}
		replacer.left = n.left
		replacer.right = n.right
		n = replacer
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
