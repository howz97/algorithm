package trietree

import (
	"strings"
	"testing"

	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/util"
)

func TestSliceNode(t *testing.T) {
	trie := NewTrie(alphabet.Ascii)
	TrieTreeTest(t, trie, DataASCII)
}

func TestTSTNode(t *testing.T) {
	trie := NewTST()
	TrieTreeTest(t, trie, DataASCII)
	trie = NewTST()
	TrieTreeTest(t, trie, DataChn)
}

func TestTSTCNode(t *testing.T) {
	trie := NewTSTC()
	TrieTreeTest(t, trie, DataASCII)
	trie = NewTSTC()
	TrieTreeTest(t, trie, DataChn)
}

func TestTSTC_Compress(t *testing.T) {
	trie := NewTSTC()
	UpsertAndFind(t, trie, DataChn.m)
	if trie.IsCompressed() {
		t.Fatalf("compress failed")
	}
	if err := trie.Compress(); err != nil {
		t.Fatal(err)
	}
	if !trie.IsCompressed() {
		t.Fatalf("compress failed")
	}
	for _, pre := range DataChn.prefix {
		KeysWithPrefix(t, trie, pre)
	}
	for _, p := range DataChn.pattern {
		KeysMatch(t, trie, p)
	}
	for _, l := range DataChn.long {
		LongestPrefixOf(t, trie, l)
	}
	UpdateAndFind(t, trie, DataChn.m)
	Delete(t, trie, DataChn.m)
}

type Data struct {
	m       map[string]string
	prefix  []string
	pattern []string
	long    []string
}

var DataASCII = Data{
	m: map[string]string{
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
		"by":        "通过",
		"byte":      "字节",
		"bytes":     "字节(复数)",
	},
	prefix:  []string{"a", "ab", "ar", "am", "arm", "b", "bit", "byte", "x", "hello"},
	pattern: []string{".", "..", "....", "a.", "b.te", "b..", "................", "hello"},
	long:    []string{"bitcoins", "byte dance", "arm", "hello"},
}

var DataChn = Data{
	m: map[string]string{
		"风":      "1",
		"风牛马":    "3",
		"风牛马不相及": "6",
		"风华绝代":   "4",
		"风骚绝代":   "4",
		"风姿绰约":   "4",
		"风度":     "2",
		"风度翩翩":   "4",
		"风蚀":     "2",
		"风云":     "2",
		"风儿吹":    "3",
		"往事随风":   "4",
		"往事":     "2",
	},
	prefix:  []string{"风", "风牛", "风度", "风蚀地貌", "往", "往事", "张"},
	pattern: []string{".", "..", "....", "风.", "风.绝代", "风..", "往.", "................", "hello"},
	long:    []string{"风牛马不相及！", "芳华绝代", "风度翩翩～"},
}

func TrieTreeTest(t *testing.T, trie *Trie, data Data) {
	UpsertAndFind(t, trie, data.m)
	for _, pre := range data.prefix {
		KeysWithPrefix(t, trie, pre)
	}
	for _, p := range data.pattern {
		KeysMatch(t, trie, p)
	}
	for _, l := range data.long {
		LongestPrefixOf(t, trie, l)
	}
	UpdateAndFind(t, trie, data.m)
	Delete(t, trie, data.m)
}

func Delete(t *testing.T, trie *Trie, m map[string]string) {
	for k := range m {
		trie.Delete(k)
	}
	for k := range m {
		if trie.Contains(k) {
			t.Fatalf("Delete(%s) failed", k)
		}
	}
}

func UpsertAndFind(t *testing.T, trie *Trie, m map[string]string) {
	for k, v := range m {
		trie.Upsert(k, v)
	}
	for k, v := range m {
		got := trie.Find(k).(string)
		if got != v {
			t.Fatalf("Find(%s)==%s, should be %s", k, got, v)
		}
	}
}

func UpdateAndFind(t *testing.T, trie *Trie, m map[string]string) {
	for k, v := range m {
		trie.Update(k, v+"#")
	}
	for k, v := range m {
		v = v + "#"
		got := trie.Find(k).(string)
		if got != v {
			t.Fatalf("Find(%s)==%s, should be %s", k, got, v)
		}

		invalidKey := k + "@invalid_suffix"
		trie.Update(invalidKey, "1")
		if trie.Find(invalidKey) == "1" {
			t.Fatalf("shoud not succeed to update")
		}
	}
}

func KeysWithPrefix(t *testing.T, trie *Trie, prefix string) {
	var correct []string
	for _, k := range trie.Keys() {
		if strings.HasPrefix(k, prefix) {
			correct = append(correct, k)
		}
	}
	got := trie.KeysWithPrefix(prefix)
	if !StringSliceEqual(correct, got) {
		t.Fatalf("KeysWithPrefix(%s) two set not equal: %v, %v", prefix, correct, got)
	}
}

func LongestPrefixOf(t *testing.T, trie *Trie, str string) {
	k := trie.LongestPrefixOf(str)
	if !strings.HasPrefix(str, k) {
		t.Fatalf("key %s is not prefix of %s", k, str)
	}
	if k != "" && !trie.Contains(k) {
		t.Fatalf("LongestPrefixOf(%s) key %s not exist", str, k)
	}

	str = strings.TrimPrefix(str, k)
	for _, r := range str {
		k += string(r)
		if trie.Contains(k) {
			t.Fatalf("longger key %s exist", k)
		}
	}
}

func KeysMatch(t *testing.T, trie *Trie, pattern string) {
	var correct []string
	for _, k := range trie.Keys() {
		if util.IsRunesMatch([]rune(pattern), []rune(k)) {
			correct = append(correct, k)
		}
	}
	got := trie.KeysMatch(pattern)
	if !StringSliceEqual(correct, got) {
		t.Fatalf("KeysMatch(%s) two set not equal: %v, %v", pattern, correct, got)
	}
}

func StringSliceEqual(s0, s1 []string) bool {
	if len(s0) != len(s1) {
		return false
	}
	for _, str := range s0 {
		if util.IndexStringSlice(s1, str) < 0 {
			return false
		}
	}
	for _, str := range s1 {
		if util.IndexStringSlice(s0, str) < 0 {
			return false
		}
	}
	return true
}
