package heap

import (
	. "github.com/howz97/algorithm/util"
)

type Elem struct {
	p   Comparable // priority
	val T
}

type Heap struct {
	len   int // amount of elems
	elems []Elem
}

func New(cap uint) *Heap {
	h := &Heap{
		elems: make([]Elem, cap+1), // ignore element at index 0
	}
	return h
}

func (h *Heap) Size() int {
	return h.len
}

func (h *Heap) Cap() int {
	return len(h.elems) - 1
}

func (h *Heap) Push(p Comparable, v T) {
	// insert into the bottom of heap
	h.len++
	if h.len < len(h.elems) {
		e := &h.elems[h.len]
		e.p = p
		e.val = v
	} else {
		h.elems = append(h.elems, Elem{p: p, val: v})
	}
	// percolate up from bottom
	h.percolateUp(h.len)
}

func (h *Heap) Pop() T {
	v := h.elems[1].val
	// move bottom element to the top position
	h.elems[1] = h.elems[h.len]
	h.len--
	// percolate down the top element
	h.percolateDown(1)
	return v
}

func (h *Heap) Del(v T) {
	i := h.find(v)
	if i <= 0 {
		return
	}
	if i == h.len {
		h.len--
		return
	}
	h.elems[i] = h.elems[h.len]
	h.len--
	h.percolateDown(i)
}

func (h *Heap) find(v T) int {
	for i := 1; i <= h.len; i++ {
		if h.elems[i].val == v {
			return i
		}
	}
	return -1
}

// Fix priority
func (h *Heap) Fix(p Comparable, v T) {
	i := h.find(v)
	if i <= 0 {
		return
	}
	switch p.Cmp(h.elems[i].p) {
	case Less:
		h.elems[i].p = p
		h.percolateUp(i)
	case More:
		h.elems[i].p = p
		h.percolateDown(i)
	}
}

func (h *Heap) IsEmpty() bool {
	return h.len == 0
}

func (h *Heap) percolateDown(vac int) {
	elem := h.elems[vac] // copy-out the element and left a vacancy
	sub := vac * 2
	for sub <= h.len {
		// choose the smaller sub-node
		if sub < h.len && h.elems[sub+1].p.Cmp(h.elems[sub].p) == Less {
			sub++
		}
		if h.elems[sub].p.Cmp(elem.p) != Less {
			break
		}
		h.elems[vac] = h.elems[sub]
		vac = sub
		sub = vac * 2
	}
	h.elems[vac] = elem
}

func (h *Heap) percolateUp(vac int) {
	elem := h.elems[vac] // copy-out the element and left a vacancy
	parent := vac / 2
	for parent > 0 && h.elems[parent].p.Cmp(elem.p) == More {
		h.elems[vac] = h.elems[parent]
		vac = parent
		parent = vac / 2
	}
	h.elems[vac] = elem
}
