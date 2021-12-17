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
}

func (q *Slice) isFull() bool {
	return q.Size() == len(q.elems)
}

func (q *Slice) Size() (size int) {
	if q.head < q.back {
		size = q.back - q.head + 1
	} else if q.head > q.back {
		size = len(q.elems) - (q.head - q.back - 1)
	} else {
		size = 0
	}
	return
}
