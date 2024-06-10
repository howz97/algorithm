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
	stdsort "sort"
	"testing"

	"github.com/howz97/algorithm/basic"
	"github.com/howz97/algorithm/sort"
	"github.com/howz97/algorithm/util"
)

func TestSCC_IsStronglyConnected(t *testing.T) {
	g := NewDigraph(13)
	g.AddEdge(0, 1)
	g.AddEdge(0, 5)
	g.AddEdge(5, 4)
	g.AddEdge(4, 3)
	g.AddEdge(4, 2)
	g.AddEdge(3, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 0)
	g.AddEdge(6, 0)
	g.AddEdge(6, 4)
	g.AddEdge(6, 9)
	g.AddEdge(9, 10)
	g.AddEdge(10, 12)
	g.AddEdge(12, 9)
	g.AddEdge(9, 11)
	g.AddEdge(11, 12)
	g.AddEdge(11, 4)
	g.AddEdge(7, 6)
	g.AddEdge(7, 8)
	g.AddEdge(8, 7)
	g.AddEdge(8, 9)
	fmt.Println("number of edge: ", g.NumEdge())
	scc := g.SCC()
	fmt.Println("number of SCC:", scc.NumComponents())
	for i := 0; i < int(g.NumVert()); i++ {
		fmt.Printf("SCC ID of vertical(%v): %v\n", i, scc.Comp(i))
	}
	if !scc.IsStronglyConn(1, 1) {
		t.Fatal()
	}
	if !scc.IsStronglyConn(0, 4) {
		t.Fatal()
	}
	if !scc.IsStronglyConn(9, 11) {
		t.Fatal()
	}
	if scc.IsStronglyConn(1, 0) {
		t.Fatal()
	}
	if scc.IsStronglyConn(11, 8) {
		t.Fatal()
	}
}

func TestDFS_Graph(t *testing.T) {
	g := NewGraph(9)
	var err error
	err = g.AddEdge(0, 1)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(0, 3)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(1, 3)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(3, 6)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(4, 7)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(4, 8)
	if err != nil {
		t.Fatal(err)
	}
	err = g.AddEdge(2, 5)
	if err != nil {
		t.Fatal(err)
	}

	dfsResults := [][]int{
		0: {0, 1, 3, 6},
		1: {0, 1, 3, 6},
		2: {2, 5},
		3: {0, 1, 3, 6},
		4: {4, 7, 8},
		5: {2, 5},
		6: {0, 1, 3, 6},
		7: {4, 7, 8},
		8: {4, 7, 8},
	}
	checkDFSResults(t, g.Digraph, dfsResults)
}

func TestDFS_Digraph(t *testing.T) {
	dg, err := LoadDigraph(".\\testdata\\dfs.yml")
	if err != nil {
		t.Fatal(err)
	}
	dfsResults := [][]int{
		0: {0, 3, 6, 7},
		1: {1, 2, 5, 7, 8},
		2: {1, 2, 5, 7, 8},
		3: {3, 6, 7},
		4: {4},
		5: {1, 2, 5, 7, 8},
		6: {6, 7},
		7: {7},
		8: {7, 8},
	}
	checkDFSResults(t, dg, dfsResults)
}

func checkDFSResults(t *testing.T, g *Digraph, dfsResults [][]int) {
	for src := range dfsResults {
		reach := g.ReachableSlice(src)
		sort.Quick(stdsort.IntSlice(reach))
		if !util.SliceEqual(reach, dfsResults[src]) {
			t.Errorf("v %d reach %v not equal %v", src, reach, dfsResults[src])
		}
	}
}

func TestRevDFS(t *testing.T) {
	g, err := LoadDigraph(".\\testdata\\dfs.yml")
	if err != nil {
		t.Fatal(err)
	}
	order := basic.NewStack[int](0)
	g.IterBDFSFrom(0, func(v int) bool {
		order.Push(v)
		return true
	})
	correct := []int{0, 3, 6, 7}
	if !util.SliceEqual(order.Drain(), correct) {
		t.Errorf("rev dfs order %v not equal %v", order, correct)
	}
}

func ExampleDigraph_FindCycle() {
	// (0)-------->(2)
	// 	| ^	        ^
	// 	|  \	    |
	// 	|	------  |
	// 	|		  \	|
	// 	v		   \|
	// (1)-------->(3)
	g := NewDigraph(4)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(1, 3)
	g.AddEdge(3, 0)
	g.AddEdge(3, 2)
	c := g.FindCycle()
	fmt.Println(c.Error())

	// Output: (distance=3): 0->1, 1->3, 3->0,
}

func ExampleDigraph_Topological() {
	dg, err := LoadDigraph(`.\testdata\no_cycle.yml`)
	if err != nil {
		panic(err)
	}
	for _, vet := range dg.Topological().Drain() {
		fmt.Printf("%d->", vet)
	}

	// Output: 5->1->3->6->4->7->0->2->
}

func ExampleDigraph_Bipartite() {
	dg, err := LoadDigraph(`.\testdata\no_cycle.yml`)
	if err != nil {
		panic(err)
	}
	fmt.Println(dg.Bipartite())

	// Output: []
}

func ExampleReachable() {
	dg, err := LoadDigraph(`.\testdata\no_cycle.yml`)
	if err != nil {
		panic(err)
	}
	reach := dg.Reachable()
	fmt.Println(reach.CanReach(5, 2))
	fmt.Println(reach.CanReach(2, 5))

	// Output:
	// true
	// false
}

func ExampleBFS() {
	dg, err := LoadDigraph(`.\testdata\no_cycle.yml`)
	if err != nil {
		panic(err)
	}
	bfs := dg.BFS(1)
	fmt.Println(bfs.CanReach(5))
	fmt.Println(bfs.ShortestPathTo(2).Str(nil))

	// Output:
	// false
	// (distance=3): 1->3, 3->6, 6->2,
}

func ExampleSCC() {
	g := NewDigraph(13)
	g.AddEdge(0, 1)
	g.AddEdge(0, 5)
	g.AddEdge(5, 4)
	g.AddEdge(4, 3)
	g.AddEdge(4, 2)
	g.AddEdge(3, 2)
	g.AddEdge(2, 3)
	g.AddEdge(2, 0)
	g.AddEdge(6, 0)
	g.AddEdge(6, 4)
	g.AddEdge(6, 9)
	g.AddEdge(9, 10)
	g.AddEdge(10, 12)
	g.AddEdge(12, 9)
	g.AddEdge(9, 11)
	g.AddEdge(11, 12)
	g.AddEdge(11, 4)
	g.AddEdge(7, 6)
	g.AddEdge(7, 8)
	g.AddEdge(8, 7)
	g.AddEdge(8, 9)
	scc := g.SCC()
	fmt.Println("amount of strongly connected component:", scc.NumComponents())
	var vertices []int
	scc.IterComponent(scc.Comp(0), func(v int) bool {
		vertices = append(vertices, v)
		return true
	})
	sort.Shell(vertices)
	fmt.Println("vertices strongly connected with 0:", vertices)
	fmt.Println(scc.IsStronglyConn(0, 6))

	// Output:
	// amount of strongly connected component: 5
	// vertices strongly connected with 0: [0 2 3 4 5]
	// false
}
