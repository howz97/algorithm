package binarytree

import (
	"github.com/howz97/algorithm/search"
	"time"
)

// BinaryTree is a simple binary search tree that does not guarantee balance
type BinaryTree struct {
	*node
	size uint
}

func New() *BinaryTree {
	return new(BinaryTree)
}

func (st *BinaryTree) Put(key search.Cmp, val search.T) {
	exist := false
	st.node, exist = st.put(key, val)
	if !exist {
		st.size++
	}
}

func (st *BinaryTree) Get(key search.Cmp) search.T {
	n := st.get(key)
	if n == nil {
		return nil
	}
	return n.value
}

func (st *BinaryTree) GetMin() search.T {
	return st.getMin().value
}

func (st *BinaryTree) GetMax() search.T {
	return st.getMax().value
}

func (st *BinaryTree) Del(key search.Cmp) {
	exist := false
	st.node, exist = st.del(key)
	if exist {
		st.size--
	}
}

func (st *BinaryTree) Size() uint {
	return st.size
}

func (st *BinaryTree) Clean() {
	st.node = nil
	st.size = 0
}

type node struct {
	value search.T
	key   search.Cmp
	left  *node
	right *node
}

func (n *node) put(k search.Cmp, v search.T) (*node, bool) {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
		return n, false
	}
	var exist bool
	switch k.Cmp(n.key) {
	case search.Less:
		n.left, exist = n.left.put(k, v)
	case search.More:
		n.right, exist = n.right.put(k, v)
	default:
		n.value = v
		exist = true
	}
	return n, exist
}

func (n *node) get(k search.Cmp) *node {
	if n == nil {
		return nil
	}
	switch k.Cmp(n.key) {
	case search.Less:
		n = n.left.get(k)
	case search.More:
		n = n.right.get(k)
	}
	return n
}

func (n *node) getMin() *node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *node) getMax() *node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *node) del(k search.Cmp) (*node, bool) {
	if n == nil {
		return nil, false
	}
	var exist bool
	switch k.Cmp(n.key) {
	case search.Less:
		n.left, exist = n.left.del(k)
	case search.More:
		n.right, exist = n.right.del(k)
	default:
		if n.left == nil {
			return n.right, true
		}
		if n.right == nil {
			return n.left, true
		}
		exist = true
		var replacer *node
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

func (n *node) Left() search.ITraversal {
	return n.left
}

func (n *node) Right() search.ITraversal {
	return n.right
}

func (n *node) IsNil() bool {
	return n == nil
}

func (n *node) Key() search.Cmp {
	return n.key
}

func (n *node) Val() search.T {
	return n.value
}
