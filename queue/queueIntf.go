package queue

import "errors"

const (
	MinCap     = 2
	DefaultCap = 10
)

var (
	ErrEmptyQ          = errors.New("queue has been empty")
	ErrFullQ           = errors.New("queue has been full")
	ErrPushBackNilElem = errors.New("can not push back nil element")
)

// Queue can not dynamically expand. Implemented by slice
type Queen struct {
	elems           []interface{}
	head, back, cap int
}

func NewQueen(cap int) *Queen {
	if cap < MinCap {
		cap = DefaultCap
	}
	return &Queen{
		elems: make([]interface{}, cap),
		cap:   cap,
	}
}

func (q *Queen) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, ErrEmptyQ
	}
	e := q.elems[q.head]
	q.elems[q.head] = nil
	q.head++
	if q.head >= q.cap {
		q.head -= q.cap
	}
	return e, nil
}

// PushBack can not insert a nil element
func (q *Queen) PushBack(e interface{}) error {
	if e == nil {
		return ErrPushBackNilElem
	}
	if q.IsFull() {
		return ErrFullQ
	}
	q.elems[q.back] = e
	q.back++
	if q.back >= q.cap {
		q.back -= q.cap
	}
	return nil
}

func (q *Queen) IsEmpty() bool {
	return q.elems[q.head] == nil
}

func (q *Queen) IsFull() bool {
	return q.elems[q.back] != nil
}
