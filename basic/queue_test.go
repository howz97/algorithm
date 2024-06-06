package basic

import (
	"testing"
)

func BenchmarkQueue_PushBack(b *testing.B) {
	q := NewQueue[int](b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
}

func BenchmarkQueue_PopFront(b *testing.B) {
	q := NewQueue[int](b.N)
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.PopFront()
	}
}

func BenchmarkLinkQueue_PushBack(b *testing.B) {
	q := NewLinkQueue[int]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
}

func BenchmarkLinkQueue_PopFront(b *testing.B) {
	q := NewLinkQueue[int]()
	for i := 0; i < b.N; i++ {
		q.PushBack(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.PopFront()
	}
}
