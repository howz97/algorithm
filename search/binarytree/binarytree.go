package binarytree

import (
	"fmt"
	"time"

	"github.com/howz97/algorithm/search"
	"golang.org/x/exp/constraints"
)

// BinaryTree is a simple binary search tree that does not guarantee balance
type BinaryTree[Ord constraints.Ordered, T any] struct {
	*node[Ord, T]
	size uint
}

func New[Ord constraints.Ordered, T any]() *BinaryTree[Ord, T] {
	return new(BinaryTree[Ord, T])
}

func (st *BinaryTree[Ord, T]) Put(key Ord, val T) {
	exist := false
	st.node, exist = st.put(key, val)
	if !exist {
		st.size++
	}
}

func (st *BinaryTree[Ord, T]) Get(key Ord) (T, bool) {
	n := st.get(key)
	if n == nil {
		var v T
		return v, false
	}
	return n.value, true
}

func (st *BinaryTree[Ord, T]) GetMin() T {
	return st.getMin().value
}

func (st *BinaryTree[Ord, T]) GetMax() T {
	return st.getMax().value
}

func (st *BinaryTree[Ord, T]) Del(key Ord) {
	exist := false
	st.node, exist = st.del(key)
	if exist {
		st.size--
	}
}

func (st *BinaryTree[Ord, T]) Size() uint {
	return st.size
}

func (st *BinaryTree[Ord, T]) Clean() {
	st.node = nil
	st.size = 0
}

type node[Ord constraints.Ordered, T any] struct {
	value T
	key   Ord
	left  *node[Ord, T]
	right *node[Ord, T]
}

func (n *node[Ord, T]) put(k Ord, v T) (*node[Ord, T], bool) {
	if n == nil {
		n = new(node[Ord, T])
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

func (n *node[Ord, T]) get(k Ord) *node[Ord, T] {
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

func (n *node[Ord, T]) getMin() *node[Ord, T] {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *node[Ord, T]) getMax() *node[Ord, T] {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *node[Ord, T]) del(k Ord) (*node[Ord, T], bool) {
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
		var replacer *node[Ord, T]
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

func (n *node[Ord, T]) Left() search.ITraversal {
	return n.left
}

func (n *node[Ord, T]) Right() search.ITraversal {
	return n.right
}

func (n *node[Ord, T]) IsNil() bool {
	return n == nil
}

func (n *node[Ord, T]) Key() Ord {
	return n.key
}

func (n *node[Ord, T]) String() string {
	return fmt.Sprint(n.value)
}
