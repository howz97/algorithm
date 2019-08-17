package sort

// QuickSort -
func QuickSort(data []int) {
	if len(data) <= 1 {
		return
	}
	if len(data) == 2 {
		if data[0] > data[1] {
			data[0], data[1] = data[1], data[0]
		}
		return
	}
	median2end(data)
	medianIdx := cutoff(data)
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
func cutoff(data []int) int {
	i, j := 0, len(data)-2
	median := data[len(data)-1]
	for i < j {
		for data[i] < median {
			i++
		}
		for data[j] > median && j > 0 {
			j--
		}
		if i < j {
			data[i], data[j] = data[j], data[i]
			i++
			j--
		}
	}
	data[i], data[len(data)-1] = data[len(data)-1], data[i]
	return i
}
