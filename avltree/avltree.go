package avltree

/*
NOTE:
	1. 使用uint类型的前提：该值永远为正，且不会做减法运算！
	   比如这里node里面的length，虽然一定为正，但是会做减法运算！那就很容易发生向下溢出
*/

import "fmt"

// AVL is an AVL tree
type AVL struct {
	root *node
}

// New new an empty AVL tree
func New() *AVL {
	return new(AVL)
}

// Insert insert a k-v pair into this dictionary
// if value is nil, is is a vitual deletion operation
func (avl *AVL) Insert(key int, value T) {
	avl.root = avl.root.insert(key, value)
}

// Find -
func (avl *AVL) Find(key int) T {
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

// FindMin -
func (avl *AVL) FindMin() T {
	return avl.root.findMin().value
}

// FindMax -
func (avl *AVL) FindMax() T {
	return avl.root.findMax().value
}

// Delete -
func (avl *AVL) Delete(key int) {
	avl.root = avl.root.delete(key)
}

// Empty return true if it is an emoty tree
func (avl *AVL) Empty() bool {
	return avl.root == nil
}

type T interface{}

type node struct {
	key    int
	value  T
	height int8
	left   *node
	right  *node
}

func (n *node) insert(k int, v T) *node {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
	} else if k == n.key {
		n.value = v
	} else if k < n.key {
		n.left = n.left.insert(k, v)
		if n.diff() > 1 {
			n = rotation(n)
		}
		n.updateHeight()
	} else {
		n.right = n.right.insert(k, v)
		if n.diff() < -1 {
			n = rotation(n)
		}
		n.updateHeight()
	}
	return n
}

func (n *node) diff() int8 {
	return height(n.left)-height(n.right)
}

func (n *node) updateHeight() {
	n.height = max(height(n.left), height(n.right)) + 1
}

func rotation(r *node) *node {
	diff := r.diff()
	switch true {
	case diff == 2:
		if height(r.left.left) > height(r.left.right) {
			r = leftSingleRotation(r)
		} else {
			r = leftDoubleRotation(r)
		}
	case diff == -2:
		if height(r.right.right) > height(r.right.left) {
			r = rightSingleRotation(r)
		} else {
			r = rightDoubleRotation(r)
		}
	default:
		panic(fmt.Sprintf("|diff| == |%v - %v| != 2", height(r.left), height(r.right)))
	}
	return r
}

func leftSingleRotation(k2 *node) *node {
	k1 := k2.left
	k2.left = k1.right
	k1.right = k2
	k2.height = max(height(k2.left), height(k2.right)) + 1
	k1.height = max(height(k1.left), height(k1.right)) + 1
	return k1
}

func rightSingleRotation(k2 *node) *node {
	k1 := k2.right
	k2.right = k1.left
	k1.left = k2
	k2.height = max(height(k2.left), height(k2.right)) + 1
	k1.height = max(height(k1.left), height(k1.right)) + 1
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

func height(n *node) int8 {
	if n == nil {
		return -1
	}
	return n.height
}

func max(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func (n *node) find(key int) *node {
	if n == nil {
		return nil
	}
	if key > n.key {
		return n.right.find(key)
	}
	if key < n.key {
		return n.left.find(key)
	}
	return n
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

func (n *node) delete(key int) *node {
	if n == nil {
		fmt.Printf("avltree: delete not existed key(%v)", key)
		return nil
	}
	if key < n.key {
		n.left = n.left.delete(key)
		n.updateHeight()
		if height(n.right)-height(n.left) > 1 {
			n = rotation(n)
		}
		return n
	}
	if key > n.key {
		n.right = n.right.delete(key)
		n.updateHeight()
		if height(n.left)-height(n.right) > 1 {
			n = rotation(n)
		}
		return n
	}

	// delete n
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
	if height(n.left)-height(n.right) > 1 {
		n = rotation(n)
	}
	return n
}
