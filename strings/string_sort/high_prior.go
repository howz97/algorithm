package string_sort

import "github.com/howz97/algorithm/strings/alphabet"

func HighPrior(data []string) {
	HighPriorWithAlphabet(alphabet.Unicode, data)
}

func HighPriorWithAlphabet(a alphabet.Interface, data []string) {
	if len(data) <= 1 {
		return
	}
	runes := make([][]rune, len(data))
	for i := range data {
		runes[i] = []rune(data[i])
	}
	aux := make([][]rune, len(data))
	highPriorSort(a, runes, aux, 0, len(data)-1, 0)
	for i := range runes {
		data[i] = string(runes[i])
	}
}

func highPriorSort(a alphabet.Interface, runes, aux [][]rune, lo, hi, depth int) {
	if lo >= hi {
		return
	}
	// count[1] -> count of rune[depth] overflow
	// count[2...R+1] -> count of rune[depth]=0...R-1
	count := make([]int, a.R()+2)
	for i := lo; i <= hi; i++ {
		count[toIndex(a, runes[i], depth)+2]++
	}
	// count[0] -> start index of rune[depth] overflow
	// count[1...R] -> start index of rune[depth]=0...R-1
	count[0] = lo
	for i := 1; i <= a.R(); i++ {
		count[i] += count[i-1]
	}

	for i := lo; i <= hi; i++ {
		aux[count[toIndex(a, runes[i], depth)+1]] = runes[i]
		count[toIndex(a, runes[i], depth)+1]++
	}
	// write back
	for i := lo; i <= hi; i++ {
		runes[i] = aux[i]
	}
	for i := 0; i < a.R(); i++ {
		highPriorSort(a, runes, aux, count[i], count[i+1]-1, depth+1)
	}
}

func toIndex(a alphabet.Interface, runes []rune, depth int) rune {
	if depth >= len(runes) {
		return -1
	}
	return a.ToIndex(runes[depth])
}
