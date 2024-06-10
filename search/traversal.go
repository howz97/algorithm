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

package search

import (
	"cmp"

	"github.com/howz97/algorithm/basic"
	"github.com/waiyva/binary-tree/btprinter"
)

type Searcher[K cmp.Ordered, V any] interface {
	Put(key K, val V)
	Get(key K) (V, bool)
	Del(key K)
	Size() uint
}

type BNode interface {
	IsNil() bool
	Left() BNode
	Right() BNode
}

// Preorder traverse nodes in pre-order recursively
func PreorderRecur[A BNode](nd A, fn func(A) bool) bool {
	if !fn(nd) {
		return false
	}
	if !nd.Left().IsNil() {
		if !PreorderRecur(nd.Left().(A), fn) {
			return false
		}
	}
	if !nd.Right().IsNil() {
		if !PreorderRecur(nd.Right().(A), fn) {
			return false
		}
	}
	return true
}

// Preorder traverse nodes in pre-order non-recursively
func Preorder[A BNode](nd A, fn func(A) bool) {
	right := basic.NewStack[A](0)
	right.Push(nd)
	for right.Size() > 0 {
		n := right.Pop()
		for !n.IsNil() {
			if !fn(n) {
				return
			}
			if !n.Right().IsNil() {
				right.Push(n.Right().(A))
			}
			n = n.Left().(A)
		}
	}
}

func Inorder[A BNode](nd A, fn func(A) bool) bool {
	if !nd.Left().IsNil() {
		if !Inorder(nd.Left().(A), fn) {
			return false
		}
	}
	if !fn(nd) {
		return false
	}
	if !nd.Right().IsNil() {
		if !Inorder(nd.Right().(A), fn) {
			return false
		}
	}
	return true
}

func Postorder[A BNode](nd A, fn func(A) bool) bool {
	if !nd.Left().IsNil() {
		if !Postorder(nd.Left().(A), fn) {
			return false
		}
	}
	if !nd.Right().IsNil() {
		if !Postorder(nd.Right().(A), fn) {
			return false
		}
	}
	return fn(nd)
}

func LevelOrder[A BNode](nd A, fn func(A) bool) {
	if nd.IsNil() {
		return
	}
	q := basic.NewLinkQueue[A]()
	q.PushBack(nd)
	for q.Size() > 0 {
		nd = q.PopFront()
		if !fn(nd) {
			break
		}
		if !nd.Left().IsNil() {
			q.PushBack(nd.Left().(A))
		}
		if !nd.Right().IsNil() {
			q.PushBack(nd.Right().(A))
		}
	}
}

func RevOrder[A BNode](nd A, fn func(A) bool) bool {
	if !nd.Right().IsNil() {
		if !RevOrder(nd.Right().(A), fn) {
			return false
		}
	}
	if !fn(nd) {
		return false
	}
	if !nd.Left().IsNil() {
		if !RevOrder(nd.Left().(A), fn) {
			return false
		}
	}
	return true
}

func PrintTree[A BNode](nd A, toStr func(A) string) {
	var sli []string
	q := basic.NewLinkQueue[A]()
	q.PushBack(nd)
	for q.Size() > 0 {
		nd = q.PopFront()
		if nd.IsNil() {
			sli = append(sli, "#")
			continue
		}
		sli = append(sli, toStr(nd))
		q.PushBack(nd.Left().(A))
		q.PushBack(nd.Right().(A))
	}
	btprinter.PrintTree(sli)
}
