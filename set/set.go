package set

import (
	"sort"
)

type Set map[int]struct{}

func New() Set {
	return make(map[int]struct{})
}

func (s Set) Add(i int) {
	s[i] = struct{}{}
}

func (s Set) Contains(i int) bool {
	_, contain := s[i]
	return contain
}

func (s Set) Remove(i int) {
	delete(s, i)
}

func (s Set) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s Set) Len() int {
	return len(s)
}

func (s Set) Traverse() []int {
	result := make([]int, 0)
	for k := range s {
		result = append(result, k)
	}
	return result
}

func (s Set) SortTraverse() []int {
	result := s.Traverse()
	sort.Ints(result)
	return result
}
