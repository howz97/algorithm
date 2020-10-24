package set

type Set map[interface{}]struct{}

func NewSet() Set {
	return make(map[interface{}]struct{})
}

func (s Set) Add(e interface{}) {
	s[e] = struct{}{}
}

func (s Set) Contains(e interface{}) bool {
	_, ok := s[e]
	return ok
}

func (s Set) Remove(e interface{}) {
	delete(s, e)
}

// TakeOne take out an element
func (s Set) TakeOne() interface{} {
	for e := range s {
		delete(s, e)
		return e
	}
	return nil
}

func (s Set) Clear() {
	for e := range s {
		delete(s, e)
	}
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Traverse() []interface{} {
	ret := make([]interface{}, 0, len(s))
	for e := range s {
		ret = append(ret, e)
	}
	return ret
}
