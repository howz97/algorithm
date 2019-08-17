package redblack

import "fmt"

const (
	red   = true
	black = false
)

// RedBlack is a dictionary implemented by read-black tree
type RedBlack struct {
	root *node
}

// New new an empty read-black tree
func New() *RedBlack {
	return new(RedBlack)
}

// Insert insert a k-v pair into this dictionary
// if value is nil, is is a vitual deletion operation
func (rb *RedBlack) Insert(key int, value interface{}) {
	rb.root = rb.root.insert(key, value)
	if rb.root.isRed() {
		rb.root.color = black
	}
}

// Find -
func (rb *RedBlack) Find(key int) interface{} {
	n := rb.root.find(key)
	if n == nil {
		return nil
	}
	if n.value == nil {
		rb.Delete(n.key) // 删除value为nil的节点
		return nil
	}
	return n.value
}

// FindMin -
func (rb *RedBlack) FindMin() interface{} {
	return rb.root.findMin().value
}

// FindMax -
func (rb *RedBlack) FindMax() interface{} {
	return rb.root.findMax().value
}

// Delete -
func (rb *RedBlack) Delete(key int) {
	rb.root = rb.root.delete(key)
}

// Empty return true if it is an emoty tree
func (rb *RedBlack) Empty() bool {
	return rb.root == nil
}

type node struct {
	key      int
	value    interface{}
	color    bool
	size     int
	leftSon  *node
	rightSon *node
}

func (n *node) insert(k int, v interface{}) *node {
	if n == nil {
		return &node{
			key:   k,
			value: v,
			color: red,
			size:  1,
		}
	}
	switch true {
	case k < n.key:
		n.leftSon = n.leftSon.insert(k, v)
	case k > n.key:
		n.rightSon = n.rightSon.insert(k, v)
	default:
		n.value = v
	}
	n.size = size(n.leftSon) + size(n.rightSon) + 1
	if !n.leftSon.isRed() && n.rightSon.isRed() {
		n = rotateLeft(n)
	}
	if n.leftSon.isRed() && n.leftSon.leftSon.isRed() {
		n = rotateRight(n)
	}
	if n.leftSon.isRed() && n.rightSon.isRed() {
		n = flipColors(n)
	}
	return n
}

func (n *node) find(key int) *node {
	if n == nil {
		return nil
	}
	if key > n.key {
		return n.rightSon.find(key)
	}
	if key < n.key {
		return n.leftSon.find(key)
	}
	return n
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

func (n *node) delete(k int) *node {
	fmt.Println("delete not support")
	return n
}

func (n *node) isRed() bool {
	if n == nil {
		return false
	}
	return n.color == red
}

func size(n *node) int {
	if n == nil {
		return 0
	}
	return n.size
}

func rotateLeft(root *node) *node {
	newRoot := root.rightSon
	root.rightSon = newRoot.leftSon
	root.color = red
	newRoot.leftSon = root
	newRoot.color = black
	root.size = size(root.leftSon) + size(root.rightSon) + 1
	newRoot.size = size(newRoot.leftSon) + size(newRoot.rightSon) + 1
	return newRoot
}

func rotateRight(root *node) *node {
	newRoot := root.leftSon
	root.leftSon = newRoot.rightSon
	root.color = red
	newRoot.rightSon = root
	newRoot.color = black
	root.size = size(root.leftSon) + size(root.rightSon) + 1
	newRoot.size = size(newRoot.leftSon) + size(newRoot.rightSon) + 1
	return newRoot
}

func flipColors(root *node) *node {
	root.leftSon.color = black
	root.rightSon.color = black
	root.color = red
	return root
}
