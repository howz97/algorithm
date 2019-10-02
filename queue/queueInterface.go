package queue

type LinkedQueue struct {
	head *elemInterface
	tail *elemInterface
}

func NewLinkedQueue() *LinkedQueue {
	return new(LinkedQueue)
}

type elemInterface struct {
	intf interface{}
	next *elemInterface
}

// Front return a nil when queue has been empty
func (q *LinkedQueue) Front() interface{} {
	if q.IsEmpty() {
		return nil
	}
	e := q.head.intf
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return e
}

// PushBack can not insert a nil element
func (q *LinkedQueue) PushBack(e interface{}) error {
	if e == nil {
		return ErrPushBackNilElem
	}
	if q.IsEmpty() {
		q.head = &elemInterface{
			intf: e,
		}
		q.tail = q.head
	} else {
		q.tail.next = &elemInterface{
			intf: e,
		}
		q.tail = q.tail.next
	}
	return nil
}

func (q *LinkedQueue) IsEmpty() bool {
	return q.head == nil
}
