package set

import (
	"sort"
)

type IntSet map[int]struct{}

func NewIntSet() IntSet {
	return make(map[int]struct{})
}

func (s IntSet) Add(i int) {
	s[i] = struct{}{}
}

func (s IntSet) Contains(i int) bool {
	_, contain := s[i]
	return contain
}

func (s IntSet) Remove(i int) {
	delete(s, i)
}

// RemoveOne is not allowed to be called when it is empty
func (s IntSet) RemoveOne() int {
	i := 0
	if s.IsEmpty() {
		panic("removing from an empty set is not allowed")
	}
	for i = range s {
		s.Remove(i)
		break
	}
	return i
}

func (s IntSet) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s IntSet) Len() int {
	return len(s)
}

func (s IntSet) IsEmpty() bool {
	return s.Len() == 0
}

func (s IntSet) Traverse() []int {
	result := make([]int, 0)
	for k := range s {
		result = append(result, k)
	}
	return result
}

func (s IntSet) Range(fn func(int) bool) {
	for e := range s {
		if !fn(e) {
			break
		}
	}
}

func (s IntSet) SortTraverse() []int {
	result := s.Traverse()
	sort.Ints(result)
	return result
}
