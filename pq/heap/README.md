Binary heap
```go
func Example() {
	pq := New2[int, string](3)
	pq.Push(1, "1")
	pq.Push(9, "9")
	pq.Push(9, "9")
	pq.Push(7, "7")
	for pq.Size() > 0 {
		fmt.Print(pq.Pop())
	}
	fmt.Println()

	pq.Push(100, "1")
	pq.Push(9, "9")
	pq.Push(9, "9")
	pq.Push(7, "7")
	pq.Push(0, "x")
	pq.Del("x")
	pq.Fix(1, "1")
	for pq.Size() > 0 {
		fmt.Print(pq.Pop())
	}

	// Output:
	// 1799
	// 1799
}
```