package queue

type LinkQ[T any] struct {
	head *elem[T]
	tail *elem[T]
	size int
}

func NewLinkQ[T any]() *LinkQ[T] {
	return new(LinkQ[T])
}

type elem[T any] struct {
	v    T
	next *elem[T]
}

func (q *LinkQ[T]) Peek() *T {
	return &q.head.v
}

func (q *LinkQ[T]) Front() T {
	e := q.head.v
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return e
}

func (q *LinkQ[T]) PushBack(e T) {
	q.size++
	if q.head == nil {
		q.head = &elem[T]{
			v: e,
		}
		q.tail = q.head
		return
	}
	q.tail.next = &elem[T]{
		v: e,
	}
	q.tail = q.tail.next
}

func (q *LinkQ[T]) Size() int {
	return q.size
}
