package tst_optimized

import (
	"errors"
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/queue"
)

type Tst2 struct {
	Alpb       alphabet.IAlphabet
	tree       *tst2
	compressed bool
}

func NewTst2(a alphabet.IAlphabet) *Tst2 {
	return &Tst2{
		Alpb: a,
	}
}

func (t *Tst2) Insert(k string, v interface{}) error {
	if t.Compressed() {
		panic("can not insert into compressed trie-tree")
	}
	t.tree = t.tree.insert(t.Alpb, []rune(k), v)
	return nil
}

func (t *Tst2) Compress() error {
	if t.Compressed() {
		return errors.New(" Duplicate compress ")
	}
	t.tree.compress()
	t.compressed = true
	return nil
}

func (t *Tst2) Compressed() bool {
	return t.compressed
}

func (t *Tst2) Find(k string) interface{} {
	f := t.tree.find(t.Alpb, []rune(k))
	if f == nil {
		// this trie-tree node not exist
		return nil
	}
	return f.v
}

func (t *Tst2) Delete(k string) {
	t.tree = t.tree.delete(t.Alpb, []rune(k))
}

func (t *Tst2) Contains(k string) bool {
	return t.tree.contains(t.Alpb, []rune(k))
}

func (t *Tst2) IsEmpty() bool {
	return t.tree == nil
}

func (t *Tst2) LongestPrefixOf(s string) string {
	runes := []rune(s)
	return string(runes[:t.tree.longestPrefixOf(t.Alpb, runes, 0, 0)])
}

func (t *Tst2) KeysWithPrefix(p string) []string {
	keys := make([]string, 0)
	keysQ := queue.NewStrQ()
	if p == "" {
		t.tree.collect(t.Alpb, p, keysQ)
	} else {
		f, i := t.tree.locate(t.Alpb, []rune(p))
		if f == nil {
			return nil
		}
		if i == len(f.rs)-1 {
			if f.v != nil {
				keysQ.PushBack(p)
			}
			f.mid.collect(t.Alpb, p, keysQ)
		} else {
			runes := []rune(p)
			f.collect(t.Alpb, string(runes[:len(runes)-i-1]), keysQ)
		}
	}
	for !keysQ.IsEmpty() {
		keys = append(keys, keysQ.Front())
	}
	return keys
}

func (t *Tst2) KeysMatch(p string) []string {
	keys := make([]string, 0)
	keysQ := queue.NewStrQ()
	t.tree.keysMatch(t.Alpb, []rune(p), []rune(""), keysQ)
	for !keysQ.IsEmpty() {
		keys = append(keys, keysQ.Front())
	}
	return keys
}

func (t *Tst2) Keys() []string {
	return t.KeysWithPrefix("")
}

type tst2 struct {
	rs               []rune
	v                interface{}
	left, mid, right *tst2
}

func (t *tst2) insert(a alphabet.IAlphabet, k []rune, v interface{}) *tst2 {
	if len(k) == 0 {
		panic("empty key")
	}
	if t == nil {
		t = &tst2{
			rs: k[:1],
		}
	}
	switch true {
	case a.ToIndex(k[0]) < a.ToIndex(t.rs[0]):
		t.left = t.left.insert(a, k, v)
	case a.ToIndex(k[0]) > a.ToIndex(t.rs[0]):
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

func (t *tst2) compress() {
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

func (t *tst2) canShrink() bool {
	return t.v == nil && t.left == nil && t.right == nil
}

func (t *tst2) delete(a alphabet.IAlphabet, k []rune) *tst2 {
	if len(k) == 0 {
		panic("empty key")
	}
	if t == nil || !equal(t.rs[:len(t.rs)-1], k[:len(t.rs)-1]) {
		return t
	}
	lastRune := len(t.rs) - 1
	switch true {
	case a.ToIndex(k[lastRune]) < a.ToIndex(t.rs[lastRune]):
		t.left = t.left.delete(a, k[lastRune:])
	case a.ToIndex(k[lastRune]) > a.ToIndex(t.rs[lastRune]):
		t.right = t.right.delete(a, k[lastRune:])
	default:
		if len(k) > len(t.rs) {
			t.mid = t.mid.delete(a, k[len(t.rs):])
		} else { // len(k) == 1
			t.v = nil
		}
	}
	if t.isEmpty() {
		t = nil
	}
	return t
}

func (t *tst2) isEmpty() bool {
	if t.v == nil && t.mid == nil && t.left == nil && t.right == nil {
		return true
	}
	return false
}

func (t *tst2) contains(a alphabet.IAlphabet, k []rune) bool {
	f := t.find(a, k)
	return f != nil && f.v != nil
}

// find 找到k对应的节点(n)，有这个节点不代表k存在，是否存在需要看n.v是否为nil
func (t *tst2) find(a alphabet.IAlphabet, k []rune) *tst2 {
	if t == nil || len(k) < len(t.rs) || !equal(t.rs[:len(t.rs)-1], k[:len(t.rs)-1]) {
		return nil
	}
	lastRune := len(t.rs) - 1
	switch true {
	case a.ToIndex(k[lastRune]) < a.ToIndex(t.rs[lastRune]):
		return t.left.find(a, k[lastRune:])
	case a.ToIndex(k[lastRune]) > a.ToIndex(t.rs[lastRune]):
		return t.right.find(a, k[lastRune:])
	default:
		if len(k) > len(t.rs) {
			return t.mid.find(a, k[len(t.rs):])
		} else { // len(k) == len(t.rs)
			return t
		}
	}
}

func (t *tst2) locate(a alphabet.IAlphabet, k []rune) (l *tst2, i int) {
	if t == nil {
		return nil, 0
	}
	if len(k) < len(t.rs) {
		if equal(k, t.rs[:len(k)]) {
			return t, len(k) - 1
		}
		return nil, 0
	}
	if !equal(t.rs[:len(t.rs)-1], k[:len(t.rs)-1]) {
		return nil, 0
	}
	lastRune := len(t.rs) - 1
	switch true {
	case a.ToIndex(k[lastRune]) < a.ToIndex(t.rs[lastRune]):
		return t.left.locate(a, k[lastRune:])
	case a.ToIndex(k[lastRune]) > a.ToIndex(t.rs[lastRune]):
		return t.right.locate(a, k[lastRune:])
	default:
		if len(k) > len(t.rs) {
			return t.mid.locate(a, k[len(t.rs):])
		} else { // len(k) == len(t.rs)
			return t, lastRune
		}
	}
}

func (t *tst2) longestPrefixOf(a alphabet.IAlphabet, s []rune, d, length int) int {
	if len(s) == 0 {
		panic("empty s")
	}
	if t == nil || !equal(s[d:d+len(t.rs)-1], t.rs[:len(t.rs)-1]) {
		return length
	}
	lastRune := len(t.rs) - 1
	switch true {
	case a.ToIndex(s[d+lastRune]) < a.ToIndex(t.rs[lastRune]):
		return t.left.longestPrefixOf(a, s, d+lastRune, length)
	case a.ToIndex(s[d+lastRune]) > a.ToIndex(t.rs[lastRune]):
		return t.right.longestPrefixOf(a, s, d+lastRune, length)
	default:
		if t.v != nil {
			length = d + len(t.rs)
		}
		if d+len(t.rs) < len(s) {
			return t.mid.longestPrefixOf(a, s, d+len(t.rs), length)
		} else {
			return length
		}
	}
}

func (t *tst2) collect(a alphabet.IAlphabet, p string, keys *queue.StrQ) {
	if t == nil {
		return
	}
	if t.v != nil {
		keys.PushBack(p + string(t.rs))
	}
	t.left.collect(a, p+string(dropLastR(t.rs)), keys)
	t.mid.collect(a, p+string(t.rs), keys)
	t.right.collect(a, p+string(dropLastR(t.rs)), keys)
}

func dropLastR(rs []rune) []rune {
	return rs[:len(rs)-1]
}

func (t *tst2) keysMatch(a alphabet.IAlphabet, pattern []rune, prefix []rune, keys *queue.StrQ) {
	if t == nil || len(pattern) == 0 || len(pattern) < len(t.rs) || !match(pattern[:len(t.rs)-1], t.rs[:len(t.rs)-1]) {
		return
	}
	lastRune := len(t.rs) - 1
	if pattern[lastRune] == '.' || a.ToIndex(pattern[lastRune]) < a.ToIndex(t.rs[lastRune]) {
		t.left.keysMatch(a, dropFirstN(pattern, len(t.rs)-1), append(prefix, t.rs[:len(t.rs)-1]...), keys)
	}
	if pattern[lastRune] == '.' || a.ToIndex(pattern[lastRune]) > a.ToIndex(t.rs[lastRune]) {
		t.right.keysMatch(a, dropFirstN(pattern, len(t.rs)-1), append(prefix, t.rs[:len(t.rs)-1]...), keys)
	}
	if pattern[lastRune] == '.' || a.ToIndex(pattern[lastRune]) == a.ToIndex(t.rs[lastRune]) {
		if len(pattern) > len(t.rs) {
			t.mid.keysMatch(a, dropFirstN(pattern, len(t.rs)), append(prefix, t.rs...), keys)
		} else if t.v != nil {
			keys.PushBack(string(append(prefix, t.rs...)))
		}
	}
}

func dropFirstN(rs []rune, n int) []rune {
	return rs[n:]
}

func equal(a, b []rune) bool {
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

func match(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] && a[i] != rune('.') && b[i] != rune('.') {
			return false
		}
	}
	return true
}
