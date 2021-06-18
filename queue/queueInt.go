package queue

type IntQ struct {
	head *elemInt
	tail *elemInt
}

func NewIntQ() *IntQ {
	return new(IntQ)
}

type elemInt struct {
	i    int
	next *elemInt
}

func (q *IntQ) Front() int {
	if q.IsEmpty() {
		return 0
	}
	e := q.head.i
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return e
}

func (q *IntQ) PushBack(i int) {
	if q.IsEmpty() {
		q.head = &elemInt{
			i: i,
		}
		q.tail = q.head
	} else {
		q.tail.next = &elemInt{
			i: i,
		}
		q.tail = q.tail.next
	}
}

func (q *IntQ) IsEmpty() bool {
	return q.head == nil
}
