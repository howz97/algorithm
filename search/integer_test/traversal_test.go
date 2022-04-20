package integer

import (
	"fmt"
	"testing"

	"github.com/howz97/algorithm/search/avlst"
	"github.com/howz97/algorithm/search/binarytree"
)

func TestTraversal(t *testing.T) {
	data := []string{"26", "-", "5", "*", "3", "+", "2"}
	avl := avlst.New[int, string]()
	for i, v := range data {
		avl.Put(i, v)
	}
	avl.Print()

	str := ""
	avl.ReverseOrder(func(t *avlst.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "2+3*5-26" {
		t.Errorf("ReverseOrder: %s", str)
	}

	str = ""
	avl.InOrder(func(t *avlst.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "26-5*3+2" {
		t.Errorf("InOrder: %s", str)
	}

	str = ""
	avl.PreOrder(func(t *avlst.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "*-265+32" {
		t.Errorf("PreOrder: %s", str)
	}

	str = ""
	avl.SufOrder(func(t *avlst.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "265-32+*" {
		t.Errorf("SufOrder: %s", str)
	}

	str = ""
	avl.LevelOrder(func(t *avlst.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "*-+26532" {
		t.Errorf("LevelOrder: %s", str)
	}
}

func ExamplePreOrderIter() {
	bt := binarytree.New[int, int]()
	nodes := []int{50, 20, 10, 15, 30, 25, 40, 35, 45, 60, 55, 70}
	for _, v := range nodes {
		bt.Put(v, v)
	}
	bt.PreOrder(func(n *binarytree.Node[int, int]) bool {
		fmt.Print(n.Key(), ", ")
		return true
	})

	// Output: 50, 20, 10, 15, 30, 25, 40, 35, 45, 60, 55, 70,
}
