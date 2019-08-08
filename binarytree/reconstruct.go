package binarytree

// Node -
type Node struct {
	key      int
	leftSon  *Node
	rightSon *Node
}

var rootOffSet = -1

/*
根据二叉树的前序遍历和中序遍历的结果，重建出该二叉树。
假设输入的前序遍历和中序遍历的结果中都不含重复的数字。
https://github.com/CyC2018/CS-Notes/blob/master/notes/%E5%89%91%E6%8C%87%20Offer%20%E9%A2%98%E8%A7%A3%20-%203~9.md#7-%E9%87%8D%E5%BB%BA%E4%BA%8C%E5%8F%89%E6%A0%91
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
		if inorder[i].key == root.key { // 要求每个节点key唯一
			break
		}
	}
	root.leftSon = ReconstructBinTree(inorder[:i], preorder)
	root.rightSon = ReconstructBinTree(inorder[i+1:], preorder)

	return &root
}
