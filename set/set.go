package set

type T interface{}

type Set map[T]struct{}

func New() Set {
	return make(map[T]struct{})
}

func (s Set) Add(e T) {
	s[e] = struct{}{}
}

func (s Set) Contains(e T) bool {
	_, ok := s[e]
	return ok
}

func (s Set) Remove(e T) {
	delete(s, e)
}

// TakeOne take out an element
func (s Set) TakeOne() (T, bool) {
	for e := range s {
		delete(s, e)
		return e, true
	}
	return nil, false
}

func (s Set) Clear() {
	for e := range s {
		delete(s, e)
	}
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) isEmpty() bool {
	return len(s) == 0
}

func (s Set) ToSlice() []T {
	ret := make([]T, 0, len(s))
	for e := range s {
		ret = append(ret, e)
	}
	return ret
}
