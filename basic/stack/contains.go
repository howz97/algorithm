package stack

type Stack2[T comparable] struct {
	Stack[T]
}

func New2[T comparable](c int) *Stack2[T] {
	return &Stack2[T]{
		Stack: *New[T](c),
	}
}

func (s *Stack2[T]) Contains(e T) bool {
	for i := 0; i < s.top; i++ {
		if s.elems[i] == e {
			return true
		}
	}
	return false
}
