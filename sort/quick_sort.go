package sort

import (
	stdsort "sort"
)

func QuickSort(data stdsort.Interface) {
	quickSort(data, 0, data.Len()-1)
}

func quickSort(data stdsort.Interface, lo, hi int) {
	if hi <= lo {
		return
	}
	if hi-lo == 1 {
		if data.Less(hi, lo) {
			data.Swap(hi, lo)
		}
		return
	}
	median2end(data, lo, hi)
	m := cutOff(data, lo, hi)
	quickSort(data, lo, m-1)
	quickSort(data, m+1, hi)
}

// 把 data[0] data[len(data)/2] data[len(data)-1] 中的中位数（枢纽元）交换到data[len(data)-1]
func median2end(data stdsort.Interface, lo, hi int) {
	m := int(uint(lo+hi) >> 1)
	if data.Less(m, lo) {
		data.Swap(m, lo)
	}
	if data.Less(hi, m) {
		data.Swap(hi, m)
		if data.Less(m, lo) {
			data.Swap(m, lo)
		}
	}
	data.Swap(m, hi)
}

// 此时枢纽元在 data[len(data)-1] , 开始分割data[:len(data)-1], 并将枢纽元交换到i最终位置
func cutOff(data stdsort.Interface, lo, hi int) int {
	i, j := lo, hi-1
	for i <= j {
		for i <= hi && data.Less(i, hi) {
			i++
		}
		for j >= lo && data.Less(hi, j) {
			j--
		}
		if i == hi {
			return hi
		}
		if j < lo {
			data.Swap(lo, hi)
			return lo
		}
		if i <= j {
			data.Swap(i, j)
			i++
			j--
		}
	}
	data.Swap(i, hi)
	return i
}

// obsoleted: 这种cutOff在随机输入下没有优势。且在输入中大量重复时复杂度达到O(n平方)
//func cutOff(data []int) int {
//	if len(data) == 0 {
//		panic("cutting off empty slice")
//	}
//	i, j := 0, len(data)-1
//	median := len(data)-1
//	for i <= j {
//		for i < len(data) && data[i] <= data[median] {
//			i++
//		}
//		for j >=0 && data[j] >= data[median] {
//			j--
//		}
//		if i == len(data) {
//			return len(data) -1
//		}
//		if j < 0 {
//			data[0], data[median] = data[median], data[0]
//			return 0
//		}
//		if i < j {
//			data[i], data[j] = data[j], data[i]
//			i++
//			j--
//		}
//	}
//	data[i], data[median] = data[median], data[i]
//	return i
//}
