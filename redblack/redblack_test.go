package redblack

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
		t.Fatal()
	}
	if redblack.FindMax().(int) != 16 {
		t.Fatal()
	}
	for i := 1; i <= 16; i++ {
		if redblack.Find(i).(int) != i {
			t.Fatal()
		}
	}

	// delMin
	for i := 2; i <= 16; i++ {
		redblack.DelMin()
		if redblack.FindMin().(int) != i {
			t.Fatal()
		}
	}
	redblack.DelMin()
	if !redblack.Empty() {
		t.Fatal()
	}

	// delMax
	for i := 1; i <= 3; i++ {
		redblack.Insert(i, i)
	}
	for i := 7; i <= 9; i++ {
		redblack.Insert(i, i)
	}
	for i := 6; i >= 4; i-- {
		redblack.Insert(i, i)
	}
	for i := 12; i >= 10; i-- {
		redblack.Insert(i, i)
	}

	for i := 11; i >= 1; i-- {
		redblack.DelMax()
		if redblack.FindMax().(int) != i {
			println()
			t.Fatal()
		}
	}
	redblack.DelMin()
	if !redblack.Empty() {
		t.Fatal()
	}

	// delete
	for i := 1; i <= 3; i++ {
		redblack.Insert(i, i)
	}
	for i := 7; i <= 9; i++ {
		redblack.Insert(i, i)
	}
	for i := 6; i >= 4; i-- {
		redblack.Insert(i, i)
	}
	for i := 12; i >= 10; i-- {
		redblack.Insert(i, i)
	}
	for i := 1; i <= 11; i++ {
		redblack.Delete(i)
		if redblack.FindMin().(int) != i+1 {
			t.Fatal()
		}
	}
	redblack.Delete(12)
	if !redblack.Empty() {
		t.Fatal()
	}
}
