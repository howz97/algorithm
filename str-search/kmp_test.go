package str_search

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestKMP_IndexAll(t *testing.T) {
	pattern := "zhanghao"
	txt := "3454zhanghao3herfzhanghaorje4h543nnr5 zhanghao4n3w ebrh4j3k2j3h4j3zhank2n^&*4zhanxnmGHJ3h4&*$Nzhh#cb3zhanghaounf2B#N$M%N^---..."
	// 			''''''''     ''''''''             ''''''''                          							 ''''''''
	kmp := NewKMP(pattern)
	all := kmp.IndexAll(txt)
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

func TestKMP_IndexAll2(t *testing.T) {
	pattern := "你好"
	txt := "3454zhanghao3herfzhanghaorje4h543nnr5哈你好呵呵哈好你啊好哈好你呵你哦哈好哈好你好呵哈你哦哈好你啊哈好好呵呵哈你哦"

	kmp := NewKMP(pattern)
	all := kmp.IndexAll(txt)
	for i := range all {
		fmt.Println(txt[all[i] : all[i]+len(pattern)])
	}
}

func TestKMP_Index(t *testing.T) {
	pattern := "It is a far, far better thing that I do, than I have ever done"
	file, err := os.Open("/Users/zhanghao/go/src/github.com/zh1014/algorithm/str-search/tale.txt")
	if err != nil {
		panic(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	txt := make([]byte, fileStat.Size())
	n, err := file.Read(txt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %v byte data from %v\n", n, fileStat.Name())
	kmp := NewKMP(pattern)
	start := time.Now()
	i := kmp.Index(string(txt))
	elapsed := time.Since(start)
	if i < 0 {
		t.Fatal()
	}
	fmt.Printf("[%v]Found at %v: %v\n", elapsed.String(), i, string(txt[i:i+kmp.lenPttrn]))
}

func TestContrast(t *testing.T) {
	pattern := "It is a far, far better thing that I do, than I have ever done"
	file, err := os.Open("/Users/zhanghao/go/src/github.com/zh1014/algorithm/str-search/tale.txt")
	if err != nil {
		panic(err)
	}
	fileStat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	txt := make([]byte, fileStat.Size())
	_, err = file.Read(txt)
	if err != nil {
		panic(err)
	}
	txtStr := string(txt)

	const testCount = 10
	kmp := NewKMP(pattern)
	fmt.Println("KMP:")
	for i := 0; i < testCount; i++ {
		start := time.Now()
		idx := kmp.Index(txtStr)
		elapsed := time.Since(start)
		fmt.Printf("%v Found at %v\n", elapsed.String(), idx)
	}

	bm := NewBM(pattern)
	fmt.Println("\nBoyerMoore:")
	for i := 0; i < testCount; i++ {
		start := time.Now()
		idx := bm.Index(txtStr)
		elapsed := time.Since(start)
		fmt.Printf("%v Found at %v\n", elapsed.String(), idx)
	}

	fmt.Println("\nRabinKarp:")
	for i := 0; i < testCount; i++ {
		start := time.Now()
		idx := IndexRabinKarp(txtStr, pattern)
		elapsed := time.Since(start)
		fmt.Printf("%v Found at %v\n", elapsed.String(), idx)
	}

	fmt.Println("\nbytes.Index(Rabin-Karp):")
	for i := 0; i < testCount; i++ {
		start := time.Now()
		idx := strings.Index(txtStr, pattern)
		elapsed := time.Since(start)
		fmt.Printf("%v Found at %v\n", elapsed.String(), idx)
	}
}
