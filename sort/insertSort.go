package sort

// InsertSort -
func InsertSort(data []int) {
	for i := 1; i < len(data); i++ {
		for j := i; j > 0 && data[j] < data[j-1]; j-- {
			data[j-1], data[j] = data[j], data[j-1]
		}
	}
}
