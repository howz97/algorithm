堆
```go
package main

import (
	"fmt"

	"github.com/howz97/algorithm/pq/heap"
)

func main() {
	pq := heap.New2[int, string](3)
	fmt.Printf("Size: %d, Cap: %d \n", pq.Size(), pq.Cap())
	pq.Push(1, "1")
	pq.Push(9, "9")
	pq.Push(9, "9")
	pq.Push(7, "7")
	fmt.Printf("Size: %d, Cap: %d. (auto re-allocate) \n", pq.Size(), pq.Cap())
	for pq.Size() > 0 {
		fmt.Print(pq.Pop())
	}
	fmt.Println()

	pq.Push(1, "1")
	pq.Push(9, "9")
	pq.Push(7, "7")
	pq.Push(99, "0")
	pq.Push(999, "z")
	pq.Del("1")
	pq.Fix(0, "0")
	for pq.Size() > 0 {
		fmt.Print(pq.Pop())
	}

	//Output:
	//Size: 0, Cap: 3
	//Size: 4, Cap: 7. (auto re-allocate)
	//1799
	//079z
}
```