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

type Stack[T any] struct {
	elems []T
}

func NewStack[T any](c int) *Stack[T] {
	return &Stack[T]{
		elems: make([]T, 0, c),
	}
}

func NewStackFrom[T any](elems []T) *Stack[T] {
	return &Stack[T]{
		elems: elems,
	}
}

func (s *Stack[T]) Size() int {
	return len(s.elems)
}

func (s *Stack[T]) Cap() int {
	return cap(s.elems)
}

func (s *Stack[T]) Back() T {
	return s.elems[len(s.elems)-1]
}

func (s *Stack[T]) PopBack() {
	s.elems = s.elems[:len(s.elems)-1]
}

func (s *Stack[T]) PushBack(e T) {
	s.elems = append(s.elems, e)
}

func (s *Stack[T]) Find(fn func(T) bool) int {
	i := 0
	s.Iterate(false, func(e T) bool {
		ok := fn(e)
		if ok {
			return false
		}
		i++
		return true
	})
	if i >= len(s.elems) {
		return -1
	}
	return i
}

func (s *Stack[T]) Iterate(popOrd bool, fn func(T) bool) {
	if popOrd {
		s.IterRange(len(s.elems)-1, 0, fn)
	} else {
		s.IterRange(0, len(s.elems)-1, fn)
	}
}

func (s *Stack[T]) IterRange(src, dst int, fn func(T) bool) {
	if src < dst {
		for ; src <= dst; src++ {
			if !fn(s.elems[src]) {
				break
			}
		}
	} else {
		for ; src >= dst; src-- {
			if !fn(s.elems[src]) {
				break
			}
		}
	}
}

func (s *Stack[T]) String() string {
	if s == nil {
		return "<nil>"
	}
	str := ""
	s.Iterate(false, func(v T) bool {
		str += fmt.Sprint(v) + "<"
		return true
	})
	str += "(top)"
	return str
}

func (s *Stack[T]) PeekAt(i int) T {
	return s.elems[i]
}

func (s *Stack[T]) Peek() T {
	return s.elems[len(s.elems)-1]
}

func (s *Stack[T]) Reverse() {
	util.Reverse(s.elems)
}

func (s *Stack[T]) ToSlice() []T {
	var elems []T
	for s.Size() > 0 {
		elems = append(elems, s.Back())
		s.PopBack()
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

func (s *StackCmp[T]) Find(e T) int {
	for i := range s.elems {
		if s.elems[i] == e {
			return i
		}
	}
	return -1
}

type Filo[T any] interface {
	Back() T
	PopBack()
	PushBack(e T)
}
