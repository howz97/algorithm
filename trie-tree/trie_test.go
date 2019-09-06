package trietree

import (
	"fmt"
	"github.com/zh1014/algorithm/alphabet"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trieT := NewTrie(alphabet.LowerCase)
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
	if !trieT.Contains("hao") {
		t.Fatal()
	}
	s := trieT.Find("zhang").(string)
	if s != "zhang" {
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
	if trieT.LongestPrefixOf("shellsela") != "shells" {
		t.Fatal()
	}
	if trieT.LongestPrefixOf("shela") != "she" {
		t.Fatal()
	}
	fmt.Println("she* :", trieT.KeysWithPrefix("she"))
	fmt.Println("all keys: ", trieT.Keys())
	trieT.Insert("hallo", "hallo")
	trieT.Insert("hillo", "hillo")
	fmt.Println("hello hallo hillo", trieT.KeysMatch("h.llo"))
}
