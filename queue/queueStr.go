package queue

type StrQ struct {
	head *elemStr
	tail *elemStr
}

func NewStrQ() *StrQ {
	return new(StrQ)
}

type elemStr struct {
	s    string
	next *elemStr
}

func (q *StrQ) Front() (string, error) {
	if q.IsEmpty() {
		return "", ErrEmptyQ
	}
	e := q.head.s
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return e, nil
}

func (q *StrQ) PushBack(s string) {
	if q.IsEmpty() {
		q.head = &elemStr{
			s: s,
		}
		q.tail = q.head
	} else {
		q.tail.next = &elemStr{
			s: s,
		}
		q.tail = q.tail.next
	}
}

func (q *StrQ) IsEmpty() bool {
	return q.head == nil
}
