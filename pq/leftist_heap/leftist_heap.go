package leftist

import . "github.com/howz97/algorithm/util"

func New() *LeftHeap {
	return new(LeftHeap)
}

type LeftHeap struct {
	root *node
	size int
}

func (lh *LeftHeap) Push(p Comparable) {
	lh.root = merge(lh.root, &node{
		p: p,
	})
	lh.size++
}

func (lh *LeftHeap) Peek() T {
	if lh.root == nil {
		return nil
	}
	return lh.root.p
}

func (lh *LeftHeap) Pop() Comparable {
	p := lh.root.p
	lh.root = merge(lh.root.left, lh.root.right)
	lh.size--
	return p
}

func (lh *LeftHeap) Merge(other *LeftHeap) {
	lh.root = merge(lh.root, other.root)
	lh.size += other.size
}

func (lh *LeftHeap) Size() int {
	return lh.size
}

type node struct {
	p     Comparable // priority
	npl   int        // null path length
	left  *node
	right *node
}

func (n *node) getNPL() int {
	if n == nil {
		return -1
	}
	return n.npl
}

func (n *node) Cmp(n2 *node) Result {
	return n.p.Cmp(n2.p)
}

func merge(n1, n2 *node) *node {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.Cmp(n2) == More {
		n1, n2 = n2, n1
	}
	n1.right = merge(n1.right, n2)
	if n1.left.getNPL() < n1.right.getNPL() {
		n1.left, n1.right = n1.right, n1.left
	}
	n1.npl = n1.right.getNPL() + 1
	return n1
}
