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
	"testing"
)

func BenchmarkQueue_PushBack(b *testing.B) {
	q := NewQueue[int](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
}

func BenchmarkQueue_PopFront(b *testing.B) {
	q := NewQueue[int](b.N)
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.PopFront()
	}
}

func BenchmarkLinkQueue_PushBack(b *testing.B) {
	q := NewList[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
}

func BenchmarkLinkQueue_PopFront(b *testing.B) {
	q := NewList[int]()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
	b.ResetTimer()
	for q.Size() > 0 {
		q.PopFront()
	}
}

func TestQueue(t *testing.T) {
	var queues = [2]Fifo[int]{
		NewQueue[int](0),
		NewList[int](),
	}
	pushFn := func(v int) {
		for _, q := range queues {
			q.PushBack(v)
		}
	}
	popFn := func() {
		v := queues[0].Front()
		for _, q := range queues {
			if v != q.Front() {
				t.Fatal()
			}
			q.PopFront()
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
