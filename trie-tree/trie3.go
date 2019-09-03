package trietree

import (
	"github.com/zh1014/algorithm/queue"
)

type Trie3 struct {
	a    alphbt
	tree *trie3
	size int
}

func NewTrie3(a alphbt) *Trie3 {
	return &Trie3{
		a: a,
	}
}

func (t *Trie3) Insert(k string, v interface{}) {
	t.tree = t.tree.insert(t.a, []rune(k), v)
	t.size++
}

func (t *Trie3) Find(k string) interface{} {
	f := t.tree.find(t.a, []rune(k))
	if f == nil {
		// this trie-tree node not exist
		return nil
	}
	return f.v
}

func (t *Trie3) Delete(k string) {
	t.tree = t.tree.delete(t.a, []rune(k))
	t.size--
}

func (t *Trie3) Contains(k string) bool {
	return t.tree.contains(t.a, []rune(k))
}

func (t *Trie3) IsEmpty() bool {
	return t.size == 0
}

func (t *Trie3) Size() int {
	return t.size
}

func (t *Trie3) LongestPrefixOf(s string) string {
	runes := []rune(s)
	return string(runes[:t.tree.longestPrefixOf(t.a, runes, 0,0)])
}

func (t *Trie3) KeysWithPrefix(p string) []string {
	keys := make([]string, 0)
	keysQ := queue.NewStrQ()
	if p == "" {
		t.tree.collect(t.a, p, keysQ)
	}else {
		t.tree.find(t.a, []rune(p)).collect(t.a, p, keysQ)
	}
	for !keysQ.IsEmpty() {
		keys = append(keys, keysQ.Front())
	}
	return keys
}

func (t *Trie3) KeysMatch(p string) []string {
	keys := make([]string, 0)
	keysQ := queue.NewStrQ()
	t.tree.keysMatch(t.a, []rune(p), "", keysQ)
	for !keysQ.IsEmpty() {
		keys = append(keys, keysQ.Front())
	}
	return keys
}

func (t *Trie3) Keys() []string {
	return t.KeysWithPrefix("")
}

type trie3 struct {
	r     rune
	left  *trie3
	mid   *trie3
	right *trie3
	v     interface{}
}

func (t *trie3) insert(a alphbt, k []rune, v interface{}) *trie3 {
	if len(k) == 0 {
		panic("empty key")
	}
	if t == nil {
		t = &trie3{
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

func (t *trie3) delete(a alphbt, k []rune) *trie3 {
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
	if t.isEmpty() {
		t = nil
	}
	return t
}

func (t *trie3) isEmpty() bool {
	if t == nil || (t.v == nil && t.left == nil && t.mid == nil && t.right == nil) {
		return true
	}
	return false
}

func (t *trie3) contains(a alphbt, k []rune) bool {
	f := t.find(a, k)
	return f != nil && f.v != nil
}

// find 找到k对应的节点(n)，有这个节点不代表k存在，是否存在需要看n.v是否为nil
func (t *trie3) find(a alphbt, k []rune) *trie3 {
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

func (t *trie3) longestPrefixOf(a alphbt, s []rune, d,length int) int {
	if len(s) == 0 {
		panic("empty s")
	}
	if t == nil {
		return length
	}
	switch true {
	case a.ToIndex(s[d]) < a.ToIndex(t.r):
		return t.left.longestPrefixOf(a, s, d,length)
	case a.ToIndex(s[d]) > a.ToIndex(t.r):
		return t.right.longestPrefixOf(a, s, d,length)
	default:
		if t.v != nil {
			length = d+1
		}
		if d < len(s)-1 {
			return t.mid.longestPrefixOf(a, s, d+1,length)
		} else { // len(k) == 1
			return length
		}
	}
}

func (t *trie3) collect(a alphbt, p string, keys *queue.StrQ) {
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

func (t *trie3) keysMatch(a alphbt, pattern []rune, prefix string, keys *queue.StrQ) {
	if len(pattern) == 0 {
		panic("empty pattern")
	}
	if t == nil {
		return
	}
	if pattern[0] == rune('.')||a.ToIndex(pattern[0]) < a.ToIndex(t.r) {
		t.left.keysMatch(a, pattern, prefix, keys)
	}
	if pattern[0] == rune('.')||a.ToIndex(pattern[0]) > a.ToIndex(t.r) {
		t.right.keysMatch(a, pattern, prefix, keys)
	}
	if pattern[0] == rune('.')||a.ToIndex(pattern[0]) == a.ToIndex(t.r) {
		if len(pattern) > 1 {
			t.mid.keysMatch(a, pattern[1:], prefix+string(t.r), keys)
		} else { // len(pattern) == 1
			keys.PushBack(prefix + string(t.r))
		}
	}
}
