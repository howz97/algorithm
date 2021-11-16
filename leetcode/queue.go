package leetcode

type queue struct {
	head *linkNode
	tail *linkNode
}

type linkNode struct {
	data *Node
	next *linkNode
}

func (q *queue) front() *Node {
	if q.head == nil {
		return nil
	}
	n := q.head.data
	if q.head.next == nil {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	return n
}

func (q *queue) pushBack(data *Node) {
	if q.tail == nil {
		q.head = &linkNode{
			data: data,
		}
		q.tail = q.head
	} else {
		q.tail.next = &linkNode{
			data: data,
		}
		q.tail = q.tail.next
	}
}

func (q *queue) empty() bool {
	return q.head == nil
}
