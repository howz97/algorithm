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
	"github.com/howz97/algorithm/pqueue"
)

func NewWGraph[T any](size uint) *WGraph[T] {
	return &WGraph[T]{
		Graph: NewGraph[T](size),
	}
}

type WGraph[T any] struct {
	*Graph[T]
}

func (g *WGraph[T]) AddEdge(src, dst Id, w Weight) error {
	return g.Graph.addWeightedEdge(src, dst, w)
}

// LazyPrim gets the minimum spanning tree by Lazy-Prim algorithm. g MUST be a connected graph
func (g *WGraph[T]) LazyPrim() (mst *WGraph[T]) {
	pq := pqueue.NewPaired[Weight, *edge](g.NumVert())
	mst = NewWGraph[T](g.NumVert())
	for _, v := range g.vertices {
		mst.AddVertex(v)
	}
	marked := make([]bool, g.NumVert())
	marked[0] = true
	g.IterWAdjacent(0, func(dst Id, w Weight) bool {
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

func lazyPrimVisit[T any](g *WGraph[T], v Id, marked []bool, pq *pqueue.Paired[Weight, *edge]) {
	marked[v] = true
	g.IterWAdjacent(v, func(a Id, w Weight) bool {
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
func (g *WGraph[T]) Prim() (mst *WGraph[T]) {
	marked := make([]bool, g.NumVert())
	pq := pqueue.NewFixable[Weight, Id](g.NumVert())
	mst = NewWGraph[T](g.NumVert())
	for _, v := range g.vertices {
		mst.AddVertex(v)
	}
	marked[0] = true
	g.IterWAdjacent(0, func(a Id, w Weight) bool {
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

func primVisit[T any](g, mst *WGraph[T], v Id, marked []bool, pq *pqueue.Fixable[Weight, Id]) {
	marked[v] = true
	g.IterWAdjacent(v, func(a Id, w Weight) bool {
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
func (g *WGraph[T]) Kruskal() (mst *WGraph[T]) {
	mst = NewWGraph[T](g.NumVert())
	for _, v := range g.vertices {
		mst.AddVertex(v)
	}
	uf := NewUnionFind(int(g.NumVert()))
	pq := pqueue.NewPaired[Weight, *edge](g.NumVert())
	g.IterWEdge(func(src, dst Id, w Weight) bool {
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

func (g *WGraph[T]) ToWDigraph() *WDigraph[T] {
	return &WDigraph[T]{
		Digraph: g.Digraph,
	}
}
