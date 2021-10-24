package str_search

const (
	byteNum = 256
)

type KMP struct {
	lenPttrn int
	dfa      [][]int
}

func NewKMP(pattern string) *KMP {
	kmp := &KMP{
		dfa: make([][]int, byteNum),
	}
	kmp.lenPttrn = len(pattern)
	for i := range kmp.dfa {
		kmp.dfa[i] = make([]int, kmp.lenPttrn)
	}
	dfa := kmp.dfa
	dfa[pattern[0]][0] = 1
	x := 0
	for i := 1; i < kmp.lenPttrn; i++ {
		for j := 0; j < byteNum; j++ {
			dfa[j][i] = dfa[j][x]
		}
		dfa[pattern[i]][i] = i + 1
		x = dfa[pattern[i]][x]
	}
	return kmp
}

func (kmp *KMP) Index(s string) int {
	lS := len(s)
	i, j := 0, 0
	for ; i < lS && j < kmp.lenPttrn; i++ {
		j = kmp.dfa[s[i]][j]
	}
	if j == kmp.lenPttrn {
		return i - j
	}
	return -1
}

func (kmp *KMP) IndexAll(s string) []int {
	indices := make([]int, 0)
	j := 0
	for i := kmp.Index(s); i >= 0; i = kmp.Index(s[j:]) {
		indices = append(indices, j+i)
		j = j + i + kmp.lenPttrn
	}
	return indices
}
