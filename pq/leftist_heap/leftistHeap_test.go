package pqueue

import (
	"testing"
)

func Test_LeftistHeap(t *testing.T) {
	lh1 := NewLeftistH()
	for i := 0; i < 10; i++ {
		lh1.Insert(i)
	}

	lh2 := NewLeftistH()
	for i := 10; i < 20; i++ {
		lh2.Insert(i)
	}
	lh1.Merge(lh2)
	if lh1.Size() != 20 {
		t.Fatal()
	}

	for i := 0; i < 20; i++ {
		k, ok := lh1.Front()
		if !ok {
			t.Fatal()
		}
		if k != i {
			println(i, k)
			t.Fatal()
		}
		lh1.DelMin()
	}
}
