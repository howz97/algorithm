package set

import (
	"sort"
)

type StrSet map[string]struct{}

func New() StrSet {
	return make(map[string]struct{})
}

func (ss StrSet) Add(str string) {
	ss[str] = struct{}{}
}

func (ss StrSet) Contains(str string) bool {
	_, contain := ss[str]
	return contain
}

func (ss StrSet) Remove(str string) {
	delete(ss, str)
}

func (ss StrSet) Clear() {
	for k := range ss {
		delete(ss, k)
	}
}

func (ss StrSet) Len() int {
	return len(ss)
}

func (ss StrSet) Traverse() []string {
	result := make([]string, 0)
	for k := range ss {
		result = append(result, k)
	}
	return result
}

func (ss StrSet) SortTraverse() []string {
	result := ss.Traverse()
	sort.Strings(result)
	return result
}
