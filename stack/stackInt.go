package stack

import "errors"

const (
	defaultInitCap = 50
)

var (
	ErrEmptyStack = errors.New("empty stack")
)

type StackInt struct {
	elemsInt []int
	top      int
}

func NewStackInt(initCap int) *StackInt {
	if initCap < 2 {
		initCap = defaultInitCap
	}
	return &StackInt{
		elemsInt: make([]int, initCap),
		top:      0,
	}
}

func (s *StackInt) Size() int {
	return s.top
}

func (s *StackInt) IsEmpty() bool {
	return s.top == 0
}

func (s *StackInt) Pop() int {
	if s.IsEmpty() {
		panic("pop from empty stack")
	}
	s.top--
	return s.elemsInt[s.top]
}

func (s *StackInt) Push(elem int) {
	if s.isFull() {
		s.elemsInt = append(s.elemsInt, elem)
	} else {
		s.elemsInt[s.top] = elem
	}
	s.top++
}

func (s *StackInt) isFull() bool {
	return s.top == len(s.elemsInt)
}
