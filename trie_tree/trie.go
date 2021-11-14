package trietree

import (
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/queue"
)

type T interface{}

type TrieNode interface {
	Value() T
	Find(a alphabet.Interface, k []rune) TrieNode
	Upsert(a alphabet.Interface, k []rune, v T)
	Delete(a alphabet.Interface, k []rune)
	LongestPrefixOf(a alphabet.Interface, s []rune, d, l int) int
	Collect(a alphabet.Interface, prefix string, keys *queue.StrQ)
	KeysMatch(a alphabet.Interface, pattern []rune, prefix string, keys *queue.StrQ)
	Keys(a alphabet.Interface, keys *queue.StrQ)
}

type Trie struct {
	a    alphabet.Interface
	tree TrieNode
	size int
}

func NewTrie(a alphabet.Interface) *Trie {
	return &Trie{
		a:    a,
		tree: newSliceNode(a.R()),
	}
}

func NewTST() *Trie {
	return &Trie{
		a:    alphabet.Unicode,
		tree: newTSTNode('z'),
	}
}

func (t *Trie) Find(k string) interface{} {
	n := t.tree.Find(t.a, []rune(k))
	if n == nil {
		return nil
	}
	return n.Value()
}

func (t *Trie) Upsert(k string, v T) {
	t.tree.Upsert(t.a, []rune(k), v)
	t.size++
}

func (t *Trie) Delete(k string) {
	t.tree.Delete(t.a, []rune(k))
	t.size--
}

func (t *Trie) Contains(k string) bool {
	return t.Find(k) != nil
}

func (t *Trie) IsEmpty() bool {
	return t.size == 0
}

func (t *Trie) LongestPrefixOf(s string) string {
	runes := []rune(s)
	l := t.tree.LongestPrefixOf(t.a, runes, 0, 0)
	return string(runes[:l])
}

func (t *Trie) KeysWithPrefix(prefix string) []string {
	if prefix == "" {
		return t.Keys()
	}
	node := t.tree.Find(t.a, []rune(prefix))
	if node == nil {
		return nil
	}
	q := queue.NewStrQ()
	node.Collect(t.a, prefix, q)
	return q.PopAll()
}

func (t *Trie) Keys() []string {
	q := queue.NewStrQ()
	t.tree.Keys(t.a, q)
	return q.PopAll()
}

func (t *Trie) KeysMatch(p string) []string {
	q := queue.NewStrQ()
	t.tree.KeysMatch(t.a, []rune(p), "", q)
	return q.PopAll()
}

func (t *Trie) Size() int {
	return t.size
}
