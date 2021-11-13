package trietree

import (
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/queue"
)

type TSTNode struct {
	r                rune
	v                interface{}
	left, mid, right *TSTNode
}

func NewTSTNode(r rune) *TSTNode {
	return &TSTNode{r: r}
}

func (t *TSTNode) Value() T {
	return t.v
}

func (t *TSTNode) Insert(a alphabet.Interface, k []rune, v T) {
	if len(k) == 0 {
		return
	}
	switch true {
	case a.ToIndex(k[0]) < a.ToIndex(t.r):
		if t.left == nil {
			t.left = NewTSTNode(k[0])
		}
		t.left.Insert(a, k, v)
	case a.ToIndex(k[0]) > a.ToIndex(t.r):
		if t.right == nil {
			t.right = NewTSTNode(k[0])
		}
		t.right.Insert(a, k, v)
	default:
		if len(k) == 1 {
			t.v = v
		} else {
			if t.mid == nil {
				t.mid = NewTSTNode(k[1])
			}
			t.mid.Insert(a, k[1:], v)
		}
	}
}

func (t *TSTNode) Delete(a alphabet.Interface, k []rune) {
	if len(k) == 0 {
		return
	}
	switch true {
	case a.ToIndex(k[0]) < a.ToIndex(t.r):
		if t.left != nil {
			t.left.Delete(a, k)
		}
	case a.ToIndex(k[0]) > a.ToIndex(t.r):
		if t.right != nil {
			t.right.Delete(a, k)
		}
	default:
		if len(k) == 1 {
			t.v = nil
		} else if t.mid != nil {
			t.mid.Delete(a, k[1:])
		}
	}
}

// find 找到k对应的节点(n)，有这个节点不代表k存在，是否存在需要看n.v是否为nil
func (t *TSTNode) Find(a alphabet.Interface, k []rune) TrieNode {
	if len(k) == 0 {
		return nil
	}
	switch true {
	case a.ToIndex(k[0]) < a.ToIndex(t.r):
		if t.left != nil {
			return t.left.Find(a, k)
		}
	case a.ToIndex(k[0]) > a.ToIndex(t.r):
		if t.right != nil {
			return t.right.Find(a, k)
		}
	default:
		if len(k) == 1 {
			return t
		} else {
			return t.mid.Find(a, k[1:])
		}
	}
	return nil
}

func (t *TSTNode) LongestPrefixOf(a alphabet.Interface, s []rune, d, l int) int {
	if len(s) == 0 {
		return l
	}
	switch true {
	case a.ToIndex(s[d]) < a.ToIndex(t.r):
		if t.left != nil {
			return t.left.LongestPrefixOf(a, s, d, l)
		}
	case a.ToIndex(s[d]) > a.ToIndex(t.r):
		if t.right != nil {
			return t.right.LongestPrefixOf(a, s, d, l)
		}
	default:
		if t.v != nil {
			l = d + 1
		}
		if d == len(s)-1 {
			return l
		} else if t.mid != nil {
			return t.mid.LongestPrefixOf(a, s, d+1, l)
		}
	}
	return l
}

func (t *TSTNode) Collect(a alphabet.Interface, prefix string, keys *queue.StrQ) {
	if t.v != nil {
		keys.PushBack(prefix + string(t.r))
	}
	if t.left != nil {
		t.left.Collect(a, prefix, keys)
	}
	if t.mid != nil {
		t.mid.Collect(a, prefix+string(t.r), keys)
	}
	if t.right != nil {
		t.right.Collect(a, prefix, keys)
	}
}

func (t *TSTNode) KeysMatch(a alphabet.Interface, pattern []rune, prefix string, keys *queue.StrQ) {
	if len(pattern) == 0 {
		return
	}
	if t.left != nil && (pattern[0] == '.' || a.ToIndex(pattern[0]) < a.ToIndex(t.r)) {
		t.left.KeysMatch(a, pattern, prefix, keys)
	}
	if t.right != nil && (pattern[0] == '.' || a.ToIndex(pattern[0]) > a.ToIndex(t.r)) {
		t.right.KeysMatch(a, pattern, prefix, keys)
	}
	if pattern[0] == '.' || a.ToIndex(pattern[0]) == a.ToIndex(t.r) {
		if len(pattern) == 1 {
			if t.v != nil {
				keys.PushBack(prefix + string(t.r))
			}
		} else if t.mid != nil {
			t.mid.KeysMatch(a, pattern[1:], prefix+string(t.r), keys)
		}
	}
}
