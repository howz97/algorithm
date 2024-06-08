package search

import (
	"cmp"
	"math/rand"
)

// BinaryTree is a simple binary search tree that does not guarantee balance
type BinaryTree[K cmp.Ordered, V any] struct {
	root *BtNode[K, V]
	size uint
}

func NewBinTree[K cmp.Ordered, V any]() *BinaryTree[K, V] {
	return new(BinaryTree[K, V])
}

func (st *BinaryTree[K, V]) Put(key K, val V) {
	exist := false
	st.root, exist = st.root.put(key, val)
	if !exist {
		st.size++
	}
}

func (st *BinaryTree[K, V]) Get(key K) (V, bool) {
	n := st.root.get(key)
	if n == nil {
		var v V
		return v, false
	}
	return n.value, true
}

func (st *BinaryTree[K, V]) GetMin() V {
	return st.root.getMin().value
}

func (st *BinaryTree[K, V]) GetMax() V {
	return st.root.getMax().value
}

func (st *BinaryTree[K, V]) Del(key K) {
	exist := false
	st.root, exist = st.root.del(key)
	if exist {
		st.size--
	}
}

func (st *BinaryTree[K, V]) Size() uint {
	return st.size
}

func (st *BinaryTree[K, V]) Root() *BtNode[K, V] {
	return st.root
}

type BtNode[K cmp.Ordered, V any] struct {
	value V
	key   K
	left  *BtNode[K, V]
	right *BtNode[K, V]
}

func (n *BtNode[K, V]) put(k K, v V) (*BtNode[K, V], bool) {
	if n == nil {
		n = new(BtNode[K, V])
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

func (n *BtNode[K, V]) get(k K) *BtNode[K, V] {
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

func (n *BtNode[K, V]) getMin() *BtNode[K, V] {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *BtNode[K, V]) getMax() *BtNode[K, V] {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *BtNode[K, V]) del(k K) (*BtNode[K, V], bool) {
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
		var replacer *BtNode[K, V]
		if rand.Intn(2) == 0 { // to make it randomly
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

func (n *BtNode[K, V]) Left() BNode {
	return n.left
}

func (n *BtNode[K, V]) Right() BNode {
	return n.right
}

func (n *BtNode[K, V]) IsNil() bool {
	return n == nil
}

func (n *BtNode[K, V]) Key() K {
	return n.key
}

func (n *BtNode[K, V]) Value() V {
	return n.value
}
