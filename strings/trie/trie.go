package trie

import (
	"github.com/howz97/algorithm/basic/queue"
	"github.com/howz97/algorithm/strings/alphabet"
)

type T interface{}

type TrieNode interface {
	SetVal(v T)
	Find(a alphabet.Interface, k []rune) T
	Upsert(a alphabet.Interface, k []rune, v T)
	Delete(a alphabet.Interface, k []rune)
	Locate(a alphabet.Interface, k []rune) (TrieNode, []rune)
	LongestPrefixOf(a alphabet.Interface, s []rune, d, l int) int
	Collect(a alphabet.Interface, prefix string, keys *queue.SliceQ[string])
	KeysMatch(a alphabet.Interface, pattern []rune, prefix string, keys *queue.SliceQ[string])
	Keys(a alphabet.Interface, keys *queue.SliceQ[string])
	Compress() error
	IsCompressed() bool
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

func NewTSTC() *Trie {
	return &Trie{
		a:    alphabet.Unicode,
		tree: &TSTC{TSTCNode: *newTSTCNode('z')},
	}
}

func (t *Trie) Find(k string) T {
	return t.tree.Find(t.a, []rune(k))
}

// Upsert update or insert a value
// Insert a new node if the node not exists
func (t *Trie) Upsert(k string, v T) {
	t.tree.Upsert(t.a, []rune(k), v)
	t.size++ // fixme: update do not inc size
}

// Update or set a value
// Can not insert any new node
func (t *Trie) Update(k string, v T) {
	node, runes := t.tree.Locate(t.a, []rune(k))
	if node == nil || len(runes) != 0 {
		return
	}
	node.SetVal(v)
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
	node, runes := t.tree.Locate(t.a, []rune(prefix))
	if node == nil {
		return nil
	}
	prefix += string(runes)
	q := queue.NewSliceQ[string](0)
	node.Collect(t.a, prefix, q)
	return q.Drain()
}

func (t *Trie) Keys() []string {
	q := queue.NewSliceQ[string](0)
	t.tree.Keys(t.a, q)
	return q.Drain()
}

func (t *Trie) KeysMatch(p string) []string {
	q := queue.NewSliceQ[string](0)
	t.tree.KeysMatch(t.a, []rune(p), "", q)
	return q.Drain()
}

func (t *Trie) Size() int {
	return t.size
}

func (t *Trie) Compress() error {
	return t.tree.Compress()
}

func (t *Trie) IsCompressed() bool {
	return t.tree.IsCompressed()
}