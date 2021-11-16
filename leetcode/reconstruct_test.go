package leetcode

import (
	"fmt"
	"testing"
)

func Test_ReconstructBinTree(t *testing.T) {
	inorder := []Node{
		Node{Key: 9},
		Node{Key: 3},
		Node{Key: 15},
		Node{Key: 20},
		Node{Key: 7},
	}
	preorder := []Node{
		Node{Key: 3},
		Node{Key: 9},
		Node{Key: 20},
		Node{Key: 15},
		Node{Key: 7},
	}

	binTree := ReconstructBinTree(inorder, preorder)
	preOrder(binTree)
	fmt.Println()
	inOrder(binTree)
	fmt.Println()
	sufOrder(binTree)
	fmt.Println()
	levelOrder(binTree)
	fmt.Println()
}
