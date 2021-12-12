package stack

type IntStack struct {
	Stack
}

func NewInt(c int) *IntStack {
	return &IntStack{
		Stack: *New(c),
	}
}

func (s *IntStack) Pop() int {
	return s.Stack.Pop().(int)
}

func (s *IntStack) ToSlice() []int {
	var sli []int
	for s.Size() > 0 {
		v := s.Pop()
		sli = append(sli, v)
	}
	return sli
}
