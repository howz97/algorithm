package heap

import (
	"fmt"
	"testing"
	"time"
)

func Test_Heap(t *testing.T) {
	h := New2[int, int](10)
	h.Push(1, 1)
	h.Push(1, 1)
	h.Push(2, 2)
	h.Push(3, 3)
	h.Push(6, 6)
	h.Push(5, 5)
	h.Push(4, 4)
	h.Push(9, 9)
	h.Push(8, 8)
	h.Push(7, 7)
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
	h := New2[int, *time.Time](9)
	t1 := time.Now()
	t2 := time.Now()
	t3 := time.Now()
	t4 := time.Now()
	h.Push(1, &t1)
	h.Push(2, &t2)
	h.Push(3, &t3)
	h.Push(4, &t4)
	fmt.Println(h.find(&t3))
}

func Example() {
	pq := New2[int, string](3)
	pq.Push(1, "1")
	pq.Push(9, "9")
	pq.Push(9, "9")
	pq.Push(7, "7")
	for pq.Size() > 0 {
		fmt.Print(pq.Pop())
	}
	fmt.Println()

	pq.Push(100, "1")
	pq.Push(9, "9")
	pq.Push(9, "9")
	pq.Push(7, "7")
	pq.Push(0, "x")
	pq.Del("x")
	pq.Fix(1, "1")
	for pq.Size() > 0 {
		fmt.Print(pq.Pop())
	}

	// Output:
	// 1799
	// 1799
}
