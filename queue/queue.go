package queue

type StrQ struct {
	head *elem
	tail *elem
}

func NewStrQ() *StrQ {
	return new(StrQ)
}

type elem struct {
	s    string
	next *elem
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
		q.head = &elem{
			s: s,
		}
		q.tail = q.head
	} else {
		q.tail.next = &elem{
			s: s,
		}
		q.tail = q.tail.next
	}
}

func (q *StrQ) IsEmpty() bool {
	return q.head == nil
}
