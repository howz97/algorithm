package pq

import "cmp"

func NewLeftist[O cmp.Ordered]() *Leftist[O] {
	return new(Leftist[O])
}

type Leftist[O cmp.Ordered] struct {
	root *lNode[O]
	size int
}

func (lh *Leftist[O]) Push(o O) {
	lh.root = lMerge(lh.root, &lNode[O]{
		p: o,
	})
	lh.size++
}

func (lh *Leftist[O]) Peek() O {
	return lh.root.p
}

func (lh *Leftist[O]) Pop() O {
	p := lh.root.p
	lh.root = lMerge(lh.root.left, lh.root.right)
	lh.size--
	return p
}

func (lh *Leftist[O]) Top() O {
	return lh.root.p
}

func (lh *Leftist[O]) Merge(other *Leftist[O]) {
	lh.root = lMerge(lh.root, other.root)
	lh.size += other.size
}

func (lh *Leftist[O]) Size() int {
	return lh.size
}

type lNode[O cmp.Ordered] struct {
	p     O   // priority
	npl   int // null path length
	left  *lNode[O]
	right *lNode[O]
}

func (n *lNode[O]) getNPL() int {
	if n == nil {
		return -1
	}
	return n.npl
}

func lMerge[O cmp.Ordered](n1, n2 *lNode[O]) *lNode[O] {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.p > n2.p {
		n1, n2 = n2, n1
	}
	n1.right = lMerge(n1.right, n2)
	if n1.left.getNPL() < n1.right.getNPL() {
		n1.left, n1.right = n1.right, n1.left
	}
	n1.npl = n1.right.getNPL() + 1
	return n1
}
