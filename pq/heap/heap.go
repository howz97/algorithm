package heap

import (
	"golang.org/x/exp/constraints"
)

type Elem[P constraints.Ordered, V any] struct {
	p   P // priority
	val V
}

type Heap[P constraints.Ordered, V any] struct {
	len   int // amount of elems
	elems []Elem[P, V]
}

// New make an empty heap. Specify a proper cap to reduce times of re-allocate elems slice
func New[P constraints.Ordered, V any](cap uint) *Heap[P, V] {
	h := &Heap[P, V]{
		elems: make([]Elem[P, V], cap+1), // ignore element at index 0
	}
	return h
}

func (h *Heap[P, V]) Size() int {
	return h.len
}

func (h *Heap[P, V]) Cap() int {
	return cap(h.elems) - 1
}

func (h *Heap[P, V]) Push(p P, v V) {
	// insert into the bottom of heap
	h.len++
	if h.len < len(h.elems) {
		e := &h.elems[h.len]
		e.p = p
		e.val = v
	} else {
		h.elems = append(h.elems, Elem[P, V]{p: p, val: v})
	}
	// percolate up from bottom
	h.swim(h.len)
}

func (h *Heap[P, V]) Pop() V {
	v := h.elems[1].val
	// move bottom element to the top position
	h.elems[1] = h.elems[h.len]
	h.len--
	// percolate down the top element
	h.sink(1)
	return v
}

func (h *Heap[P, V]) sink(vac int) {
	elem := h.elems[vac] // copy-out the element and left a vacancy
	sub := vac * 2
	for sub <= h.len {
		// choose the smaller sub-node
		if sub < h.len && h.elems[sub+1].p < h.elems[sub].p {
			sub++
		}
		if h.elems[sub].p >= elem.p {
			break
		}
		h.elems[vac] = h.elems[sub]
		vac = sub
		sub = vac * 2
	}
	h.elems[vac] = elem
}

func (h *Heap[P, V]) swim(vac int) {
	elem := h.elems[vac] // copy-out the element and left a vacancy
	parent := vac / 2
	for parent > 0 && h.elems[parent].p > elem.p {
		h.elems[vac] = h.elems[parent]
		vac = parent
		parent = vac / 2
	}
	h.elems[vac] = elem
}
