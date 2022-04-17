package integer

import (
	"fmt"
	"testing"

	"github.com/howz97/algorithm/search/avltree"
)

func TestTraversal(t *testing.T) {
	data := []string{"26", "-", "5", "*", "3", "+", "2"}
	avl := avltree.New[int, string]()
	for i, v := range data {
		avl.Put(i, v)
	}
	avl.Print()

	str := ""
	avl.ReverseOrder(func(t *avltree.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "2+3*5-26" {
		t.Errorf("ReverseOrder: %s", str)
	}

	str = ""
	avl.InOrder(func(t *avltree.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "26-5*3+2" {
		t.Errorf("InOrder: %s", str)
	}

	str = ""
	avl.PreOrder(func(t *avltree.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "*-265+32" {
		t.Errorf("PreOrder: %s", str)
	}

	str = ""
	avl.SufOrder(func(t *avltree.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "265-32+*" {
		t.Errorf("SufOrder: %s", str)
	}

	str = ""
	avl.LevelOrder(func(t *avltree.Node[int, string]) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	if str != "*-+26532" {
		t.Errorf("LevelOrder: %s", str)
	}
}
