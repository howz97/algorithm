package tst

import (
	"github.com/zh1014/algorithm/queue"
)

type Tst struct {
	a    alphbt
	tree *tst
	size int
}

func NewTst(a alphbt) *Tst {
	return &Tst{
		a: a,
	}
}

func (t *Tst) Insert(k string, v interface{}) {
	t.tree = t.tree.insert(t.a, []rune(k), v)
	t.size++
}

func (t *Tst) Find(k string) interface{} {
	f := t.tree.find(t.a, []rune(k))
	if f == nil {
		// this trie-tree node not exist
		return nil
	}
	return f.v
}

func (t *Tst) Delete(k string) {
	t.tree = t.tree.delete(t.a, []rune(k))
	t.size--
}

func (t *Tst) Contains(k string) bool {
	return t.tree.contains(t.a, []rune(k))
}

func (t *Tst) IsEmpty() bool {
	return t.size == 0
}

func (t *Tst) Size() int {
	return t.size
}

func (t *Tst) LongestPrefixOf(s string) string {
	runes := []rune(s)
	return string(runes[:t.tree.longestPrefixOf(t.a, runes, 0, 0)])
}

func (t *Tst) KeysWithPrefix(p string) []string {
	keys := make([]string, 0)
	keysQ := queue.NewStrQ()
	if p == "" {
		t.tree.collect(t.a, p, keysQ)
	} else {
		f := t.tree.find(t.a, []rune(p))
		if f == nil {
			return nil
		}
		if f.v != nil {
			keysQ.PushBack(p)
		}
		f.mid.collect(t.a, p, keysQ)
	}
	for !keysQ.IsEmpty() {
		keys = append(keys, keysQ.Front())
	}
	return keys
}

func (t *Tst) KeysMatch(p string) []string {
	keys := make([]string, 0)
	keysQ := queue.NewStrQ()
	t.tree.keysMatch(t.a, []rune(p), "", keysQ)
	for !keysQ.IsEmpty() {
		keys = append(keys, keysQ.Front())
	}
	return keys
}

func (t *Tst) Keys() []string {
	return t.KeysWithPrefix("")
}

type tst struct {
	r                rune
	v                interface{}
	left, mid, right *tst
}

func (t *tst) insert(a alphbt, k []rune, v interface{}) *tst {
	if len(k) == 0 {
		panic("empty key")
	}
	if t == nil {
		t = &tst{
			r: k[0],
		}
	}
	switch true {
	case a.ToIndex(k[0]) < a.ToIndex(t.r):
		t.left = t.left.insert(a, k, v)
	case a.ToIndex(k[0]) > a.ToIndex(t.r):
		t.right = t.right.insert(a, k, v)
	default:
		if len(k) > 1 {
			t.mid = t.mid.insert(a, k[1:], v)
		} else { // len(k) == 1
			t.v = v
		}
	}
	return t
}

func (t *tst) delete(a alphbt, k []rune) *tst {
	if len(k) == 0 {
		panic("empty key")
	}
	if t == nil {
		return nil
	}
	switch true {
	case a.ToIndex(k[0]) < a.ToIndex(t.r):
		t.left = t.left.delete(a, k)
	case a.ToIndex(k[0]) > a.ToIndex(t.r):
		t.right = t.right.delete(a, k)
	default:
		if len(k) > 1 {
			t.mid = t.mid.delete(a, k[1:])
		} else { // len(k) == 1
			t.v = nil
		}
	}
	if t.isEmpty() { // 该节点没有存东西，也没有后继节点，类似BST删除操作
		if t.left == nil {
			t = t.right
		} else if t.right == nil {
			t = t.left
		}
	}
	return t
}

func (t *tst) isEmpty() bool {
	if t.v == nil && t.mid == nil {
		return true
	}
	return false
}

func (t *tst) contains(a alphbt, k []rune) bool {
	f := t.find(a, k)
	return f != nil && f.v != nil
}

// find 找到k对应的节点(n)，有这个节点不代表k存在，是否存在需要看n.v是否为nil
func (t *tst) find(a alphbt, k []rune) *tst {
	if len(k) == 0 {
		panic("empty key")
	}
	if t == nil {
		return nil
	}
	switch true {
	case a.ToIndex(k[0]) < a.ToIndex(t.r):
		return t.left.find(a, k)
	case a.ToIndex(k[0]) > a.ToIndex(t.r):
		return t.right.find(a, k)
	default:
		if len(k) > 1 {
			return t.mid.find(a, k[1:])
		} else { // len(k) == 1
			return t
		}
	}
}

func (t *tst) longestPrefixOf(a alphbt, s []rune, d, length int) int {
	if len(s) == 0 {
		panic("empty s")
	}
	if t == nil {
		return length
	}
	switch true {
	case a.ToIndex(s[d]) < a.ToIndex(t.r):
		return t.left.longestPrefixOf(a, s, d, length)
	case a.ToIndex(s[d]) > a.ToIndex(t.r):
		return t.right.longestPrefixOf(a, s, d, length)
	default:
		if t.v != nil {
			length = d + 1
		}
		if d < len(s)-1 {
			return t.mid.longestPrefixOf(a, s, d+1, length)
		} else { // len(k) == 1
			return length
		}
	}
}

func (t *tst) collect(a alphbt, p string, keys *queue.StrQ) {
	if t == nil {
		return
	}
	if t.v != nil {
		keys.PushBack(p + string(t.r))
	}
	t.left.collect(a, p, keys)
	t.mid.collect(a, p+string(t.r), keys)
	t.right.collect(a, p, keys)
}

func (t *tst) keysMatch(a alphbt, pattern []rune, prefix string, keys *queue.StrQ) {
	if len(pattern) == 0 {
		panic("empty pattern")
	}
	if t == nil {
		return
	}
	if pattern[0] == rune('.') || a.ToIndex(pattern[0]) < a.ToIndex(t.r) {
		t.left.keysMatch(a, pattern, prefix, keys)
	}
	if pattern[0] == rune('.') || a.ToIndex(pattern[0]) > a.ToIndex(t.r) {
		t.right.keysMatch(a, pattern, prefix, keys)
	}
	if pattern[0] == rune('.') || a.ToIndex(pattern[0]) == a.ToIndex(t.r) {
		if len(pattern) > 1 {
			t.mid.keysMatch(a, pattern[1:], prefix+string(t.r), keys)
		} else if t.v != nil { // len(pattern) == 1
			keys.PushBack(prefix + string(t.r))
		}
	}
}
