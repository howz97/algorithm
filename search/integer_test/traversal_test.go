package integer

import (
	"fmt"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/avltree"
	"testing"
)

func TestTraversal(t *testing.T) {
	data := []string{"26", "-", "5", "*", "3", "+", "2"}
	avl := avltree.New()
	for i, v := range data {
		avl.Put(search.Integer(i), v)
	}
	search.PrintBinaryTree(avl)

	str := ""
	search.ReverseOrder(avl, func(_ search.Cmp, v search.T) bool {
		str += fmt.Sprint(v)
		return true
	})
	t.Logf("ReverseOrder: %s", str)

	str = ""
	search.InOrder(avl, func(_ search.Cmp, v search.T) bool {
		str += fmt.Sprint(v)
		return true
	})
	t.Logf("InOrder: %s", str)

	str = ""
	search.PreOrder(avl, func(_ search.Cmp, v search.T) bool {
		str += fmt.Sprint(v)
		return true
	})
	t.Logf("PreOrder: %s", str)

	str = ""
	search.SufOrder(avl, func(_ search.Cmp, v search.T) bool {
		str += fmt.Sprint(v)
		return true
	})
	t.Logf("SufOrder: %s", str)

	str = ""
	search.LevelOrder(avl, func(_ search.Cmp, v search.T) bool {
		str += fmt.Sprint(v)
		return true
	})
	t.Logf("LevelOrder: %s", str)
}
