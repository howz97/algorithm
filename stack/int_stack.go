package stack

type IntStack struct {
	Stack
}

func NewInt(c int) *IntStack {
	return &IntStack{
		Stack: *New(c),
	}
}

func (s *IntStack) Pop() (int, bool) {
	e, ok := s.Stack.Pop()
	if !ok {
		return 0, false
	}
	return e.(int), true
}
