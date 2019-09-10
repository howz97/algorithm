package str_search

const (
	byteNum = 256
)

type KMP struct {
	lPttrn int
	dfa [][]int
}

func New(pattern string) *KMP {
	kmp := &KMP{
		dfa: make([][]int, byteNum),
	}
	kmp.lPttrn = len(pattern)
	for i := range kmp.dfa {
		kmp.dfa[i] = make([]int, kmp.lPttrn)
	}
	dfa := kmp.dfa
	dfa[pattern[0]][0] = 1
	x := 0
	for i := 1; i < kmp.lPttrn; i++ {
		for j := 0; j < byteNum; j++ {
			dfa[j][i] = dfa[j][x]
		}
		dfa[pattern[i]][i] = i+1
		x = dfa[pattern[i]][x]
	}
	return kmp
}

func (kmp *KMP) Index(s string) int {
	lS := len(s)
	i, j := 0, 0
	for ; i < lS && j < kmp.lPttrn; i++ {
		j = kmp.dfa[s[i]][j]
	}
	if j == kmp.lPttrn {
		return i - j
	}
	return -1
}

func (kmp *KMP) IndexAll(s string) []int {
	indices := make([]int, 0)
	j := 0
	for i := kmp.Index(s);i >=0 ;i = kmp.Index(s[j:]){
		indices = append(indices, j+i)
		j = j + i + kmp.lPttrn
	}
	return indices
}