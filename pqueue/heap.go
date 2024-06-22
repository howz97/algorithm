// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pqueue

import (
	"cmp"
)

type Elem[P cmp.Ordered, V any] struct {
	p   P // priority
	val V
}

type Paired[P cmp.Ordered, V any] struct {
	len   int // amount of elems
	elems []Elem[P, V]
}

// NewPaired make an empty heap. Specify a proper cap to reduce times of re-allocate elems slice
func NewPaired[P cmp.Ordered, V any](cap uint) *Paired[P, V] {
	h := &Paired[P, V]{
		elems: make([]Elem[P, V], cap+1), // ignore element at index 0
	}
	return h
}

func (h *Paired[P, V]) Size() int {
	return h.len
}

func (h *Paired[P, V]) Cap() int {
	return cap(h.elems) - 1
}

func (h *Paired[P, V]) PushPair(p P, v V) {
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

func (h *Paired[P, V]) PopPair() (P, V) {
	e := h.elems[1]
	// move bottom element to the top position
	h.elems[1] = h.elems[h.len]
	h.len--
	// percolate down the top element
	h.sink(1)
	return e.p, e.val
}

func (h *Paired[P, V]) Pop() V {
	_, v := h.PopPair()
	return v
}

func (h *Paired[P, V]) Top() V {
	return h.elems[1].val
}

func (h *Paired[P, V]) TopPair() (P, V) {
	return h.elems[1].p, h.elems[1].val
}

func (h *Paired[P, V]) FixTop(p P) {
	h.elems[1].p = p
	h.sink(1)
}

func (h *Paired[P, V]) sink(vac int) {
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

func (h *Paired[P, V]) swim(vac int) {
	elem := h.elems[vac] // copy-out the element and left a vacancy
	parent := vac / 2
	for parent > 0 && h.elems[parent].p > elem.p {
		h.elems[vac] = h.elems[parent]
		vac = parent
		parent = vac / 2
	}
	h.elems[vac] = elem
}

type Fixable[P cmp.Ordered, V comparable] struct {
	Paired[P, V]
}

func NewFixable[P cmp.Ordered, V comparable](cap uint) *Fixable[P, V] {
	return &Fixable[P, V]{
		Paired: *NewPaired[P, V](cap),
	}
}

func (h *Fixable[P, V]) Del(v V) {
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

func (h *Fixable[P, V]) find(v V) int {
	for i := 1; i <= h.len; i++ {
		if h.elems[i].val == v {
			return i
		}
	}
	return -1
}

// Fix priority
func (h *Fixable[P, V]) Fix(p P, v V) {
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

type Heap[P cmp.Ordered] struct {
	Paired[P, struct{}]
}

func NewHeap[P cmp.Ordered](cap uint) *Heap[P] {
	return &Heap[P]{
		Paired: *NewPaired[P, struct{}](cap),
	}
}

func (h *Heap[P]) Push(v P) {
	h.Paired.PushPair(v, struct{}{})
}

func (h *Heap[P]) Pop() P {
	p, _ := h.Paired.PopPair()
	return p
}

func (h *Heap[P]) Top() P {
	p, _ := h.Paired.TopPair()
	return p
}
