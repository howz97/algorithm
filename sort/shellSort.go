package sort

// ShellSort -
func ShellSort(data []int) {
	// 使用希尔增量
	for incre := len(data) / 2; incre > 0; incre /= 2 {
		for i := incre; i < len(data); i++ {
			for j := i; j >= incre && data[j] < data[j-incre]; j -= incre {
				data[j-incre], data[j] = data[j], data[j-incre]
			}
		}
	}
}
