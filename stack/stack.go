package stack

const (
	minCap = 2
)

type T interface{}

type Stack struct {
	elems []T
	top   int
}

func New(c int) *Stack {
	if c < minCap {
		c = minCap
	}
	return &Stack{
		elems: make([]T, c),
		top:   0,
	}
}

func (s *Stack) Size() int {
	return s.top
}

func (s *Stack) Cap() int {
	return len(s.elems)
}

func (s *Stack) Pop() (T, bool) {
	if s.Size() <= 0 {
		return nil, false
	}
	s.top--
	return s.elems[s.top], true
}

func (s *Stack) Push(e T) {
	if s.isFull() {
		s.elems = append(s.elems, e)
	} else {
		s.elems[s.top] = e
	}
	s.top++
}

func (s *Stack) isFull() bool {
	return s.top == len(s.elems)
}
