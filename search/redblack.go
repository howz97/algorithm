// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package search

import (
	"cmp"
	"fmt"
)

type RbNode[K cmp.Ordered, V any] struct {
	key            K
	value          V
	isRed          bool
	p, left, right *RbNode[K, V]
}

func (n *RbNode[K, V]) Left() BNode {
	return n.left
}

func (n *RbNode[K, V]) Right() BNode {
	return n.right
}

func (n *RbNode[K, V]) IsNil() bool {
	return n.left == nil && n.right == nil
}

func (n *RbNode[K, V]) Key() K {
	return n.key
}

func (n *RbNode[K, V]) Value() V {
	return n.value
}

func (n *RbNode[K, V]) IsRed() bool {
	return n.isRed
}

func (n *RbNode[K, V]) String() string {
	if n.IsNil() {
		return "Nil"
	}
	if n.isRed {
		return fmt.Sprintf("red[%v]", n.key)
	} else {
		return fmt.Sprintf("(%v)", n.key)
	}
}

type RBTree[K cmp.Ordered, V any] struct {
	root *RbNode[K, V]
	null *RbNode[K, V]
	size uint
}

func (rbt *RBTree[K, V]) Put(key K, val V) {
	p := rbt.null
	x := rbt.root
	for x != rbt.null {
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
	rbt.size++
	in := &RbNode[K, V]{
		key:   key,
		value: val,
		isRed: true,
		p:     p,
		left:  rbt.null,
		right: rbt.null,
	}
	if p == rbt.null {
		// Inset root node into an empty tree
		rbt.root = in
		// fixInsert only need to set the root to black
	} else if in.key < p.key {
		p.left = in
	} else {
		p.right = in
	}
	rbt.fixInsert(in)
}

func (rbt *RBTree[K, V]) fixInsert(n *RbNode[K, V]) {
	// only 2 possible problem:
	// 1. root is red
	// 2. both n and it's parent is red
	// let us solve them and keep the other properties of RedBlack tree
	for n.p.isRed {
		if n.p == n.p.p.left {
			uncle := n.p.p.right
			if uncle.isRed {
				// case 1: swim red-attribute of parent and uncle to grandparent.
				n.p.isRed, uncle.isRed, n.p.p.isRed = false, false, true
				n = n.p.p
				// continue loop to fix grandparent
			} else {
				// uncle is black
				if n == n.p.right {
					// case 2: -> case3
					n = n.p
					rbt.leftRotate(n)
				}
				// case 3
				n.p.isRed, n.p.p.isRed = false, true
				rbt.rightRotate(n.p.p)
				// break loop
			}
		} else {
			// symmetrical to above
			uncle := n.p.p.left
			if uncle.isRed {
				n.p.isRed, uncle.isRed, n.p.p.isRed = false, false, true
				n = n.p.p
			} else {
				if n == n.p.left {
					n = n.p
					rbt.rightRotate(n)
				}
				n.p.isRed, n.p.p.isRed = false, true
				rbt.leftRotate(n.p.p)
			}
		}
	}
	rbt.root.isRed = false
}

// clockwise
func (rbt *RBTree[K, V]) rightRotate(n *RbNode[K, V]) {
	top := n.left
	n.left = top.right
	if n.left != rbt.null {
		n.left.p = n
	}
	rbt.transplant(n, top)
	top.right = n
	n.p = top
}

// counterclockwise
func (rbt *RBTree[K, V]) leftRotate(n *RbNode[K, V]) {
	top := n.right
	n.right = top.left
	if n.right != rbt.null {
		n.right.p = n
	}
	rbt.transplant(n, top)
	top.left = n
	n.p = top
}

func (rbt *RBTree[K, V]) transplant(a, b *RbNode[K, V]) {
	b.p = a.p
	if b.p == rbt.null {
		rbt.root = b
	} else if a == a.p.left {
		b.p.left = b
	} else {
		b.p.right = b
	}
}

func (rbt *RBTree[K, V]) find(key K) *RbNode[K, V] {
	cur := rbt.root
loop:
	for cur != rbt.null {
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

func (rbt *RBTree[K, V]) Del(key K) {
	del := rbt.find(key)
	if del == rbt.null {
		return
	}
	rbt.size--
	del2 := del
	d2isRed := del2.isRed
	var replacer *RbNode[K, V]
	if del.left == rbt.null {
		replacer = del.right
		rbt.transplant(del, replacer)
	} else if del.right == rbt.null {
		replacer = del.left
		rbt.transplant(del, replacer)
	} else {
		del2 = rbt.getMin(del.right)
		d2isRed = del2.isRed
		replacer = del2.right
		if del2.p == del {
			// replacer may be tree.null, but fixDelete require replacer.p non-null
			replacer.p = del2
		} else {
			rbt.transplant(del2, replacer)
			del2.right = del.right
			del2.right.p = del2
		}
		rbt.transplant(del, del2)
		del2.left = del.left
		del2.left.p = del2
		del2.isRed = del.isRed
	}
	if !d2isRed {
		rbt.fixDelete(replacer)
	}
}

func (rbt *RBTree[K, V]) fixDelete(n *RbNode[K, V]) {
	for n != rbt.root && !n.isRed {
		if n == n.p.left {
			sibling := n.p.right
			if sibling.isRed {
				// case 1: convert sibling to black (-> case2,3,4)
				sibling.isRed, sibling.p.isRed = false, true
				rbt.leftRotate(sibling.p)
				sibling = n.p.right
			}
			// sibling is black
			if !sibling.left.isRed && !sibling.right.isRed {
				// case 2: swim black of n and sibling to parent
				sibling.isRed = true
				n = n.p
				// continue loop to fix parent
			} else {
				if !sibling.right.isRed {
					// case 3: -> case4
					sibling.left.isRed, sibling.isRed = false, true
					rbt.rightRotate(sibling)
					sibling = n.p.right
				}
				// case 4
				sibling.right.isRed, sibling.isRed, sibling.p.isRed = false, sibling.p.isRed, false
				rbt.leftRotate(sibling.p)
				n = rbt.root
				// break loop
			}
		} else {
			// symmetrical to above
			sibling := n.p.left
			if sibling.isRed {
				sibling.isRed, sibling.p.isRed = false, true
				rbt.rightRotate(n.p)
				sibling = n.p.left
			}
			if !sibling.left.isRed && !sibling.right.isRed {
				sibling.isRed = true
				n = n.p
			} else {
				if !sibling.left.isRed {
					sibling.right.isRed, sibling.isRed = false, true
					rbt.leftRotate(sibling)
					sibling = n.p.left
				}
				sibling.left.isRed, sibling.isRed, sibling.p.isRed = false, sibling.p.isRed, false
				rbt.rightRotate(sibling.p)
				n = rbt.root
			}
		}
	}
	n.isRed = false
}

func (rbt *RBTree[K, V]) getMin(n *RbNode[K, V]) (min *RbNode[K, V]) {
	if n.IsNil() {
		return n
	}
	min = n
	for min.left != rbt.null {
		min = min.left
	}
	return
}

func (rbt *RBTree[K, V]) Get(key K) (V, bool) {
	n := rbt.find(key)
	if n == rbt.null {
		var v V
		return v, false
	}
	return n.value, true
}

func (rbt *RBTree[K, V]) Size() uint {
	return rbt.size
}

func (rbt *RBTree[K, V]) Root() *RbNode[K, V] {
	return rbt.root
}

func (rbt *RBTree[K, V]) CheckValid() {
	if rbt.root.isRed {
		panic("root is red")
	}
	cnt := uint(0)
	blackHeight := -1
	Inorder(rbt.Root(), func(nd *RbNode[K, V]) bool {
		cnt++
		if nd.left == rbt.null && nd.right == rbt.null {
			// this is leaf
			bh := 0
			if !nd.isRed {
				bh++
			}
			childIsRed := nd.isRed
			nd = nd.p
			for nd != rbt.null {
				if nd.isRed && childIsRed {
					panic("red next to red")
				}
				if !nd.isRed {
					bh++
				}
				childIsRed = nd.isRed
				nd = nd.p
			}
			if blackHeight < 0 {
				blackHeight = bh
			} else if bh != blackHeight {
				panic("black height not euqal")
			}
		}
		return true
	})
	if cnt != rbt.size {
		panic("size not match")
	}
}

func NewRBTree[K cmp.Ordered, V any]() *RBTree[K, V] {
	null := new(RbNode[K, V])
	null.isRed = false
	null.p = null
	return &RBTree[K, V]{
		root: null,
		null: null,
	}
}
