package leftist

import (
	"fmt"
	. "github.com/howz97/algorithm/util"
	"testing"
)

func TestExample(t *testing.T) {
	ExampleLeftHeap()
}

func ExampleLeftHeap() {
	b := New()
	b.Push(Int(1))
	b.Push(Int(9))
	b.Push(Int(9))
	b.Push(Int(7))
	b2 := New()
	b2.Push(Int(13))
	b2.Push(Int(11))
	b.Merge(b2)
	for b.Size() > 0 {
		fmt.Print(b.Pop(), ",")
	}

	//Output: 1,7,9,9,11,13,
}
