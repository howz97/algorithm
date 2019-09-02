package strsort

import (
	"fmt"
	"github.com/zh1014/algorithm/alphabet"
	"testing"
)

func Test_HighPrior(t *testing.T) {
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
	HighPrior(alphabet.LowerCase, strs)
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
	HighPrior(alphabet.Unicode, strs)
	fmt.Println(strs)
}

func Test_Quick3(t *testing.T) {
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
	Quick3(alphabet.LowerCase, strs)
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
	Quick3(alphabet.Unicode, strs)
	fmt.Println(strs)
}
