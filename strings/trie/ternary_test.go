package trie

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/howz97/algorithm/util"
)

func TestTernary(t *testing.T) {
	trie := NewTernary[string]()
	Ternary_Integer(t, trie, DataASCII)
	trie = NewTernary[string]()
	Ternary_Integer(t, trie, DataChn)
}

func TestCompress(t *testing.T) {
	trie := NewTernary[string]()
	Ternary_UpsertAndFind(t, trie, DataChn.m)
	trie.Compress()
	for _, pre := range DataChn.prefix {
		Ternary_KeysWithPrefix(t, trie, pre)
	}
	for _, p := range DataChn.pattern {
		Ternary_KeysMatch(t, trie, p)
	}
	for _, l := range DataChn.long {
		Ternary_LongestPrefixOf(t, trie, l)
	}
	Ternary_UpdateAndFind(t, trie, DataChn.m)
	Ternary_Delete(t, trie, DataChn.m)
}

func Ternary_Integer(t *testing.T, trie *Ternary[string], data Data) {
	Ternary_UpsertAndFind(t, trie, data.m)
	for _, pre := range data.prefix {
		Ternary_KeysWithPrefix(t, trie, pre)
	}
	for _, p := range data.pattern {
		Ternary_KeysMatch(t, trie, p)
	}
	for _, l := range data.long {
		Ternary_LongestPrefixOf(t, trie, l)
	}
	Ternary_UpdateAndFind(t, trie, data.m)
	Ternary_Delete(t, trie, data.m)
}

func Ternary_Delete(t *testing.T, trie *Ternary[string], m map[string]string) {
	for k := range m {
		trie.Delete(k)
	}
	for k := range m {
		if trie.Find(k) != nil {
			t.Fatalf("Delete(%s) failed", k)
		}
	}
}

func Ternary_UpsertAndFind(t *testing.T, trie *Ternary[string], m map[string]string) {
	for k, v := range m {
		trie.Upsert(k, v)
	}
	for k, v := range m {
		got := *trie.Find(k)
		if got != v {
			t.Fatalf("Find(%s)==%s, should be %s", k, got, v)
		}
	}
}

func Ternary_UpdateAndFind(t *testing.T, trie *Ternary[string], m map[string]string) {
	for k, v := range m {
		if !trie.Upsert(k, v+"#") {
			t.Fatalf("can not update key(%s)", k)
		}
	}
	for k, v := range m {
		v = v + "#"
		got := *trie.Find(k)
		if got != v {
			t.Fatalf("Find(%s)==%s, should be %s", k, got, v)
		}
	}
}

func Ternary_KeysWithPrefix(t *testing.T, trie *Ternary[string], prefix string) {
	var correct []string
	for _, k := range trie.KeysWithPrefix("") {
		if strings.HasPrefix(k, prefix) {
			correct = append(correct, k)
		}
	}
	got := trie.KeysWithPrefix(prefix)
	if !StringSliceEqual(correct, got) {
		t.Fatalf("KeysWithPrefix(%s) two set not equal: %v, %v", prefix, correct, got)
	}
}

func Ternary_LongestPrefixOf(t *testing.T, trie *Ternary[string], str string) {
	k := trie.LongestPrefixOf(str)
	if !strings.HasPrefix(str, k) {
		t.Fatalf("key %s is not prefix of %s", k, str)
	}
	if k != "" && trie.Find(k) == nil {
		t.Fatalf("LongestPrefixOf(%s) key %s not exist", str, k)
	}

	str = strings.TrimPrefix(str, k)
	for _, r := range str {
		k += string(r)
		if trie.Find(k) != nil {
			t.Fatalf("longger key %s exist", k)
		}
	}
}

func Ternary_KeysMatch(t *testing.T, trie *Ternary[string], pattern string) {
	var correct []string
	for _, k := range trie.KeysWithPrefix("") {
		if util.IsRunesMatch([]rune(pattern), []rune(k)) {
			correct = append(correct, k)
		}
	}
	got := trie.KeysMatch(pattern)
	if !StringSliceEqual(correct, got) {
		t.Fatalf("KeysMatch(%s) two set not equal: %v, %v", pattern, correct, got)
	}
}

var dict = map[string]string{
	"a":         "一个",
	"an":        "一个",
	"abandon":   "遗弃",
	"abnormal":  "反常的",
	"apollo":    "阿波罗",
	"archive":   "存档",
	"are":       "是",
	"am":        "是",
	"automatic": "自动的",
	"best":      "最佳的",
	"bit":       "一点",
	"bite":      "咬",
	"bitcoin":   "比特币",
	"byte":      "字节",
}

func ExampleTernary() {
	trie := NewTernary[string]()
	for k, v := range dict {
		trie.Upsert(k, v)
	}
	result := trie.KeysWithPrefix("")
	sort.Strings(result)
	fmt.Println("all keys:", result)

	pattern := "b.te"
	result = trie.KeysMatch(pattern)
	sort.Strings(result)
	fmt.Printf("keys match '%s': %v\n", pattern, result)

	prefix := "bi"
	result = trie.KeysWithPrefix(prefix)
	sort.Strings(result)
	fmt.Printf("keys with prefix '%s': %v\n", prefix, result)

	str := "bitcoins"
	fmt.Printf("longest key with prefix '%s': %s\n", str, trie.LongestPrefixOf(str))

	// Output:
	// all keys: [a abandon abnormal am an apollo archive are automatic best bit bitcoin bite byte]
	// keys match 'b.te': [bite byte]
	// keys with prefix 'bi': [bit bitcoin bite]
	// longest key with prefix 'bitcoins': bitcoin
}
