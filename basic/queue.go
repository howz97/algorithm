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

type Queue[T any] struct {
	elems      []T
	head, back int
	size       int
}

func NewQueue[T any](cap int) *Queue[T] {
	if cap < 2 {
		cap = 2
	}
	return &Queue[T]{
		elems: make([]T, cap),
	}
}

func QueueFrom[T any](elems []T) *Queue[T] {
	return &Queue[T]{
		elems: elems[:cap(elems)],
		back:  len(elems),
		size:  len(elems),
	}
}

func (q *Queue[T]) Front() T {
	if q.size <= 0 {
		panic("empty queue")
	}
	return q.elems[q.head]
}

func (q *Queue[T]) Back() T {
	if q.size <= 0 {
		panic("empty queue")
	}
	back := q.back - 1
	if back < 0 {
		back = len(q.elems) - 1
	}
	return q.elems[back]
}

func (q *Queue[T]) PopFront() {
	if q.size <= 0 {
		panic("empty queue")
	}
	q.size--
	q.head++
	if q.head == len(q.elems) {
		q.head = 0
	}
}

func (q *Queue[T]) PopBack() {
	if q.size <= 0 {
		panic("empty queue")
	}
	q.size--
	q.back--
	if q.back < 0 {
		q.back = len(q.elems) - 1
	}
}

func (q *Queue[T]) PushBack(e T) {
	if q.size <= 0 {
		q.elems[0] = e
		q.head = 0
		q.back = 1
		q.size = 1
		return
	}
	if q.size == len(q.elems) {
		expand := make([]T, 2*len(q.elems))
		n := q.size
		for i := 0; i < n; i++ {
			expand[i] = q.elems[(q.head+i)%len(q.elems)]
		}
		q.elems = expand
		q.head = 0
		q.back = n
	}
	q.size++
	q.elems[q.back] = e
	q.back++
	if q.back == len(q.elems) {
		q.back = 0
	}
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) Clone() []T {
	elems := make([]T, 0, q.size)
	for q.Size() > 0 {
		elems = append(elems, q.Front())
	}
	return elems
}
