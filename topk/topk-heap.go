package topk

// ByHeap put biggest number in top k to data[:k] by using min-heap
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
	for i := (len(heap) - 1) / 2; i >= 0; i-- {
		adjustMinHeap(heap, i)
	}
}

func adjustMinHeap(heap []int, adj int) {
	minSon := -1
	if leftSonOf(adj) < len(heap)-1 {
		minSon = leftSonOf(adj)
	}
	if rightSonOf(adj) < len(heap)-1 && heap[rightSonOf(adj)] < heap[leftSonOf(adj)] {
		minSon = rightSonOf(adj)
	}

	if minSon > 0 && heap[minSon] < heap[adj] {
		heap[adj], heap[minSon] = heap[minSon], heap[adj]
		adjustMinHeap(heap, minSon)
	}
}

func parentOf(i int) int {
	return (i - 1) / 2
}

func leftSonOf(i int) int {
	return 2*i + 1
}

func rightSonOf(i int) int {
	return 2*i + 2
}
