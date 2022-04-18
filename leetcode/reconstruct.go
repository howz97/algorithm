package leetcode

var rootOffSet = -1

type Node struct {
	Key      int
	leftSon  *Node
	rightSon *Node
}

/*
根据二叉树的前序遍历和中序遍历的结果，重建出该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
*/

// ReconstructBinTree -
func ReconstructBinTree(inorder, preorder []Node) *Node {
	if len(inorder) == 0 {
		return nil
	}
	rootOffSet++ // 当前inorder中根结点在preorder中的位置

	root := preorder[rootOffSet]

	i := 0
	for i = range inorder {
		if inorder[i].Key == root.Key { // 要求每个节点key唯一
			break
		}
	}
	root.leftSon = ReconstructBinTree(inorder[:i], preorder)
	root.rightSon = ReconstructBinTree(inorder[i+1:], preorder)

	return &root
}
