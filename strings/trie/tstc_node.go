package trie

import (
	"github.com/howz97/algorithm/basic/queue"
	"github.com/howz97/algorithm/strings/alphabet"
	"github.com/howz97/algorithm/util"
)

// Three direction trie tree that is compressible
type TSTC struct {
	TSTCNode
	compressed bool
}

func (t *TSTC) Compress() error {
	t.compress()
	t.compressed = true
	return nil
}

func (t *TSTC) IsCompressed() bool {
	return t.compressed
}

func (t *TSTC) Upsert(_ alphabet.Interface, k []rune, v T) {
	if t.compressed {
		panic("can not upsert after compress")
	}
	t.TSTCNode.Upsert(nil, k, v)
}

type TSTCNode struct {
	rs               []rune
	v                T
	left, mid, right *TSTCNode
}

func newTSTCNode(r rune) *TSTCNode {
	return &TSTCNode{rs: []rune{r}}
}

func (t *TSTCNode) Value() T {
	return t.v
}

func (t *TSTCNode) Upsert(_ alphabet.Interface, k []rune, v T) {
	if len(k) == 0 {
		panic("empty key")
	}
	switch true {
	case k[0] < t.rs[0]:
		if t.left == nil {
			t.left = newTSTCNode(k[0])
		}
		t.left.Upsert(nil, k, v)
	case k[0] > t.rs[0]:
		if t.right == nil {
			t.right = newTSTCNode(k[0])
		}
		t.right.Upsert(nil, k, v)
	default:
		if len(k) == 1 {
			t.v = v
		} else {
			if t.mid == nil {
				t.mid = newTSTCNode(k[1])
			}
			t.mid.Upsert(nil, k[1:], v)
		}
	}
}

func (t *TSTCNode) compress() {
	if t == nil {
		return
	}
	if t.canShrink() {
		t.rs = append(t.rs, t.mid.rs...)
		t.left = t.mid.left
		t.right = t.mid.right
		t.v = t.mid.v
		t.mid = t.mid.mid
		t.compress()
	} else {
		t.left.compress()
		t.mid.compress()
		t.right.compress()
	}
	return
}

func (t *TSTCNode) canShrink() bool {
	return t.v == nil && t.left == nil && t.right == nil
}

func (t *TSTCNode) Delete(_ alphabet.Interface, k []rune) {
	if len(k) < len(t.rs) {
		return
	}
	if !isRunesEqual(t.rs, k, 0, len(t.rs)-1) {
		return
	}
	i := len(t.rs) - 1
	switch true {
	case k[i] < t.rs[i]:
		if t.left != nil {
			t.left.Delete(nil, k[i:])
		}
	case k[i] > t.rs[i]:
		if t.right != nil {
			t.right.Delete(nil, k[i:])
		}
	default:
		if len(k) == len(t.rs) {
			t.v = nil
		} else if t.mid != nil {
			t.mid.Delete(nil, k[len(t.rs):])
		}
	}
}

// Find 找到k对应的节点(n)，有这个节点不代表k存在，是否存在需要看n.v是否为nil
func (t *TSTCNode) Find(_ alphabet.Interface, k []rune) T {
	if len(k) < len(t.rs) || !isRunesEqual(t.rs, k, 0, len(t.rs)-1) {
		return nil
	}
	i := len(t.rs) - 1
	switch true {
	case k[i] < t.rs[i]:
		if t.left != nil {
			return t.left.Find(nil, k[i:])
		}
	case k[i] > t.rs[i]:
		if t.right != nil {
			return t.right.Find(nil, k[i:])
		}
	default:
		if len(k) == len(t.rs) {
			return t.v
		} else if t.mid != nil {
			return t.mid.Find(nil, k[len(t.rs):])
		}
	}
	return nil
}

func (t *TSTCNode) Locate(_ alphabet.Interface, k []rune) (TrieNode, []rune) {
	if len(k) < len(t.rs) {
		if isRunesEqual2(k, t.rs[:len(k)]) {
			return t, t.rs[len(k):]
		}
		return nil, nil
	}
	if !isRunesEqual(t.rs, k, 0, len(t.rs)-1) {
		return nil, nil
	}
	i := len(t.rs) - 1
	switch true {
	case k[i] < t.rs[i]:
		if t.left != nil {
			return t.left.Locate(nil, k[i:])
		}
	case k[i] > t.rs[i]:
		if t.right != nil {
			return t.right.Locate(nil, k[i:])
		}
	default:
		if len(k) == len(t.rs) {
			return t, nil
		} else if t.mid != nil {
			return t.mid.Locate(nil, k[len(t.rs):])
		}
	}
	return nil, nil
}

func (t *TSTCNode) LongestPrefixOf(_ alphabet.Interface, s []rune, d, l int) int {
	if len(s) < d+len(t.rs) {
		return l
	}
	if !isRunesEqual(s[d:], t.rs, 0, len(t.rs)-1) {
		return l
	}
	i := len(t.rs) - 1
	switch true {
	case s[d+i] < t.rs[i]:
		if t.left != nil {
			return t.left.LongestPrefixOf(nil, s, d+i, l)
		}
	case s[d+i] > t.rs[i]:
		if t.right != nil {
			return t.right.LongestPrefixOf(nil, s, d+i, l)
		}
	default:
		if t.v != nil {
			l = d + len(t.rs)
		}
		if len(s) == d+len(t.rs) {
			return l
		} else if t.mid != nil {
			return t.mid.LongestPrefixOf(nil, s, d+len(t.rs), l)
		}
	}
	return l
}

func (t *TSTCNode) Collect(_ alphabet.Interface, prefix string, keys *queue.SliceQ[string]) {
	if t.v != nil {
		keys.PushBack(prefix)
	}
	if t.mid != nil {
		t.mid.collect(nil, prefix, keys)
	}
}

func (t *TSTCNode) collect(_ alphabet.Interface, prefix string, keys *queue.SliceQ[string]) {
	if t.v != nil {
		keys.PushBack(prefix + string(t.rs))
	}
	if t.left != nil {
		t.left.collect(nil, prefix+string(t.rs[:len(t.rs)-1]), keys)
	}
	if t.mid != nil {
		t.mid.collect(nil, prefix+string(t.rs), keys)
	}
	if t.right != nil {
		t.right.collect(nil, prefix+string(t.rs[:len(t.rs)-1]), keys)
	}
}

func (t *TSTCNode) KeysMatch(_ alphabet.Interface, pattern []rune, prefix string, keys *queue.SliceQ[string]) {
	if len(pattern) < len(t.rs) || !util.IsRunesMatch(pattern[:len(t.rs)-1], t.rs[:len(t.rs)-1]) {
		return
	}
	i := len(t.rs) - 1
	if t.left != nil && (pattern[i] == '.' || pattern[i] < t.rs[i]) {
		t.left.KeysMatch(nil, pattern[len(t.rs)-1:], prefix+string(t.rs[:len(t.rs)-1]), keys)
	}
	if t.right != nil && (pattern[i] == '.' || pattern[i] > t.rs[i]) {
		t.right.KeysMatch(nil, pattern[len(t.rs)-1:], prefix+string(t.rs[:len(t.rs)-1]), keys)
	}
	if util.IsRuneMatch(pattern[i], t.rs[i]) {
		if len(pattern) == len(t.rs) {
			if t.v != nil {
				keys.PushBack(prefix + string(t.rs))
			}
		} else if t.mid != nil {
			t.mid.KeysMatch(nil, pattern[len(t.rs):], prefix+string(t.rs), keys)
		}
	}
}

func (t *TSTCNode) Keys(_ alphabet.Interface, keys *queue.SliceQ[string]) {
	t.collect(nil, "", keys)
}

func (t *TSTCNode) IsCompressed() bool {
	panic("should not be called")
}

func (t *TSTCNode) Compress() error {
	panic("should not be called")
}

func (t *TSTCNode) SetVal(v T) {
	t.v = v
}

func isRunesEqual2(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func isRunesEqual(a, b []rune, lo, hi int) bool {
	for ; lo < hi; lo++ {
		if a[lo] != b[lo] {
			return false
		}
	}
	return true
}
