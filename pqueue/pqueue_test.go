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

/*
go test -bench="^Benchmark.*_Push" -cpu=1 -benchtime=10000000x github.com/howz97/algorithm/pqueue
go test -bench="^Benchmark.*_Pop" -cpu=1 -benchtime=1000000x github.com/howz97/algorithm/pqueue
go test -bench="^Benchmark.*_Merge" -cpu=1 -benchtime=10000x github.com/howz97/algorithm/pqueue
*/
package pqueue

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkLeftist_Merge(b *testing.B) {
	pq := NewLeftist[int]()
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pq2 := NewLeftist[int]()
		pq2Size := rnd.Intn(b.N)
		for j := 0; j < pq2Size; j++ {
			pq2.Push(rnd.Int())
		}
		b.StartTimer()

		pq.Merge(pq2)
	}
}

func BenchmarkBinomial_Merge(b *testing.B) {
	pq := NewBinomial[int]()
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pq2 := NewBinomial[int]()
		pq2Size := rnd.Intn(b.N)
		for j := 0; j < pq2Size; j++ {
			pq2.Push(rnd.Int())
		}
		b.StartTimer()

		pq.Merge(pq2)
	}
}

func BenchmarkHeap_Push(b *testing.B) {
	pq := NewPaired[int, int](0)
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.PushPair(rnd.Intn(b.N), i)
	}
}

func BenchmarkLeftist_Push(b *testing.B) {
	pq := NewLeftist[int]()
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N))
	}
}

func BenchmarkBinomial_Push(b *testing.B) {
	pq := NewBinomial[int]()
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N))
	}
}

func BenchmarkHeap_Pop(b *testing.B) {
	pq := NewPaired[int, int](0)
	rnd := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		pq.PushPair(rnd.Intn(b.N), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}

func BenchmarkLeftist_Pop(b *testing.B) {
	pq := NewLeftist[int]()
	rnd := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}

func BenchmarkBinomial_Pop(b *testing.B) {
	pq := NewBinomial[int]()
	rnd := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}

func TestInteger(t *testing.T) {
	rnd := rand.New(rand.NewSource(0))
	var pqs [3]PriorQueue[int]
	pqs[0] = NewHeap[int](0)
	pqs[1] = NewLeftist[int]()
	pqs[2] = NewBinomial[int]()

	pushFn := func(v int) {
		for i := range pqs {
			pqs[i].Push(v)
		}
		var topv [3]int
		for i := range pqs {
			topv[i] = pqs[i].Top()
		}
		if topv[0] != topv[1] || topv[1] != topv[2] {
			t.Fatalf("%v", topv)
		}
	}

	popFn := func() {
		var poped [3]int
		for i := range pqs {
			poped[i] = pqs[i].Pop()
		}
		if poped[0] != poped[1] || poped[1] != poped[2] {
			t.Fatalf("%v", poped)
		}
	}

	const round = 500
	for n := round; n >= 0; n-- {
		for j := 0; j < n; j++ {
			pushFn(rnd.Intn(round * 10))
		}
		for j := 0; j < (round - n); j++ {
			popFn()
		}
		var sizes [3]int
		for i := range pqs {
			sizes[i] = pqs[i].Size()
		}
		if sizes[0] != sizes[1] || sizes[1] != sizes[2] {
			t.Fatalf("%v", sizes)
		}
	}
}

func ExampleFixable() {
	pq := NewFixable[int, string](3)
	pq.PushPair(1, "1")
	pq.PushPair(9, "9")
	pq.PushPair(9, "9")
	pq.PushPair(7, "7")
	for pq.Size() > 0 {
		fmt.Print(pq.Pop())
	}
	fmt.Println()

	pq.PushPair(100, "1")
	pq.PushPair(9, "9")
	pq.PushPair(9, "9")
	pq.PushPair(7, "7")
	pq.PushPair(0, "x")
	pq.Del("x")
	pq.Fix(1, "1")
	for pq.Size() > 0 {
		fmt.Print(pq.Pop())
	}

	// Output:
	// 1799
	// 1799
}

func ExampleLeftist() {
	b := NewLeftist[int]()
	b.Push(1)
	b.Push(9)
	b.Push(9)
	b.Push(7)
	b2 := NewLeftist[int]()
	b2.Push(13)
	b2.Push(11)
	b.Merge(b2)
	for b.Size() > 0 {
		fmt.Print(b.Pop(), ",")
	}

	//Output: 1,7,9,9,11,13,
}

func ExampleBinomial() {
	b := NewBinomial[int]()
	b.Push(1)
	b.Push(9)
	b.Push(9)
	b.Push(7)
	b2 := NewBinomial[int]()
	b2.Push(13)
	b2.Push(11)
	b.Merge(b2)
	for b.Size() > 0 {
		fmt.Print(b.Pop(), ",")
	}

	// Output: 1,7,9,9,11,13,
}
