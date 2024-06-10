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

func Shell[Ord cmp.Ordered](data []Ord) {
	// shell increment
	for incre := len(data) >> 1; incre > 1; incre >>= 1 {
		// sort $incre arries with interval of $incre in one loop
		for i := incre; i < len(data); i++ {
			insrtNum := data[i]
			j := i - incre
			for ; j >= 0 && data[j] > insrtNum; j -= incre {
				data[j+incre] = data[j]
			}
			data[j+incre] = insrtNum
		}
	}
	Insert(data)
}

//func ShellSort(data []int) {
//	// shell increment
//	for incre := len(data) >> 1; incre > 0; incre >>= 1 {
//		for i := incre; i < len(data); i++ {
//			insrtNum := data[i]
//			j := i-incre
//			for ; j >= 0 && data[j] > insrtNum; j -= incre {
//				data[j+incre] = data[j]
//			}
//			data[j+incre] = insrtNum
//		}
//	}
//}
