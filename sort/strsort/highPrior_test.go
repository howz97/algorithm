package strsort

import (
	"fmt"
	"testing"
	"github.com/zh1014/algorithm/alphabet"
)

func Test_HighPriorSort(t *testing.T) {
	strs := []string{
		"she",
		"sells",
		"seashells",
		"by",
		"the",
		"seashore",
		"the",
		"shells",
		"she",
		"sells",
		"are",
		"surely",
		"seashells",
		"xyz",
	}
	HighPriorSort(alphabet.LowerCase, strs)
	fmt.Println(strs)

	strs = append(strs,
		"你好1024",
		"你好1025",
		"%中国",
		"%中 国",
		"&……%#",
		"CHINA",
		"ℹChina",
	)
	HighPriorSort(alphabet.Unicode, strs)
	fmt.Println(strs)
}
