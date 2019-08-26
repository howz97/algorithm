package pqueue

// LeftistH -
type LeftistH struct {
	h    *node
	size int
}

// NewLeftistH -
func NewLeftistH() *LeftistH {
	return new(LeftistH)
}

// Insert -
func (lh *LeftistH) Insert(k int) {
	n := new(node)
	n.k = k
	lh.h = lh.h.insert(n)
	lh.size++
}

// Front -
func (lh *LeftistH) Front() (k int, ok bool) {
	if lh.h == nil {
		return 0, false
	}
	return lh.h.k, true
}

// DelMin -
func (lh *LeftistH) DelMin() {
	lh.h = lh.h.delMin()
	lh.size--
}

// Merge -
func (lh *LeftistH) Merge(lh1 *LeftistH) {
	lh.h = merge(lh.h, lh1.h)
	lh.size += lh1.size
}

// Size -
func (lh *LeftistH) Size() int {
	return lh.size
}

type node struct {
	k     int // priority key
	left  *node
	right *node
	npl   int
}

func (h *node) insert(h1 *node) *node {
	return merge(h, h1)
}

func (h *node) front() int {
	return h.k
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
	if h1.k < h2.k {
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
	if h.left.k < h.right.k {
		h.left, h.right = h.right, h.left
	}
	h.npl = h.right.npl + 1
}
