package trietree

import (
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/queue"
)

type SliceNode struct {
	val  T
	next []*SliceNode
}

func NewNodeSlice(size int) *SliceNode {
	return &SliceNode{next: make([]*SliceNode, size)}
}

func (t *SliceNode) Value() T {
	return t.val
}

func (t *SliceNode) Find(a alphabet.Interface, k []rune) TrieNode {
	if len(k) == 0 {
		return t
	}
	next := t.next[a.ToIndex(k[0])]
	if next == nil {
		return nil
	}
	return next.Find(a, k[1:])
}

func (t *SliceNode) Insert(a alphabet.Interface, k []rune, v T) {
	if len(k) == 0 {
		t.val = v
		return
	}

	i := a.ToIndex(k[0])
	next := t.next[i]
	if next == nil {
		next = NewNodeSlice(a.R())
		t.next[i] = next
	}
	next.Insert(a, k[1:], v)
}

func (t *SliceNode) Delete(a alphabet.Interface, k []rune) {
	if len(k) == 0 {
		t.val = nil
		return
	}
	i := a.ToIndex(k[0])
	next := t.next[i]
	if next != nil {
		next.Delete(a, k[1:])
	}
}

// LongestPrefixOf 找出t的所有匹配s[d:]的前缀的key中最长的那一个
// 返回值length代表s的前l个rune就是这个要找的key
func (t *SliceNode) LongestPrefixOf(a alphabet.Interface, s []rune, d int, l int) int {
	if t.val != nil {
		l = d
	}
	if len(s) == d {
		return l
	}
	i := a.ToIndex(s[d])
	next := t.next[i]
	if next == nil {
		return l
	}
	return next.LongestPrefixOf(a, s, d+1, l)
}

// Collect collects all keys of t and put them into StrQ
// p is the prefix record
func (t *SliceNode) Collect(a alphabet.Interface, prefix string, keys *queue.StrQ) {
	if t.val != nil {
		keys.PushBack(prefix)
	}
	for i, next := range t.next {
		if next == nil {
			continue
		}
		next.Collect(a, prefix+string(a.ToRune(rune(i))), keys)
	}
}

func (t *SliceNode) KeysMatch(a alphabet.Interface, pattern []rune, prefix string, keys *queue.StrQ) {
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
			next.KeysMatch(a, pattern[1:], prefix+string(a.ToRune(rune(i))), keys)
		}
	} else {
		next := t.next[a.ToIndex(pattern[0])]
		if next != nil {
			prefix = prefix + string(pattern[0])
			next.KeysMatch(a, pattern[1:], prefix, keys)
		}
	}
}
