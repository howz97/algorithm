package wgraph

import (
	"fmt"
	"github.com/howz97/algorithm/graphs/graph"
	pqueue "github.com/howz97/algorithm/pqueue/binaryheap"
	"github.com/howz97/algorithm/queue"
	unionfind "github.com/howz97/algorithm/union-find"
)

type WGraph struct {
	graph.Graph
}

func NewWGraph(size int) *WGraph {
	return &WGraph{
		Graph: *graph.New(size),
	}
}

func (g WGraph) AddEdge(src, dst int, w float64) error {
	return g.Graph.AddWEdge(src, dst, w)
}

func (g WGraph) LazyPrim() *MSTForest {
	marked := make([]bool, g.NumVertical())
	f := newMSTForest()
	for i, b := range marked {
		if !b {
			f.addMST(lazyPrim(g, i, marked))
		}
	}
	return f
}

func lazyPrim(g WGraph, v int, marked []bool) *queue.LinkedQueue {
	pq := pqueue.NewBinHeap(g.NumEdge())
	marked[v] = true
	mst := queue.NewLinkedQueue()
	g.IterateAdj(v, func(src int, dst int, w float64) bool {
		pq.Insert(int(w), &Edge{ // fixme pq Cmp
			from:   src,
			to:     dst,
			weight: int(w),
		})
		return true
	})
	for !pq.IsEmpty() {
		m := pq.DelMin()
		e := m.(*Edge)
		if marked[e.from] && marked[e.to] {
			continue
		}
		mst.PushBack(e)
		if !marked[e.from] {
			lazyPrimVisit(g, e.from, marked, pq)
		}
		if !marked[e.to] {
			lazyPrimVisit(g, e.to, marked, pq)
		}
	}
	return mst
}

func lazyPrimVisit(g WGraph, v int, marked []bool, pq *pqueue.BinHeap) {
	marked[v] = true
	g.IterateAdj(v, func(_ int, a int, w float64) bool {
		if !marked[a] {
			pq.Insert(int(w), &Edge{
				from:   v,
				to:     a,
				weight: int(w),
			})
		}
		return true
	})
}

func (g WGraph) Prim() *MSTForest {
	marked := make([]bool, g.NumVertical())
	edgeTo := make([]*Edge, g.NumVertical())
	f := newMSTForest()
	for i, b := range marked {
		if !b {
			f.addMST(prim(g, i, marked, edgeTo))
		}
	}
	return f
}

func prim(g WGraph, v int, marked []bool, edgeTo []*Edge) *queue.LinkedQueue {
	pq := pqueue.NewBinHeap(g.NumVertical() - 1)
	marked[v] = true
	g.IterateAdj(v, func(_ int, a int, w float64) bool {
		pq.Insert(int(w), a)
		edgeTo[a] = &Edge{
			from:   v,
			to:     a,
			weight: int(w),
		}
		return true
	})
	mst := queue.NewLinkedQueue()
	for !pq.IsEmpty() {
		m := pq.DelMin()
		w := m.(int)
		mst.PushBack(edgeTo[w])
		primVisit(g, w, marked, pq, edgeTo)
	}
	return mst
}

func primVisit(g WGraph, v int, marked []bool, pq *pqueue.BinHeap, edgeTo []*Edge) {
	marked[v] = true
	g.IterateAdj(v, func(_ int, a int, wt float64) bool {
		if marked[a] {
			return true
		}
		e := &Edge{
			from:   v,
			to:     a,
			weight: int(wt),
		}
		if edgeTo[a] == nil {
			pq.Insert(int(wt), a)
			edgeTo[a] = e
		} else if e.weight < edgeTo[a].weight {
			pq.Update(e.weight, a)
			edgeTo[a] = e
		}
		return true
	})
}

// Kruskal 该实现仅支持连通图
func (g WGraph) Kruskal() *queue.Queen {
	mst := queue.NewQueen(g.NumVertical() - 1)
	uf := unionfind.NewUF(g.NumVertical())
	pq := pqueue.NewBinHeap(g.NumEdge())
	g.IterateWEdge(func(src int, dst int, w float64) bool {
		pq.Insert(int(w), &Edge{ // fixme pq Cmp Key
			from:   src,
			to:     dst,
			weight: int(w),
		})
		return true
	})
	for !mst.IsFull() {
		min := pq.DelMin()
		minE := min.(*Edge)
		if uf.IsConnected(minE.from, minE.to) {
			continue
		}
		uf.Union(minE.from, minE.to)
		mst.PushBack(minE)
	}
	return mst
}

type MSTForest []*queue.LinkedQueue

func newMSTForest() *MSTForest {
	f := make(MSTForest, 0, 1)
	return &f
}

// 根据连通分量的id获取它的最小生成树
func (f *MSTForest) MST(cc int) *queue.LinkedQueue {
	if cc < 0 || cc >= len(*f) {
		panic(fmt.Sprintf("subgraph %v does not exist", cc))
	}
	return (*f)[cc]
}

func (f *MSTForest) addMST(mst *queue.LinkedQueue) {
	*f = append(*f, mst)
}

func (f *MSTForest) NumConnectedComponent() int {
	return len(*f)
}

type Edge struct {
	from, to int
	weight   int
}
