package avltree

import (
	"fmt"
	. "github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/util"
)

type AVL struct {
	root *node
}

func New() *AVL {
	return new(AVL)
}

// Insert insert a k-v pair into this dictionary
// if value is nil, is is a vitual deletion operation
func (avl *AVL) Insert(key Cmp, value T) {
	avl.root = avl.root.insert(key, value)
}

func (avl *AVL) Find(key Cmp) T {
	n := avl.root.find(key)
	if n == nil {
		return nil
	}
	if n.value == nil {
		avl.Delete(n.key) // 删除value为nil的节点
		return nil
	}
	return n.value
}

func (avl *AVL) FindMin() T {
	return avl.root.findMin().value
}

func (avl *AVL) FindMax() T {
	return avl.root.findMax().value
}

func (avl *AVL) Delete(key Cmp) {
	avl.root = avl.root.delete(key)
}

func (avl *AVL) IsEmpty() bool {
	return avl.root == nil
}

type node struct {
	key   Cmp
	value T
	h     int8
	left  *node
	right *node
}

func (n *node) insert(k Cmp, v T) *node {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
		return n
	}
	switch k.Cmp(n.key) {
	case Less:
		n.left = n.left.insert(k, v)
		if n.diff() > 1 {
			n = rotation(n)
		}
		n.updateHeight()
	case More:
		n.right = n.right.insert(k, v)
		if n.diff() < -1 {
			n = rotation(n)
		}
		n.updateHeight()
	default:
		n.value = v
	}
	return n
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
		if r.left.left.height() > r.left.right.height() {
			r = leftSingleRotation(r)
		} else {
			r = leftDoubleRotation(r)
		}
	case diff == -2:
		if r.right.right.height() > r.right.left.height() {
			r = rightSingleRotation(r)
		} else {
			r = rightDoubleRotation(r)
		}
	default:
		panic(fmt.Sprintf("|diff| == |%v - %v| != 2", r.left.height(), r.right.height()))
	}
	return r
}

func leftSingleRotation(k2 *node) *node {
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	k2.updateHeight()
	k1.updateHeight()
	return k1
}

func rightSingleRotation(k2 *node) *node {
	k1 := k2.right
	k2.right = k1.left
	k1.left = k2
	k2.updateHeight()
	k1.updateHeight()
	return k1
}

func leftDoubleRotation(k3 *node) *node {
	k3.left = rightSingleRotation(k3.left)
	return leftSingleRotation(k3)
}

func rightDoubleRotation(k3 *node) *node {
	k3.right = leftSingleRotation(k3.right)
	return rightSingleRotation(k3)
}

func (n *node) height() int8 {
	if n == nil {
		return -1
	}
	return n.h
}

func (n *node) find(k Cmp) *node {
	if n == nil {
		return nil
	}
	switch k.Cmp(n.key) {
	case Less:
		return n.left.find(k)
	case More:
		return n.right.find(k)
	default:
		return n
	}
}

func (n *node) findMin() *node {
	if n == nil {
		return nil
	}
	for n.left != nil {
		n = n.left
	}
	return n
}

func (n *node) findMax() *node {
	if n == nil {
		return nil
	}
	for n.right != nil {
		n = n.right
	}
	return n
}

func (n *node) delete(k Cmp) *node {
	if n == nil {
		return nil
	}
	switch k.Cmp(n.key) {
	case Less:
		n.left = n.left.delete(k)
		n.updateHeight()
		if n.diff() < -1 {
			n = rotation(n)
		}
	case More:
		n.right = n.right.delete(k)
		n.updateHeight()
		if n.diff() > 1 {
			n = rotation(n)
		}
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		deleted := n
		n = n.right.findMin()
		n.right = deleted.right.delete(n.key)
		n.left = deleted.left
		n.updateHeight()
		if n.diff() > 1 {
			n = rotation(n)
		}
	}
	return n
}
