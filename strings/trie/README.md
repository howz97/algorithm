* R-way Trie 
* Ternary Search Trie
```go
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

func ExampleTrie() {
	trie := NewTrie[string](alphabet.NewAlphabet(alphabet.ASCII))
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
```