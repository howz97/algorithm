package integer

import (
	"fmt"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
	"github.com/howz97/algorithm/util"
	"testing"
)

func TestTraversal(t *testing.T) {
	data := []string{"26", "-", "5", "*", "3", "+", "2"}
	avl := avltree.New()
	for i, v := range data {
		avl.Put(util.Int(i), v)
	}
	search.PrintBinaryTree(avl)

	str := ""
	search.ReverseOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	t.Logf("ReverseOrder: %s", str)

	str = ""
	search.InOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	t.Logf("InOrder: %s", str)

	str = ""
	search.PreOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	t.Logf("PreOrder: %s", str)

	str = ""
	search.SufOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	t.Logf("SufOrder: %s", str)

	str = ""
	search.LevelOrder(avl, func(t search.ITraversal) bool {
		str += fmt.Sprint(t.String())
		return true
	})
	t.Logf("LevelOrder: %s", str)
}
