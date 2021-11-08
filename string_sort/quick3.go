package string_sort

import "github.com/howz97/algorithm/alphabet"

func Quick3(a alphabet.Interface, strs []string) {
	// convert string to []rune
	strsRune := make([][]rune, len(strs))
	for i := range strsRune {
		strsRune[i] = []rune(strs[i])
	}

	// sort [][]rune which equal to strs
	quick3(a, strsRune, 0, len(strs)-1, 0)

	// convert []rune to string
	for i := range strs {
		strs[i] = string(strsRune[i])
	}
}

func quick3(a alphabet.Interface, strs [][]rune, lo, hi, d int) {
	if lo >= hi {
		return
	}
	v := toIndex(a, strs[lo], d)
	tailV, i := lo, lo+1
	end := hi

	for i <= end {
		switch true {
		case toIndex(a, strs[i], d) == v:
			i++
		case toIndex(a, strs[i], d) < v:
			strs[tailV], strs[i] = strs[i], strs[tailV]
			tailV++
			i++
		default:
			strs[i], strs[end] = strs[end], strs[i]
			end--
		}
	}

	quick3(a, strs, lo, tailV-1, d)
	if v >= 0 {
		quick3(a, strs, tailV, end, d+1)
	}
	quick3(a, strs, end+1, hi, d)
}
