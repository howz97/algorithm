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
