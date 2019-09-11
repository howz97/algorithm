package str_search

import (
	"fmt"
	"testing"
)

func TestBoyerMoore_IndexAll(t *testing.T) {
	pattern := "zhanghao"
	txt := "3454zhanghao3herfzhanghaorje4h54haorje4h543nnr5 zhanghao4n3w ebrh4j3k2jnghao4n3w3h4j3zhank2n^&*4zhanxnmGHJ3h4&*$Nzhh#cb3zhanghaounf2B#N$M%N^---..."
	bm := NewBM(pattern)
	all := bm.IndexAll(txt)
	dots := make([]byte, len(txt))
	for i := range dots {
		dots[i] = ' '
	}
	for i := range all {
		for j := all[i]; j < all[i]+len(pattern); j++ {
			dots[j] = '\''
		}
	}
	fmt.Println(txt)
	fmt.Println(string(dots))
}

func TestBoyerMoore_IndexAll2(t *testing.T) {
	pattern := "你好"
	txt := "3454zhanghao3herfzhanghaorje4h543nnr5哈你好呵呵哈好你啊好你啊好哈好好你呵你哦哈好哈你呵你哦哈好哈好你啊好你啊好哈好好你好呵哈你哦哈好你啊哈好好呵呵哈你哦"

	bm := NewBM(pattern)
	all := bm.IndexAll(txt)
	for i := range all {
		if txt[all[i]:all[i]+len(pattern)] != pattern {
			t.Fatal()
		}
	}

	pattern = "你啊"
	bm = NewBM(pattern)
	all = bm.IndexAll(txt)
	for i := range all {
		if txt[all[i]:all[i]+len(pattern)] != pattern {
			t.Fatal()
		}
	}
}
