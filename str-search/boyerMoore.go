package str_search

type BoyerMoore struct {
	pttrn    string
	lenPttrn int
	right    []int
}

func NewBM(pattern string) *BoyerMoore {
	bm := &BoyerMoore{
		pttrn:    pattern,
		lenPttrn: len(pattern),
		right:    make([]int, byteNum),
	}
	for i := 0; i < len(pattern); i++ {
		bm.right[pattern[i]] = i
	}
	return bm
}

func (bm *BoyerMoore) Index(s string) int {
	lenS := len(s)
	i := bm.lenPttrn - 1
	for i < lenS {
		j := bm.lenPttrn - 1
		k := 0
		for k = 0; k < bm.lenPttrn && s[i-k] == bm.pttrn[j]; k++ {
			j--
		}
		if j < 0 {
			return i - bm.lenPttrn + 1
		}
		iadd := j - bm.right[s[i-k]]
		if iadd < 1 {
			iadd = 1
		}
		i += iadd
	}
	return -1
}

func (bm *BoyerMoore) IndexAll(s string) []int {
	indices := make([]int, 0)
	j := 0
	for i := bm.Index(s); i >= 0; i = bm.Index(s[j:]) {
		indices = append(indices, j+i)
		j = j + i + bm.lenPttrn
	}
	return indices
}
