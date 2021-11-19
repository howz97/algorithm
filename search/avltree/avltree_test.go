package avltree

import (
	"github.com/howz97/algorithm/search"
	"testing"
)

func Test_AVL(t *testing.T) {
	avlTree := New()
	for i := 1; i <= 7; i++ {
		avlTree.Insert(search.Integer(i), i)
	}
	for i := 16; i >= 10; i-- {
		avlTree.Insert(search.Integer(i), i)
	}
	avlTree.Insert(search.Integer(8), 8)
	avlTree.Insert(search.Integer(9), 9)

	if avlTree.FindMin().(int) != 1 {
		t.Fatal("FindMin failed")
	}
	if avlTree.FindMax().(int) != 16 {
		t.Fatal("FindMax failed")
	}
	for i := 1; i <= 16; i++ {
		if avlTree.Find(search.Integer(i)).(int) != i {
			t.Fatal("Find failed")
		}
		avlTree.Delete(search.Integer(i))
		if avlTree.Find(search.Integer(i)) != nil {
			t.Fatal("Delete failed")
		}
	}
}
