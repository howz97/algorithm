package basic

type LinkQueue[T any] struct {
	head *elem[T]
	tail *elem[T]
	size int
}

func NewLinkQueue[T any]() *LinkQueue[T] {
	return new(LinkQueue[T])
}

type elem[T any] struct {
	v    T
	next *elem[T]
}

func (q *LinkQueue[T]) Peek() *T {
	return &q.head.v
}

func (q *LinkQueue[T]) PopFront() T {
	e := q.head.v
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	q.size--
	return e
}

func (q *LinkQueue[T]) PushBack(e T) {
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

func (q *LinkQueue[T]) Size() int {
	return q.size
}
