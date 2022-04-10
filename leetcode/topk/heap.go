package topk

// ByHeap put biggest k number to data[:k] by using min-heap
func ByHeap(data []int, k int) {
	if k < 0 {
		panic("k < 0")
	}
	if k == 0 {
		return
	}
	if data == nil || len(data) < k {
		panic("data == nil || len(data) < k")
	}
	if len(data) == k {
		return
	}
	buildMinHeap(data[:k])
	for i := k; i < len(data); i++ {
		if data[i] > data[0] {
			data[0], data[i] = data[i], data[0]
			adjustMinHeap(data[:k], 0)
		}
	}
}

func buildMinHeap(heap []int) {
	for i := (len(heap) - 2) / 2; i >= 0; i-- {
		adjustMinHeap(heap, i)
	}
}

func adjustMinHeap(heap []int, adj int) {
	// fmt.Println("\ncurrent heap:", heap)
	// fmt.Println("adjusting", adj, "th element: ", heap[adj])
	// defer fmt.Println("after adjust: ", heap)
	minSonIdx := -1
	if hasLeftSon(heap, adj) {
		minSonIdx = leftSonIdx(adj)
	}
	if hasRightSon(heap, adj) && heap[rightSonIdx(adj)] < heap[leftSonIdx(adj)] {
		minSonIdx = rightSonIdx(adj)
	}

	if minSonIdx > 0 && heap[minSonIdx] < heap[adj] {
		heap[adj], heap[minSonIdx] = heap[minSonIdx], heap[adj]
		adjustMinHeap(heap, minSonIdx)
	}
}

func parentOf(i int) int {
	return (i - 1) / 2
}

func leftSonIdx(i int) int {
	return 2*i + 1
}

func rightSonIdx(i int) int {
	return 2*i + 2
}

func hasLeftSon(heap []int, adj int) bool {
	return leftSonIdx(adj) < len(heap)
}

func hasRightSon(heap []int, adj int) bool {
	return rightSonIdx(adj) < len(heap)
}
