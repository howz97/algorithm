package trietree

import (
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/queue"
)

type T interface{}

type TrieNode interface {
	Value() T
	Find(a alphabet.Interface, k []rune) TrieNode
	Insert(a alphabet.Interface, k []rune, v T)
	Delete(a alphabet.Interface, k []rune)
	LongestPrefixOf(a alphabet.Interface, s []rune, d, l int) int
	Collect(a alphabet.Interface, prefix string, keys *queue.StrQ)
	KeysMatch(a alphabet.Interface, pattern []rune, prefix string, keys *queue.StrQ)
}

type Trie struct {
	a    alphabet.Interface
	tree TrieNode
	size int
}

func NewTrie(a alphabet.Interface, node TrieNode) *Trie {
	return &Trie{
		a:    a,
		tree: node,
	}
}

func (t *Trie) Find(k string) interface{} {
	n := t.tree.Find(t.a, []rune(k))
	if n == nil {
		return nil
	}
	return n.Value()
}

func (t *Trie) Insert(k string, v interface{}) {
	t.tree.Insert(t.a, []rune(k), v)
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

func (t *Trie) KeysWithPrefix(prefix string) (keys []string) {
	node := t.tree.Find(t.a, []rune(prefix))
	if node == nil {
		return
	}
	q := queue.NewStrQ()
	node.Collect(t.a, prefix, q)
	for !q.IsEmpty() {
		keys = append(keys, q.Front())
	}
	return
}

func (t *Trie) Keys() []string {
	return t.KeysWithPrefix("")
}

func (t *Trie) KeysMatch(p string) (keys []string) {
	q := queue.NewStrQ()
	t.tree.KeysMatch(t.a, []rune(p), "", q)
	for !q.IsEmpty() {
		keys = append(keys, q.Front())
	}
	return
}

func (t *Trie) Size() int {
	return t.size
}
