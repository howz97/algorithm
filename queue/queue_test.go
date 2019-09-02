package queue

import (
	"strconv"
	"testing"
)

func Test_StrQueue(t *testing.T) {
	q := NewStrQ()
	if !q.IsEmpty() {
		t.Fatal()
	}
	for i := 0; i <= 10; i++ {
		q.PushBack(strconv.Itoa(i))
	}
	for i := 0; i <= 10; i++ {
		n, _ := strconv.Atoi(q.Front())
		if n != i {
			t.Fatal()
		}
	}
}
