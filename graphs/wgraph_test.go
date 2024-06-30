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

package graphs

import (
	"testing"
)

func TestMST_Prim(t *testing.T) {
	g := NewWGraph[int](8) // 算法4th 图4.3.10 (P399)
	WPopulate(g, map[int]map[int]Weight{
		0: {2: 26, 4: 38, 6: 58},
		1: {2: 36, 3: 29, 5: 32, 7: 19},
		2: {3: 17, 6: 40, 7: 34},
		3: {6: 52},
		4: {5: 35, 6: 93, 7: 37},
		5: {7: 28},
	})
	w0 := g.LazyPrim().TotalWeight()
	w1 := g.Prim().TotalWeight()
	if w0 != w1 {
		t.Fatalf("weight %v not equal %v", w0, w1)
	}
	w2 := g.Kruskal().TotalWeight()
	if w0 != w2 {
		t.Fatalf("weight %v not equal %v", w0, w2)
	}
	t.Logf("MST %v:\n%s \n", w0, g.LazyPrim().String())
}
