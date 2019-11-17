package leetcode

// https://github.com/CyC2018/CS-Notes/blob/master/notes/54.%20%E4%BA%8C%E5%8F%89%E6%9F%A5%E6%89%BE%E6%A0%91%E7%9A%84%E7%AC%AC%20K%20%E4%B8%AA%E7%BB%93%E7%82%B9.md
func bstKthNode(tree *node, k int) *node {
	if k <= 0 {
		return nil
	}
	k2 := k
	return bstKthNodeCore(tree, &k2)
}

func bstKthNodeCore(root *node, pk *int) *node {
	if root == nil {
		return nil
	}
	target := bstKthNodeCore(root.left, pk)
	if target != nil {
		return target
	}
	*pk--
	if *pk == 0 {
		return root
	}
	return bstKthNodeCore(root.right, pk)
}
