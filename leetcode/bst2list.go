package leetcode

type node struct {
	v     int
	left  *node
	right *node
}

func bst2list(t *node) (head, tail *node) {
	if t == nil {
		return nil, nil
	}
	head, tail = t, t
	lHead, lTail := bst2list(t.left)
	if lHead != nil {
		head = lHead
		lTail.right = t
		t.left = lTail
	}
	rHead, rTail := bst2list(t.right)
	if rHead != nil {
		tail = rTail
		rHead.left = t
		t.right = rHead
	}
	return head, tail
}
