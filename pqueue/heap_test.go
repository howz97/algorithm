package pqueue

import (
	"testing"
)

func Test_Heap(t *testing.T) {
	h := NewBinHeap(10)
	if !h.Insert(1) ||
		!h.Insert(1) ||
		!h.Insert(2) ||
		!h.Insert(3) ||
		!h.Insert(6) ||
		!h.Insert(5) ||
		!h.Insert(4) ||
		!h.Insert(9) ||
		!h.Insert(8) ||
		!h.Insert(7) {
		t.Fatal("Insert failed")
	}
	if h.Insert(9) {
		t.Fatal("Insert failed")
	}
	if h.DelMin() != 1 {
		t.Fatal("DelMin failed")
	}
	for i := 1; i <= 9; i++ {
		if h.DelMin() != i {
			t.Fatal("DelMin failed")
		}
	}

	arry := []int{4, 5, 6, 3, 2, 1, 7, 8, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	h = NewBinHeapWitArray(arry, 9)
	for i := 1; i <= 9; i++ {
		if h.DelMin() != i {
			t.Fatal("DelMin failed")
		}
	}

	h = NewBinHeapWitArray(arry, 20)
	if h.Cap() != 20 || h.Size() != 18 {
		t.Fatal("NewBinHeapWitArray failed")
	}
	for i := 1; h.Size() != 0; i++ {
		h.DelMin()
	}
	if h.Size() != 0 {
		t.Fatal("DelMin failed")
	}
}

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
