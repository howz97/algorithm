package stack

import (
	"fmt"
	. "github.com/howz97/algorithm/util"
)

const (
	minCap = 2
)

type Stack struct {
	elems Elems
	top   int
}

func New(c int) *Stack {
	if c < minCap {
		c = minCap
	}
	return &Stack{
		elems: make(Elems, c),
		top:   0,
	}
}

func (s *Stack) Size() int {
	return s.top
}

func (s *Stack) Cap() int {
	return len(s.elems)
}

func (s *Stack) Pop() T {
	s.top--
	return s.elems[s.top]
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

func (s *Stack) Contains(e T) bool {
	for i := 0; i < s.top; i++ {
		if s.elems[i] == e {
			return true
		}
	}
	return false
}

func (s *Stack) Index(fn func(T) bool) int {
	i := 0
	s.Iterate(func(e T) bool {
		ok := fn(e)
		if ok {
			return false
		}
		i++
		return true
	})
	if i >= s.top {
		return -1
	}
	return i
}

func (s *Stack) Iterate(fn func(T) bool) {
	s.IterateRange(0, s.Size(), fn)
}

func (s *Stack) IterateRange(lo, hi int, fn func(T) bool) {
	if lo < 0 {
		lo = 0
	}
	if hi > s.top {
		hi = s.top
	}
	for ; lo < hi; lo++ {
		if !fn(s.elems[lo]) {
			break
		}
	}
}

func (s *Stack) String() string {
	if s == nil {
		return "<nil>"
	}
	str := ""
	s.Iterate(func(v T) bool {
		str += fmt.Sprint(v) + "<"
		return true
	})
	str += "(top)"
	return str
}

func (s *Stack) Peek(i int) T {
	if s.top <= 0 {
		return nil
	}
	if i < 0 {
		i = 0
	}
	if i >= s.top {
		i = s.top - 1
	}
	return s.elems[i]
}

func (s *Stack) Reverse() {
	ReverseSlice(s.elems)
}

func SizeOf(s *Stack) int {
	if s == nil {
		return -1
	}
	return s.Size()
}
