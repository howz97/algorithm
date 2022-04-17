package strings

type BoyerMoore struct {
	pattern string
	right   [byteNum]int
}

func NewBoyerMoore(pattern string) *BoyerMoore {
	bm := &BoyerMoore{
		pattern: pattern,
	}
	for i := 0; i < len(pattern); i++ {
		bm.right[pattern[i]] = i
	}
	return bm
}

func (bm *BoyerMoore) Index(s string) int {
	for i := 0; i < len(s)-bm.LenP(); {
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

func (bm *BoyerMoore) IndexAll(s string) (indices []int) {
	j := 0
	for {
		i := bm.Index(s)
		if i < 0 {
			break
		}
		indices = append(indices, j+i)
		j = j + i + bm.LenP()
		s = s[i+bm.LenP():]
	}
	return
}

func (bm *BoyerMoore) LenP() int {
	return len(bm.pattern)
}
