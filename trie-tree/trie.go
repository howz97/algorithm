package trietree

import "github.com/zh1014/algorithm/queue"

type Trie struct {
	a alphbt
	tree *trie
	size int
}

func New(a alphbt) *Trie {
	return &Trie{
		a:    a,
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
	t.tree=t.tree.delete(t.a, []rune(k))
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

func (t *Trie)KeysMatch(p string) []string {
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

type trie struct {
	val   interface{}
	next []*trie
}

func (t *trie)find(a alphbt, k []rune) *trie {
	if t == nil {
		return nil
	}
	if len(k) == 0 {
		return t
	}
	return t.next[a.ToIndex(k[0])].find(a, k[1:])
}

func (t *trie)insert(a alphbt, k []rune, v interface{}) *trie {
	if t == nil {
		t = &trie{
			next: make([]*trie, a.R()),
		}
	}
	if len(k) == 0 {
		t.val = v
	}else {
		t.next[a.ToIndex(k[0])] = t.next[a.ToIndex(k[0])].insert(a, k[1:], v)
	}
	return t
}

func (t *trie)delete(a alphbt, k []rune) *trie {
	if t != nil {
		if len(k) == 0  {
			t.val = nil
		}else {
			t.next[a.ToIndex(k[0])] = t.next[a.ToIndex(k[0])].delete(a, k[1:])
		}
	}
	if t.isEmpty() {
		return nil
	}
	return t
}

func (t *trie)isEmpty() bool {
	if t == nil {
		return true
	}
	if t.val != nil {
		return false
	}
	for i := range t.next {
		if t.next[i] != nil {
			return false
		}
	}
	return true
}

func (t *trie)contains(a alphbt, k []rune) bool {
	if t == nil {
		return false
	}
	if len(k) == 0 {
		return t.val != nil
	}else {
		return t.next[a.ToIndex(k[0])].contains(a, k[1:])
	}
}

// longestPrefixOf 找出t的所有匹配s[d:]的前缀的key中最长的那一个
// 返回值length代表s的前length个rune就是这个要找的key
func (t *trie)longestPrefixOf(a alphbt,s []rune, d int, length int) int {
	if t == nil {
		return length
	}
	if t.val != nil {
		length = d
	}
	if len(s) == d {
		return length
	}
	return t.next[a.ToIndex(s[d])].longestPrefixOf(a, s, d+1, length)
}

func (t *trie)keysWithPrefix(a alphbt, p string) *queue.StrQ {
	keysQ := queue.NewStrQ()
	t.find(a, []rune(p)).collect(a, p, keysQ)
	return keysQ
}

// collect collects all keys of t and put them into StrQ
// p is the prefix record
func (t *trie)collect(a alphbt, p string, keys *queue.StrQ){
	if t == nil {
		return
	}
	if t.val != nil {
		keys.PushBack(p)
	}
	for i := range t.next {
		t.next[i].collect(a, p+string(a.ToRune(i)), keys)
	}
}

func (t *trie)keysMatch(a alphbt,pattern []rune, prefix string,keys *queue.StrQ) {
	if t == nil {
		return
	}
	if len(pattern) == 0 {
		if t.val != nil {
			keys.PushBack(prefix)
		}
		return
	}
	if pattern[0] == rune('.') {
		for i := range t.next {
			t.next[i].keysMatch(a, pattern[1:], prefix+string(a.ToRune(i)), keys)
		}
	}else {
		t.next[a.ToIndex(pattern[0])].keysMatch(a, pattern[1:], prefix+string(pattern[0]), keys)
	}
	return
}