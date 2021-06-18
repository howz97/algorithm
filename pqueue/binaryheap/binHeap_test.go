package pqueue

import (
	"fmt"
	"testing"
	"time"
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
	fmt.Println(h.find(5))
	if m := h.DelMin(); m != 1 {
		t.Fatal("DelMin failed")
	}
	for i := 1; i <= 9; i++ {
		if m := h.DelMin(); m != i {
			t.Fatal("DelMin failed")
		}
	}
}

func TestBinHeap_Delete(t *testing.T) {
	h := NewBinHeap(9)
	t1 := time.Now()
	t2 := time.Now()
	t3 := time.Now()
	t4 := time.Now()
	h.Insert(1, &t1)
	h.Insert(2, &t2)
	h.Insert(3, &t3)
	h.Insert(4, &t4)
	fmt.Println(h.find(&t3))
}
