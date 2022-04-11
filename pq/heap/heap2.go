package heap

import "golang.org/x/exp/constraints"

type Heap2[P constraints.Ordered, V comparable] struct {
	Heap[P, V]
}

func New2[P constraints.Ordered, V comparable](cap uint) *Heap2[P, V] {
	return &Heap2[P, V]{
		Heap: *New[P, V](cap),
	}
}

func (h *Heap2[P, V]) Del(v V) {
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
	h.sink(i)
}

func (h *Heap2[P, V]) find(v V) int {
	for i := 1; i <= h.len; i++ {
		if h.elems[i].val == v {
			return i
		}
	}
	return -1
}

// Fix priority
func (h *Heap2[P, V]) Fix(p P, v V) {
	i := h.find(v)
	if i <= 0 {
		return
	}
	if p < h.elems[i].p {
		h.elems[i].p = p
		h.swim(i)
	} else if p > h.elems[i].p {
		h.elems[i].p = p
		h.sink(i)
	}
}
