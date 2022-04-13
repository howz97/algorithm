前缀树
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/trie"
)

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

func main() {
	//trie := trie.NewTrie(alphabet.Ascii)
	//trie := trie.NewTST()
	trie := trie.NewTSTC()
	for k, v := range dict {
		trie.Upsert(k, v)
	}
	fmt.Println("all keys:", trie.Keys())

	pattern := "b.te"
	fmt.Printf("keys match '%s': %v\n", pattern, trie.KeysMatch(pattern))

	prefix := "bi"
	fmt.Printf("keys with prefix '%s': %v\n", prefix, trie.KeysWithPrefix(prefix))

	str := "bitcoins"
	fmt.Printf("longest key with prefix '%s': %s\n", str, trie.LongestPrefixOf(str))
}

/*
all keys: [a abandon abnormal am an apollo are archive automatic best bit bitcoin bite byte]
keys match 'b.te': [byte bite]
keys with prefix 'bi': [bit bitcoin bite]
longest key with prefix 'bitcoins': bitcoin
*/
```