package set

type Set[T comparable] map[T]struct{}

func New[T comparable]() Set[T] {
	return make(map[T]struct{})
}

func (s Set[T]) Add(e T) {
	s[e] = struct{}{}
}

func (s Set[T]) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set[T]) Remove(e T) {
	delete(s, e)
}

// TakeOne take out an element
func (s Set[T]) TakeOne() T {
	for e := range s {
		delete(s, e)
		return e
	}
	panic("set is empty")
}

func (s Set[T]) Clear() {
	for e := range s {
		delete(s, e)
	}
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Drain() []T {
	ret := make([]T, 0, len(s))
	for e := range s {
		ret = append(ret, e)
	}
	return ret
}
