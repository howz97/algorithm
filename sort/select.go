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

func Select[Ord cmp.Ordered](data []Ord) {
	for i := 0; i < len(data)-1; i++ {
		idxMin := i
		for j := i + 1; j < len(data); j++ {
			if data[j] < data[idxMin] {
				idxMin = j
			}
		}
		data[i], data[idxMin] = data[idxMin], data[i]
	}
}
