package leftist

import . "github.com/howz97/algorithm/util"

func New() *LeftHeap {
	return new(LeftHeap)
}

type LeftHeap struct {
	root *node
	size int
}

func (lh *LeftHeap) Push(p Comparable, v T) {
	lh.root = merge(lh.root, &node{
		Comparable: p,
		val:        v,
	})
	lh.size++
}

func (lh *LeftHeap) Peek() T {
	if lh.root == nil {
		return nil
	}
	return lh.root.val
}

func (lh *LeftHeap) Pop() T {
	v := lh.root.val
	lh.root = merge(lh.root.left, lh.root.right)
	lh.size--
	return v
}

func (lh *LeftHeap) Merge(other *LeftHeap) {
	lh.root = merge(lh.root, other.root)
	lh.size += other.size
}

func (lh *LeftHeap) Size() int {
	return lh.size
}

type node struct {
	Comparable // priority
	val        T
	npl        int // null path length
	left       *node
	right      *node
}

func (h *node) getNPL() int {
	if h == nil {
		return -1
	}
	return h.npl
}

func merge(h1, h2 *node) *node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.Cmp(h2) == More {
		h1, h2 = h2, h1
	}
	h1.right = merge(h1.right, h2)
	if h1.left.getNPL() < h1.right.getNPL() {
		h1.left, h1.right = h1.right, h1.left
	}
	h1.npl = h1.right.getNPL() + 1
	return h1
}
