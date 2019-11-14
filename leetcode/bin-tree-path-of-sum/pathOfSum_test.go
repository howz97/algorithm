package pathofsum

import "testing"

func TestPathOfSum(t *testing.T) {
	leaf1 := &node{
		v: 1,
	}
	leaf2 := &node{
		v: 2,
	}
	leaf3 := &node{
		v: 3,
	}
	h1node1 := &node{
		v:     3,
		left:  leaf1,
		right: leaf2,
	}
	h1node2 := &node{
		v:     20,
		left:  leaf1,
		right: leaf2,
	}
	h1node3 := &node{
		v:     5,
		left:  leaf1,
		right: leaf3,
	}
	h1node4 := &node{
		v:     6,
		left:  leaf1,
		right: leaf2,
	}
	h2node1 := &node{
		v:     -11,
		left:  h1node1,
		right: h1node2,
	}
	h2node2 := &node{
		v:     4,
		left:  h1node3,
		right: h1node4,
	}
	tree := &node{
		v:     8,
		left:  h2node1,
		right: h2node2,
	}
	printPathOfSum(tree, 18, make([]int, 0, 3))
}
