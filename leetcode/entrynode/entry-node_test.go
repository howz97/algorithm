package entrynode

import "testing"

func Test_entryNode(t *testing.T) {
	head := new(node)

	node1 := new(node)
	head.next = node1

	node2 := new(node)
	node1.next = node2

	node3 := new(node)
	node2.next = node3

	node4 := new(node)
	node3.next = node4

	node4.next = node2

	enode := entryNode(head)
	if enode != node2 {
		t.Fail()
	}
	println(enode)
}
