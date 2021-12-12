package stack

import (
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	testTimes := 100
	s := New(10)
	for i := 0; i < testTimes; i++ {
		s.Push(i)
	}
	for i := 0; i < testTimes; i++ {
		e := s.Pop()
		fmt.Print(e, " ")
	}
}
