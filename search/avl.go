package search

import (
	"cmp"
	"fmt"

	"github.com/howz97/algorithm/util"
)

// AVL is a strictly balanced binary search tree
type AVL[K cmp.Ordered, V any] struct {
	root *AvlNode[K, V]
	size uint
}

func NewAVL[K cmp.Ordered, V any]() *AVL[K, V] {
	return new(AVL[K, V])
}

func (avl *AVL[K, V]) Put(key K, val V) {
	exist := false
	avl.root, exist = avl.root.put(key, val)
	if !exist {
		avl.size++
	}
}

func (avl *AVL[K, V]) Get(key K) (V, bool) {
	n := avl.root.get(key)
	if n == nil {
		var v V
		return v, false
	}
	return n.value, true
}

func (avl *AVL[K, V]) GetMin() V {
	return avl.root.getMin().value
}

func (avl *AVL[K, V]) GetMax() V {
	return avl.root.getMax().value
}

func (avl *AVL[K, V]) Del(key K) {
	exist := false
	avl.root, exist = avl.root.del(key)
	if exist {
		avl.size--
	}
}

func (avl *AVL[K, V]) Size() uint {
	return avl.size
}

func (avl *AVL[K, V]) Root() *AvlNode[K, V] {
	return avl.root
}

type AvlNode[K cmp.Ordered, V any] struct {
	key   K
	value V
	h     int8
	left  *AvlNode[K, V]
	right *AvlNode[K, V]
}

func (n *AvlNode[K, V]) put(k K, v V) (*AvlNode[K, V], bool) {
	if n == nil {
		n = new(AvlNode[K, V])
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

func (n *AvlNode[K, V]) diff() int8 {
	return n.left.height() - n.right.height()
}

func (n *AvlNode[K, V]) updateHeight() {
	n.h = util.Max(n.left.height(), n.right.height()) + 1
}

func rotation[K cmp.Ordered, V any](r *AvlNode[K, V]) *AvlNode[K, V] {
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

func leftRotation[K cmp.Ordered, V any](n *AvlNode[K, V]) *AvlNode[K, V] {
	replacer := n.left
	n.left = replacer.right
	replacer.right = n
	n.updateHeight()
	replacer.updateHeight()
	return replacer
}

func rightRotation[K cmp.Ordered, V any](n *AvlNode[K, V]) *AvlNode[K, V] {
	replacer := n.right
	n.right = replacer.left
	replacer.left = n
	n.updateHeight()
	replacer.updateHeight()
	return replacer
}

func (n *AvlNode[K, V]) height() int8 {
	if n == nil {
		return -1
	}
	return n.h
}

func (n *AvlNode[K, V]) get(k K) *AvlNode[K, V] {
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

func (n *AvlNode[K, V]) getMin() *AvlNode[K, V] {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *AvlNode[K, V]) getMax() *AvlNode[K, V] {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *AvlNode[K, V]) del(k K) (*AvlNode[K, V], bool) {
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

func (n *AvlNode[K, V]) isBalance() bool {
	diff := n.diff()
	return diff > -2 && diff < 2
}

func (n *AvlNode[K, V]) Left() BNode {
	return n.left
}

func (n *AvlNode[K, V]) Right() BNode {
	return n.right
}

func (n *AvlNode[K, V]) IsNil() bool {
	return n == nil
}

func (n *AvlNode[K, V]) Key() K {
	return n.key
}

func (n *AvlNode[K, V]) Value() V {
	return n.value
}
