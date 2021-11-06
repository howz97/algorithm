package string_search

type BoyerMoore struct {
	pattern string
	right   []int
}

func NewBM(pattern string) *BoyerMoore {
	bm := &BoyerMoore{
		pattern: pattern,
		right:   make([]int, byteNum),
	}
	for i := 0; i < len(pattern); i++ {
		bm.right[pattern[i]] = i
	}
	return bm
}

func (bm *BoyerMoore) Index(s string) int {
	lenS := len(s)
	i := 0
	for i < lenS-bm.LenP() {
		j := bm.LenP() - 1
		for ; j >= 0 && s[i+j] == bm.pattern[j]; j-- {
		}
		if j < 0 {
			return i
		}
		skip := j - bm.right[s[i+j]]
		if skip < 1 {
			skip = 1
		}
		i += skip
	}
	return -1
}

func (bm *BoyerMoore) IndexAll(s string) []int {
	indices := make([]int, 0)
	j := 0
	for i := bm.Index(s); i >= 0; i = bm.Index(s[j:]) {
		indices = append(indices, j+i)
		j = j + i + bm.LenP()
	}
	return indices
}

func (bm *BoyerMoore) LenP() int {
	return len(bm.pattern)
}
