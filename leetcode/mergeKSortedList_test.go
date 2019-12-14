package leetcode

import (
	"testing"
)

func TestMergeKList(t *testing.T) {
	lists := make([]*ListNode, 3)
	lists[0] = &ListNode{
		Val: 1,
	}
	lists[0].Next = &ListNode{
		Val: 4,
	}
	lists[0].Next.Next = &ListNode{
		Val: 5,
	}

	lists[1] = &ListNode{Val: 1}
	lists[1].Next = &ListNode{Val: 3}
	lists[1].Next.Next = &ListNode{Val: 4}

	lists[2] = &ListNode{Val: 2}
	lists[2].Next = &ListNode{Val: 6}

	printList(mergeKLists(lists))
}
