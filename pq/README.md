堆
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
	for pq.Size() > 0 {
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
	for pq.Size() > 0 {
		fmt.Print(pq.Pop().(string))
	}
	
	//Output:
	//Size: 0, Cap: 3
	//Size: 4, Cap: 7. (auto re-allocate)
	//1799
	//079z
}
```

左式堆
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/pq/leftist_heap"
	. "github.com/howz97/algorithm/util"
)

func main() {
	b := leftist.New()
	b.Push(Int(1))
	b.Push(Int(9))
	b.Push(Int(9))
	b.Push(Int(7))
	b2 := leftist.New()
	b2.Push(Int(13))
	b2.Push(Int(11))
	b.Merge(b2)
	for b.Size() > 0 {
		fmt.Print(b.Pop(), ",")
	}

	//Output: 1,7,9,9,11,13,
}
```

二项队列
```go
package main

import (
	"fmt"
	"github.com/howz97/algorithm/pq/binomial"
	. "github.com/howz97/algorithm/util"
)

func main() {
	b := binomial.New()
	b.Push(Int(1))
	b.Push(Int(9))
	b.Push(Int(9))
	b.Push(Int(7))
	b2 := binomial.New()
	b2.Push(Int(13))
	b2.Push(Int(11))
	b.Merge(b2)
	for b.Size() > 0 {
		fmt.Print(b.Pop(), ",")
	}

	//Output: 1,7,9,9,11,13,
}
```


