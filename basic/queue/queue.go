package queue

const (
	MinCap = 4
)

type Queue[T any] struct {
	elems      []T
	head, back int
	size       int
}

func NewQueue[T any](cap int) *Queue[T] {
	if cap < MinCap {
		cap = MinCap
	}
	return &Queue[T]{
		elems: make([]T, cap),
	}
}

func QueueFrom[T any](elems []T) *Queue[T] {
	return &Queue[T]{
		elems: elems[:cap(elems)],
		back:  len(elems) - 1,
		size:  len(elems),
	}
}

func (q *Queue[T]) Peek() *T {
	if q.size <= 0 {
		return nil
	}
	return &q.elems[q.head]
}

func (q *Queue[T]) PopFront() T {
	if q.size <= 0 {
		panic("empty queue")
	}
	e := q.elems[q.head]
	q.head++
	if q.head == len(q.elems) {
		q.head = 0
	}
	q.size--
	return e
}

func (q *Queue[T]) PushBack(e T) {
	if q.size <= 0 {
		q.elems[0] = e
		q.head = 0
		q.back = 0
		q.size = 1
		return
	}
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

func (q *Queue[T]) isFull() bool {
	return q.size == len(q.elems)
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Clone() []T {
	elems := make([]T, 0, q.size)
	for q.size > 0 {
		elems = append(elems, q.PopFront())
	}
	return elems
}
