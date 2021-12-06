package pqueue

import "container/heap"

type Elem struct {
	p int
	v interface{}
}

// BinHeap -
type BinHeap struct {
	size  int
	elems []Elem
}

// NewBinHeap -
func NewBinHeap(cap uint) *BinHeap {
	if cap < 1 {
		panic("capacity less than 1")
	}
	h := &BinHeap{
		elems: make([]Elem, cap+1),
	}
	h.elems[0].p = -1 << 63
	return h
}

// NewBinHeapWitArray -
func NewBinHeapWitArray(arry []Elem, cap int) *BinHeap {
	if cap < 1 {
		panic("capacity less than 1")
	}
	h := &BinHeap{
		elems: make([]Elem, cap+1),
	}
	h.elems[0].p = -1 << 63
	if len(arry) >= cap {
		copy(h.elems[1:], arry[:cap])
		h.size = cap
	} else {
		copy(h.elems[1:], arry[:])
		h.size = len(arry)
	}
	for i := h.size / 2; i > 0; i-- {
		h.percolateDown(i)
	}
	return h
}

// Size return the total amount of element in heap
func (h *BinHeap) Size() int {
	heap.Init()
	return h.size
}

// Cap return the upper limit of the number of element in heap
func (h *BinHeap) Cap() int {
	return len(h.elems) - 1
}

func (h *BinHeap) Insert(p int, v interface{}) (ok bool) {
	if h.size >= h.Cap() {
		return false
	}
	h.size++
	h.elems[h.size] = Elem{p: p, v: v}
	h.percolateUp(h.size)
	return true
}

func (h *BinHeap) DelMin() interface{} {
	if h.IsEmpty() {
		panic("delete from empty heap")
	}
	del := h.elems[1]
	h.elems[1] = h.elems[h.size]
	h.size--
	h.percolateDown(1)
	return del.v
}

func (h *BinHeap) Delete(v interface{}) {
	i := h.find(v)
	if i == -1 {
		return
	}
	h.elems[i] = h.elems[h.size]
	h.size--
	h.percolateDown(i)
}

func (h *BinHeap) find(v interface{}) int {
	for i, e := range h.elems {
		if e.v == v {
			return i
		}
	}
	return -1
}

func (h *BinHeap) Update(p int, v interface{}) {
	i := h.find(v)
	if i == -1 {
		return
	}
	if p > h.elems[i].p {
		h.elems[i].p = p
		h.percolateDown(i)
	} else {
		h.elems[i].p = p
		h.percolateUp(i)
	}
}

func (h *BinHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *BinHeap) IsFull() bool {
	return h.size == h.Cap()
}

func (h *BinHeap) percolateDown(i int) {
	arry := h.elems
	downingElem := arry[i]
	cavPointer := i
	for {
		if cavPointer*2 > h.size {
			break
		}
		smallC := cavPointer * 2
		if smallC != h.size && arry[smallC+1].p < arry[smallC].p {
			smallC++
		}
		if arry[smallC].p > downingElem.p {
			break
		}
		arry[cavPointer] = arry[smallC]
		cavPointer = smallC
	}
	arry[cavPointer] = downingElem
}

func (h *BinHeap) percolateUp(i int) {
	arry := h.elems
	uppingElem := arry[i]
	cavPointer := i
	for ; arry[cavPointer/2].p > uppingElem.p; cavPointer /= 2 {
		h.elems[cavPointer] = h.elems[cavPointer/2]
	}
	arry[cavPointer] = uppingElem
}
