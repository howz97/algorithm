package sort

import (
	"fmt"
	"testing"
)

var disorderData = []int{3, 4, 6, 3, 2, 0, 1, 5, 678, 3, 56, 78, 2, 45, 34, 6, 6, 6, 6, 87, 34, 23, 12, 23, 33, 66, 77, 9}

func Test_QuickSort(t *testing.T) {
	QuickSort(disorderData)
	fmt.Println(disorderData)
	if !isInOrder(disorderData) {
		t.Fatal("data not in order after sorted")
	}
}

func Test_cutoff(t *testing.T) {
	cutoff(disorderData)
	fmt.Println(disorderData)
}

func Test_SelectSort(t *testing.T) {
	SelectSort(disorderData)
	fmt.Println(disorderData)
	if !isInOrder(disorderData) {
		t.Fatal("data not in order after sorted")
	}
}

func isInOrder(data []int) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i] > data[i+1] {
			return false
		}
	}
	return true
}
