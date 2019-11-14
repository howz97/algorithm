package leetcode

import "testing"

func TestBst2list(t *testing.T) {
	leaf1 := &node{
		v: 1,
	}
	leaf2 := &node{
		v: 3,
	}
	leaf3 := &node{
		v: 7,
	}
	h1node1 := &node{
		v:     2,
		left:  leaf1,
		right: leaf2,
	}
	h1node2 := &node{
		v:     6,
		right: leaf3,
	}
	h2node1 := &node{
		v:     4,
		left:  h1node1,
		right: h1node2,
	}
	tree := &node{
		v:    8,
		left: h2node1,
	}
	head, tail := bst2list(tree)
	for head != nil {
		print(head.v)
		head = head.right
	}
	println()
	for tail != nil {
		print(tail.v)
		tail = tail.left
	}
	println()
}
