package redblack

import (
	"fmt"
	. "github.com/howz97/algorithm/search"
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

func (n *node) Left() ITraversal {
	return n.left
}

func (n *node) Right() ITraversal {
	return n.right
}

func (n *node) IsNil() bool {
	return n.left == nil && n.right == nil
}

func (n *node) String() string {
	if n.IsNil() {
		return "Nil"
	}
	var color string
	if n.color == red {
		color = "r"
	} else {
		color = "b"
	}
	return fmt.Sprintf("%s[%v]", color, n.key)
}

type Tree struct {
	root *node
	null *node
	size uint
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
	tree.size++
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
			if uncle == nil {
				fmt.Println("uncle is nil ")
				PrintBinaryTree(tree.root) // todo
			}
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
	tree.root.color = black
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
	cur := tree.root
loop:
	for cur != tree.null {
		switch key.Cmp(cur.key) {
		case Less:
			cur = cur.left
		case More:
			cur = cur.right
		case Equal:
			break loop
		}
	}
	return cur
}

func (tree *Tree) Del(key Comparable) {
	del := tree.find(key)
	if del == tree.null {
		return
	}
	tree.size--
	del2 := del
	d2Orig := del2.color
	var rep *node
	if del.left == tree.null {
		rep = del.right
		tree.transplant(del, rep)
	} else if del.right == tree.null {
		rep = del.left
		tree.transplant(del, rep)
	} else {
		del2 = tree.getMin(del.right)
		d2Orig = del2.color
		rep = del2.right
		if del2.p == del {
			rep.p = del2 // todo: why ?
		} else {
			tree.transplant(del2, rep)
			del2.right = del.right
			del2.right.p = del2
		}
		tree.transplant(del, del2)
		del2.left = del.left
		del2.left.p = del2
		del2.color = del.color
	}
	if d2Orig == black {
		tree.fixDelete(rep)
	}
}

func (tree *Tree) fixDelete(n *node) {
	for n != tree.root && n.color == black {
		if n == n.p.left {
			sibling := n.p.right
			if sibling.color == red {
				sibling.color = black
				n.p.color = red
				tree.leftRotate(n.p)
				sibling = n.p.right
			}
			// sibling is black
			if sibling.left.color == black && sibling.right.color == black {
				sibling.color = red
				n = n.p
				// continue
			} else {
				if sibling.right.color == black {
					sibling.left.color = black
					sibling.color = red
					tree.rightRotate(sibling)
					sibling = n.p.right
				}
				// sibling.right is red
				n.p.color, sibling.color = sibling.color, n.p.color
				sibling.right.color = black
				tree.leftRotate(n.p)
				n = tree.root
			}
		} else {
			sibling := n.p.left
			if sibling.color == red {
				sibling.color = black
				n.p.color = red
				tree.rightRotate(n.p)
				sibling = n.p.left
			}
			// sibling is black
			if sibling.left.color == black && sibling.right.color == black {
				sibling.color = red
				n = n.p
				// continue
			} else {
				if sibling.left.color == black {
					sibling.right.color = black
					sibling.color = red
					tree.leftRotate(sibling)
					sibling = n.p.left
				}
				// sibling.right is red
				n.p.color, sibling.color = sibling.color, n.p.color
				sibling.left.color = black
				tree.rightRotate(n.p)
				n = tree.root
			}
		}
	}
	n.color = black
}

func (tree *Tree) getMin(n *node) (min *node) {
	if n.IsNil() {
		return n
	}
	min = n
	for min.left != tree.null {
		min = min.left
	}
	return
}

func (tree *Tree) Get(key Comparable) T {
	n := tree.find(key)
	if n == tree.null {
		return nil
	}
	return n.value
}

func (tree *Tree) Size() uint {
	return tree.size
}

func (tree *Tree) Clean() {
	tree.root = tree.null
	tree.size = 0
}

func New() *Tree {
	null := new(node)
	null.color = black
	null.p = null
	return &Tree{
		root: null,
		null: null,
	}
}
