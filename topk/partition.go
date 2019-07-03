package topk

import "fmt"

// ByPartition return biggest k number of data
func ByPartition(data []int, k int) []int {
	if k < 0 {
		panic("k < 0")
	}
	if k == 0 {
		return make([]int, 0)
	}
	if len(data) < k {
		fmt.Println("len(data) is", len(data), "< k")
		return nil
	}
	if len(data) == k {
		return data // should not return origin array, but I want to make code simpler
	}
	mid := data[0]
	midcount := 0
	bigger := make([]int, 0, len(data))
	smaller := make([]int, 0, len(data))

	for i := range data {
		if data[i] > mid {
			bigger = append(bigger, data[i])
		} else if data[i] < mid {
			smaller = append(smaller, data[i])
		} else {
			midcount++
		}
	}
	// recycle data
	data = nil

	switch true {
	case len(bigger) == k:
		return bigger
	case len(bigger) > k:
		return ByPartition(bigger, k)
	default:
		stillNeed := k - len(bigger)
		if stillNeed <= midcount {
			for i := 0; i < stillNeed; i++ {
				bigger = append(bigger, mid)
			}
			return bigger
		}
		for i := 0; i < midcount; i++ {
			bigger = append(bigger, mid)
		}
		return append(bigger, ByPartition(smaller, stillNeed-midcount)...)
	}
}
