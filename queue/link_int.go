package queue

type LinkInt struct {
	Linked
}

func NewLinkInt() *LinkInt {
	return new(LinkInt)
}

func (q *LinkInt) Front() int {
	return q.Linked.Front().(int)
}

func (q *LinkInt) PushBack(i int) {
	q.Linked.PushBack(i)
}
