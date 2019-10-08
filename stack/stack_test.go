package stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	testTimes := 100
	s := NewStack(10)
	for i := 0; i < testTimes; i++ {
		s.Push(i)
	}
	for i := 0; i < testTimes; i++ {
		fmt.Print(s.Pop(), " ")
	}
}
