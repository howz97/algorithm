package tst

import (
	"fmt"
	"github.com/zh1014/algorithm/alphabet"
	"math/rand"
	"testing"
)

func TestTst_Insert(t *testing.T) {
	trieT := NewTst(alphabet.LowerCase)
	if !trieT.IsEmpty() {
		t.Fatal()
	}
	trieT.Insert("hello", "hello")
	trieT.Insert("hi", "hi")
	trieT.Insert("she", "she")
	trieT.Insert("shells", "shells")
	trieT.Insert("shell", "shell")
	trieT.Insert("zhang", "zhang")
	trieT.Insert("hao", "hao")
	trieT.Insert("zh", "zh")
	trieT.Insert("alloc", "alloc")
	trieT.Insert("milk", "milk")
	trieT.Insert("memory", "memory")
	if trieT.tree.isEmpty() {
		t.Fatal()
	}
	if trieT.Size() != 11 {
		t.Fatal()
	}
	s := trieT.Find("zhang").(string)
	if s != "zhang" {
		t.Fatal()
	}
	if !trieT.Contains("hi") {
		t.Fatal()
	}
	trieT.Delete("zhang")
	_, ok := trieT.Find("zhang").(string)
	if ok {
		t.Fatal()
	}
	if trieT.Size() != 10 {
		t.Fatal()
	}
	if trieT.LongestPrefixOf("shela") != "she" {
		t.Fatal()
	}
	if trieT.LongestPrefixOf("shellsela") != "shells" {
		t.Fatal()
	}
	fmt.Println("she* :", trieT.KeysWithPrefix("she"))
	fmt.Println("all keys: ", trieT.Keys())
	trieT.Insert("hallo", "hallo")
	trieT.Insert("hillo", "hillo")
	fmt.Println("h* :", trieT.KeysWithPrefix("h"))
	fmt.Println("h.llo (should be hello hallo hillo) :", trieT.KeysMatch("h.llo"))
}

func TestTst_Delete(t *testing.T) {
	trie := NewTst(alphabet.Unicode)
	items := []string{"你好", "你好看", "你好好看", "你是？", "你是谁", "你是谁啊", "你是胖虎",
		"你还问？", "你还看", "你还不说", "你还不睡",
		"你真蠢", "你真是个弟弟", "你真好看", "你真的是胖虎？", "你真是个弟弟啊2333",
		"我是你爸", "我是纯甄小蛮腰", "我😍你",
		"abc", "123", "abb", "ab13", "w2f", "2d2wd", "s2qd", "2s2", "$%^&", "....", "1w2r3tyd", "3f", "s2qd",
	}
	for i := range items {
		trie.Insert(items[i], i)
	}
	midPoint := 10
	for i := midPoint; i >= 0; i-- {
		trie.Delete(items[i])
		if trie.Contains(items[i]) {
			t.Fatal()
		}
		trie.Delete(items[i])
	}
	for i := midPoint + 1; i < len(items); i++ {
		trie.Delete(items[i])
		if trie.Contains(items[i]) {
			t.Fatal()
		}
		trie.Delete(items[i])
	}
	trie.IsEmpty()
}

func TestTst_Contains(t *testing.T) {
	trie := NewTst(alphabet.Unicode)
	items := []string{"你好", "你好看", "你好好看", "你是？", "你是谁", "你是谁啊", "你是胖虎",
		"你还问？", "你还看", "你还不说", "你还不睡",
		"你真蠢", "你真是个弟弟", "你真好看", "你真的是胖虎？", "你真是个弟弟啊2333",
		"我是你爸", "我是纯甄小蛮腰", "我😍你",
		"abc", "123", "abb", "ab13", "w2f", "2d2wd", "s2qd", "2s2", "$%^&", "....", "1w2r3tyd", "3f", "s2qd",
	}
	for i := range items {
		trie.Insert(items[i], i)
	}
	for i := range items {
		if !trie.Contains(items[i]) {
			t.Fatal()
		}
	}
	for i := len(items) - 1; i >= 0; i-- {
		if !trie.Contains(items[i]) {
			t.Fatal()
		}
	}
	for i := 0; i < 100; i++ {
		r := rand.Int() % len(items)
		if !trie.Contains(items[r]) {
			t.Fatal()
		}
	}
}

func TestTst_KeysWithPrefix(t *testing.T) {
	trie := NewTst(alphabet.Unicode)
	items := []string{"你好", "你好看", "你好好看", "你是？", "你是谁", "你是谁啊", "你是胖虎",
		"你还问？", "你还看", "你还不说", "你还不睡",
		"你真蠢", "你真是个弟弟", "你真好看", "你真的是胖虎？", "你真是个弟弟啊2333",
		"我是你爸", "我是纯甄小蛮腰", "我😍你",
	}
	for i := range items {
		trie.Insert(items[i], i)
	}
	fmt.Println("你*: ", trie.KeysWithPrefix("你"))
	fmt.Println("你好*: ", trie.KeysWithPrefix("你好"))
	fmt.Println("你是*: ", trie.KeysWithPrefix("你是"))
	fmt.Println("你还好马？？？*: ", trie.KeysWithPrefix("你还好马？？？"))
	fmt.Println("你真*: ", trie.KeysWithPrefix("你真"))
	fmt.Println("你真的*: ", trie.KeysWithPrefix("你真的"))
	fmt.Println("你真是*: ", trie.KeysWithPrefix("你真是"))
	fmt.Println("我*: ", trie.KeysWithPrefix("我"))
	fmt.Println("我是*: ", trie.KeysWithPrefix("我是"))
}

func TestTst_KeysMatch(t *testing.T) {
	trie := NewTst(alphabet.Unicode)
	items := []string{"你好", "你好看", "你好好看", "你是？", "你是谁", "你是谁啊", "你是胖虎",
		"你还问？", "你还看", "你还不说", "你还不睡",
		"你真蠢", "你真是个弟弟", "你真好看", "你真的是胖虎？", "你真是个弟弟啊2333",
		"我是你爸", "我是纯甄小蛮腰", "我😍你",
	}
	for i := range items {
		trie.Insert(items[i], i)
	}
	fmt.Println("你: ", trie.KeysMatch("你"))
	fmt.Println("你.: ", trie.KeysMatch("你."))
	fmt.Println("你..: ", trie.KeysMatch("你.."))
	fmt.Println("你...: ", trie.KeysMatch("你..."))
	fmt.Println("你....: ", trie.KeysMatch("你...."))
	fmt.Println("你.....: ", trie.KeysMatch("你....."))
	fmt.Println("你......: ", trie.KeysMatch("你......"))
	fmt.Println(".是..: ", trie.KeysMatch(".是.."))
	fmt.Println("你.是个.弟: ", trie.KeysMatch("你.是个.弟"))
	fmt.Println("....小.腰: ", trie.KeysMatch("....小.腰"))
	fmt.Println("....: ", trie.KeysMatch("...."))
	fmt.Println("..你.: ", trie.KeysMatch("..你."))
}

func TestTst_LongestPrefixOf(t *testing.T) {
	trie := NewTst(alphabet.Unicode)
	items := []string{"你好", "你好看", "你好好看", "你是？", "你是谁", "你是谁啊", "你是胖虎",
		"你还问？", "你还看", "你还不说", "你还不睡",
		"你真蠢", "你真是个弟弟", "你真好看", "你真的是胖虎？", "你真是个弟弟啊2333",
		"我是你爸", "我是纯甄小蛮腰", "我😍你",
	}
	for i := range items {
		trie.Insert(items[i], i)
	}
	if trie.LongestPrefixOf("你好看!") != "你好看" {
		t.Fatal()
	}
	if trie.LongestPrefixOf("你是") != "" {
		t.Fatal()
	}
	if trie.LongestPrefixOf("你真是个弟弟啊2333~~~~~") != "你真是个弟弟啊2333" {
		t.Fatal()
	}
}
