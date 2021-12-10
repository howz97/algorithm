package graphs

import (
	"github.com/howz97/algorithm/pq/heap"
	"github.com/howz97/algorithm/queue"
	unionfind "github.com/howz97/algorithm/union-find"
	"github.com/howz97/algorithm/util"
)

type WGraph struct {
	Graph
}

func NewWGraph(size uint) *WGraph {
	return &WGraph{
		Graph: *NewGraph(size),
	}
}

func (g *WGraph) AddEdge(src, dst int, w float64) error {
	return g.Graph.addWeightedEdge(src, dst, w)
}

func (g *WGraph) LazyPrim() (mst *WGraph) {
	pq := heap.New(g.NumVertical())
	mst = NewWGraph(g.NumVertical())
	marked := make([]bool, g.NumVertical())
	marked[0] = true
	g.iterateAdj(0, func(src int, dst int, w float64) bool {
		pq.Push(util.Float(w), &Edge{
			from:   src,
			to:     dst,
			weight: w,
		})
		return true
	})
	for !pq.IsEmpty() {
		e := pq.Pop().(*Edge)
		if marked[e.to] {
			continue
		}
		mst.AddEdge(e.from, e.to, e.weight)
		lazyPrimVisit(g, e.to, marked, pq)
	}
	return mst
}

func lazyPrimVisit(g *WGraph, v int, marked []bool, pq *heap.Heap) {
	marked[v] = true
	g.iterateAdj(v, func(_ int, a int, w float64) bool {
		if !marked[a] {
			pq.Push(util.Float(w), &Edge{
				from:   v,
				to:     a,
				weight: w,
			})
		}
		return true
	})
}

func (g *WGraph) Prim() (mst *WGraph) {
	marked := make([]bool, g.NumVertical())
	edgeTo := make([]*Edge, g.NumVertical())
	pq := heap.New(g.NumVertical())
	mst = NewWGraph(g.NumVertical())
	marked[0] = true
	g.iterateAdj(0, func(_ int, a int, w float64) bool {
		pq.Push(util.Float(w), a)
		edgeTo[a] = &Edge{
			from:   0,
			to:     a,
			weight: w,
		}
		return true
	})
	for !pq.IsEmpty() {
		w := pq.Pop().(int)
		e := edgeTo[w]
		mst.AddEdge(e.from, e.to, e.weight)
		primVisit(g, w, marked, pq, edgeTo, mst)
	}
	return mst
}

func primVisit(g *WGraph, v int, marked []bool, pq *heap.Heap, mst *WGraph) {
	marked[v] = true
	g.iterateAdj(v, func(_ int, a int, wt float64) bool {
		if marked[a] {
			return true
		}
		adj := mst.Adjacent(a)
		if len(adj) == 0 {
			pq.Push(util.Float(wt), a)
			mst.AddEdge(v, a, wt)
		} else {
			w2 := mst.GetWeight(adj, a)
			if wt < w2 {
				pq.Fix(util.Float(wt), a)
				edgeTo[a] = e // repalce edge
			}
		}
		return true
	})
}

// Kruskal 该实现仅支持连通图
func (g *WGraph) Kruskal() *queue.Queen {
	mst := queue.NewQueen(int(g.NumVertical()))
	uf := unionfind.NewUF(int(g.NumVertical()))
	pq := heap.New(g.NumEdge())
	g.IterateWEdge(func(src int, dst int, w float64) bool {
		pq.Push(util.Float(w), &Edge{
			from:   src,
			to:     dst,
			weight: w,
		})
		return true
	})
	for !mst.IsFull() {
		min := pq.Pop()
		minE := min.(*Edge)
		if uf.IsConnected(minE.from, minE.to) {
			continue
		}
		uf.Union(minE.from, minE.to)
		mst.PushBack(minE)
	}
	return mst
}

type Edge struct {
	from, to int
	weight   float64
}
