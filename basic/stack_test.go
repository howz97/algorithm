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
	"testing"
)

func ExampleStack() {
	testTimes := 100
	s := NewStack[int](10)
	for i := 0; i < testTimes; i++ {
		s.PushBack(i)
	}
	for i := 0; i < testTimes; i++ {
		e := s.Back()
		s.PopBack()
		fmt.Print(e, " ")
	}
}

func TestStack(t *testing.T) {
	var stacks = [2]Filo[int]{
		NewStack[int](0),
		NewQueue[int](0),
	}
	pushFn := func(v int) {
		for _, q := range stacks {
			q.PushBack(v)
		}
	}
	popFn := func() {
		v := stacks[0].Back()
		for _, stk := range stacks {
			if v != stk.Back() {
				t.Fatal()
			}
			stk.PopBack()
		}
	}
	for i := 0; i < 1000; i++ {
		pushFn(i)
	}
	for i := 0; i < 500; i++ {
		popFn()
	}
	for i := 0; i < 2000; i++ {
		pushFn(i)
	}
	for i := 0; i < 1500; i++ {
		popFn()
	}
	for i := 0; i < 1000; i++ {
		pushFn(i)
	}
	for i := 0; i < 2000; i++ {
		popFn()
	}
	for i := 0; i < 100; i++ {
		pushFn(i)
	}
	for i := 0; i < 100; i++ {
		popFn()
	}
}
