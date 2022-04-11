package leftist

import "golang.org/x/exp/constraints"

func New[O constraints.Ordered]() *Heap[O] {
	return new(Heap[O])
}

type Heap[O constraints.Ordered] struct {
	root *node[O]
	size int
}

func (lh *Heap[O]) Push(o O) {
	lh.root = merge(lh.root, &node[O]{
		p: o,
	})
	lh.size++
}

func (lh *Heap[O]) Peek() O {
	return lh.root.p
}

func (lh *Heap[O]) Pop() O {
	p := lh.root.p
	lh.root = merge(lh.root.left, lh.root.right)
	lh.size--
	return p
}

func (lh *Heap[O]) Merge(other *Heap[O]) {
	lh.root = merge(lh.root, other.root)
	lh.size += other.size
}

func (lh *Heap[O]) Size() int {
	return lh.size
}

type node[O constraints.Ordered] struct {
	p     O   // priority
	npl   int // null path length
	left  *node[O]
	right *node[O]
}

func (n *node[O]) getNPL() int {
	if n == nil {
		return -1
	}
	return n.npl
}

func merge[O constraints.Ordered](n1, n2 *node[O]) *node[O] {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	if n1.p > n2.p {
		n1, n2 = n2, n1
	}
	n1.right = merge(n1.right, n2)
	if n1.left.getNPL() < n1.right.getNPL() {
		n1.left, n1.right = n1.right, n1.left
	}
	n1.npl = n1.right.getNPL() + 1
	return n1
}
