package sort

import "golang.org/x/exp/constraints"

func Bubble[Ord constraints.Ordered](data []Ord) {
	for i := len(data) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
