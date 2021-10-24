package main

import (
	"fmt"
	"github.com/howz97/algorithm/hash_table"
)

func main() {
	ht := hash_table.New(1)
	ht.Put(hash_table.Str("a"), "A")
	ht.Put(hash_table.Str("b"), "B")
	ht.Put(hash_table.Str("c"), "C")
	ht.Put(hash_table.Str("d"), "D")
	ht.Put(hash_table.Str("e"), "E")
	ht.Put(hash_table.Str("f"), "F")
	ht.Range(func(key hash_table.Key, val hash_table.T) bool {
		fmt.Println(key, "->", val)
		return true
	})

	ht.Delete(hash_table.Str("d"))
	ht.Delete(hash_table.Str("f"))
	ht.Delete(hash_table.Str("x"))
	fmt.Println("after delete (d/f/x) ...")
	ht.Range(func(key hash_table.Key, val hash_table.T) bool {
		fmt.Println(key, "->", val)
		return true
	})

	ht.Range(func(key hash_table.Key, _ hash_table.T) bool {
		ht.Delete(key)
		return true
	})
	fmt.Println("after delete all ...", ht.Size(), ht.TblSize())
	ht.Range(func(key hash_table.Key, val hash_table.T) bool {
		fmt.Println(key, "->", val)
		return true
	})
}
