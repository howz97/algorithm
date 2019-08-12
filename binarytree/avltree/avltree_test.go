package avltree

import "testing"

func Test_AVL(t *testing.T) {
	avlTree := New()
	for i := 1; i <= 7; i++ {
		avlTree.Insert(i, i)
	}
	for i := 16; i >= 10; i-- {
		avlTree.Insert(i, i)
	}
	avlTree.Insert(8, 8)
	avlTree.Insert(9, 9)

	if avlTree.FindMin().(int) != 1 {
		t.Fatal("FindMin failed")
	}
	if avlTree.FindMax().(int) != 16 {
		t.Fatal("FindMax failed")
	}
	for i := 1; i <= 16; i++ {
		if avlTree.Find(i).(int) != i {
			t.Fatal("Find failed")
		}
		avlTree.Delete(i)
		if avlTree.Find(i) != nil {
			t.Fatal("Delete failed")
		}
	}
}
