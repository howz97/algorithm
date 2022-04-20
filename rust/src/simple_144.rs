use std::rc::Rc;
use std::cell::RefCell;
use std::collections::Vec;

impl Solution {
    pub fn preorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut order = Vec::new();
        let mut right = Vec::new();
        right.push(root);
        while right.len() > 0 {
            let mut opt_n = right.pop();
            while {
                if let Some(n) = opt_n {
                    order.push(n.key);
                    if n.right.is_none() {
                        right.push(n);
                    }
                    opt_n = n.left;
                } else {
                    break
                }
            }
        }
        return order
    }
}

// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
  pub val: i32,
  pub left: Option<Rc<RefCell<TreeNode>>>,
  pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
  #[inline]
  pub fn new(val: i32) -> Self {
    TreeNode {
      val,
      left: None,
      right: None
    }
  }
}