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
	"math/rand"
	"testing"
)

func TestDijkstra_Simple(t *testing.T) {
	g, err := LoadSymbWDigraph(testDir + "w_digraph.yml")
	fatalIfErr(t, err)
	sps, err := g.SearcherDijkstra()
	fatalIfErr(t, err)
	CheckSearcher(t, sps, g.WDigraph)
}

func TestTopological_Simple(t *testing.T) {
	g, err := LoadSymbWDigraph(testDir + "no_cycle_w.yml")
	fatalIfErr(t, err)
	sps, err := g.SearcherTopological()
	fatalIfErr(t, err)
	CheckSearcher(t, sps, g.WDigraph)
}

func TestBellmanFord_Simple(t *testing.T) {
	g, err := LoadSymbWDigraph(testDir + "w_digraph.yml")
	if err != nil {
		t.Fatal(err)
	}
	sps, err := g.SearcherBellmanFord()
	if err != nil {
		t.Fatal(err)
	}
	CheckSearcher(t, sps, g.WDigraph)
}

func TestNegativeCycle(t *testing.T) {
	// g, err := LoadSymbWDigraph(testDir + "negative_cycle.yml")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// _, err = g.SearcherBellmanFord()
	// if err == nil {
	// 	t.Fatalf("negative cycle exist, error should be received")
	// }
	// fmt.Println("negative cycle detected:", err)
}

func CheckSearcher[T any](t *testing.T, s *Searcher[T], dg *WDigraph[T]) {
	for _, spt := range s.spt {
		CheckSPT(t, spt, dg)
	}
}

func CheckSPT[T any](t *testing.T, tree *PathTree[T], dg *WDigraph[T]) {
	dg.IterWEdge(func(from, to Id, w Weight) bool {
		d1 := tree.distTo[from]
		d2 := tree.distTo[to]
		if d1 < DistanceMax && d1+w < d2 {
			t.Errorf("edge %d->%d should belong to SPT: %v + %v < %v", from, to, d1, w, d2)
			return false
		}
		return true
	})
}

func TestTopological(t *testing.T) {
	LoopTestSearcher(t, 2, func(wd *WDigraph[int]) (*Searcher[int], error) {
		return wd.SearcherTopological()
	})
}

func TestDijkstra(t *testing.T) {
	LoopTestSearcher(t, 10, func(wd *WDigraph[int]) (*Searcher[int], error) {
		return wd.SearcherDijkstra()
	})
}

func TestBellmanFord(t *testing.T) {
	LoopTestSearcher(t, 10, func(wd *WDigraph[int]) (*Searcher[int], error) {
		return wd.SearcherBellmanFord()
	})
}

func TestSearcher(t *testing.T) {
	LoopTestSearcher(t, 10, func(wd *WDigraph[int]) (*Searcher[int], error) {
		return wd.Searcher()
	})
}

func LoopTestSearcher(t *testing.T, edgeLimit int,
	fn func(*WDigraph[int]) (*Searcher[int], error)) {
	for i := 0; i < 10; i++ {
		wd := RandWDigraph(edgeLimit)
		sps, err := fn(wd)
		if err != nil {
			t.Log(err)
			continue
		}
		CheckSearcher(t, sps, wd)
	}
}

func RandWDigraph(edgeLimit int) (wd *WDigraph[int]) {
	wd = NewWDigraph[int](0)
	nv := 200 + rand.Intn(800)
	for i := 0; i < nv; i++ {
		wd.AddVertex(i)
	}
	for from := 0; from < nv; from++ {
		ne := rand.Intn(edgeLimit)
		for j := 0; j < ne; j++ {
			to := rand.Intn(nv)
			if from == to {
				continue
			}
			wd.AddEdge(Id(from), Id(to), Weight(rand.Intn(10000)))
		}
	}
	return
}

func ExampleWDigraph() {
	g, err := LoadSymbWDigraph(testDir + "no_cycle_w.yml")
	panicIfErr(err)

	searcher, err := g.SearcherDijkstra()
	panicIfErr(err)
	fmt.Println(searcher.GetPath(g.IdOf("B"), g.IdOf("C")))

	searcher, err = g.SearcherTopological()
	panicIfErr(err)
	fmt.Println(searcher.GetPath(g.IdOf("B"), g.IdOf("C")))

	searcher, err = g.SearcherBellmanFord()
	panicIfErr(err)
	fmt.Println(searcher.GetPath(g.IdOf("B"), g.IdOf("C")))

	// Output:
	// [Distance=85] B->D(29) D->G(52) G->C(4)
	// [Distance=85] B->D(29) D->G(52) G->C(4)
	// [Distance=85] B->D(29) D->G(52) G->C(4)
}

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func fatalIfErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
