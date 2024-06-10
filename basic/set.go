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

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
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
