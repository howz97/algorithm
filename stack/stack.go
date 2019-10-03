package stack

import "errors"

const (
	defaultInitCap = 50
)

var (
	ErrEmptyStack = errors.New("empty stack")
)

type Stack struct {
	elems []int
	top   int
}

func New(initCap int) *Stack {
	if initCap < 2 {
		initCap = defaultInitCap
	}
	return &Stack{
		elems: make([]int, initCap),
		top:   0,
	}
}

func (s *Stack) Size() int {
	return s.top
}

func (s *Stack) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("pop from empty stack")
	}
	s.top--
	return s.elems[s.top]
}

func (s *Stack) Push(elem int) {
	if s.isFull() {
		s.elems = append(s.elems, elem)
	} else {
		s.elems[s.top] = elem
	}
	s.top++
}

func (s *Stack) isFull() bool {
	return s.top == len(s.elems)
}
