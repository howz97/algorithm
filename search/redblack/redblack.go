package redblack

import . "github.com/howz97/algorithm/search"

const (
	red   = true
	black = false
)

type RedBlack struct {
	root *node
}

func New() *RedBlack {
	return new(RedBlack)
}

func (rb *RedBlack) Put(key Cmp, value T) {
	rb.root = rb.root.insert(key, value)
	if rb.root.isRed() {
		rb.root.color = black
	}
}

func (rb *RedBlack) Get(key Cmp) T {
	n := rb.root.find(key)
	if n == nil {
		return nil
	}
	if n.value == nil {
		rb.Del(n.key) // 删除value为nil的节点
		return nil
	}
	return n.value
}

func (rb *RedBlack) GetMin() T {
	return rb.root.findMin().value
}

func (rb *RedBlack) GetMax() T {
	return rb.root.findMax().value
}

func (rb *RedBlack) DelMin() {
	if rb.root == nil {
		return
	}
	if !rb.root.leftSon.isRed() && !rb.root.rightSon.isRed() {
		rb.root.color = red
	}
	rb.root = rb.root.delMin()
	if rb.root != nil {
		rb.root.color = black
	}
}

func (rb *RedBlack) DelMax() {
	if rb.root == nil {
		return
	}
	if !rb.root.leftSon.isRed() && !rb.root.rightSon.isRed() {
		rb.root.color = red
	}
	rb.root = rb.root.delMax()
	if rb.root != nil {
		rb.root.color = black
	}
}

func (rb *RedBlack) Del(key Cmp) {
	if rb.root == nil {
		return
	}
	if !rb.root.leftSon.isRed() && !rb.root.rightSon.isRed() {
		rb.root.color = red
	}
	rb.root = rb.root.delete(key)
	if rb.root != nil {
		rb.root.color = black
	}
}

func (rb *RedBlack) Empty() bool {
	return rb.root == nil
}

func (rb *RedBlack) Size() int {
	if rb.root == nil {
		return 0
	}
	return rb.root.size
}

func (rb *RedBlack) Clean() {
	rb.root = nil
}

/*=============================================================================*/

type node struct {
	key      Cmp
	value    T
	color    bool
	size     int
	leftSon  *node
	rightSon *node
}

func (n *node) insert(k Cmp, v T) *node {
	if n == nil {
		return &node{
			key:   k,
			value: v,
			color: red,
			size:  1,
		}
	}
	switch k.Cmp(n.key) {
	case Less:
		n.leftSon = n.leftSon.insert(k, v)
	case More:
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
		flipColors(n)
	}
	return n
}

func (n *node) find(key Cmp) *node {
	if n == nil {
		return nil
	}
	switch key.Cmp(n.key) {
	case Less:
		return n.leftSon.find(key)
	case More:
		return n.rightSon.find(key)
	default:
		return n
	}
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

func (n *node) delMin() *node {
	if n.leftSon == nil {
		return nil
	}
	if !n.leftSon.isRed() && !n.leftSon.leftSon.isRed() {
		n = moveRedLeft(n)
	}
	n.leftSon = n.leftSon.delMin()
	return balance(n)
}

func (n *node) delMax() *node {
	if n.leftSon.isRed() {
		n = rotateRight(n)
	}
	if n.rightSon == nil {
		return nil
	}
	if !n.rightSon.isRed() && !n.rightSon.rightSon.isRed() {
		n = moveRedRight(n)
	}
	n.rightSon = n.rightSon.delMax()
	return balance(n)
}

func moveRedRight(r *node) *node {
	flipColors2(r)
	if !r.leftSon.leftSon.isRed() {
		r = rotateRight(r)
	}
	return r
}

func moveRedLeft(r *node) *node {
	flipColors2(r)
	if r.rightSon.leftSon.isRed() {
		r.rightSon = rotateRight(r.rightSon)
		r = rotateLeft(r)
	}
	return r
}

func balance(r *node) *node {
	if r.rightSon.isRed() {
		r = rotateLeft(r)
	}
	r.size = size(r.leftSon) + size(r.rightSon) + 1
	if !r.leftSon.isRed() && r.rightSon.isRed() {
		r = rotateLeft(r)
	}
	if r.leftSon.isRed() && r.leftSon.leftSon.isRed() {
		r = rotateRight(r)
	}
	if r.leftSon.isRed() && r.rightSon.isRed() {
		flipColors(r) // ?
	}
	return r
}

func flipColors2(r *node) {
	r.color = black
	r.leftSon.color = red
	r.rightSon.color = red
}

func (n *node) delete(k Cmp) *node {
	if k.Cmp(n.key) == Less {
		if !n.leftSon.isRed() && !n.leftSon.leftSon.isRed() {
			n = moveRedLeft(n)
		}
		n.leftSon = n.leftSon.delete(k)
	} else {
		if n.leftSon.isRed() {
			n = rotateRight(n)
		}
		if k.Cmp(n.key) == Equal && n.rightSon == nil {
			return nil
		}
		if !n.rightSon.isRed() && !n.rightSon.leftSon.isRed() {
			n = moveRedRight(n)
		}
		if k.Cmp(n.key) == Equal {
			min := n.rightSon.findMin()
			n.value = min.value
			n.key = min.key
			n.rightSon = n.rightSon.delMin()
		} else {
			n.rightSon = n.rightSon.delete(k)
		}
	}
	return balance(n)
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
	newRoot.leftSon = root
	root.size = size(root.leftSon) + size(root.rightSon) + 1
	newRoot.size = size(newRoot.leftSon) + size(newRoot.rightSon) + 1
	if root.color == black {
		newRoot.color = black
		root.color = red
	}
	return newRoot
}

func rotateRight(root *node) *node {
	newRoot := root.leftSon
	root.leftSon = newRoot.rightSon
	newRoot.rightSon = root
	root.size = size(root.leftSon) + size(root.rightSon) + 1
	newRoot.size = size(newRoot.leftSon) + size(newRoot.rightSon) + 1
	if root.color == black {
		newRoot.color = black
		root.color = red
	}
	return newRoot
}

func flipColors(root *node) {
	root.leftSon.color = black
	root.rightSon.color = black
	root.color = red
}
