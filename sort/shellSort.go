package sort

func ShellSort(data []int) {
	// 使用希尔增量
	for incre := len(data) >> 1; incre > 1; incre >>= 1 {
		for i := incre; i < len(data); i++ {
			insrtNum := data[i]
			j := i - incre
			for ; j >= 0 && data[j] > insrtNum; j -= incre {
				data[j+incre] = data[j]
			}
			data[j+incre] = insrtNum
		}
	}
	InsertSort(data)
}

//func ShellSort(data []int) {
//	// 使用希尔增量
//	for incre := len(data) >> 1; incre > 0; incre >>= 1 {
//		for i := incre; i < len(data); i++ {
//			insrtNum := data[i]
//			j := i-incre
//			for ; j >= 0 && data[j] > insrtNum; j -= incre {
//				data[j+incre] = data[j]
//			}
//			data[j+incre] = insrtNum
//		}
//	}
//}
