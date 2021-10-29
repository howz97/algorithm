package sort

// HeapSort -
func HeapSort(data []int) {
	adjust(data)
	for i := len(data) - 1; i > 0; i-- {
		data[0], data[i] = data[i], data[0]
		percolateDown(data[:i], 0)
	}
}

func adjust(h []int) {
	for i := (len(h) - 2) / 2; i >= 0; i-- {
		percolateDown(h, i)
	}
}

func percolateDown(h []int, i int) {
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
