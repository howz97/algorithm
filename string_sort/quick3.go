package string_sort

import "github.com/howz97/algorithm/alphabet"

func Quick3(a alphabet.Interface, data []string) {
	runes := make([][]rune, len(data))
	for i := range runes {
		runes[i] = []rune(data[i])
	}
	quick3(a, runes, 0, len(data)-1, 0)
	for i := range data {
		data[i] = string(runes[i])
	}
}

func quick3(a alphabet.Interface, runes [][]rune, lo, hi, depth int) {
	if lo >= hi {
		return
	}
	middleV := toIndex(a, runes[lo], depth)
	tail, i := lo, lo+1
	head := hi
	for i <= head {
		v := toIndex(a, runes[i], depth)
		switch true {
		case v < middleV:
			runes[tail], runes[i] = runes[i], runes[tail]
			tail++
			i++
		case v > middleV:
			runes[i], runes[head] = runes[head], runes[i]
			head--
		default:
			i++
		}
	}
	// runes[0...tail] < middleV
	// runes[tail...head] = middleV
	// runes[head...] > middleV

	quick3(a, runes, lo, tail-1, depth)
	if middleV >= 0 {
		quick3(a, runes, tail, head, depth+1)
	}
	quick3(a, runes, head+1, hi, depth)
}
