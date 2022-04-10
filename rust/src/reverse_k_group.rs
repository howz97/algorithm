// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
  pub val: i32,
  pub next: Option<Box<ListNode>>
}

// impl ListNode {
//   #[inline]
//   fn new(val: i32) -> Self {
//     ListNode {
//       next: None,
//       val
//     }
//   }
// }

/*
作者：navy_d
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/rustdi-gui-javadi-gui-by-navy_d/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

pub struct RecursionSolution {}

impl RecursionSolution {
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

/*
作者：yescallop
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group/solution/die-dai-unsafe-zhi-zhen-jiao-huan-by-yescallop/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

unsafe fn swap<T>(x: *mut T, y: *mut T, t: *mut T) {
    use std::ptr::copy_nonoverlapping as copy;
    copy(x, t, 1);
    copy(y, x, 1);
    copy(t, y, 1);
}

pub fn reverse_k_group(mut head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
    if k <= 1 || head.is_none() { return head }
    let mut p: *mut _ = &mut head;
    let mut last = [std::ptr::null_mut(); 2];
    let mut t = std::mem::MaybeUninit::uninit();
    let t = t.as_mut_ptr();
    'outer: loop { unsafe {
        let start = p;
        let mut cur = p;
        for i in 0..k {
            cur = if let Some(cur) = (*cur).as_mut() {
                &mut cur.next
            } else { break 'outer };
            if i == 0 { p = cur }
            if i >= k - 2 { last[(i - k + 2) as usize] = cur }
        }
        cur = start;
        let mut i = 0;
        for _ in 0..k - 1 {
            let next = &mut (*cur).as_mut().unwrap().next;
            swap(cur, last[i], t);
            cur = next;
            i ^= 1;
        }
        if i == 1 { swap(last[0], last[1], t) }
    }}
    head
}