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

func Insert[Ord constraints.Ordered](data []Ord) {
	for i := 1; i < len(data); i++ {
		insrtNum := data[i]
		j := i - 1
		for ; j >= 0 && data[j] > insrtNum; j-- {
			data[j+1] = data[j]
		}
		data[j+1] = insrtNum
	}
}
