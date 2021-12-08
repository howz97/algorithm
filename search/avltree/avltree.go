package avltree

import (
	"fmt"
	. "github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/util"
)

// AVL is a strictly balanced binary search tree
type AVL struct {
	*node
	size uint
}

func New() *AVL {
	return new(AVL)
}

func (avl *AVL) Put(key util.Comparable, val util.T) {
	exist := false
	avl.node, exist = avl.put(key, val)
	if !exist {
		avl.size++
	}
}

func (avl *AVL) Get(key util.Comparable) util.T {
	n := avl.get(key)
	if n == nil {
		return nil
	}
	return n.value
}

func (avl *AVL) GetMin() util.T {
	return avl.getMin().value
}

func (avl *AVL) GetMax() util.T {
	return avl.getMax().value
}

func (avl *AVL) Del(key util.Comparable) {
	exist := false
	avl.node, exist = avl.del(key)
	if exist {
		avl.size--
	}
}

func (avl *AVL) Size() uint {
	return avl.size
}

func (avl *AVL) Clean() {
	avl.node = nil
	avl.size = 0
}

type node struct {
	key   util.Comparable
	value util.T
	h     int8
	left  *node
	right *node
}

func (n *node) put(k util.Comparable, v util.T) (*node, bool) {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
		return n, false
	}
	var exist bool
	switch k.Cmp(n.key) {
	case util.Less:
		n.left, exist = n.left.put(k, v)
		if n.diff() > 1 {
			n = rotation(n)
		}
		n.updateHeight()
	case util.More:
		n.right, exist = n.right.put(k, v)
		if n.diff() < -1 {
			n = rotation(n)
		}
		n.updateHeight()
	default:
		n.value = v
		exist = true
	}
	return n, exist
}

func (n *node) diff() int8 {
	return n.left.height() - n.right.height()
}

func (n *node) updateHeight() {
	n.h = util.MaxInt8(n.left.height(), n.right.height()) + 1
}

func rotation(r *node) *node {
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

func leftRotation(n *node) *node {
	replacer := n.left
	n.left = replacer.right
	replacer.right = n
	n.updateHeight()
	replacer.updateHeight()
	return replacer
}

func rightRotation(n *node) *node {
	replacer := n.right
	n.right = replacer.left
	replacer.left = n
	n.updateHeight()
	replacer.updateHeight()
	return replacer
}

func (n *node) height() int8 {
	if n == nil {
		return -1
	}
	return n.h
}

func (n *node) get(k util.Comparable) *node {
	if n == nil {
		return nil
	}
	switch k.Cmp(n.key) {
	case util.Less:
		return n.left.get(k)
	case util.More:
		return n.right.get(k)
	default:
		return n
	}
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

func (n *node) del(k util.Comparable) (*node, bool) {
	if n == nil {
		return nil, false
	}
	var exist bool
	switch k.Cmp(n.key) {
	case util.Less:
		n.left, exist = n.left.del(k)
	case util.More:
		n.right, exist = n.right.del(k)
	default:
		if n.left == nil {
			return n.right, true
		}
		if n.right == nil {
			return n.left, true
		}
		exist = true
		replacer := n.right.getMin()
		n.right, _ = n.right.del(replacer.key)
		replacer.left = n.left
		replacer.right = n.right
		n = replacer
	}
	n.updateHeight()
	if !n.isBalance() {
		n = rotation(n)
	}
	return n, exist
}

func (n *node) Left() ITraversal {
	return n.left
}

func (n *node) Right() ITraversal {
	return n.right
}

func (n *node) IsNil() bool {
	return n == nil
}

func (n *node) Key() util.Comparable {
	return n.key
}

func (n *node) Val() util.T {
	return n.value
}

func (n *node) isBalance() bool {
	diff := n.diff()
	return diff > -2 && diff < 2
}
