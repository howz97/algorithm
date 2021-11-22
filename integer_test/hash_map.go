package main

import (
	"fmt"
	"github.com/howz97/algorithm/search/hash_map"
)

func main() {
	ht := hash_map.New(1)
	ht.Put(hash_map.Str("a"), "A")
	ht.Put(hash_map.Str("b"), "B")
	ht.Put(hash_map.Str("c"), "C")
	ht.Put(hash_map.Str("d"), "D")
	ht.Put(hash_map.Str("e"), "E")
	ht.Put(hash_map.Str("f"), "F")
	ht.Range(func(key hash_map.Key, val hash_map.T) bool {
		fmt.Println(key, "->", val)
		return true
	})

	ht.Del(hash_map.Str("d"))
	ht.Del(hash_map.Str("f"))
	ht.Del(hash_map.Str("x"))
	fmt.Println("after delete (d/f/x) ...")
	ht.Range(func(key hash_map.Key, val hash_map.T) bool {
		fmt.Println(key, "->", val)
		return true
	})

	ht.Range(func(key hash_map.Key, _ hash_map.T) bool {
		ht.Del(key)
		return true
	})
	fmt.Println("after delete all ...", ht.Size(), ht.TblSize())
	ht.Range(func(key hash_map.Key, val hash_map.T) bool {
		fmt.Println(key, "->", val)
		return true
	})
}
