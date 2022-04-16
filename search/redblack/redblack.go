package redblack

import (
	"fmt"

	. "github.com/howz97/algorithm/search"
	"golang.org/x/exp/constraints"
)

const (
	red   = true
	black = false
)

type node[Ord constraints.Ordered, T any] struct {
	key            Ord
	value          T
	color          bool
	p, left, right *node[Ord, T]
}

func (n *node[Ord, T]) Left() ITraversal {
	return n.left
}

func (n *node[Ord, T]) Right() ITraversal {
	return n.right
}

func (n *node[Ord, T]) IsNil() bool {
	return n.left == nil && n.right == nil
}

func (n *node[Ord, T]) String() string {
	if n.IsNil() {
		return "Nil"
	}
	if n.color == red {
		return fmt.Sprintf("red[%v]", n.key)
	} else {
		return fmt.Sprintf("(%v)", n.key)
	}
}

type Tree[Ord constraints.Ordered, T any] struct {
	*node[Ord, T]
	null *node[Ord, T]
	size uint
}

func (tree *Tree[Ord, T]) Put(key Ord, val T) {
	p := tree.null
	x := tree.node
	for x != tree.null {
		p = x
		if key < x.key {
			x = x.left
		} else if key > x.key {
			x = x.right
		} else {
			// update value, no insertion
			x.value = val
			return
		}
	}
	// p is parent of in.
	tree.size++
	in := &node[Ord, T]{
		key:   key,
		value: val,
		color: red,
		p:     p,
		left:  tree.null,
		right: tree.null,
	}
	if p == tree.null {
		// Inset root node into an empty tree
		tree.node = in
		// fixInsert only need to set the root to black
	} else if in.key < p.key {
		p.left = in
	} else {
		p.right = in
	}
	tree.fixInsert(in)
}

func (tree *Tree[Ord, T]) fixInsert(n *node[Ord, T]) {
	// only 2 possible problem:
	// 1. root is red
	// 2. both n and it's parent is red
	// let us solve them and keep the other properties of RedBlack tree
	for n.p.color == red {
		if n.p == n.p.p.left {
			uncle := n.p.p.right
			if uncle.color == red {
				// case 1: swim red-attribute of parent and uncle to grandparent.
				n.p.color, uncle.color, n.p.p.color = black, black, red
				n = n.p.p
				// continue loop to fix grandparent
			} else {
				// uncle is black
				if n == n.p.right {
					// case 2: -> case3
					n = n.p
					tree.leftRotate(n)
				}
				// case 3
				n.p.color, n.p.p.color = black, red
				tree.rightRotate(n.p.p)
				// break loop
			}
		} else {
			// symmetrical to above
			uncle := n.p.p.left
			if uncle.color == red {
				n.p.color, uncle.color, n.p.p.color = black, black, red
				n = n.p.p
			} else {
				if n == n.p.left {
					n = n.p
					tree.rightRotate(n)
				}
				n.p.color, n.p.p.color = black, red
				tree.leftRotate(n.p.p)
			}
		}
	}
	tree.node.color = black
}

// clockwise
func (tree *Tree[Ord, T]) rightRotate(n *node[Ord, T]) {
	top := n.left
	n.left = top.right
	if n.left != tree.null {
		n.left.p = n
	}
	tree.transplant(n, top)
	top.right = n
	n.p = top
}

// counterclockwise
func (tree *Tree[Ord, T]) leftRotate(n *node[Ord, T]) {
	top := n.right
	n.right = top.left
	if n.right != tree.null {
		n.right.p = n
	}
	tree.transplant(n, top)
	top.left = n
	n.p = top
}

func (tree *Tree[Ord, T]) transplant(a, b *node[Ord, T]) {
	b.p = a.p
	if b.p == tree.null {
		tree.node = b
	} else if a == a.p.left {
		b.p.left = b
	} else {
		b.p.right = b
	}
}

func (tree *Tree[Ord, T]) find(key Ord) *node[Ord, T] {
	cur := tree.node
loop:
	for cur != tree.null {
		if key < cur.key {
			cur = cur.left
		} else if key > cur.key {
			cur = cur.right
		} else {
			break loop
		}
	}
	return cur
}

func (tree *Tree[Ord, T]) Del(key Ord) {
	del := tree.find(key)
	if del == tree.null {
		return
	}
	tree.size--
	del2 := del
	d2color := del2.color
	var replacer *node[Ord, T]
	if del.left == tree.null {
		replacer = del.right
		tree.transplant(del, replacer)
	} else if del.right == tree.null {
		replacer = del.left
		tree.transplant(del, replacer)
	} else {
		del2 = tree.getMin(del.right)
		d2color = del2.color
		replacer = del2.right
		if del2.p == del {
			// replacer may be tree.null, but fixDelete require replacer.p non-null
			replacer.p = del2
		} else {
			tree.transplant(del2, replacer)
			del2.right = del.right
			del2.right.p = del2
		}
		tree.transplant(del, del2)
		del2.left = del.left
		del2.left.p = del2
		del2.color = del.color
	}
	if d2color == black {
		tree.fixDelete(replacer)
	}
}

func (tree *Tree[Ord, T]) fixDelete(n *node[Ord, T]) {
	for n != tree.node && n.color == black {
		if n == n.p.left {
			sibling := n.p.right
			if sibling.color == red {
				// case 1: convert sibling to black (-> case2,3,4)
				sibling.color, sibling.p.color = black, red
				tree.leftRotate(sibling.p)
				sibling = n.p.right
			}
			// sibling is black
			if sibling.left.color == black && sibling.right.color == black {
				// case 2: swim black of n and sibling to parent
				sibling.color = red
				n = n.p
				// continue loop to fix parent
			} else {
				if sibling.right.color == black {
					// case 3: -> case4
					sibling.left.color, sibling.color = black, red
					tree.rightRotate(sibling)
					sibling = n.p.right
				}
				// case 4
				sibling.right.color, sibling.color, sibling.p.color = black, sibling.p.color, black
				tree.leftRotate(sibling.p)
				n = tree.node
				// break loop
			}
		} else {
			// symmetrical to above
			sibling := n.p.left
			if sibling.color == red {
				sibling.color, sibling.p.color = black, red
				tree.rightRotate(n.p)
				sibling = n.p.left
			}
			if sibling.left.color == black && sibling.right.color == black {
				sibling.color = red
				n = n.p
			} else {
				if sibling.left.color == black {
					sibling.right.color, sibling.color = black, red
					tree.leftRotate(sibling)
					sibling = n.p.left
				}
				sibling.left.color, sibling.color, sibling.p.color = black, sibling.p.color, black
				tree.rightRotate(sibling.p)
				n = tree.node
			}
		}
	}
	n.color = black
}

func (tree *Tree[Ord, T]) getMin(n *node[Ord, T]) (min *node[Ord, T]) {
	if n.IsNil() {
		return n
	}
	min = n
	for min.left != tree.null {
		min = min.left
	}
	return
}

func (tree *Tree[Ord, T]) Get(key Ord) (T, bool) {
	n := tree.find(key)
	if n == tree.null {
		var v T
		return v, false
	}
	return n.value, true
}

func (tree *Tree[Ord, T]) Size() uint {
	return tree.size
}

func (tree *Tree[Ord, T]) Clean() {
	tree.node = tree.null
	tree.size = 0
}

func New[Ord constraints.Ordered, T any]() *Tree[Ord, T] {
	null := new(node[Ord, T])
	null.color = black
	null.p = null
	return &Tree[Ord, T]{
		node: null,
		null: null,
	}
}
