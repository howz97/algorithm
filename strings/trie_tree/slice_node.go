package trietree

import (
	"errors"

	"github.com/howz97/algorithm/basic/queue"
	"github.com/howz97/algorithm/strings/alphabet"
)

// R direction trie tree
type SliceNode struct {
	val  T
	next []*SliceNode
}

func newSliceNode(size int) *SliceNode {
	return &SliceNode{next: make([]*SliceNode, size)}
}

func (t *SliceNode) Find(a alphabet.Interface, k []rune) T {
	node, _ := t.Locate(a, k)
	if node == nil {
		return nil
	}
	return node.(*SliceNode).val
}

func (t *SliceNode) Locate(a alphabet.Interface, k []rune) (TrieNode, []rune) {
	if len(k) == 0 {
		return t, nil
	}
	next := t.next[a.ToIndex(k[0])]
	if next == nil {
		return nil, nil
	}
	return next.Locate(a, k[1:])
}

func (t *SliceNode) Upsert(a alphabet.Interface, k []rune, v T) {
	if len(k) == 0 {
		t.val = v
		return
	}

	i := a.ToIndex(k[0])
	next := t.next[i]
	if next == nil {
		next = newSliceNode(a.R())
		t.next[i] = next
	}
	next.Upsert(a, k[1:], v)
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

// Collect collects all keys of t and put them into SliStr
// p is the prefix record
func (t *SliceNode) Collect(a alphabet.Interface, prefix string, keys *queue.SliceQ[string]) {
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

func (t *SliceNode) KeysMatch(a alphabet.Interface, pattern []rune, prefix string, keys *queue.SliceQ[string]) {
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

func (t *SliceNode) Keys(a alphabet.Interface, keys *queue.SliceQ[string]) {
	t.Collect(a, "", keys)
}

func (t *SliceNode) Compress() error {
	return errors.New("compress not support")
}

func (t *SliceNode) IsCompressed() bool {
	return false
}

func (t *SliceNode) SetVal(v T) {
	t.val = v
}
