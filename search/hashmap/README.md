链式哈希表
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/search"
	"github.com/howz97/algorithm/search/hashmap"
)

func main() {
	hm := hashmap.New(1)
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
	hm.Range(func(key hashmap.Key, _ search.T) bool {
		hm.Del(key)
		return true
	})
	fmt.Println(hm.String())
}

/*
size=9
bucket0: (b:B) -> (d:D) -> (f:F) -> (h:H) -> nil
bucket1: (a:A) -> (c:C) -> (e:E) -> (g:G) -> (i:I) -> nil

delete (d/f/g) ...
size=6
bucket0: (b:B) -> (h:H) -> nil
bucket1: (a:A) -> (c:C) -> (e:E) -> (i:I) -> nil

delete all ...
size=0
bucket0: nil
*/
```