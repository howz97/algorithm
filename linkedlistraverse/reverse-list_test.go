package linkedlistreverse

import "testing"

func Test_printListReverse(t *testing.T) {
	head := &node{
		v: "i",
	}
	node2 := &node{
		v: "am",
	}
	head.next = node2
	node3 := &node{
		v: "testing",
	}
	node2.next = node3
	node4 := &node{
		v: "code",
	}
	node3.next = node4

	printListReverse(head)
	printListReverse(node4.next)
}
