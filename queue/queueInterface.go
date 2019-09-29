package queue

type InterfaceQ struct {
	head *elemInterface
	tail *elemInterface
}

func NewInterfaceQ() *InterfaceQ {
	return new(InterfaceQ)
}

type elemInterface struct {
	intf interface{}
	next *elemInterface
}

// Front return a nil when queue has been empty
func (q *InterfaceQ) Front() interface{} {
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
func (q *InterfaceQ) PushBack(e interface{}) error {
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

func (q *InterfaceQ) IsEmpty() bool {
	return q.head == nil
}
