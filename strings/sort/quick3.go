package sort

import (
	"github.com/howz97/algorithm/strings/alphabet"
)

// Quick3 seems to be a faster string sorting algorithm than the sort.Strings method in the standard library
func Quick3(strings []string) {
	quick3(strings, 0, len(strings)-1, 0)
}

func quick3(strings []string, lo, hi, depth int) {
	if lo+1 >= hi {
		if lo >= hi {
			return
		}
		for depth < len(strings[lo]) && depth < len(strings[hi]) &&
			byteAt(strings[lo], depth) == byteAt(strings[hi], depth) {
			depth++
		}
		if byteAt(strings[lo], depth) > byteAt(strings[hi], depth) {
			strings[lo], strings[hi] = strings[hi], strings[lo]
		}
		return
	}
	median(strings, lo, hi, depth)
	middleV := byteAt(strings[lo], depth)
	tail, i := lo, lo+1
	head := hi
	for i <= head {
		v := byteAt(strings[i], depth)
		switch true {
		case v < middleV:
			strings[tail], strings[i] = strings[i], strings[tail]
			tail++
			i++
		case v > middleV:
			strings[i], strings[head] = strings[head], strings[i]
			head--
		default:
			i++
		}
	}
	// bytes[0...tail] < middleV
	// bytes[tail...head] = middleV
	// bytes[head...] > middleV

	quick3(strings, lo, tail-1, depth)
	if middleV >= 0 {
		quick3(strings, tail, head, depth+1)
	}
	quick3(strings, head+1, hi, depth)
}

func byteAt(str string, depth int) int {
	if depth >= len(str) {
		return -1
	}
	return int(str[depth])
}

func median(bytes []string, lo, hi, depth int) {
	m := int(uint(lo+hi) >> 1)
	if byteAt(bytes[m], depth) < byteAt(bytes[lo], depth) {
		bytes[m], bytes[lo] = bytes[lo], bytes[m]
	}
	if byteAt(bytes[hi], depth) < byteAt(bytes[m], depth) {
		bytes[hi], bytes[m] = bytes[m], bytes[hi]
		if byteAt(bytes[m], depth) < byteAt(bytes[lo], depth) {
			bytes[lo], bytes[m] = bytes[m], bytes[lo]
		}
	}
	bytes[hi], bytes[m] = bytes[m], bytes[hi]
}

func Quick3Alp(a alphabet.IAlp, data []string) {
	runes := make([][]rune, len(data))
	for i := range runes {
		runes[i] = []rune(data[i])
	}
	quick3alp(a, runes, 0, len(data)-1, 0)
	for i := range data {
		data[i] = string(runes[i])
	}
}

func quick3alp(a alphabet.IAlp, runes [][]rune, lo, hi, depth int) {
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
	medianAlp(a, runes, lo, hi, depth)
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

	quick3alp(a, runes, lo, tail-1, depth)
	if middleV >= 0 {
		quick3alp(a, runes, tail, head, depth+1)
	}
	quick3alp(a, runes, head+1, hi, depth)
}

func medianAlp(a alphabet.IAlp, runes [][]rune, lo, hi, depth int) {
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
