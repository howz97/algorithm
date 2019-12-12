package sort

// QuickSort -
func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	median2end(data)
	medianIdx := cutOff(data)
	QuickSort(data[:medianIdx])
	QuickSort(data[medianIdx+1:])
}

// 把 data[0] data[len(data)/2] data[len(data)-1] 中的中位数（枢纽元）交换到data[len(data)-1]
func median2end(data []int) {
	if data[0] > data[len(data)/2] {
		data[0], data[len(data)/2] = data[len(data)/2], data[0]
	}
	if data[len(data)/2] > data[len(data)-1] {
		data[len(data)/2], data[len(data)-1] = data[len(data)-1], data[len(data)/2]
	}
	if data[0] > data[len(data)-1] {
		data[0], data[len(data)-1] = data[len(data)-1], data[0]
	}
	data[len(data)/2], data[len(data)-1] = data[len(data)-1], data[len(data)/2]
}

// 此时枢纽元在 data[len(data)-1] , 开始分割data[:len(data)-1], 并将枢纽元交换到i最终位置
func cutOff(data []int) int {
	if len(data) == 0 {
		panic("cutting off empty slice")
	}
	if len(data) == 1 {
		return 0
	}
	i, j := 0, len(data)-2
	median := len(data) - 1
	for i <= j {
		for i < len(data) && data[i] < data[median] {
			i++
		}
		for j >= 0 && data[j] > data[median] {
			j--
		}
		if i == median {
			return len(data) - 1
		}
		if j < 0 {
			data[0], data[median] = data[median], data[0]
			return 0
		}
		if i <= j {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		}
	}
	data[i], data[median] = data[median], data[i]
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
