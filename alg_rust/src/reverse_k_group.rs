/*
作者：navy_d
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/rustdi-gui-javadi-gui-by-navy_d/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

impl ListNode {
  #[inline]
  fn new(val: i32) -> Self {
    ListNode {
      next: None,
      val
    }
  }
}

pub struct Solution {}

impl Solution {
    pub fn reverse_k_group(mut head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
        let mut next_head = &mut head;
        for _ in 0..k {
            if let Some(node) = next_head.as_mut() {
                next_head = &mut node.next;
            } else {
                return head;
            }
        }
        let next_head = Self::reverse_k_group(next_head.take(), k);
        Self::reverse(head, next_head)
    }

    // head -> ... -> tail -x-> next_head -> ... 反转链接 ... <- next_head <- head <- ... <- tail 并返回new_head(tail)
    fn reverse(mut head: Option<Box<ListNode>>, mut next_head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        while let Some(mut node) = head {
            head = node.next.take();
            // link head -> next_head
            node.next = next_head.take();
            // as tail
            next_head = Some(node);
        }
        next_head
    }
}

// pub fn reverse_k_group(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
//     if k <= 1 {
//         return head
//     }
//     let mut length = 2;
//     let mut left_left: &mut Option<Box<ListNode>> = &mut None;
//     let mut left = &mut head;
//     let mut right = match left {
//         Some(l) => &mut l.next,
//         None => &None
//     };
//     let mut right_right: &mut Option<Box<ListNode>> = &mut None;
//     loop {
//         while length < k {
//             right = match right {
//                 Some(mut b) => &mut b.next,
//                 None => return head
//             };
//             length += 1;
//         }
//         right_right = match right {
//             Some(b) => &mut b.next,
//             None => return head
//         };
//         match left_left {
//             Some(b) => {b.next = right.take()},
//             None => ()
//         }

//     }
//     None
// }