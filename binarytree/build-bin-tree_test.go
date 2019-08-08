package binarytree

import (
	"fmt"
	"testing"
)

func Test_BuildBinaryTree(t *testing.T) {
	inorder := []Node{
		Node{key: 9},
		Node{key: 3},
		Node{key: 15},
		Node{key: 20},
		Node{key: 7},
	}
	preorder := []Node{
		Node{key: 3},
		Node{key: 9},
		Node{key: 20},
		Node{key: 15},
		Node{key: 7},
	}

	binTree := BuildBinaryTree(inorder, preorder)
	fmt.Println(binTree)
	fmt.Println(binTree.leftSon)
	fmt.Println(binTree.rightSon)
}
