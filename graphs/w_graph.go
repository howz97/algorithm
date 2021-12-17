package graphs

import (
	"github.com/howz97/algorithm/pq/heap"
	unionfind "github.com/howz97/algorithm/union-find"
	"github.com/howz97/algorithm/util"
)

func NewWGraph(size uint) *WGraph {
	return &WGraph{
		Graph: *NewGraph(size),
	}
}

type WGraph struct {
	Graph
}

func (g *WGraph) AddEdge(src, dst int, w float64) error {
	return g.Graph.addWeightedEdge(src, dst, w)
}

// LazyPrim gets the minimum spanning tree by Lazy-Prim algorithm. g MUST be a connected graph
func (g *WGraph) LazyPrim() (mst *WGraph) {
	pq := heap.New(g.NumVert())
	mst = NewWGraph(g.NumVert())
	marked := make([]bool, g.NumVert())
	marked[0] = true
	g.iterateWAdj(0, func(dst int, w float64) bool {
		pq.Push(util.Float(w), &edge{
			from:   0,
			to:     dst,
			weight: w,
		})
		return true
	})
	for pq.Size() > 0 {
		e := pq.Pop().(*edge)
		if marked[e.to] {
			continue
		}
		mst.AddEdge(e.from, e.to, e.weight)
		lazyPrimVisit(g, e.to, marked, pq)
	}
	return
}

func lazyPrimVisit(g *WGraph, v int, marked []bool, pq *heap.Heap) {
	marked[v] = true
	g.iterateWAdj(v, func(a int, w float64) bool {
		if !marked[a] {
			pq.Push(util.Float(w), &edge{
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
	pq := heap.New(g.NumVert())
	mst = NewWGraph(g.NumVert())
	marked[0] = true
	g.iterateWAdj(0, func(a int, w float64) bool {
		pq.Push(util.Float(w), a)
		mst.AddEdge(0, a, w)
		return true
	})
	for pq.Size() > 0 {
		v := pq.Pop().(int)
		from := mst.Adjacent(v)[0]
		mst.AddEdge(from, v, g.getWeightMust(from, v))
		primVisit(g, mst, v, marked, pq)
	}
	return
}

func primVisit(g, mst *WGraph, v int, marked []bool, pq *heap.Heap) {
	marked[v] = true
	g.iterateWAdj(v, func(a int, w float64) bool {
		if marked[a] {
			return true
		}
		orig := mst.Adjacent(a)
		if len(orig) == 0 {
			pq.Push(util.Float(w), a)
			mst.AddEdge(v, a, w)
		} else if w < mst.getWeightMust(orig[0], a) {
			pq.Fix(util.Float(w), a)
			mst.DelEdge(orig[0], a)
			mst.AddEdge(v, a, w)
		}
		return true
	})
}

// Kruskal gets the minimum spanning tree by Kruskal algorithm. g MUST be a connected graph
func (g *WGraph) Kruskal() (mst *WGraph) {
	mst = NewWGraph(g.NumVert())
	uf := unionfind.NewUF(int(g.NumVert()))
	pq := heap.New(g.NumVert())
	g.IterateWEdge(func(src int, dst int, w float64) bool {
		pq.Push(util.Float(w), &edge{
			from:   src,
			to:     dst,
			weight: w,
		})
		return true
	})
	for mst.NumEdge() < mst.NumVert()-1 {
		minE := pq.Pop().(*edge)
		if uf.IsConnected(minE.from, minE.to) {
			continue
		}
		uf.Union(minE.from, minE.to)
		mst.AddEdge(minE.from, minE.to, minE.weight)
	}
	return
}
