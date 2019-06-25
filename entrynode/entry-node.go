package entrynode

type node struct {
	next *node
}

func entryNode(head *node) *node {
	mnode := meetNode(head)
	node := head
	for node != mnode {
		node, mnode = node.next, mnode.next
	}
	return node
}

func meetNode(head *node) *node {
	if head == nil || head.next == nil || head.next.next == nil {
		return nil
	}
	slow := head.next
	fast := head.next.next

	for slow != fast {
		if fast.next == nil || fast.next.next == nil {
			return nil
		}
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
