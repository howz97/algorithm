package main

import (
	"fmt"
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/trie_tree"
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
	trie := trietree.NewTrie(alphabet.Ascii)
	//trie := trietree.NewTST()
	for k, v := range dict {
		trie.Upsert(k, v)
	}
	fmt.Println("all keys:", trie.Keys())

	pattern := "a......"
	fmt.Printf("keys match '%s': %v\n", pattern, trie.KeysMatch(pattern))

	prefix := "bi"
	fmt.Printf("keys with prefix '%s': %v\n", prefix, trie.KeysWithPrefix(prefix))

	str := "bitcoins"
	fmt.Printf("longest key with prefix '%s': %s\n", str, trie.LongestPrefixOf(str))
}
