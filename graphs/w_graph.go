package graphs

import (
	unionfind "github.com/howz97/algorithm/basic/union-find"
	"github.com/howz97/algorithm/pq/heap"
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
	pq := heap.New[float64, *edge](g.NumVert())
	mst = NewWGraph(g.NumVert())
	marked := make([]bool, g.NumVert())
	marked[0] = true
	g.IterateWAdj(0, func(dst int, w float64) bool {
		pq.Push(w, &edge{
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

func lazyPrimVisit(g *WGraph, v int, marked []bool, pq *heap.Heap[float64, *edge]) {
	marked[v] = true
	g.IterateWAdj(v, func(a int, w float64) bool {
		if !marked[a] {
			pq.Push(w, &edge{
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
	pq := heap.New2[float64, int](g.NumVert())
	mst = NewWGraph(g.NumVert())
	marked[0] = true
	g.IterateWAdj(0, func(a int, w float64) bool {
		pq.Push(w, a)
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

func primVisit(g, mst *WGraph, v int, marked []bool, pq *heap.Heap2[float64, int]) {
	marked[v] = true
	g.IterateWAdj(v, func(a int, w float64) bool {
		if marked[a] {
			return true
		}
		orig := mst.Adjacent(a)
		if len(orig) == 0 {
			pq.Push(w, a)
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
	uf := unionfind.NewUF(int(g.NumVert()))
	pq := heap.New[float64, *edge](g.NumVert())
	g.IterateWEdge(func(src int, dst int, w float64) bool {
		pq.Push(w, &edge{
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
