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

import "cmp"

func Merge[Ord cmp.Ordered](data []Ord) {
	aux := make([]Ord, len(data))
	mergeSort(data, aux)
}

func mergeSort[Ord cmp.Ordered](data []Ord, aux []Ord) {
	if len(data) < 2 {
		return
	}
	mid := len(data) >> 1
	mergeSort(data[:mid], aux[:mid])
	mergeSort(data[mid:], aux[mid:])
	merge(data, mid, aux)
}

func merge[Ord cmp.Ordered](data []Ord, mid int, aux []Ord) {
	i, j, k := 0, mid, 0
	for i < mid && j < len(data) {
		if data[i] < data[j] {
			aux[k] = data[i]
			i++
		} else {
			aux[k] = data[j]
			j++
		}
		k++
	}
	if i < mid {
		copy(data[k:], data[i:mid])
	}
	copy(data[:k], aux[:k])
}
