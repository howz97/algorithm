package trietree

import (
	"errors"
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/queue"
)

// Three direction trie tree
type TSTNode struct {
	r                rune
	v                T
	left, mid, right *TSTNode
}

func newTSTNode(r rune) *TSTNode {
	return &TSTNode{r: r}
}

func (t *TSTNode) Upsert(_ alphabet.Interface, k []rune, v T) {
	if len(k) == 0 {
		return
	}
	switch true {
	case k[0] < t.r:
		if t.left == nil {
			t.left = newTSTNode(k[0])
		}
		t.left.Upsert(nil, k, v)
	case k[0] > t.r:
		if t.right == nil {
			t.right = newTSTNode(k[0])
		}
		t.right.Upsert(nil, k, v)
	default:
		if len(k) == 1 {
			t.v = v
		} else {
			if t.mid == nil {
				t.mid = newTSTNode(k[1])
			}
			t.mid.Upsert(nil, k[1:], v)
		}
	}
}

func (t *TSTNode) Delete(_ alphabet.Interface, k []rune) {
	if len(k) == 0 {
		return
	}
	switch true {
	case k[0] < t.r:
		if t.left != nil {
			t.left.Delete(nil, k)
		}
	case k[0] > t.r:
		if t.right != nil {
			t.right.Delete(nil, k)
		}
	default:
		if len(k) == 1 {
			t.v = nil
		} else if t.mid != nil {
			t.mid.Delete(nil, k[1:])
		}
	}
}

func (t *TSTNode) Find(a alphabet.Interface, k []rune) T {
	node, _ := t.Locate(a, k)
	if node == nil {
		return nil
	}
	return node.(*TSTNode).v
}

// find 找到k对应的节点(n)，有这个节点不代表k存在，是否存在需要看n.v是否为nil
func (t *TSTNode) Locate(_ alphabet.Interface, k []rune) (TrieNode, []rune) {
	if len(k) == 0 {
		return nil, nil
	}
	switch true {
	case k[0] < t.r:
		if t.left != nil {
			return t.left.Locate(nil, k)
		}
	case k[0] > t.r:
		if t.right != nil {
			return t.right.Locate(nil, k)
		}
	default:
		if len(k) == 1 {
			return t, nil
		} else if t.mid != nil {
			return t.mid.Locate(nil, k[1:])
		}
	}
	return nil, nil
}

func (t *TSTNode) LongestPrefixOf(_ alphabet.Interface, s []rune, d, l int) int {
	if len(s) == 0 {
		return l
	}
	switch true {
	case s[d] < t.r:
		if t.left != nil {
			return t.left.LongestPrefixOf(nil, s, d, l)
		}
	case s[d] > t.r:
		if t.right != nil {
			return t.right.LongestPrefixOf(nil, s, d, l)
		}
	default:
		if t.v != nil {
			l = d + 1
		}
		if d == len(s)-1 {
			return l
		} else if t.mid != nil {
			return t.mid.LongestPrefixOf(nil, s, d+1, l)
		}
	}
	return l
}

func (t *TSTNode) Collect(_ alphabet.Interface, prefix string, keys *queue.StrQ) {
	if t.v != nil {
		keys.PushBack(prefix)
	}
	if t.mid != nil {
		t.mid.collect(nil, prefix, keys)
	}
}

func (t *TSTNode) collect(_ alphabet.Interface, prefix string, keys *queue.StrQ) {
	if t.v != nil {
		keys.PushBack(prefix + string(t.r))
	}
	if t.left != nil {
		t.left.collect(nil, prefix, keys)
	}
	if t.mid != nil {
		t.mid.collect(nil, prefix+string(t.r), keys)
	}
	if t.right != nil {
		t.right.collect(nil, prefix, keys)
	}
}

func (t *TSTNode) KeysMatch(_ alphabet.Interface, pattern []rune, prefix string, keys *queue.StrQ) {
	if len(pattern) == 0 {
		return
	}
	if t.left != nil && (pattern[0] == '.' || pattern[0] < t.r) {
		t.left.KeysMatch(nil, pattern, prefix, keys)
	}
	if t.right != nil && (pattern[0] == '.' || pattern[0] > t.r) {
		t.right.KeysMatch(nil, pattern, prefix, keys)
	}
	if pattern[0] == '.' || pattern[0] == t.r {
		if len(pattern) == 1 {
			if t.v != nil {
				keys.PushBack(prefix + string(t.r))
			}
		} else if t.mid != nil {
			t.mid.KeysMatch(nil, pattern[1:], prefix+string(t.r), keys)
		}
	}
}

func (t *TSTNode) Keys(_ alphabet.Interface, keys *queue.StrQ) {
	t.collect(nil, "", keys)
}

func (t *TSTNode) Compress() error {
	return errors.New("compress not support")
}

func (t *TSTNode) IsCompressed() bool {
	return false
}

func (t *TSTNode) SetVal(v T) {
	t.v = v
}
