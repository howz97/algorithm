/*
go test -bench="^Benchmark.*_Push" -cpu=1 -benchtime=10000000x github.com/howz97/algorithm/pq
go test -bench="^Benchmark.*_Pop" -cpu=1 -benchtime=1000000x github.com/howz97/algorithm/pq
go test -bench="^Benchmark.*_Merge" -cpu=1 -benchtime=10000x github.com/howz97/algorithm/pq
*/
package pq

import (
	"math/rand"
	"testing"

	"github.com/howz97/algorithm/pq/binomial"
	"github.com/howz97/algorithm/pq/heap"
	"github.com/howz97/algorithm/pq/leftist"
)

func BenchmarkLeftist_Merge(b *testing.B) {
	pq := leftist.New[int]()
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pq2 := leftist.New[int]()
		pq2Size := rnd.Intn(b.N)
		for j := 0; j < pq2Size; j++ {
			pq2.Push(rnd.Int())
		}
		b.StartTimer()

		pq.Merge(pq2)
	}
}

func BenchmarkBinomial_Merge(b *testing.B) {
	pq := binomial.New[int]()
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		pq2 := binomial.New[int]()
		pq2Size := rnd.Intn(b.N)
		for j := 0; j < pq2Size; j++ {
			pq2.Push(rnd.Int())
		}
		b.StartTimer()

		pq.Merge(pq2)
	}
}

func BenchmarkHeap_Push(b *testing.B) {
	pq := heap.New[int, int](0)
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N), i)
	}
}

func BenchmarkLeftist_Push(b *testing.B) {
	pq := leftist.New[int]()
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N))
	}
}

func BenchmarkBinomial_Push(b *testing.B) {
	pq := binomial.New[int]()
	rnd := rand.New(rand.NewSource(0))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N))
	}
}

func BenchmarkHeap_Pop(b *testing.B) {
	pq := heap.New[int, int](0)
	rnd := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N), i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}

func BenchmarkLeftist_Pop(b *testing.B) {
	pq := leftist.New[int]()
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
	pq := binomial.New[int]()
	rnd := rand.New(rand.NewSource(0))
	for i := 0; i < b.N; i++ {
		pq.Push(rnd.Intn(b.N))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pq.Pop()
	}
}
