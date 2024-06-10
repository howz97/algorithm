// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package basic

import (
	"fmt"

	"github.com/howz97/algorithm/util"
)

const (
	MinCap = 2
)

type Stack[T any] struct {
	elems []T
	top   int
}

func NewStack[T any](c int) *Stack[T] {
	if c < MinCap {
		c = MinCap
	}
	return &Stack[T]{
		elems: make([]T, c),
		top:   0,
	}
}

func (s *Stack[T]) Size() int {
	return s.top
}

func (s *Stack[T]) Cap() int {
	return len(s.elems)
}

func (s *Stack[T]) Pop() T {
	s.top--
	return s.elems[s.top]
}

func (s *Stack[T]) Push(e T) {
	if s.isFull() {
		s.elems = append(s.elems, e)
	} else {
		s.elems[s.top] = e
	}
	s.top++
}

func (s *Stack[T]) isFull() bool {
	return s.top == len(s.elems)
}

func (s *Stack[T]) Index(fn func(T) bool) int {
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

func (s *Stack[T]) Iterate(fn func(T) bool) {
	s.IterateRange(0, s.Size(), fn)
}

func (s *Stack[T]) IterateRange(lo, hi int, fn func(T) bool) {
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

func (s *Stack[T]) String() string {
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

func (s *Stack[T]) PeekIndex(i int) T {
	if i >= s.top {
		panic(fmt.Sprintf("%d out of bound", i))
	}
	return s.elems[i]
}

func (s *Stack[T]) Peek() T {
	return s.PeekIndex(s.top - 1)
}

func (s *Stack[T]) Reverse() {
	util.Reverse(s.elems)
}

func (s *Stack[T]) Drain() []T {
	var elems []T
	for s.Size() > 0 {
		elems = append(elems, s.Pop())
	}
	return elems
}

type StackCmp[T comparable] struct {
	Stack[T]
}

func NewStackCmp[T comparable](c int) *StackCmp[T] {
	return &StackCmp[T]{
		Stack: *NewStack[T](c),
	}
}

func (s *StackCmp[T]) Contains(e T) bool {
	for i := 0; i < s.top; i++ {
		if s.elems[i] == e {
			return true
		}
	}
	return false
}
