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

	"github.com/howz97/algorithm/basic"
	"github.com/howz97/algorithm/sort"
	"github.com/howz97/algorithm/util"
)

const testDir = "../assets/graphs/"

func TestSCC_IsStronglyConnected(t *testing.T) {
	g := NewDigraph[string](0)
	m := Populate(g, map[string][]string{
		"0":  {"1", "5"},
		"2":  {"0", "3"},
		"3":  {"2"},
		"4":  {"2", "3"},
		"5":  {"4"},
		"6":  {"0", "4", "9"},
		"7":  {"6", "8"},
		"8":  {"7", "9"},
		"9":  {"10", "11"},
		"10": {"12"},
		"11": {"4", "12"},
		"12": {"9"},
	})
	scc := g.SCC()
	t.Log("number of SCC:", scc.NumComponents())
	for i := Id(0); uint(i) < g.NumVert(); i++ {
		t.Logf("SCC ID of vertical(%v): %v\n", i, scc.Comp(i))
	}
	if !scc.IsStronglyConn(m["1"], m["1"]) {
		t.Fatal()
	}
	if !scc.IsStronglyConn(m["0"], m["4"]) {
		t.Fatal()
	}
	if !scc.IsStronglyConn(m["9"], m["11"]) {
		t.Fatal()
	}
	if scc.IsStronglyConn(m["1"], m["0"]) {
		t.Fatal()
	}
	if scc.IsStronglyConn(m["11"], m["8"]) {
		t.Fatal()
	}
}

func ExampleSCC() {
	g := NewDigraph[int](13)
	for i := 0; i < 13; i++ {
		g.AddVertex(i)
	}
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
	var vertices []Id
	scc.IterComponent(scc.Comp(0), func(v Id) bool {
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

func TestDFS_Graph(t *testing.T) {
	g := NewGraph[int](9)
	for i := 0; i < 9; i++ {
		g.AddVertex(i)
	}
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

	dfsResults := [][]Id{
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

func checkDFSResults[T any](t *testing.T, g *Digraph[T], dfsResults [][]Id) {
	for src := range dfsResults {
		reach := g.ReachableSlice(Id(src))
		sort.Quick(reach)
		if !util.SliceEqual(reach, dfsResults[src]) {
			t.Errorf("v %d reach %v not equal %v", src, reach, dfsResults[src])
		}
	}
}

func TestRevDFS(t *testing.T) {
	g, err := LoadSymbDigraph(testDir + "dfs.yml")
	if err != nil {
		t.Fatal(err)
	}
	order := basic.NewStack[Id](0)
	g.VetBackDfsFrom(g.IdOf("a"), func(v Id) bool {
		order.PushBack(v)
		return true
	})
	correct := []Id{g.IdOf("a"), g.IdOf("d"), g.IdOf("g"), g.IdOf("h")}
	if !util.SliceEqual(order.ToSlice(), correct) {
		t.Errorf("rev dfs order %v not equal %v", order, correct)
	}
}

func ExampleDigraph_FindCycle() {
	// (A)--->(C)
	// 	| ^	   ^
	// 	|  \   |
	// 	|	\  |
	// 	|	 \ |
	// 	v	  \|
	// (B)--->(D)
	g := NewDigraph[string](0)
	Populate(g, map[string][]string{
		"A": {"B", "C"},
		"B": {"D"},
		"C": {},
		"D": {"A", "C"},
	})
	c := g.FindCycle()
	for _, id := range c {
		fmt.Print(g.Vertex(id), " -> ")
	}
	fmt.Println(g.Vertex(c[0]))
}

func ExampleDigraph_Topological() {
	dg, err := LoadSymbDigraph(testDir + "no_cycle.yml")
	if err != nil {
		panic(err)
	}
	for _, vet := range dg.Topological().ToSlice() {
		fmt.Printf("%v -> ", dg.Vertex(vet))
	}

	// Output: F -> B -> D -> G -> E -> H -> A -> C ->
}

func ExampleDigraph_Bipartite() {
	dg, err := LoadSymbDigraph(testDir + "no_cycle.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println(dg.Bipartite())

	// Output: []
}

func ExampleReachable() {
	dg, err := LoadSymbDigraph(testDir + "no_cycle.yml")
	if err != nil {
		panic(err)
	}
	reach := dg.Reachable()
	fmt.Println(reach.CanReach(dg.IdOf("F"), dg.IdOf("C")))
	fmt.Println(reach.CanReach(dg.IdOf("C"), dg.IdOf("F")))

	// Output:
	// true
	// false
}

func ExampleBFS() {
	dg, err := LoadSymbDigraph(testDir + "no_cycle.yml")
	if err != nil {
		panic(err)
	}
	bfs := dg.BFS(dg.IdOf("B"))
	fmt.Println(bfs.CanReach(dg.IdOf("F")))
	fmt.Println(bfs.ShortestPathTo(dg.IdOf("C")).Str())
}
