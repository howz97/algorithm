package pqueue

import (
	"testing"
)

func Test_Heap(t *testing.T) {
	h := NewBinHeap(10)
	if !h.Insert(1, 1) ||
		!h.Insert(1, 1) ||
		!h.Insert(2, 2) ||
		!h.Insert(3, 3) ||
		!h.Insert(6, 6) ||
		!h.Insert(5, 5) ||
		!h.Insert(4, 4) ||
		!h.Insert(9, 9) ||
		!h.Insert(8, 8) ||
		!h.Insert(7, 7) {
		t.Fatal("Insert failed")
	}
	if h.Insert(9, 9) {
		t.Fatal("Insert failed")
	}
	if _, m := h.DelMin(); m != 1 {
		t.Fatal("DelMin failed")
	}
	for i := 1; i <= 9; i++ {
		if _, m := h.DelMin(); m != i {
			t.Fatal("DelMin failed")
		}
	}
}
