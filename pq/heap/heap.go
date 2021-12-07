package heap

type T interface{}

type Elem struct {
	p int
	v T
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

// NewBinHeapWitArray -
func NewBinHeapWitArray(arry []Elem, cap int) *Heap {
	if cap < 1 {
		panic("capacity less than 1")
	}
	h := &Heap{
		elems: make([]Elem, cap+1),
	}
	h.elems[0].p = -1 << 63
	if len(arry) >= cap {
		copy(h.elems[1:], arry[:cap])
		h.len = cap
	} else {
		copy(h.elems[1:], arry[:])
		h.len = len(arry)
	}
	for i := h.len / 2; i > 0; i-- {
		h.percolateDown(i)
	}
	return h
}

func (h *Heap) Size() int {
	return h.len
}

func (h *Heap) Cap() int {
	return len(h.elems) - 1
}

func (h *Heap) Push(p int, v T) {
	// insert into the bottom of heap
	h.len++
	if h.len < len(h.elems) {
		e := &h.elems[h.len]
		e.p = p
		e.v = v
	} else {
		h.elems = append(h.elems, Elem{p: p, v: v})
	}
	// percolate up from bottom
	h.percolateUp(h.len)
}

func (h *Heap) Pop() T {
	v := h.elems[1].v
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
		if h.elems[i].v == v {
			return i
		}
	}
	return -1
}

// Fix priority
func (h *Heap) Fix(p int, v T) {
	i := h.find(v)
	if i <= 0 {
		return
	}
	p0 := h.elems[i].p
	if p > p0 {
		h.elems[i].p = p
		h.percolateDown(i)
	} else if p < p0 {
		h.elems[i].p = p
		h.percolateUp(i)
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

func (h *Heap) percolateUp(vac int) {
	elem := h.elems[vac] // copy-out the element and left a vacancy
	p := vac / 2
	for p > 0 && h.elems[p].p > elem.p {
		h.elems[vac] = h.elems[p]
		vac = p
		p = vac / 2
	}
	h.elems[vac] = elem
}
