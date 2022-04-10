package topk

import (
	"fmt"
	"testing"
)

func Test_ByPartition(t *testing.T) {
	data := []int{4, 1, 7, 3, 87, 5, 46, 7, 23, 5, 67, 8, 8, 7, 6, 5, 4, 37, 84, 56, 76, 5, 43, 25, 6, 76, 5, 43}
	k := 27
	fmt.Println(ByPartition(data, k))
}

func Test_Compare(t *testing.T) {
	data := []int{4, 1, 7, 3, 87, 5, 46, 7, 23, 5, 67, 100, 46, 55, 14, 5, 8, 8, 7, 6, 5, 4, 37, 84, 56, 76, 5, 43, 25, 6, 76, 5, 43}
	k := 7
	topkByPartition := ByPartition(data, k)
	checksum := 0
	for i := range topkByPartition {
		checksum += topkByPartition[i]
	}
	fmt.Println("Top K By Partition: ", topkByPartition, "--checksum=", checksum)
	ByHeap(data, k)
	checksum = 0
	for i := range data[:k] {
		checksum += data[i]
	}
	fmt.Println("Top K By Heap: ", data[:k], "--checksum=", checksum)
}
