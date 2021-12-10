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

//func (g *WGraph) Prim() *MSTForest {
//	marked := make([]bool, g.NumVertical())
//	edgeTo := make([]*Edge, g.NumVertical())
//	f := newMSTForest()
//	for i, b := range marked {
//		if !b {
//			f.addMST(prim(g, i, marked, edgeTo))
//		}
//	}
//	return f
//}

func prim(g WGraph, v int, marked []bool, edgeTo []*Edge) *queue.LinkedQueue {
	pq := heap.New(g.NumVertical())
	marked[v] = true
	g.iterateAdj(v, func(_ int, a int, w float64) bool {
		pq.Push(util.Float(w), a)
		edgeTo[a] = &Edge{
			from:   v,
			to:     a,
			weight: w,
		}
		return true
	})
	mst := queue.NewLinkedQueue()
	for !pq.IsEmpty() {
		w := pq.Pop().(int)
		mst.PushBack(edgeTo[w])
		primVisit(g, w, marked, pq, edgeTo)
	}
	return mst
}

func primVisit(g WGraph, v int, marked []bool, pq *heap.Heap, edgeTo []*Edge) {
	marked[v] = true
	g.iterateAdj(v, func(_ int, a int, wt float64) bool {
		if marked[a] {
			return true
		}
		e := &Edge{
			from:   v,
			to:     a,
			weight: wt,
		}
		if edgeTo[a] == nil {
			pq.Push(util.Float(wt), a)
			edgeTo[a] = e
		} else if e.weight < edgeTo[a].weight {
			pq.Fix(util.Float(e.weight), a)
			edgeTo[a] = e
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
