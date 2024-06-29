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
	"fmt"
	"testing"
)

func TestMST_Prim(t *testing.T) {
	g := NewWGraph[int](8) // 算法4th 图4.3.10 (P399)
	g.AddEdge(0, 2, 0.26)
	g.AddEdge(0, 4, 0.38)
	g.AddEdge(0, 6, 0.58)
	g.AddEdge(0, 7, 0.17)
	g.AddEdge(1, 2, 0.36)
	g.AddEdge(1, 3, 0.29)
	g.AddEdge(1, 5, 0.32)
	g.AddEdge(1, 7, 0.19)
	g.AddEdge(2, 3, 0.17)
	g.AddEdge(2, 6, 0.40)
	g.AddEdge(2, 7, 0.34)
	g.AddEdge(3, 6, 0.52)
	g.AddEdge(4, 5, 0.35)
	g.AddEdge(4, 6, 0.93)
	g.AddEdge(4, 7, 0.37)
	g.AddEdge(5, 7, 0.28)
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

func Example() {
	g, err := LoadSymbWGraph(testDir + "mst.yml")
	if err != nil {
		panic(err)
	}
	mst := g.Prim()
	//mst := g.LazyPrim()
	//mst := g.Kruskal()
	fmt.Println(mst.String())

	// Output:
	// possible output:
	// 0 : 2 7
	// 1 : 7
	// 2 : 0 3 6
	// 3 : 2
	// 4 : 5
	// 5 : 7 4
	// 6 : 2
	// 7 : 0 1 5
}
