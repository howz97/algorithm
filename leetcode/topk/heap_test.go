package topk

import (
	"fmt"
	"testing"
)

func Test_ByHeap(t *testing.T) {
	data := []int{4, 2, 7, 3, 87, 5, 46, 7, 23, 5, 67, 8, 8, 7, 6, 5, 4, 37, 84, 56, 76, 5, 43, 25, 6, 76, 5, 43}
	k := 20
	ByHeap(data, k)
	fmt.Println(data[:k])
}
