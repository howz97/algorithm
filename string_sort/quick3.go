package string_sort

import (
	"github.com/howz97/algorithm/alphabet"
)

func Quick3(data []string) {
	Quick3WithAlphabet(alphabet.Unicode, data)
}

func Quick3WithAlphabet(a alphabet.Interface, data []string) {
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
	if lo+1 >= hi {
		if lo >= hi {
			return
		}
		for toIndex(a, runes[lo], depth) >= 0 && toIndex(a, runes[hi], depth) >= 0 &&
			toIndex(a, runes[lo], depth) == toIndex(a, runes[hi], depth) {
			depth++
		}
		if toIndex(a, runes[lo], depth) > toIndex(a, runes[hi], depth) {
			runes[lo], runes[hi] = runes[hi], runes[lo]
		}
		return
	}
	medianOfTree(a, runes, lo, hi, depth)
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

func medianOfTree(a alphabet.Interface, runes [][]rune, lo, hi, depth int) {
	m := int(uint(lo+hi) >> 1)
	if toIndex(a, runes[m], depth) < toIndex(a, runes[lo], depth) {
		runes[m], runes[lo] = runes[lo], runes[m]
	}
	if toIndex(a, runes[hi], depth) < toIndex(a, runes[m], depth) {
		runes[hi], runes[m] = runes[m], runes[hi]
		if toIndex(a, runes[m], depth) < toIndex(a, runes[lo], depth) {
			runes[lo], runes[m] = runes[m], runes[lo]
		}
	}
	runes[hi], runes[m] = runes[m], runes[hi]
}
