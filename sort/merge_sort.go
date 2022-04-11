package sort

import "golang.org/x/exp/constraints"

func MergeSort[Ord constraints.Ordered](data []Ord) {
	aux := make([]Ord, len(data))
	mergeSort(data, aux)
}

func mergeSort[Ord constraints.Ordered](data []Ord, aux []Ord) {
	if len(data) < 2 {
		return
	}
	mid := len(data) >> 1
	mergeSort(data[:mid], aux[:mid])
	mergeSort(data[mid:], aux[mid:])
	merge(data, mid, aux)
}

func merge[Ord constraints.Ordered](data []Ord, mid int, aux []Ord) {
	i, j, k := 0, mid, 0
	for i < mid && j < len(data) {
		if data[i] < data[j] {
			aux[k] = data[i]
			i++
		} else {
			aux[k] = data[j]
			j++
		}
		k++
	}
	if i < mid {
		copy(data[k:], data[i:mid])
	}
	copy(data[:k], aux[:k])
}
