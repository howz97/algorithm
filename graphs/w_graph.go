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
	"github.com/howz97/algorithm/basic"
	"github.com/howz97/algorithm/pq"
)

func NewWGraph(size uint) *WGraph {
	return &WGraph{
		Graph: NewGraph(size),
	}
}

type WGraph struct {
	*Graph
}

func (g *WGraph) AddEdge(src, dst int, w float64) error {
	return g.Graph.addWeightedEdge(src, dst, w)
}

// LazyPrim gets the minimum spanning tree by Lazy-Prim algorithm. g MUST be a connected graph
func (g *WGraph) LazyPrim() (mst *WGraph) {
	pq := pq.NewPaired[float64, *edge](g.NumVert())
	mst = NewWGraph(g.NumVert())
	marked := make([]bool, g.NumVert())
	marked[0] = true
	g.IterWAdjacent(0, func(dst int, w float64) bool {
		pq.PushPair(w, &edge{
			from:   0,
			to:     dst,
			weight: w,
		})
		return true
	})
	for pq.Size() > 0 {
		e := pq.Pop()
		if marked[e.to] {
			continue
		}
		mst.AddEdge(e.from, e.to, e.weight)
		lazyPrimVisit(g, e.to, marked, pq)
	}
	return
}

func lazyPrimVisit(g *WGraph, v int, marked []bool, pq *pq.Paired[float64, *edge]) {
	marked[v] = true
	g.IterWAdjacent(v, func(a int, w float64) bool {
		if !marked[a] {
			pq.PushPair(w, &edge{
				from:   v,
				to:     a,
				weight: w,
			})
		}
		return true
	})
}

// Prim gets the minimum spanning tree by Prim algorithm. g MUST be a connected graph
func (g *WGraph) Prim() (mst *WGraph) {
	marked := make([]bool, g.NumVert())
	pq := pq.NewFixable[float64, int](g.NumVert())
	mst = NewWGraph(g.NumVert())
	marked[0] = true
	g.IterWAdjacent(0, func(a int, w float64) bool {
		pq.PushPair(w, a)
		mst.AddEdge(0, a, w)
		return true
	})
	for pq.Size() > 0 {
		v := pq.Pop()
		from := mst.Adjacent(v)[0]
		mst.AddEdge(from, v, g.GetWeight(from, v))
		primVisit(g, mst, v, marked, pq)
	}
	return
}

func primVisit(g, mst *WGraph, v int, marked []bool, pq *pq.Fixable[float64, int]) {
	marked[v] = true
	g.IterWAdjacent(v, func(a int, w float64) bool {
		if marked[a] {
			return true
		}
		orig := mst.Adjacent(a)
		if len(orig) == 0 {
			pq.PushPair(w, a)
			mst.AddEdge(v, a, w)
		} else if w < mst.GetWeight(orig[0], a) {
			pq.Fix(w, a)
			mst.DelEdge(orig[0], a)
			mst.AddEdge(v, a, w)
		}
		return true
	})
}

// Kruskal gets the minimum spanning tree by Kruskal algorithm. g MUST be a connected graph
func (g *WGraph) Kruskal() (mst *WGraph) {
	mst = NewWGraph(g.NumVert())
	uf := basic.NewUnionFind(int(g.NumVert()))
	pq := pq.NewPaired[float64, *edge](g.NumVert())
	g.IterWEdge(func(src int, dst int, w float64) bool {
		pq.PushPair(w, &edge{
			from:   src,
			to:     dst,
			weight: w,
		})
		return true
	})
	for mst.NumEdge() < mst.NumVert()-1 {
		minE := pq.Pop()
		if uf.IsConnected(minE.from, minE.to) {
			continue
		}
		uf.Union(minE.from, minE.to)
		mst.AddEdge(minE.from, minE.to, minE.weight)
	}
	return
}

func (g *WGraph) ToWDigraph() *WDigraph {
	return &WDigraph{
		Digraph: g.Digraph,
	}
}
