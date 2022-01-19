package v2

import (
	. "github.com/howz97/algorithm/util"
)

const (
	red   = true
	black = false
)

type node struct {
	key            Comparable
	value          T
	color          bool
	p, left, right *node
}

func (n *node) Cmp(a Comparable) Result {
	return n.key.Cmp(a.(*node).key)
}

func (n *node) getMin() *node {

}

type Tree struct {
	root *node
	null *node
}

func (tree *Tree) Put(key Comparable, val T) {
	p := tree.null
	x := tree.root
	for x != tree.null {
		p = x
		switch key.Cmp(x.key) {
		case Less:
			x = x.left
		case More:
			x = x.right
		case Equal:
			x.value = val
			return
		}
	}
	in := &node{
		key:   key,
		value: val,
		color: red,
		p:     p,
		left:  tree.null,
		right: tree.null,
	}
	if p == tree.null {
		tree.root = in
	} else if in.Cmp(p) == Less {
		p.left = in
	} else {
		p.right = in
	}
	tree.fixInsert(in)
}

func (tree *Tree) fixInsert(n *node) {
	for n.p.color == red {
		if n.p == n.p.p.left {
			uncle := n.p.p.right
			if uncle.color == red {
				n.p.color, uncle.color = black, black
				n = n.p.p
				n.color = red
				// continue
			} else {
				// uncle is black
				if n == n.p.right {
					n = n.p
					tree.leftRotate(n)
				}
				n.p.color = black
				n.p.p.color = red
				tree.rightRotate(n.p.p)
			}
		} else {
			uncle := n.p.p.left
			if uncle.color == red {
				n.p.color, uncle.color = black, black
				n = n.p.p
				n.color = red
				// continue
			} else {
				// uncle is black
				if n == n.p.left {
					n = n.p
					tree.rightRotate(n)
				}
				n.p.color = black
				n.p.p.color = red
				tree.leftRotate(n.p.p)
			}
		}
	}
}

func (tree *Tree) rightRotate(n *node) {
	top := n.left
	n.left = top.right
	if n.left != tree.null {
		n.left.p = n
	}
	tree.transplant(n, top)
	top.right = n
	n.p = top
}

func (tree *Tree) leftRotate(n *node) {
	top := n.right
	n.right = top.left
	if n.right != tree.null {
		n.right.p = n
	}
	tree.transplant(n, top)
	top.left = n
	n.p = top
}

func (tree *Tree) transplant(a, b *node) {
	b.p = a.p
	if b.p == tree.null {
		tree.root = b
	} else if a == a.p.left {
		b.p.left = b
	} else {
		b.p.right = b
	}
}

func (tree *Tree) find(key Comparable) *node {

}

func (tree *Tree) Del(key Comparable) {
	z := tree.find(key)
	if z == nil {
		return
	}
	y := z
	yOrig := y.color
	var x *node
	if z.left == tree.null {
		x = z.right
		tree.transplant(z, x)
	} else if z.right == tree.null {
		x = z.left
		tree.transplant(z, x)
	} else {
		y = z.right.getMin()
		yOrig = y.color
		x = y.right
		if y.p != z {
			tree.transplant(y, x)
			y.right = z.right
			y.right.p = y
		}
		tree.transplant(z, y)
		y.left = z.left
		y.left.p = y
		y.color = z.color
	}
	if yOrig == black {
		tree.fixDelete(x)
	}
}

func (tree *Tree) fixDelete(n *node) {
	for n != tree.root && n.color == black {

	}
	n.color = black
}
