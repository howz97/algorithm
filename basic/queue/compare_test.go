package queue

import (
	"fmt"
	"testing"
	"time"
)

const (
	testTimes = 30000000
)

func TestInterfaceQ(t *testing.T) {
	qSlice := NewSliceQ[int](0)
	start := time.Now()
	for i := 0; i < testTimes; i++ {
		qSlice.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		qSlice.Front()
	}
	elapsed := time.Since(start)
	fmt.Printf("Slice cost [%v]\n", elapsed.String())

	qLinked := NewLinkQ[int]()
	start = time.Now()
	for i := 0; i < testTimes; i++ {
		qLinked.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		qLinked.Front()
	}
	elapsed = time.Since(start)
	fmt.Printf("Linked [%v]\n", elapsed.String())
}
