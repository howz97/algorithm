package main

import (
	"fmt"

	"github.com/howz97/algorithm/pq/binomial"
)

func main() {
	b := binomial.New[int]()
	b.Push(1)
	b.Push(9)
	b.Push(9)
	b.Push(7)
	b2 := binomial.New[int]()
	b2.Push(13)
	b2.Push(11)
	b.Merge(b2)
	for b.Size() > 0 {
		fmt.Print(b.Pop(), ",")
	}

	//Output: 1,7,9,9,11,13,
}
