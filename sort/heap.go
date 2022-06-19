package sort

import "golang.org/x/exp/constraints"

func Heap[Ord constraints.Ordered](data []Ord) {
	// construct big top heap
	// leaf do not need to sink
	for i := (len(data) - 2) / 2; i >= 0; i-- {
		percolateDown(data, i)
	}
	// pop biggest value in heap to end position
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		percolateDown(data[:i], 0)
	}
}

// big top heap
func percolateDown[Ord constraints.Ordered](h []Ord, i int) {
	k := h[i]
	cavIdx := i
	for {
		if cavIdx*2+1 > len(h)-1 {
			break
		}
		bigC := cavIdx*2 + 1
		if bigC != len(h)-1 && h[bigC+1] > h[bigC] {
			bigC++
		}
		if h[bigC] < k {
			break
		}
		h[cavIdx] = h[bigC]
		cavIdx = bigC
	}
	h[cavIdx] = k
}
