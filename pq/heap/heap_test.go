package heap

import (
	"fmt"
	"github.com/howz97/algorithm/util"
	"testing"
	"time"
)

func Test_Heap(t *testing.T) {
	h := New(10)
	h.Push(util.Integer(1), 1)
	h.Push(util.Integer(1), 1)
	h.Push(util.Integer(2), 2)
	h.Push(util.Integer(3), 3)
	h.Push(util.Integer(6), 6)
	h.Push(util.Integer(5), 5)
	h.Push(util.Integer(4), 4)
	h.Push(util.Integer(9), 9)
	h.Push(util.Integer(8), 8)
	h.Push(util.Integer(7), 7)
	fmt.Println(h.find(5))
	if m := h.Pop(); m != 1 {
		t.Fatal("Pop failed")
	}
	for i := 1; i <= 9; i++ {
		if m := h.Pop(); m != i {
			t.Fatal("Pop failed")
		}
	}
}

func TestBinHeap_Delete(t *testing.T) {
	h := New(9)
	t1 := time.Now()
	t2 := time.Now()
	t3 := time.Now()
	t4 := time.Now()
	h.Push(util.Integer(1), &t1)
	h.Push(util.Integer(2), &t2)
	h.Push(util.Integer(3), &t3)
	h.Push(util.Integer(4), &t4)
	fmt.Println(h.find(&t3))
}
