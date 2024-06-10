// Copyright 2024 Hao Zhang
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strings

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestBoyerMoore_IndexAll(t *testing.T) {
	pattern := "zhanghao"
	txt := "3454zhanghao3herfzhanghaorje4h54haorje4h543nnr5 zhanghao4n3w ebrh4j3k2jnghao4n3w3h4j3zhank2n^&*4zhanxnmGHJ3h4&*$Nzhh#cb3zhanghaounf2B#N$M%N^---..."
	bm := NewBoyerMoore(pattern)
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

	bm := NewBoyerMoore(pattern)
	all := bm.IndexAll(txt)
	for i := range all {
		if txt[all[i]:all[i]+len(pattern)] != pattern {
			t.Fatal()
		}
	}

	pattern = "你啊"
	bm = NewBoyerMoore(pattern)
	all = bm.IndexAll(txt)
	for i := range all {
		if txt[all[i]:all[i]+len(pattern)] != pattern {
			t.Fatal()
		}
	}
}

func TestBoyerMoore_Index(t *testing.T) {
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
	bm := NewBoyerMoore(pattern)
	start := time.Now()
	i := bm.Index(string(txt))
	elapsed := time.Since(start)
	if i < 0 {
		t.Fatal()
	}
	fmt.Printf("[%v]Found at %v: %v\n", elapsed.String(), i, string(txt[i:i+bm.LenP()]))
}
