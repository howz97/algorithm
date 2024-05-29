package trie

import (
	"github.com/howz97/algorithm/basic/queue"
	"github.com/howz97/algorithm/util"
)

func NewTernary[T any]() *Ternary[T] {
	return &Ternary[T]{
		root: newTernary[T]('r'),
	}
}

type Ternary[T any] struct {
	root *tnode[T]
}

func (t *Ternary[T]) Find(key string) *T {
	node := t.root.find([]rune(key))
	if node == nil {
		return nil
	}
	return node.val
}

func (t *Ternary[T]) Upsert(key string, val T) bool {
	return t.root.upsert([]rune(key), val)
}

func (t *Ternary[T]) Delete(key string) {
	node := t.root.find([]rune(key))
	if node != nil {
		node.val = nil
		// todo: remove empty node
	}
}

func (t *Ternary[T]) LongestPrefixOf(s string) string {
	runes := []rune(s)
	l := t.root.longestPrefixOf(runes, 0, 0)
	return string(runes[:l])
}

func (t *Ternary[T]) KeysWithPrefix(prefix string) []string {
	q := queue.NewQueue[string](0)
	var node *tnode[T]
	if prefix == "" {
		node = t.root
	} else {
		var i int
		node, i = t.root.locate([]rune(prefix))
		if node == nil {
			return nil
		}
		if i < len(node.rs)-1 {
			prefix += string(node.rs[i+1:])
		}
		if node.val != nil {
			q.PushBack(prefix)
		}
		node = node.mid
	}
	if node != nil {
		node.collect(prefix, q)
	}
	return q.Clone()
}

func (t *Ternary[T]) KeysMatch(p string) []string {
	q := queue.NewQueue[string](0)
	t.root.keysMatch([]rune(p), "", q)
	return q.Clone()
}

func (t *Ternary[T]) Compress() {
	t.root.compress()
}

type tnode[T any] struct {
	rs               []rune
	val              *T
	left, mid, right *tnode[T]
}

func newTernary[T any](r rune) *tnode[T] {
	return &tnode[T]{rs: []rune{r}}
}

func (t *tnode[T]) find(k []rune) *tnode[T] {
	node, i := t.locate(k)
	if node == nil || i != len(node.rs)-1 {
		return nil
	}
	return node
}

func (t *tnode[T]) locate(k []rune) (*tnode[T], int) {
	if len(k) < len(t.rs) {
		if string(k) == string(t.rs[:len(k)]) {
			// target node has already been compressed
			return t, len(k) - 1
		}
		return nil, 0
	}
	if !isEqual(t.rs, k, 0, len(t.rs)-1) {
		return nil, 0
	}
	i := len(t.rs) - 1
	switch true {
	case k[i] < t.rs[i]:
		if t.left != nil {
			return t.left.locate(k[i:])
		}
	case k[i] > t.rs[i]:
		if t.right != nil {
			return t.right.locate(k[i:])
		}
	default:
		if len(k) == len(t.rs) {
			return t, len(t.rs) - 1
		} else if t.mid != nil {
			return t.mid.locate(k[i+1:])
		}
	}
	return nil, 0
}

func (t *tnode[T]) upsert(k []rune, v T) (ok bool) {
	if len(k) < len(t.rs) || !isEqual(t.rs, k, 0, len(t.rs)-1) {
		// can not upsert compressed node
		return false
	}
	i := len(t.rs) - 1
	if k[i] < t.rs[i] {
		if t.left == nil {
			t.left = newTernary[T](k[i])
		}
		ok = t.left.upsert(k[i:], v)
	} else if k[i] > t.rs[i] {
		if t.right == nil {
			t.right = newTernary[T](k[i])
		}
		ok = t.right.upsert(k[i:], v)
	} else {
		if len(k) == len(t.rs) {
			t.val = &v
			ok = true
		} else {
			if t.mid == nil {
				t.mid = newTernary[T](k[i+1])
			}
			ok = t.mid.upsert(k[i+1:], v)
		}
	}
	return
}

func (t *tnode[T]) compress() {
	if t == nil {
		return
	}
	if t.canShrink() {
		t.rs = append(t.rs, t.mid.rs...)
		t.left = t.mid.left
		t.right = t.mid.right
		t.val = t.mid.val
		t.mid = t.mid.mid
		t.compress()
	} else {
		t.left.compress()
		t.mid.compress()
		t.right.compress()
	}
}

func (t *tnode[T]) canShrink() bool {
	return t.val == nil && t.left == nil && t.right == nil
}

func (t *tnode[T]) longestPrefixOf(s []rune, d, l int) int {
	if len(s) < d+len(t.rs) || !isEqual(s[d:], t.rs, 0, len(t.rs)-1) {
		return l
	}
	i := len(t.rs) - 1
	switch true {
	case s[d+i] < t.rs[i]:
		if t.left != nil {
			return t.left.longestPrefixOf(s, d+i, l)
		}
	case s[d+i] > t.rs[i]:
		if t.right != nil {
			return t.right.longestPrefixOf(s, d+i, l)
		}
	default:
		if t.val != nil {
			l = d + len(t.rs)
		}
		if len(s) > d+len(t.rs) && t.mid != nil {
			return t.mid.longestPrefixOf(s, d+len(t.rs), l)
		}
	}
	return l
}

func (t *tnode[T]) collect(prefix string, keys *queue.Queue[string]) {
	if t.val != nil {
		keys.PushBack(prefix + string(t.rs))
	}
	if t.left != nil {
		t.left.collect(prefix+string(t.rs[:len(t.rs)-1]), keys)
	}
	if t.mid != nil {
		t.mid.collect(prefix+string(t.rs), keys)
	}
	if t.right != nil {
		t.right.collect(prefix+string(t.rs[:len(t.rs)-1]), keys)
	}
}

func (t *tnode[T]) keysMatch(pattern []rune, prefix string, keys *queue.Queue[string]) {
	if len(pattern) < len(t.rs) || !util.IsRunesMatch(pattern[:len(t.rs)-1], t.rs[:len(t.rs)-1]) {
		return
	}
	i := len(t.rs) - 1
	if t.left != nil && (pattern[i] == '.' || pattern[i] < t.rs[i]) {
		t.left.keysMatch(pattern[len(t.rs)-1:], prefix+string(t.rs[:len(t.rs)-1]), keys)
	}
	if t.right != nil && (pattern[i] == '.' || pattern[i] > t.rs[i]) {
		t.right.keysMatch(pattern[len(t.rs)-1:], prefix+string(t.rs[:len(t.rs)-1]), keys)
	}
	if util.IsRuneMatch(pattern[i], t.rs[i]) {
		if len(pattern) == len(t.rs) {
			if t.val != nil {
				keys.PushBack(prefix + string(t.rs))
			}
		} else if t.mid != nil {
			t.mid.keysMatch(pattern[len(t.rs):], prefix+string(t.rs), keys)
		}
	}
}

func isEqual(a, b []rune, lo, hi int) bool {
	for ; lo < hi; lo++ {
		if a[lo] != b[lo] {
			return false
		}
	}
	return true
}
