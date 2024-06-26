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

import "fmt"

func ExampleQuick3() {
	data := []string{
		"ABABC",
		"ABAAA",
		"BBAAA",
		"CA",
		"HUAWEI",
		"AA",
		"张豪",
		"张三",
		"李四",
		"王麻子",
		"李二狗",
	}
	Quick3(data)
	fmt.Println(data)
	// Output: [AA ABAAA ABABC BBAAA CA HUAWEI 张三 张豪 李二狗 李四 王麻子]
}
