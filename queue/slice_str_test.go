package queue

import (
	"strconv"
	"testing"
)

func Test_StrQueue(t *testing.T) {
	q := NewSliceStr(0)
	if q.Size() > 0 {
		t.Fatal()
	}
	for i := 0; i <= 10; i++ {
		q.PushBack(strconv.Itoa(i))
	}
	for i := 0; i <= 10; i++ {
		s := q.Front()
		n, _ := strconv.Atoi(s)
		if n != i {
			t.Fatal()
		}
	}
}
