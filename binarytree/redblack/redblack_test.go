package	redblack

import "testing"

func Test_RedBlack(t *testing.T) {
	redblack := New()
	for i := 1; i <= 7; i++ {
		redblack.Insert(i, i)
	}
	for i := 16; i >= 10; i-- {
		redblack.Insert(i, i)
	}
	redblack.Insert(8, 8)
	redblack.Insert(9, 9)

	if redblack.FindMin().(int) != 1 {
		t.Fatal("FindMin failed")
	}
	if redblack.FindMax().(int) != 16 {
		t.Fatal("FindMax failed")
	}
	for i := 1; i <= 16; i++ {
		if redblack.Find(i).(int) != i {
			t.Fatal("Find failed")
		}
	}
}
