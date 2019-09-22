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
	if q.head == nil {
		return 0
	}
	i := q.head.i
	if q.head.next == nil {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	return i
}

func (q *IntQ) PushBack(i int) {
	if q.tail == nil {
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
