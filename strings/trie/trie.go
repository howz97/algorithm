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

package trie

import (
	"github.com/howz97/algorithm/basic"
	"github.com/howz97/algorithm/strings/alphabet"
)

func NewTrie[T any](alp alphabet.IAlp) *Trie[T] {
	return &Trie[T]{
		alp:  alp,
		root: newNode[T](alp.R()),
	}
}

type Trie[T any] struct {
	alp  alphabet.IAlp
	root *node[T]
}

func (t *Trie[T]) Find(key string) *T {
	node := t.root.locate(t.alp.ToIndices(key))
	if node == nil {
		return nil
	}
	return node.val
}

func (t *Trie[T]) Upsert(key string, v T) {
	t.root.upsert(t.alp.ToIndices(key), v, t.alp.R())
}

func (t *Trie[T]) Delete(key string) {
	node := t.root.locate(t.alp.ToIndices(key))
	if node != nil {
		node.val = nil
	}
}

func (t *Trie[T]) LongestPrefixOf(s string) string {
	l := t.root.longestPrefixOf(t.alp.ToIndices(s), 0, 0)
	return string([]rune(s)[:l])
}

func (t *Trie[T]) KeysWithPrefix(prefix string) []string {
	node := t.root.locate(t.alp.ToIndices(prefix))
	if node == nil {
		return nil
	}
	q := basic.NewQueue[string](0)
	node.collect(t.alp, prefix, q)
	return q.Clone()
}

func (t *Trie[T]) KeysMatch(p string) []string {
	q := basic.NewQueue[string](0)
	t.root.keysMatch(t.alp, []rune(p), "", q)
	return q.Clone()
}

type node[T any] struct {
	val  *T
	next []*node[T]
}

func newNode[T any](size int) *node[T] {
	return &node[T]{next: make([]*node[T], size)}
}

func (t *node[T]) locate(k []rune) *node[T] {
	if len(k) == 0 {
		return t
	}
	next := t.next[k[0]]
	if next == nil {
		return nil
	}
	return next.locate(k[1:])
}

func (t *node[T]) upsert(k []rune, v T, r int) {
	if len(k) == 0 {
		t.val = &v
		return
	}
	next := t.next[k[0]]
	if next == nil {
		next = newNode[T](r)
		t.next[k[0]] = next
	}
	next.upsert(k[1:], v, r)
}

func (t *node[T]) longestPrefixOf(s []rune, d int, l int) int {
	if t.val != nil {
		l = d
	}
	if len(s) == d {
		return l
	}
	next := t.next[s[d]]
	if next == nil {
		return l
	}
	return next.longestPrefixOf(s, d+1, l)
}

func (t *node[T]) collect(a alphabet.IAlp, prefix string, keys *basic.Queue[string]) {
	if t.val != nil {
		keys.PushBack(prefix)
	}
	for i, next := range t.next {
		if next == nil {
			continue
		}
		next.collect(a, prefix+string(a.ToRune(rune(i))), keys)
	}
}

func (t *node[T]) keysMatch(a alphabet.IAlp, pattern []rune, prefix string, keys *basic.Queue[string]) {
	if len(pattern) == 0 {
		if t.val != nil {
			keys.PushBack(prefix)
		}
		return
	}
	if pattern[0] == '.' {
		for i, next := range t.next {
			if next == nil {
				continue
			}
			next.keysMatch(a, pattern[1:], prefix+string(a.ToRune(rune(i))), keys)
		}
	} else {
		next := t.next[a.ToIndex(pattern[0])]
		if next != nil {
			prefix = prefix + string(pattern[0])
			next.keysMatch(a, pattern[1:], prefix, keys)
		}
	}
}
