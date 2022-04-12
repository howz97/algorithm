package sort

import "golang.org/x/exp/constraints"

func Insert[Ord constraints.Ordered](data []Ord) {
	for i := 1; i < len(data); i++ {
		insrtNum := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > insrtNum; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = insrtNum
	}
}
