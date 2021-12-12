package leftist

import . "github.com/howz97/algorithm/util"

func New() *LeftHeap {
	return new(LeftHeap)
}

type LeftHeap struct {
	h    *node
	size int
}

func (lh *LeftHeap) Push(p Comparable, v T) {
	n := new(node)
	n.priority = p
	n.val = v
	lh.h = lh.h.insert(n)
	lh.size++
}

func (lh *LeftHeap) Peek() T {
	return lh.h.val
}

func (lh *LeftHeap) Pop() {
	lh.h = lh.h.delMin()
	lh.size--
}

func (lh *LeftHeap) Merge(other *LeftHeap) {
	lh.h = merge(lh.h, other.h)
	lh.size += other.size
}

func (lh *LeftHeap) Size() int {
	return lh.size
}

type node struct {
	priority Comparable
	val      T
	npl      int // null path length
	left     *node
	right    *node
}

func (h *node) insert(other *node) *node {
	return merge(h, other)
}

func (h *node) delMin() *node {
	if h == nil {
		return nil
	}
	return merge(h.left, h.right)
}

func merge(h1, h2 *node) *node {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.priority.Cmp(h2.priority) == Less {
		h1.merge1(h2)
		return h1
	}
	h2.merge1(h1)
	return h2
}

func (h *node) merge1(h1 *node) {
	if h.left == nil {
		h.left = h1
		return
	}
	h.right = merge(h.right, h1)
	if h.left.priority.Cmp(h.right.priority) == Less {
		h.left, h.right = h.right, h.left
	}
	h.npl = h.right.npl + 1
}
