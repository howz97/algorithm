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
	qSlice := NewSlice(0)
	start := time.Now()
	for i := 0; i < testTimes; i++ {
		qSlice.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		qSlice.Front()
	}
	elapsed := time.Since(start)
	fmt.Printf("Slice cost [%v]\n", elapsed.String())

	qSliStr := NewSliStr(testTimes)
	start = time.Now()
	for i := 0; i < testTimes; i++ {
		qSliStr.PushBack("x")
	}
	for i := 0; i < testTimes; i++ {
		qSliStr.Front()
	}
	elapsed = time.Since(start)
	fmt.Printf("Slice-Str cost [%v]\n", elapsed.String())

	qLinked := NewLinked()
	start = time.Now()
	for i := 0; i < testTimes; i++ {
		qLinked.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		qLinked.Front()
	}
	elapsed = time.Since(start)
	fmt.Printf("Linked [%v]\n", elapsed.String())

	qInt := NewLinkInt()
	start = time.Now()
	for i := 0; i < testTimes; i++ {
		qInt.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		qInt.Front()
	}
	elapsed = time.Since(start)
	fmt.Printf("Linked-Int [%v]\n", elapsed.String())
}
