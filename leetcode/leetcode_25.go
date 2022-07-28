package leetcode

/**
 * Definition for singly-linked list.
 type ListNode struct {
    Val int
    Next *ListNode
}
*/

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	nextHead := head
	for i := 1; i <= k; i++ {
		nextHead = nextHead.Next
		if nextHead == nil {
			if i == k {
				return reverse(head, nil)
			}
			return head
		}
	}
	ret := reverse(head, nextHead)
	for {
		preTail := head
		head = nextHead
		for i := 1; i <= k; i++ {
			nextHead = nextHead.Next
			if nextHead == nil {
				if i == k {
					preTail.Next = reverse(head, nextHead)
				}
				return ret
			}
		}
		preTail.Next = reverse(head, nextHead)
	}
}

func reverse(head, nextHead *ListNode) *ListNode {
	to := nextHead
	for {
		tmp := head.Next
		head.Next = to
		if tmp == nextHead {
			break
		}
		to = head
		head = tmp
	}
	return head
}
