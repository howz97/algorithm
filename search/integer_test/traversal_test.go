package integer

import (
	"fmt"
	"testing"

	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
)

func TestTraversal(t *testing.T) {
	data := []string{"26", "-", "5", "*", "3", "+", "2"}
	avl := avltree.New[int, string]()
	for i, v := range data {
		avl.Put(i, v)
	}
	search.PrintBinaryTree(avl)

	str := ""
	search.ReverseOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "2+3*5-26" {
		t.Errorf("ReverseOrder: %s", str)
	}

	str = ""
	search.InOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "26-5*3+2" {
		t.Errorf("InOrder: %s", str)
	}

	str = ""
	search.PreOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "*-265+32" {
		t.Errorf("PreOrder: %s", str)
	}

	str = ""
	search.SufOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "265-32+*" {
		t.Errorf("SufOrder: %s", str)
	}

	str = ""
	search.LevelOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "*-+26532" {
		t.Errorf("LevelOrder: %s", str)
	}
}
