前缀树
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/alphabet"
	"github.com/howz97/algorithm/trie_tree/tst"
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
	tst1 := tst.NewTst(alphabet.Ascii)
	for k, v := range dict {
		tst1.Insert(k, v)
	}
	fmt.Println("all keys:", tst1.Keys())

	pattern := "a......"
	fmt.Printf("keys match '%s': %v\n", pattern, tst1.KeysMatch(pattern))

	prefix := "bi"
	fmt.Printf("keys with prefix '%s': %v\n", prefix, tst1.KeysWithPrefix(prefix))

	str := "bitcoins"
	fmt.Printf("longest key with prefix '%s': %s\n", str, tst1.LongestPrefixOf(str))
}

/*
all keys: [a abandon abnormal an am apollo are archive automatic best bit bitcoin bite byte]
keys match 'a......': [archive abandon]
keys with prefix 'bi': [bit bitcoin bite]
longest key with prefix 'bitcoins': bitcoin
*/
```