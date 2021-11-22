package avltree

import (
	"fmt"
	. "github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/util"
)

type AVL struct {
	*node
	size uint
}

func New() *AVL {
	return new(AVL)
}

func (avl *AVL) Put(key Cmp, val T) {
	var newNode bool
	avl.node, newNode = avl.insert(key, val)
	if newNode {
		avl.size++
	}
}

func (avl *AVL) Get(key Cmp) T {
	n := avl.find(key)
	if n == nil {
		return nil
	}
	return n.value
}

func (avl *AVL) GetMin() T {
	return avl.findMin().value
}

func (avl *AVL) GetMax() T {
	return avl.findMax().value
}

func (avl *AVL) Del(key Cmp) {
	avl.node = avl.delete(key)
	avl.size--
}

func (avl *AVL) Size() uint {
	return avl.size
}

func (avl *AVL) Clean() {
	avl.node = nil
}

type node struct {
	key   Cmp
	value T
	h     int8
	left  *node
	right *node
}

func (n *node) insert(k Cmp, v T) (*node, bool) {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
		return n, true
	}
	var newNode bool
	switch k.Cmp(n.key) {
	case Less:
		n.left, newNode = n.left.insert(k, v)
		if n.diff() > 1 {
			n = rotation(n)
		}
		n.updateHeight()
	case More:
		n.right, newNode = n.right.insert(k, v)
		if n.diff() < -1 {
			n = rotation(n)
		}
		n.updateHeight()
	default:
		n.value = v
	}
	return n, newNode
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

func leftRotation(k2 *node) *node {
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	k2.updateHeight()
	k1.updateHeight()
	return k1
}

func rightRotation(k2 *node) *node {
	k1 := k2.right
	k2.right = k1.left
	k1.left = k2
	k2.updateHeight()
	k1.updateHeight()
	return k1
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
	case More:
		n.right = n.right.delete(k)
	default:
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		candi := n.right.findMin()
		n.right = n.right.delete(candi.key)
		candi.left = n.left
		candi.right = n.right
		n = candi
	}
	n.updateHeight()
	if !n.isBalance() {
		n = rotation(n)
	}
	return n
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

func (n *node) Key() Cmp {
	return n.key
}

func (n *node) Val() T {
	return n.value
}

func (n *node) isBalance() bool {
	diff := n.diff()
	return diff > -2 && diff < 2
}
