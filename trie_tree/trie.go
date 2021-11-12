package trietree

import "github.com/howz97/algorithm/queue"

type Interface interface {
}

type Trie struct {
	a    alphbt
	tree *trie
	size int
}

func NewTrie(a alphbt) *Trie {
	return &Trie{
		a: a,
	}
}

func (t *Trie) Find(k string) interface{} {
	n := t.tree.find(t.a, []rune(k))
	if n == nil {
		return nil
	}
	return n.val
}

func (t *Trie) Insert(k string, v interface{}) {
	t.tree = t.tree.insert(t.a, []rune(k), v)
	t.size++
}

func (t *Trie) Delete(k string) {
	t.tree = t.tree.delete(t.a, []rune(k))
	t.size--
}

func (t *Trie) Contains(k string) bool {
	return t.tree.contains(t.a, []rune(k))
}

func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

func (t *Trie) LongestPrefixOf(s string) string {
	runes := []rune(s)
	l := t.tree.longestPrefixOf(t.a, runes, 0, 0)
	return string(runes[:l])
}

func (t *Trie) KeysWithPrefix(p string) []string {
	keys := make([]string, 0)
	keysQ := queue.NewStrQ()
	t.tree.find(t.a, []rune(p)).collect(t.a, p, keysQ)
	for !keysQ.IsEmpty() {
		keys = append(keys, keysQ.Front())
	}
	return keys
}

func (t *Trie) KeysMatch(p string) []string {
	keys := make([]string, 0)
	keysQ := queue.NewStrQ()
	t.tree.keysMatch(t.a, []rune(p), "", keysQ)
	for !keysQ.IsEmpty() {
		keys = append(keys, keysQ.Front())
	}
	return keys
}

func (t *Trie) Keys() []string {
	return t.KeysWithPrefix("")
}

func (t *Trie) Size() int {
	return t.size
}
