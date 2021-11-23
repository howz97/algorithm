package main

import (
	"fmt"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/hash_map"
)

func main() {
	hm := hash_map.New(1)
	hm.Put(search.Str("a"), "A")
	hm.Put(search.Str("b"), "B")
	hm.Put(search.Str("c"), "C")
	hm.Put(search.Str("d"), "D")
	hm.Put(search.Str("e"), "E")
	hm.Put(search.Str("f"), "F")
	hm.Put(search.Str("g"), "G")
	hm.Put(search.Str("h"), "H")
	hm.Put(search.Str("i"), "I")
	fmt.Println(hm.String())

	fmt.Println("delete (d/f/g/x) ...")
	hm.Del(search.Str("d"))
	hm.Del(search.Str("f"))
	hm.Del(search.Str("g"))
	hm.Del(search.Str("x"))
	fmt.Println(hm.String())

	fmt.Println("delete all ...")
	hm.Range(func(key hash_map.Key, _ search.T) bool {
		hm.Del(key)
		return true
	})
	fmt.Println(hm.String())
}
