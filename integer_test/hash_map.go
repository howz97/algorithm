package main

import (
	"fmt"
	"github.com/howz97/algorithm/search/hash_map"
	"github.com/howz97/algorithm/util"
)

func main() {
	hm := hash_map.New(1)
	hm.Put(util.Str("a"), "A")
	hm.Put(util.Str("b"), "B")
	hm.Put(util.Str("c"), "C")
	hm.Put(util.Str("d"), "D")
	hm.Put(util.Str("e"), "E")
	hm.Put(util.Str("f"), "F")
	hm.Put(util.Str("g"), "G")
	hm.Put(util.Str("h"), "H")
	hm.Put(util.Str("i"), "I")
	fmt.Println(hm.String())

	fmt.Println("delete (d/f/g/x) ...")
	hm.Del(util.Str("d"))
	hm.Del(util.Str("f"))
	hm.Del(util.Str("g"))
	hm.Del(util.Str("x"))
	fmt.Println(hm.String())

	fmt.Println("delete all ...")
	hm.Range(func(key hash_map.Key, _ util.T) bool {
		hm.Del(key)
		return true
	})
	fmt.Println(hm.String())
}
