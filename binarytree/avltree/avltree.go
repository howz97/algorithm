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
func (avl *AVL) Insert(key int, value interface{}) {
	avl.root = avl.root.insert(key, value)
}

// Find -
func (avl *AVL) Find(key int) interface{} {
	return avl.root.find(key)
}

// FindMin -
func (avl *AVL) FindMin() interface{} {
	return avl.root.findMin().value
}

// FindMax -
func (avl *AVL) FindMax() interface{} {
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

type node struct {
	value    interface{}
	key      int
	height   int8
	leftSon  *node
	rightSon *node
}

func (n *node) insert(k int, v interface{}) *node {
	if n == nil {
		n = new(node)
		n.key = k
		n.value = v
	} else if k == n.key {
		n.value = v
	} else if k < n.key {
		n.leftSon = n.leftSon.insert(k, v)
		if height(n.leftSon)-height(n.rightSon) > 1 {
			n = rotation(n)
		}
		n.height = max(height(n.leftSon), height(n.rightSon)) + 1
	} else {
		n.rightSon = n.rightSon.insert(k, v)
		if height(n.rightSon)-height(n.leftSon) > 1 {
			n = rotation(n)
		}
		n.height = max(height(n.leftSon), height(n.rightSon)) + 1
	}
	return n
}

func rotation(r *node) *node {
	switch true {
	case height(r.leftSon)-height(r.rightSon) == 2:
		if height(r.leftSon.leftSon) > height(r.leftSon.rightSon) {
			r = leftSingelRotation(r)
		} else {
			r = leftDoubleRotation(r)
		}
	case height(r.rightSon)-height(r.leftSon) == 2:
		if height(r.rightSon.rightSon) > height(r.rightSon.leftSon) {
			r = rightSingelRotation(r)
		} else {
			r = rightDoubleRotation(r)
		}
	default:
		panic(fmt.Sprintf("rotation: |height(leftSon) - height(rightSon)| == |%v - %v| != 2", height(r.leftSon), height(r.rightSon)))
	}
	return r
}

func leftSingelRotation(k2 *node) *node {
	k1 := k2.leftSon
	k2.leftSon = k1.rightSon
	k1.rightSon = k2
	k2.height = max(height(k2.leftSon), height(k2.rightSon)) + 1
	k1.height = max(height(k1.leftSon), height(k1.rightSon)) + 1
	return k1
}

func rightSingelRotation(k2 *node) *node {
	k1 := k2.rightSon
	k2.rightSon = k1.leftSon
	k1.leftSon = k2
	k2.height = max(height(k2.leftSon), height(k2.rightSon)) + 1
	k1.height = max(height(k1.leftSon), height(k1.rightSon)) + 1
	return k1
}

func leftDoubleRotation(k3 *node) *node {
	k3.leftSon = rightSingelRotation(k3.leftSon)
	return leftSingelRotation(k3)
}

func rightDoubleRotation(k3 *node) *node {
	k3.rightSon = leftSingelRotation(k3.rightSon)
	return rightSingelRotation(k3)
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

func (n *node) find(key int) interface{} {
	if n == nil {
		return nil
	}
	if key > n.key {
		return n.rightSon.find(key)
	}
	if key < n.key {
		return n.leftSon.find(key)
	}
	return n.value
}

func (n *node) findMin() *node {
	if n == nil {
		return nil
	}
	for n.leftSon != nil {
		n = n.leftSon
	}
	return n
}

func (n *node) findMax() *node {
	if n == nil {
		return nil
	}
	for n.rightSon != nil {
		n = n.rightSon
	}
	return n
}

func (n *node) delete(key int) *node {
	if n == nil {
		fmt.Printf("avltree: delete not existed key(%v)", key)
		return nil
	}
	if key < n.key {
		n.leftSon = n.leftSon.delete(key)
		n.height = max(height(n.leftSon), height(n.rightSon)) + 1
		if height(n.rightSon)-height(n.leftSon) > 1 {
			n = rotation(n)
		}
		return n
	}
	if key > n.key {
		n.rightSon = n.rightSon.delete(key)
		n.height = max(height(n.leftSon), height(n.rightSon)) + 1
		if height(n.leftSon)-height(n.rightSon) > 1 {
			n = rotation(n)
		}
		return n
	}

	// delete n
	if n.leftSon == nil {
		return n.rightSon
	}
	if n.rightSon == nil {
		return n.leftSon
	}
	deleted := n
	n = n.rightSon.findMin()
	n.rightSon = deleted.rightSon.delete(n.key)
	n.leftSon = deleted.leftSon
	n.height = max(height(n.leftSon), height(n.rightSon)) + 1
	if height(n.leftSon)-height(n.rightSon) > 1 {
		n = rotation(n)
	}
	return n
}
