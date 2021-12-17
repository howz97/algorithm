package queue

import (
	. "github.com/howz97/algorithm/util"
)

const (
	MinCap = 2
)

type Slice struct {
	elems      []T
	head, back int
	size       int
}

func NewSlice(cap int) *Slice {
	if cap < MinCap {
		cap = MinCap
	}
	return &Slice{
		elems: make([]T, cap),
	}
}

func (q *Slice) Front() T {
	e := q.elems[q.head]
	q.head++
	if q.head == len(q.elems) {
		q.head = 0
	}
	q.size--
	return e
}

func (q *Slice) PushBack(e T) {
	if q.isFull() {
		expand := make([]T, 2*len(q.elems))
		n := q.Size()
		for i := 0; i < n; i++ {
			expand[i] = q.elems[(q.head+i)%len(q.elems)]
		}
		q.elems = expand
		q.head = 0
		q.back = n - 1
	}
	q.back++
	if q.back == len(q.elems) {
		q.back = 0
	}
	q.elems[q.back] = e
	q.size++
}

func (q *Slice) isFull() bool {
	return q.Size() == len(q.elems)
}

func (q *Slice) Size() int {
	return q.size
}
