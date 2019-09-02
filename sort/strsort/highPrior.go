package strsort

var aux [][]rune

// HighPriorSort -
func HighPrior(a alphbt, strs []string) {
	aux = make([][]rune, len(strs))

	// convert string to []rune
	strsRune := make([][]rune, len(strs))
	for i := range strsRune {
		strsRune[i] = []rune(strs[i])
	}

	// sort [][]rune which equal to strs
	highPriorSort(a, strsRune, 0, len(strs)-1, 0)

	// convert []rune to string
	for i := range strs {
		strs[i] = string(strsRune[i])
	}
}

func highPriorSort(a alphbt, strs [][]rune, lo, hi, d int) {
	if lo >= hi {
		return
	}
	count := make([]int, a.R()+2)
	// start counting
	for i := lo; i <= hi; i++ {
		count[toIndex(a, strs[i], d)+2]++
	}

	// convert count to inserting index
	// insrtingIdx[0]是已经结束的runes的插入位置，insrtingIdx[1]~insrtingIdx[R]对应首rune的index为0 ~ R-1 的runes的插入位置
	count[0] = lo
	for i := 1; i <= a.R(); i++ {
		count[i] += count[i-1]
	}
	insrtingIdx := count

	// inserting accord to head rune
	for i := lo; i <= hi; i++ {
		aux[insrtingIdx[toIndex(a, strs[i], d)+1]] = strs[i]
		insrtingIdx[toIndex(a, strs[i], d)+1]++
	}

	// write back
	for i := lo; i <= hi; i++ {
		strs[i] = aux[i]
	}

	// recursion
	for i := 0; i < a.R(); i++ {
		highPriorSort(a, strs, insrtingIdx[i], insrtingIdx[i+1]-1, d+1)
	}
}

// toIndex convert rune to index.
// when d is out of the range of the runes, -1 returned
// otherwise, it is equal to ToIndex
func toIndex(a alphbt, runes []rune, d int) int {
	if d >= len(runes) {
		return -1
	}
	return a.ToIndex(runes[d])
}
