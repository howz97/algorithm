package stack

type Stack struct {
	elems []interface{}
	top   int
}

func NewStack(initCap int) *Stack {
	if initCap < 2 {
		initCap = defaultInitCap
	}
	return &Stack{
		elems: make([]interface{}, initCap),
		top:   0,
	}
}

func (s *Stack) Size() int {
	return s.top
}

func (s *Stack) IsEmpty() bool {
	return s == nil || s.top == 0
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		panic("pop from empty stack")
	}
	s.top--
	return s.elems[s.top]
}

func (s *Stack) Push(elem interface{}) {
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
