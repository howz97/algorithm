package binaryheap

// Heap -
type Heap struct {
	size int
	arry []int
}

// New -
func New(cap int) *Heap {
	h := &Heap{
		arry: make([]int, cap+1),
	}
	h.arry[0] = -1 << 63
	return h
}

// NewWitArray -
func NewWitArray(arry []int, cap int) *Heap {
	h := &Heap{
		arry: make([]int, cap+1),
	}
	h.arry[0] = -1 << 63
	if len(arry) >= cap {
		copy(h.arry[1:], arry[:cap])
		h.size = cap
	} else {
		copy(h.arry[1:], arry[:])
		h.size = len(arry)
	}
	for i := h.size / 2; i > 0; i-- {
		h.percolateDown(i)
	}
	return h
}

// Size return the total amount of element in heap
func (h *Heap) Size() int {
	return h.size
}

// Cap return the upper limit of the number of element in heap
func (h *Heap) Cap() int {
	return len(h.arry) - 1
}

// Insert -
func (h *Heap) Insert(k int) (ok bool) {
	if h.size >= h.Cap() {
		return false
	}
	h.size++
	h.arry[h.size] = k
	h.percolateUp(h.size)
	return true
}

// DelMin -
func (h *Heap) DelMin() int {
	del := h.arry[1]
	h.arry[1] = h.arry[h.size]
	h.size--
	h.percolateDown(1)
	return del
}

func (h *Heap) percolateDown(i int) {
	arry := h.arry
	k := arry[i]
	cavIdx := i
	for {
		if cavIdx*2 > h.size {
			break
		}
		smallC := cavIdx * 2
		if smallC != h.size && arry[smallC+1] < arry[smallC] {
			smallC++
		}
		if arry[smallC] > k {
			break
		}
		arry[cavIdx] = arry[smallC]
		cavIdx = smallC
	}
	arry[cavIdx] = k
}

func (h *Heap) percolateUp(i int) {
	arry := h.arry
	k := arry[i]
	cavIdx := i
	for ; arry[cavIdx/2] > k; cavIdx /= 2 {
		h.arry[cavIdx] = h.arry[cavIdx/2]
	}
	arry[cavIdx] = k
}
