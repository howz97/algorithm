二叉堆
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/pq/heap"
	. "github.com/howz97/algorithm/util"
)

func main() {
	pq := heap.New(3)
	fmt.Printf("Size: %d, Cap: %d \n", pq.Size(), pq.Cap())
	pq.Push(Int(1), "1")
	pq.Push(Int(9), "9")
	pq.Push(Int(9), "9")
	pq.Push(Int(7), "7")
	fmt.Printf("Size: %d, Cap: %d. (auto re-allocate) \n", pq.Size(), pq.Cap())
	for !pq.IsEmpty() {
		fmt.Print(pq.Pop().(string))
	}
	fmt.Println()

	pq.Push(Int(1), "1")
	pq.Push(Int(9), "9")
	pq.Push(Int(7), "7")
	pq.Push(Int(99), "0")
	pq.Push(Int(999), "z")
	pq.Del("1")
	pq.Fix(Int(0), "0")
	for !pq.IsEmpty() {
		fmt.Print(pq.Pop().(string))
	}

	/*
		Size: 0, Cap: 3
		Size: 4, Cap: 7. (auto re-allocate)
		1799
		079z
	*/
}
```