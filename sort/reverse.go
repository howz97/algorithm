package sort

func Reverse(data []int) {
	swapTimes := len(data) / 2
	for i := 0; i < swapTimes; i++ {
		data[i], data[len(data)-1-i] = data[len(data)-1-i], data[i]
	}
}
