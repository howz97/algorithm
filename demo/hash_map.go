package main

import (
	"fmt"

	"github.com/howz97/algorithm/search/hashmap"
	. "github.com/howz97/algorithm/util"
)

func main() {
	hm := hashmap.New[Str, string]()
	hm.Put(Str("a"), "A")
	hm.Put(Str("b"), "B")
	hm.Put(Str("c"), "C")
	hm.Put(Str("d"), "D")
	hm.Put(Str("e"), "E")
	hm.Put(Str("f"), "F")
	hm.Put(Str("g"), "G")
	hm.Put(Str("h"), "H")
	hm.Put(Str("i"), "I")
	fmt.Println(hm.String())

	fmt.Println("delete (d/f/g/x) ...")
	hm.Del(Str("d"))
	hm.Del(Str("f"))
	hm.Del(Str("g"))
	hm.Del(Str("x"))
	fmt.Println(hm.String())

	fmt.Println("delete all ...")
	hm.Range(func(key Str, _ string) bool {
		hm.Del(key)
		return true
	})
	fmt.Println(hm.String())

	// size=9
	// bucket0: (b:B) -> (d:D) -> (f:F) -> (h:H) -> nil
	// bucket1: (a:A) -> (c:C) -> (e:E) -> (g:G) -> (i:I) -> nil

	// delete (d/f/g/x) ...
	// size=6
	// bucket0: (b:B) -> (h:H) -> nil
	// bucket1: (a:A) -> (c:C) -> (e:E) -> (i:I) -> nil

	// delete all ...
	// size=0
	// bucket0: nil
}
