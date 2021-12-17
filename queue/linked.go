package queue

import . "github.com/howz97/algorithm/util"

type Linked struct {
	head *elem
	tail *elem
	size int
}

func NewLinked() *Linked {
	return new(Linked)
}

type elem struct {
	v    T
	next *elem
}

func (q *Linked) Front() T {
	q.size--
	e := q.head.v
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return e
}

// PushBack can not insert a nil element
func (q *Linked) PushBack(e T) {
	q.size++
	if q.head == nil {
		q.head = &elem{
			v: e,
		}
		q.tail = q.head
		return
	}
	q.tail.next = &elem{
		v: e,
	}
	q.tail = q.tail.next
}

func (q *Linked) Size() int {
	return q.size
}
