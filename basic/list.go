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

type List[T any] struct {
	head *elem[T]
	tail *elem[T]
	size int
}

func NewList[T any]() *List[T] {
	return new(List[T])
}

type elem[T any] struct {
	v    T
	next *elem[T]
}

func (q *List[T]) Front() T {
	return q.head.v
}

func (q *List[T]) PopFront() {
	q.size--
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
}

func (q *List[T]) PushBack(e T) {
	q.size++
	if q.head == nil {
		q.head = &elem[T]{
			v: e,
		}
		q.tail = q.head
		return
	}
	q.tail.next = &elem[T]{
		v: e,
	}
	q.tail = q.tail.next
}

func (q *List[T]) Size() int {
	return q.size
}

type Fifo[T any] interface {
	Front() T
	PopFront()
	PushBack(v T)
}
