package leetcode

import (
	"testing"
)

func Test_queue(t *testing.T) {
	q := new(queue)
	if !q.empty() {
		t.Fatal("q.empty() failed")
	}
	for i := 0; i <= 10; i++ {
		q.pushBack(&Node{Key: i})
	}
	for i := 0; i <= 10; i++ {
		if q.front().Key != i {
			t.Fatal("q.front() failed")
		}
	}
}
