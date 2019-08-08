package sort

import (
	"fmt"
	"testing"
)

func Test_QuickSort(t *testing.T) {
	data := []int{3, 4, 6, 3, 2, 0, 1, 5, 678, 3, 56, 78, 2, 45, 34, 6, 6, 6, 6, 87, 34, 23, 12, 23, 33, 66, 77, 9}
	QuickSort(data)
	fmt.Println(data)
}

func Test_cutoff(t *testing.T) {
	data := []int{3, 4, 6, 3, 2, 0, 1, 5, 678, 3, 56, 78, 2, 45, 34, 6, 6, 6, 6, 87, 34, 23, 12, 23, 33, 66, 77, 9}
	cutoff(data)
	fmt.Println(data)

	data = []int{3, 4, 4, 4, 4, 4}
	cutoff(data)
	fmt.Println(data)

	data = []int{3, 4, 1}
	cutoff(data)
	fmt.Println(data)
}
