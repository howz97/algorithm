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

func (q *StrQ) Front() string {
	if q.head == nil {
		return ""
	}
	s := q.head.s
	if q.head.next == nil {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	return s
}

func (q *StrQ) PushBack(s string) {
	if q.tail == nil {
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
