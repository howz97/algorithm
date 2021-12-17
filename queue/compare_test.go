package queue

import (
	"fmt"
	"testing"
	"time"
)

const (
	testTimes = 10000000
)

func TestInterfaceQ(t *testing.T) {
	qSlice := NewQueen(testTimes)
	qLinked := NewLinked()
	qInt := NewLinkInt()

	start := time.Now()
	for i := 0; i < testTimes; i++ {
		qSlice.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		qSlice.Front()
	}
	elapsed := time.Since(start)
	fmt.Printf("Queen implemented by array cost [%v]\n", elapsed.String())

	start = time.Now()
	for i := 0; i < testTimes; i++ {
		qLinked.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		qLinked.Front()
	}
	elapsed = time.Since(start)
	fmt.Printf("Queen implemented by linked-list cost [%v]\n", elapsed.String())

	start = time.Now()
	for i := 0; i < testTimes; i++ {
		qInt.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		qInt.Front()
	}
	elapsed = time.Since(start)
	fmt.Printf("Queen implemented by linked-list and direct type cost [%v]\n", elapsed.String())
}
