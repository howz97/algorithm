package sort

import "golang.org/x/exp/constraints"

func Shell[Ord constraints.Ordered](data []Ord) {
	// shell increment
	for incre := len(data) >> 1; incre > 1; incre >>= 1 {
		// sort $incre arries with interval of $incre in one loop
		for i := incre; i < len(data); i++ {
			insrtNum := data[i]
			j := i - incre
			for ; j >= 0 && data[j] > insrtNum; j -= incre {
				data[j+incre] = data[j]
			}
			data[j+incre] = insrtNum
		}
	}
	Insert(data)
}

//func ShellSort(data []int) {
//	// shell increment
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
