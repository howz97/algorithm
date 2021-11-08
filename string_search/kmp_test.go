package string_search

import (
	"fmt"
	"io/ioutil"
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
	file, err := os.Open("./tale.txt")
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
	fmt.Printf("[%v]Found at %v: %v\n", elapsed.String(), i, string(txt[i:i+kmp.stateCnt]))
}

const (
	testCount     = 20
	KMPAlg        = "KMP       "
	BoyerMooreAlg = "BoyerMoore"
	RabinKarpAlg  = "RabinKarp "
	Stdlib        = "Stdlib    "
)

func TestPerformance(t *testing.T) {
	content, err := ioutil.ReadFile("./tale.txt")
	if err != nil {
		t.Fatal(err)
	}
	txtStr := string(content)
	pattern := "It is a far, far better thing that I do, than I have ever done"
	idx := strings.Index(txtStr, pattern)
	result := make(map[string][]time.Duration)
	kmp := NewKMP(pattern)
	bm := NewBM(pattern)
	for i := 0; i < testCount; i++ {
		start := time.Now()
		if kmp.Index(txtStr) != idx {
			t.Fatal("wrong result")
		}
		elapsed := time.Since(start)
		result[KMPAlg] = append(result[KMPAlg], elapsed)

		start = time.Now()
		if bm.Index(txtStr) != idx {
			t.Fatal("wrong result")
		}
		elapsed = time.Since(start)
		result[BoyerMooreAlg] = append(result[BoyerMooreAlg], elapsed)

		start = time.Now()
		if i := IndexRabinKarp(txtStr, pattern); i != idx {
			t.Fatalf("wrong result %d, should be %d", i, idx)
		}
		elapsed = time.Since(start)
		result[RabinKarpAlg] = append(result[RabinKarpAlg], elapsed)

		start = time.Now()
		if strings.Index(txtStr, pattern) != idx {
			t.Fatal("wrong result")
		}
		elapsed = time.Since(start)
		result[Stdlib] = append(result[Stdlib], elapsed)
	}

	for alg, sli := range result {
		var avg time.Duration
		for _, dur := range sli {
			avg += dur
		}
		avg /= time.Duration(len(sli))
		t.Logf("%s: avg=%s %v", alg, avg, sli)
	}
}
