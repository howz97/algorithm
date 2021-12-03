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

func (s *IntStack) ToSlice() []int {
	var sli []int
	for {
		v, ok := s.Pop()
		if !ok {
			break
		}
		sli = append(sli, v)
	}
	return sli
}
